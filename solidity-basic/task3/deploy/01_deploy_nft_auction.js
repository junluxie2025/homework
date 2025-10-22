const { deployments, upgrades, ethers } = require("hardhat");

const fs = require("fs");
const path = require("path");

module.exports = async({ deployments, getNamedAccounts }) => {
    const { save } = deployments;
    const { deployer } = await getNamedAccounts();

    console.log("部署用户地址1:", deployer);

    // 获取合约工厂
    const NftAuction = await ethers.getContractFactory("NftAuction");

    // 部署可升级代理合约
    const nftAuctionProxy = await upgrades.deployProxy(NftAuction, {
        initializer: "initialize"
    });

    // 等待部署完成
    await nftAuctionProxy.waitForDeployment();

    const proxyAddress = await nftAuctionProxy.getAddress();
    console.log("nftAuctionProxy address:", proxyAddress);

    // 获取实现合约地址
    const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
    console.log("实现合约地址", implAddress);

    // 保存代理合约地址
    await save("NftAuctionProxy", {
        abi: NftAuction.interface.format("json"),
        address: proxyAddress,
        // args: [],
        // log: true,
    })

};

module.exports.tags = ["depolyNftAuction"];