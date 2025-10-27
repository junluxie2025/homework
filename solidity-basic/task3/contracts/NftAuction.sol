// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

import "hardhat/console.sol";

contract NftAuction is Initializable ,UUPSUpgradeable{
    struct Auction {
        address seller; //卖家
        uint256 duration;
        uint256 startPrice; //起始价格
        uint256 startTime;
        bool ended;
        address highestBidder; //最高出价者
        uint256 highestBid; //最高出价
        address nftContract; //NFT合约地址
        uint256 tokenId;
        address tokenAddress; //参与竞价的资产类型0x地址表示eth 其他地址表示erc20
    }

    //状态变量
    mapping(uint256 => Auction) public auctions;
    //下一个竞拍Id
    uint256 public nextAuctionId;
    //管理员地址
    address public admin;

    mapping(address => AggregatorV3Interface) public priceFeeds;

    function initialize() public initializer {
        admin = msg.sender;
    }

    function setPriceFeed(address tokenAddress, address _priceFeed) public {
        priceFeeds[tokenAddress] = AggregatorV3Interface(_priceFeed);
    }

    function getChainlinkDataFeedLatestPrice(
        address tokenAddress
    ) public view returns (int) {
        AggregatorV3Interface dataFeed = priceFeeds[tokenAddress];
        // prettier-ignore
        (
            /* uint80 roundId */,
            int256 answer,
            /*uint256 startedAt*/,
            /*uint256 updatedAt*/,
            /*uint80 answeredInRound*/
        ) = dataFeed.latestRoundData();
        return answer;
    }

    //创建拍卖
    function createAuction(
        uint256 _duration,
        uint256 _startPrice,
        address _nftAdress,
        uint256 _tokenId
    ) public {
        require(msg.sender==admin,"Only admin can create auctions");
        require(_duration>5,"Duration must be greater than 10s");
        require(_startPrice>0,"Start price must be greater than 0");

        IERC721(_nftAdress).transferFrom(
            msg.sender,
            address(this),
            _tokenId
        );

        auctions[nextAuctionId] = Auction({
            seller: msg.sender,
            duration: _duration,
            startPrice: _startPrice,
            startTime: block.timestamp,
            ended: false,
            highestBidder: address(0),
            highestBid: 0,
            nftContract: _nftAdress,
            tokenId: _tokenId,
            tokenAddress: address(0)
        });
        nextAuctionId++;
    }

    //买家参与买单
    function placeBid(uint256 _auctionID,uint256 amount,address _tokenAddress) public payable{

        Auction storage auction = auctions[_auctionID];
        require(block.timestamp < auction.startTime + auction.duration, "Auction already ended.");
        
        uint payValue;
        if(_tokenAddress!=address(0)){
            payValue = amount*uint(getChainlinkDataFeedLatestPrice(_tokenAddress));
        }else{
            payValue = amount*uint(getChainlinkDataFeedLatestPrice(address(0)));
        }

        uint startPriceValue = auction.startPrice * uint(getChainlinkDataFeedLatestPrice(auction.tokenAddress));
        uint highestBidValue = auction.highestBid * uint(getChainlinkDataFeedLatestPrice(auction.tokenAddress));
        require(payValue >= startPriceValue && payValue > highestBidValue, "There already is a higher or equal bid.");

        //转移ERC20到合约
        if(_tokenAddress!=address(0)){
            IERC20(_tokenAddress).transferFrom(msg.sender,address(this),amount);
        }

        if(auction.highestBid>0){
            if (auction.tokenAddress!=address(0)){
                IERC20(auction.tokenAddress).transfer(auction.highestBidder,auction.highestBid);
            }else{
                payable(auction.highestBidder).transfer(auction.highestBid);
            }
        }
        
        auction.highestBidder = msg.sender;
        auction.highestBid = amount;
        auction.tokenAddress = _tokenAddress;
    }

    //结束拍卖
    function endAuction(uint256 _auctionID) public{
        Auction storage auction = auctions[_auctionID];
        require(block.timestamp >= auction.startTime + auction.duration, "Auction not yet ended.");

        console.log("endAuction",auction.startTime,auction.duration,block.timestamp);
        
        require(!auction.ended&&(auction.startTime+ auction.duration)<=block.timestamp, "endAuction has already been called.");

        //转移NFT给最高出价者
        IERC721(auction.nftContract).safeTransferFrom(address(this),auction.highestBidder,auction.tokenId);

        auction.ended = true;
    }


    function _authorizeUpgrade(address) internal view override{
        // 只有管理员可以升级合约
        require(msg.sender == admin, "Only admin can upgrade");
    }

    //接收NFT回调函数
     function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external pure returns (bytes4) {
        return this.onERC721Received.selector;
    }
}
