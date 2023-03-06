# How to deploy VRF for Cosmos enterprise use cases through IRITA blockchain instance

## 1. **What is VRF?**

VRF (Verifiable Random Function) is a cryptographic function that can be used to generate random numbers like RNG (Random Number Generator). Unlike regular RNGs, VRF's main feature includes fairness, tamper-proof, verifiability, and unpredictability.

VRF has now been adopted in various business scenarios. For example, as an industrial leading blockchain oracle service provider, Chainlink has integrated VRF as one of its main features. The following are some common use cases of VRF:

* Gaming

Build better games by leveraging random outcomes in blockchain gaming applications. map generation, critical hits (battling games), matchmaking (multiplayer games), card draw order, and random encounters/events are now possible with VRF.

* NFTs

Distribute rare non-fungible tokens (NFTs) and assign randomized attributes to NFTs, providing players access to auditable evidence that their NFT assets are created and awarded using tamper-proof randomness.

* Ordering for processes

Distribute highly-coveted items like ticketed events, select participants in a popular public sale, and choose presale winners for luxury items like rare shoes.

* Random entity selection

Generate rich experiences with challenging and unpredictable scenarios and environments, and find the perfect mix between strategy and fun by using randomness to determine outcomes in PvP battles and other scenarios.

[By courtesy of Chainlink](https://chain.link/education-hub/rng-in-blockchain-use-cases)


## 2. **Enterprise Use Cases of VRF**

In the enterprise blockchain scenarios, crypto-currency is not involved due to regulatory requirements, and VRF service can be provided through a trusted centralized service provider. Hence certain modifications are needed for compliance consideration as well as efficiency and performance reasons. In this how-to instruction, we explain the modifications we made to Chainlink's original VRF design to address the enterprise requirements as well, we also share the deployment guide using IRITA ( a Cosmos SDK based enterprise blockchain framework).  

This work was inspired by the business requirements of Wenchang Chain during the 2022 World Cup to support activities such as bracket challenges, lucky draws, etc. This allows for verifiable, tamper-proof, and low-cost random numbers and encryption certificates to support DApps on WenChang Chain.

>*Wenchang Chain is an Open-Permissioned Blockchain developed by Bianjie.AI using the technical framework IRITA - the Cosmos enterprise-level consortium chain product.*
Since its release, the VRF service on Wenchang Chain has supported various applications such as metaverse gaming, customer loyalty program, and a variety of other businesses.

## 3. **Deployment guide for enterprise VRF use cases**

With reference to Chainlink's best practice, the enterprise VRF service within the trusted consortium chain environment can simplify the fee module in Chainlink, and customize cost-control strategies demanded in various business scenarios.

In short, it helps with:

* cost-control with no crypto involved
* eliminating the need of building a decentralized network that could cost tons of processing power
* Deployment Guide:
    * [vrf contract](./vrf)
    * [vrf provider](./vrf-provider)

