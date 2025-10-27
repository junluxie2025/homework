// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Client} from "@chainlink/contracts-ccip/contracts/libraries/Client.sol";
import {CCIPReceiver} from "@chainlink/contracts-ccip/contracts/applications/CCIPReceiver.sol";

contract CrossChainTokenTransferReceiver is CCIPReceiver {
    using SafeERC20 for IERC20;
    
    // 事件记录
    event TokensReceived(
        bytes32 indexed messageId,
        uint64 indexed sourceChainSelector,
        address sender,
        address token,
        uint256 tokenAmount
    );

    constructor(address router) CCIPReceiver(router) {}
    
    /**
     * 接收来自源链的Token
     */
    function _ccipReceive(
        Client.Any2EVMMessage memory message
    ) internal override {
        // 解析消息
        address sender = abi.decode(message.sender, (address));
        uint64 sourceChainSelector = message.sourceChainSelector;
        
        // 处理接收到的Token
        for (uint256 i = 0; i < message.destTokenAmounts.length; i++) {
            Client.EVMTokenAmount memory tokenAmount = message.destTokenAmounts[i];
            
            // 将Token转给消息发送者指定的接收者
            IERC20(tokenAmount.token).safeTransfer(
                sender,
                tokenAmount.amount
            );
            
            emit TokensReceived(
                message.messageId,
                sourceChainSelector,
                sender,
                tokenAmount.token,
                tokenAmount.amount
            );
        }
    }
}