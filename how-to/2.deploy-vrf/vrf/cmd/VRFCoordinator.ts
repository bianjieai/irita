import "@nomiclabs/hardhat-web3";
import { task } from "hardhat/config";
import { BigNumber } from "ethers";

task("deployVRFCoordinator", "Deploy VRFCoordinator")
    .addParam("minrequestconfirmations", "minimum request confirmations")
    .addParam("maxgaslimit", "max gas limit")
    .setAction(async (args, hre) => {
        const minRequestConfirmations = BigNumber.from(args.minrequestconfirmations)
        const maxGasLimit = BigNumber.from(args.maxgaslimit)
        const vrfCoordinatorFactory = await hre.ethers.getContractFactory('VRFCoordinator')
        const vrfCoordinator = await hre.upgrades.deployProxy(vrfCoordinatorFactory,[
            minRequestConfirmations,
            maxGasLimit,
        ]);
        await vrfCoordinator.deployed();
        console.log("vrfCoordinator deployed to:", vrfCoordinator.address);
    });