import "@openzeppelin/hardhat-upgrades"
import "@nomicfoundation/hardhat-toolbox";
import "./cmd/VRFCoordinator"

module.exports = {
  defaultNetwork: 'hardhat',
  networks: {
    hardhat: {
      allowUnlimitedContractSize: true,
    },
    wctest: {
      url: 'http://testnet.bianjie.ai:8545',
      gasPrice: 1,
      chainId: 12231,
      gas: 4000000,
      accounts: ['3f2ca07c1f351caed872317dba6693ef917393121331fefdfa56012e1cbb1e5c'],
    },
  },
  solidity: {
    version: '0.8.4',
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    }
  },
  contractSizer: {
    alphaSort: true,
    runOnCompile: true,
    disambiguatePaths: false,
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  }
}
