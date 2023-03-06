// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./VRF.sol";
import "./VRFConsumerBase.sol";
import "./VRFCoordinatorV2Interface.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract VRFCoordinator is VRF, VRFCoordinatorV2Interface, OwnableUpgradeable {
    // We use the config for the mgmt APIs
    struct Subscription {
        address owner; // Owner can fund/withdraw/cancel the sub.
        uint64 reqCount;
        // Maintains the list of keys in s_consumers.
        // We do this for 2 reasons:
        // 1. To be able to clean up all keys from s_consumers when canceling a subscription.
        // 2. To be able to return the list of all consumers in getSubscription.
        // Note that we need the s_consumers map to be able to directly check if a
        // consumer is valid without reading all the consumers from storage.
        address[] consumers;
    }
    struct RequestCommitment {
        uint64 blockNum;
        uint64 subId;
        uint32 callbackGasLimit;
        uint32 numWords;
        address sender;
    }
    struct Config {
        uint16 minimumRequestConfirmations;
        uint32 maxGasLimit;
        // Reentrancy protection.
        bool reentrancyLock;
    }

    // We need to maintain a list of consuming addresses.
    // This bound ensures we are able to loop over them as needed.
    // Should a user require more consumers, they can use multiple subscriptions.
    uint16 public constant MAX_CONSUMERS = 100;
    // Set this maximum to 200 to give us a 56 block window to fulfill
    // the request before requiring the block hash feeder.
    uint16 public constant MAX_REQUEST_CONFIRMATIONS = 200;
    uint32 public constant MAX_NUM_WORDS = 500;
    // 5k is plenty for an EXTCODESIZE call (2600) + warm CALL (100)
    // and some arithmetic operations.
    uint256 private constant GAS_FOR_CALL_EXACT_CHECK = 5_000;

    error TooManyConsumers();
    error InvalidConsumer(uint64 subId, address consumer);
    error InvalidSubscription();
    error MustBeSubOwner(address owner);
    error PendingRequestExists();
    error InvalidRequestConfirmations(uint16 have, uint16 min, uint16 max);
    error GasLimitTooBig(uint32 have, uint32 want);
    error NumWordsTooBig(uint32 have, uint32 want);
    error NoSuchProvingKey(bytes32 keyHash);
    error NoCorrespondingRequest();
    error IncorrectCommitment();
    error Reentrant();
    error ProvingKeyAlreadyRegistered(bytes32 keyHash);
    error BlockhashNotFound(uint256 blockNum);

    event ProvingKeyRegistered(bytes32 keyHash);
    event ProvingKeyDeregistered(bytes32 keyHash);
    event SubscriptionCreated(uint64 indexed subId, address owner);
    event SubscriptionConsumerAdded(uint64 indexed subId, address consumer);
    event SubscriptionConsumerRemoved(uint64 indexed subId, address consumer);
    event SubscriptionCanceled(
        uint64 indexed subId,
        address to,
        uint256 amount
    );
    event RandomWordsFulfilled(
        uint256 indexed requestId,
        uint256 outputSeed,
        uint96 payment,
        bool success
    );
    event RandomWordsRequested(
        bytes32 indexed keyHash,
        uint256 requestId,
        uint256 preSeed,
        uint64 indexed subId,
        uint16 minimumRequestConfirmations,
        uint32 callbackGasLimit,
        uint32 numWords,
        address indexed sender
    );

    // We make the sub count public so that its possible to
    // get all the current subscriptions via getSubscription.
    uint64 private s_currentSubId;
    bytes32[] private s_provingKeyHashes;
    // Note a nonce of 0 indicates an the consumer is not assigned to that subscription.
    // consumer -> subId -> nonce
    mapping(address => mapping(uint64 => uint64)) private s_consumers;
    // subId -> subscription
    mapping(uint64 => Subscription) private s_subscriptions;
    // keyHash -> bool
    mapping(bytes32 => bool) private s_provingKeys;
    // requestID -> commitment
    mapping(uint256 => bytes32) private s_requestCommitments;
    // block height -> block hash
    mapping(uint256 => bytes32) private s_blockStore;

    Config private s_config;

    modifier onlySubOwner(uint64 subId) {
        address owner = s_subscriptions[subId].owner;
        if (owner == address(0)) {
            revert InvalidSubscription();
        }
        if (msg.sender != owner) {
            revert MustBeSubOwner(owner);
        }
        _;
    }

    modifier onlySubOwnerOrAdmin(uint64 subId) {
        if (msg.sender == owner()) {
            _;
        } else {
            address subOwner = s_subscriptions[subId].owner;
            if (subOwner == address(0)) {
                revert InvalidSubscription();
            }
            if (msg.sender != subOwner) {
                revert MustBeSubOwner(subOwner);
            }
            _;
        }
    }

    modifier nonReentrant() {
        if (s_config.reentrancyLock) {
            revert Reentrant();
        }
        _;
    }

    function initialize(uint16 minRequestConfirmations, uint32 maxGasLimit) public initializer {
        __Ownable_init();
        s_config = Config({
            minimumRequestConfirmations: minRequestConfirmations,
            maxGasLimit: maxGasLimit,
            reentrancyLock: false
        });
    }

    /**
     * @notice Returns the proving key hash key associated with this public key
     * @param publicKey the key to return the hash of
     */
    function hashOfKey(uint256[2] memory publicKey)
        public
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(publicKey));
    }

    /**
     * @notice Registers a proving key to an oracle.
     * @param publicProvingKey key that oracle can use to submit vrf fulfillments
     */
    function registerProvingKey(uint256[2] calldata publicProvingKey)
        external
        onlyOwner
    {
        bytes32 kh = hashOfKey(publicProvingKey);
        if (s_provingKeys[kh]) {
            revert ProvingKeyAlreadyRegistered(kh);
        }
        s_provingKeys[kh] = true;
        s_provingKeyHashes.push(kh);
        emit ProvingKeyRegistered(kh);
    }

    /**
     * @notice upload blockhash.
     * @param height block height
     * @param hash block hash
     */
    function storeBlockHash(uint256 height ,bytes32 hash)
        external
        onlyOwner
    {
       s_blockStore[height] = hash;
    }

    /**
     * @notice get blockhash.
     * @param height block height
     */
    function getBlockHash(uint256 height)
        public
        view
        returns (bytes32)
    {
       return s_blockStore[height];
    }

    /**
     * @notice Deregisters a proving key to an oracle.
     * @param publicProvingKey key that oracle can use to submit vrf fulfillments
     */
    function deregisterProvingKey(uint256[2] calldata publicProvingKey)
        external
        onlyOwner
    {
        bytes32 kh = hashOfKey(publicProvingKey);
        if (!s_provingKeys[kh]) {
            revert NoSuchProvingKey(kh);
        }
        delete s_provingKeys[kh];
        for (uint256 i = 0; i < s_provingKeyHashes.length; i++) {
            if (s_provingKeyHashes[i] == kh) {
                bytes32 last = s_provingKeyHashes[
                    s_provingKeyHashes.length - 1
                ];
                // Copy last element and overwrite kh to be deleted with it
                s_provingKeyHashes[i] = last;
                s_provingKeyHashes.pop();
            }
        }
        emit ProvingKeyDeregistered(kh);
    }

    /**
     * @notice Sets the configuration of the vrfv2 coordinator
     * @param minimumRequestConfirmations global min for request confirmations
     * @param maxGasLimit global max for request gas limit
     */
    function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit)
        external
        onlyOwner
    {
        if (minimumRequestConfirmations > MAX_REQUEST_CONFIRMATIONS) {
            revert InvalidRequestConfirmations(
                minimumRequestConfirmations,
                minimumRequestConfirmations,
                MAX_REQUEST_CONFIRMATIONS
            );
        }
        s_config = Config({
            minimumRequestConfirmations: minimumRequestConfirmations,
            maxGasLimit: maxGasLimit,
            reentrancyLock: false
        });
    }

    function getConfig()
        external
        view
        returns (uint16 minimumRequestConfirmations, uint32 maxGasLimit)
    {
        return (s_config.minimumRequestConfirmations, s_config.maxGasLimit);
    }

    /**
     * @inheritdoc VRFCoordinatorV2Interface
     */
    function getRequestConfig()
        external
        view
        override
        returns (
            uint16,
            uint32,
            bytes32[] memory
        )
    {
        return (
            s_config.minimumRequestConfirmations,
            s_config.maxGasLimit,
            s_provingKeyHashes
        );
    }

    function requestRandomWords(
        bytes32 keyHash,
        uint64 subId,
        uint16 requestConfirmations,
        uint32 callbackGasLimit,
        uint32 numWords
    ) external override nonReentrant returns (uint256) {
        // Input validation using the subscription storage.
        if (s_subscriptions[subId].owner == address(0)) {
            revert InvalidSubscription();
        }
        // Its important to ensure that the consumer is in fact who they say they
        // are, otherwise they could use someone else's subscription balance.
        // A nonce of 0 indicates consumer is not allocated to the sub.
        uint64 currentNonce = s_consumers[msg.sender][subId];
        if (currentNonce == 0) {
            revert InvalidConsumer(subId, msg.sender);
        }
        // Input validation using the config storage word.
        if (
            requestConfirmations < s_config.minimumRequestConfirmations ||
            requestConfirmations > MAX_REQUEST_CONFIRMATIONS
        ) {
            revert InvalidRequestConfirmations(
                requestConfirmations,
                s_config.minimumRequestConfirmations,
                MAX_REQUEST_CONFIRMATIONS
            );
        }
        // No lower bound on the requested gas limit. A user could request 0
        // and they would simply be billed for the proof verification and wouldn't be
        // able to do anything with the random value.
        if (callbackGasLimit > s_config.maxGasLimit) {
            revert GasLimitTooBig(callbackGasLimit, s_config.maxGasLimit);
        }
        if (numWords > MAX_NUM_WORDS) {
            revert NumWordsTooBig(numWords, MAX_NUM_WORDS);
        }
        // Note we do not check whether the keyHash is valid to save gas.
        // The consequence for users is that they can send requests
        // for invalid keyHashes which will simply not be fulfilled.
        uint64 nonce = currentNonce + 1;
        (uint256 requestId, uint256 preSeed) = computeRequestId(
            keyHash,
            msg.sender,
            subId,
            nonce
        );

        s_requestCommitments[requestId] = keccak256(
            abi.encode(
                requestId,
                block.number,
                subId,
                callbackGasLimit,
                numWords,
                msg.sender
            )
        );
        emit RandomWordsRequested(
            keyHash,
            requestId,
            preSeed,
            subId,
            requestConfirmations,
            callbackGasLimit,
            numWords,
            msg.sender
        );
        s_consumers[msg.sender][subId] = nonce;

        return requestId;
    }

    /**
     * @notice Get request commitment
     * @param requestId id of request
     * @dev used to determine if a request is fulfilled or not
     */
    function getCommitment(uint256 requestId) external view returns (bytes32) {
        return s_requestCommitments[requestId];
    }

    function computeRequestId(
        bytes32 keyHash,
        address sender,
        uint64 subId,
        uint64 nonce
    ) private pure returns (uint256, uint256) {
        uint256 preSeed = uint256(
            keccak256(abi.encode(keyHash, sender, subId, nonce))
        );
        return (uint256(keccak256(abi.encode(keyHash, preSeed))), preSeed);
    }

    /**
     * @dev calls target address with exactly gasAmount gas and data as calldata
     * or reverts if at least gasAmount gas is not available.
     */
    function callWithExactGas(
        uint256 gasAmount,
        address target,
        bytes memory data
    ) private returns (bool success) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            let g := gas()
            // Compute g -= GAS_FOR_CALL_EXACT_CHECK and check for underflow
            // The gas actually passed to the callee is min(gasAmount, 63//64*gas available).
            // We want to ensure that we revert if gasAmount >  63//64*gas available
            // as we do not want to provide them with less, however that check itself costs
            // gas.  GAS_FOR_CALL_EXACT_CHECK ensures we have at least enough gas to be able
            // to revert if gasAmount >  63//64*gas available.
            if lt(g, GAS_FOR_CALL_EXACT_CHECK) {
                revert(0, 0)
            }
            g := sub(g, GAS_FOR_CALL_EXACT_CHECK)
            // if g - g//64 <= gasAmount, revert
            // (we subtract g//64 because of EIP-150)
            if iszero(gt(sub(g, div(g, 64)), gasAmount)) {
                revert(0, 0)
            }
            // solidity calls check that a contract actually exists at the destination, so we do the same
            if iszero(extcodesize(target)) {
                revert(0, 0)
            }
            // call and return whether we succeeded. ignore return data
            // call(gas,addr,value,argsOffset,argsLength,retOffset,retLength)
            success := call(
                gasAmount,
                target,
                0,
                add(data, 0x20),
                mload(data),
                0,
                0
            )
        }
        return success;
    }

    function getRandomnessFromProof(
        Proof memory proof,
        RequestCommitment memory rc
    )
        private
        view
        returns (
            bytes32 keyHash,
            uint256 requestId,
            uint256 randomness
        )
    {
        keyHash = hashOfKey(proof.pk);
        // Only registered proving keys are permitted.
        if (!s_provingKeys[keyHash]) {
            revert NoSuchProvingKey(keyHash);
        }
        requestId = uint256(keccak256(abi.encode(keyHash, proof.seed)));
        bytes32 commitment = s_requestCommitments[requestId];
        if (commitment == 0) {
            revert NoCorrespondingRequest();
        }
        if (
            commitment !=
            keccak256(
                abi.encode(
                    requestId,
                    rc.blockNum,
                    rc.subId,
                    rc.callbackGasLimit,
                    rc.numWords,
                    rc.sender
                )
            )
        ) {
            revert IncorrectCommitment();
        }

        bytes32 blockHash = blockhash(rc.blockNum);
        if (blockHash == bytes32(0)) {
            blockHash = s_blockStore[rc.blockNum];
            if (blockHash == bytes32(0)) {
                revert BlockhashNotFound(rc.blockNum);
            }
        }
        // The seed actually used by the VRF machinery, mixing in the blockhash
        uint256 actualSeed = uint256(
            keccak256(abi.encodePacked(proof.seed, blockHash))
        );
        randomness = VRF.randomValueFromVRFProof(proof, actualSeed); // Reverts on failure
    }

    /*
     * @notice Fulfill a randomness request
     * @param proof contains the proof and randomness
     * @param rc request commitment pre-image, committed to at request time
     * @return payment amount billed to the subscription
     * @dev simulated offchain to determine if sufficient balance is present to fulfill the request
     */
    function fulfillRandomWords(Proof memory proof, RequestCommitment memory rc)
        external
        nonReentrant
        returns (uint96)
    {
        (, uint256 requestId, uint256 randomness) = getRandomnessFromProof(
            proof,
            rc
        );

        uint256[] memory randomWords = new uint256[](rc.numWords);
        for (uint256 i = 0; i < rc.numWords; i++) {
            randomWords[i] = uint256(keccak256(abi.encode(randomness, i)));
        }

        delete s_requestCommitments[requestId];
        VRFConsumerBase v;
        bytes memory resp = abi.encodeWithSelector(
            v.rawFulfillRandomWords.selector,
            requestId,
            randomWords
        );
        // Call with explicitly the amount of callback gas requested
        // Important to not let them exhaust the gas budget and avoid oracle payment.
        // Do not allow any non-view/non-pure coordinator functions to be called
        // during the consumers callback code via reentrancyLock.
        // Note that callWithExactGas will revert if we do not have sufficient gas
        // to give the callee their requested amount.
        s_config.reentrancyLock = true;
        bool success = callWithExactGas(rc.callbackGasLimit, rc.sender, resp);
        s_config.reentrancyLock = false;

        //remove unusde data
        delete s_blockStore[rc.blockNum];
        s_subscriptions[rc.subId].reqCount += 1;
        // Include payment in the event for tracking costs.
        emit RandomWordsFulfilled(requestId, randomness, 0, success);
        return 0;
    }

    function getSubscription(uint64 subId)
        external
        view
        override
        returns (
            uint96 balance,
            uint64 reqCount,
            address owner,
            address[] memory consumers
        )
    {
        if (s_subscriptions[subId].owner == address(0)) {
            revert InvalidSubscription();
        }
        return (
            0,
            s_subscriptions[subId].reqCount,
            s_subscriptions[subId].owner,
            s_subscriptions[subId].consumers
        );
    }

    /**
     * @inheritdoc VRFCoordinatorV2Interface
     */
    function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner)
        external
        override
        onlySubOwner(subId)
        nonReentrant
    {}

    /**
     * @inheritdoc VRFCoordinatorV2Interface
     */
    function acceptSubscriptionOwnerTransfer(uint64 subId)
        external
        override
        nonReentrant
    {}

    /**
     * @notice Create a VRF subscription.
     * @return subId - A unique subscription id.
     * @dev You can manage the consumer set dynamically with addConsumer/removeConsumer.
     * @dev Note to fund the subscription, use transferAndCall. For example
     * @dev  LINKTOKEN.transferAndCall(
     * @dev    address(COORDINATOR),
     * @dev    amount,
     * @dev    abi.encode(subId));
     */
    function createSubscription()
        external
        override
        nonReentrant
        returns (uint64 subId)
    {
        s_currentSubId++;
        uint64 currentSubId = s_currentSubId;
        address[] memory consumers = new address[](0);
        s_subscriptions[currentSubId] = Subscription({
            owner: msg.sender,
            reqCount: 0,
            consumers: consumers
        });

        emit SubscriptionCreated(currentSubId, msg.sender);
        return currentSubId;
    }

    /**
     * @notice Create a VRF subscription.
     * @return subId - A unique subscription id.
     * @dev You can manage the consumer set dynamically with addConsumer/removeConsumer.
     * @dev Note to fund the subscription, use transferAndCall. For example
     * @dev  LINKTOKEN.transferAndCall(
     * @dev    address(COORDINATOR),
     * @dev    amount,
     * @dev    abi.encode(subId));
     */
    function createSubscription(address subscriber)
        external
        nonReentrant
        onlyOwner
        returns (uint64 subId)
    {
        s_currentSubId++;
        uint64 currentSubId = s_currentSubId;
        address[] memory consumers = new address[](0);
        s_subscriptions[currentSubId] = Subscription({
            owner: subscriber,
            reqCount: 0,
            consumers: consumers
        });

        emit SubscriptionCreated(currentSubId, subscriber);
        return currentSubId;
    }

    /*
     * @notice remove subscriber consumption contract
     * @param subId the id  of subscriber
     * @param consumer the address  of consumption contract
     */
    function removeConsumer(uint64 subId, address consumer)
        external
        override
        onlySubOwnerOrAdmin(subId)
        nonReentrant
    {
        if (s_consumers[consumer][subId] == 0) {
            revert InvalidConsumer(subId, consumer);
        }
        // Note bounded by MAX_CONSUMERS
        address[] memory consumers = s_subscriptions[subId].consumers;
        uint256 lastConsumerIndex = consumers.length - 1;
        for (uint256 i = 0; i < consumers.length; i++) {
            if (consumers[i] == consumer) {
                address last = consumers[lastConsumerIndex];
                // Storage write to preserve last element
                s_subscriptions[subId].consumers[i] = last;
                // Storage remove last element
                s_subscriptions[subId].consumers.pop();
                break;
            }
        }
        delete s_consumers[consumer][subId];
        emit SubscriptionConsumerRemoved(subId, consumer);
    }

    /*
     * @notice add subscriber consumption contract
     * @param subId the id  of subscriber
     * @param consumer the address  of consumption contract
     */
    function addConsumer(uint64 subId, address consumer)
        external
        override
        onlySubOwnerOrAdmin(subId)
        nonReentrant
    {
        // Already maxed, cannot add any more consumers.
        if (s_subscriptions[subId].consumers.length == MAX_CONSUMERS) {
            revert TooManyConsumers();
        }
        if (s_consumers[consumer][subId] != 0) {
            // Idempotence - do nothing if already added.
            // Ensures uniqueness in s_subscriptions[subId].consumers.
            return;
        }
        // Initialize the nonce to 1, indicating the consumer is allocated.
        s_consumers[consumer][subId] = 1;
        s_subscriptions[subId].consumers.push(consumer);

        emit SubscriptionConsumerAdded(subId, consumer);
    }

    /*
     * @notice user unsubscribes
     * @param subId the id  of subscriber
     * @param to the address of user receives the redemption fund
     */
    function cancelSubscription(uint64 subId, address to)
        external
        override
        onlySubOwnerOrAdmin(subId)
        nonReentrant
    {
        if (pendingRequestExists(subId)) {
            revert PendingRequestExists();
        }
        cancelSubscriptionHelper(subId, to);
    }

    function cancelSubscriptionHelper(uint64 subId, address to)
        private
        nonReentrant
    {
        Subscription memory subConfig = s_subscriptions[subId];
        // Note bounded by MAX_CONSUMERS;
        // If no consumers, does nothing.
        for (uint256 i = 0; i < subConfig.consumers.length; i++) {
            delete s_consumers[subConfig.consumers[i]][subId];
        }
        delete s_subscriptions[subId];
        emit SubscriptionCanceled(subId, to, 0);
    }

    function pendingRequestExists(uint64 subId)
        public
        view
        override
        returns (bool)
    {
        Subscription memory subConfig = s_subscriptions[subId];
        for (uint256 i = 0; i < subConfig.consumers.length; i++) {
            for (uint256 j = 0; j < s_provingKeyHashes.length; j++) {
                (uint256 reqId, ) = computeRequestId(
                    s_provingKeyHashes[j],
                    subConfig.consumers[i],
                    subId,
                    s_consumers[subConfig.consumers[i]][subId]
                );
                if (s_requestCommitments[reqId] != 0) {
                    return true;
                }
            }
        }
        return false;
    }
}
