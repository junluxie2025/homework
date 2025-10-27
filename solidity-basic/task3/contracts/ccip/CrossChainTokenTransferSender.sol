pragma solidity ^0.8.7;

import {IRouterClient} from "@chainlink/contracts-ccip/contracts/interfaces/IRouterClient.sol";
import {OwnerIsCreator} from "@chainlink/contracts/src/v0.8/shared/access/OwnerIsCreator.sol";
import {Client} from "@chainlink/contracts-ccip/contracts/libraries/Client.sol";
import {LinkTokenInterface} from "@chainlink/contracts/src/v0.8/shared/interfaces/LinkTokenInterface.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


contract CrossChainTokenTransferSender is OwnerIsCreator {
    using SafeERC20 for IERC20;
    
    // 事件记录
    event TokensSent(
        bytes32 indexed messageId,
        uint64 indexed destinationChainSelector,
        address receiver,
        address token,
        uint256 tokenAmount,
        address feeToken,
        uint256 fees
    );

    IRouterClient private  s_router;
    LinkTokenInterface private s_linkToken;

    constructor(address _router,address _link){
        s_router = IRouterClient(_router);
        s_linkToken = LinkTokenInterface(_link);
    }
    
    /**
     * 发送ERC20 Token到目标链
     * @param destinationChainSelector 目标链的ChainSelector
     * @param receiver 目标链上的接收合约地址
     * @param token ERC20 Token地址
     * @param amount 发送数量
     */
    function sendToken(
        uint64 destinationChainSelector,
        address receiver,
        address token,
        uint256 amount
    ) external returns (bytes32 messageId) {
        // 转移Token到本合约
        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        
        // 授权给路由器使用Token
        IERC20(token).approve(address(s_router), amount);
        
        // 构造CCIP消息
        Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
        tokenAmounts[0] = Client.EVMTokenAmount({
            token: token,
            amount: amount
        });
        
        Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
            receiver: abi.encode(receiver),
            data: "",
            tokenAmounts: tokenAmounts,
            extraArgs: "",
            feeToken: address(0) // 使用原生代币支付费用
        });
        
        // 获取费用估算
        uint256 fee = s_router.getFee(destinationChainSelector, message);
        
        // 发送跨链消息
        messageId = s_router.ccipSend(destinationChainSelector, message);
        
        emit TokensSent(
            messageId,
            destinationChainSelector,
            receiver,
            token,
            amount,
            address(0),
            fee
        );
    }
}