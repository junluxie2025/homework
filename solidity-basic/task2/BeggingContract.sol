// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8;

contract BeggingContract{

    //记录每个捐赠者的捐赠金额
    mapping(address=>uint256) private _donors;
    address _self;
    //记录捐赠者
    address[] private _addrs;
    //记录合约所有者所有金额
    uint256 private _totalAmount;

    constructor(){
        _self = msg.sender;
    }

    receive() external payable{
        _donate();
    }

    fallback() external payable{
        require(
            msg.value>0,
            "BeggingContract: The donation amount must be greater than 0"
        );
        _donors[msg.sender] += msg.value;
        _addrs.push(msg.sender);
        _totalAmount +=  msg.value;

    }

    function  _donate() public payable{
         require(
            msg.value > 0,
            "BeggingContract: The donation amount must be greater than 0"
        );
        _donors[msg.sender] += msg.value;
        _addrs.push(msg.sender);
        _totalAmount +=  msg.value;
    }

    modifier onlyOwner(){
        require(_self==msg.sender,"only owner can withdraw");
        _;
    }

    function withdraw() external payable onlyOwner{
        uint256 len = _addrs.length;
        for(uint i=0;i<len;i++){
            _donors[_addrs[i]]=0;
        }  
        require(_totalAmount > 0, "BeggingContract: No amount can be withdraw"); 
        payable(_self).transfer(_totalAmount);
    }

    function getDonation(address addr) external view returns(uint256){
        return _donors[addr];
    }
}