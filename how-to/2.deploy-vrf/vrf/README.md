# VRF Contract

EVM Based VRF Contract

## 1. Deoloy `VRFCoordinator` contract

```bash
yarn hardhat deployVRFCoordinator --minrequestconfirmations <minrequestconfirmations> --maxgaslimit <maxgaslimit> --network <network>
```

## 2. Apply for a subscriber id

Users need to call the mehtod `createSubscription` of `VRFCoordinator` contract to register an administrator to manage their consumer contracts.

## 3. Deploy the consumer contract

After the registration is completed, the `VRFCoordinator` contract will return a `subId` to the user, user deploys consumer contract with `subId`.
**_Note, before deploying the contract, please confirm the `keyHash` with the manager of the `VRFCoordinator` contract._**

## 4. Add the consumer contract

The administrator calls the `addConsumer` method of the `VRFCoordinator` contract to add the consumer contract deployed in the previous step.
