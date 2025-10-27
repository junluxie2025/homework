const { ethers, upgrades } = require("hardhat")
const fs = require("fs")
const path = require("path")

module.exports = async function({ getNamedAccounts, deployments }) {

    // await deployments.fixture(["depolyNftAuction"]);
    const { save } = deployments
    //  获取部署账户
    const { deployer } = await getNamedAccounts()
    console.log("部署用户地址：", deployer)

    // 读取之前保存的代理合约地址和abi
    const nftAuctionProxy = await deployments.get("NftAuctionProxy");
    const proxyAddress = nftAuctionProxy.address;
    const abi = nftAuctionProxy.abi;

    //
    const NftAuctionV2 = await ethers.getContractFactory("NftAuctionV2")

    // 升级代理合约
    const nftAuctionProxyV2 = await upgrades.upgradeProxy(proxyAddress, NftAuctionV2, { call: "admin" })
    await nftAuctionProxyV2.waitForDeployment()

    const proxyAddressV2 = await nftAuctionProxyV2.getAddress()

    console.log("nftAuctionProxyV2 address:", proxyAddressV2);

    const implAddress2 = await upgrades.erc1967.getImplementationAddress(proxyAddressV2);
    console.log("实现合约地址2", implAddress2);

    await save("NftAuctionProxyV2", {
        abi: NftAuctionV2.interface.format("json"),
        address: proxyAddressV2,
    })
}


module.exports.tags = ["upgradeNftAuction"]