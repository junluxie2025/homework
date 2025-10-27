// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8;

contract BeggingContract{

    event Donation(address from,uint256 amount);
    //记录每个捐赠者的捐赠金额
    mapping(address=>uint256) private _donors;
    address _self;
    //记录捐赠者
    address[] private _addrs;
    //记录合约所有者所有金额
    uint256 private _totalAmount;
    //时间限制
    uint256 public _startTime;
    uint256 public _endTime;
    //捐赠金额top3
    address[3] private _top3Donors;

    constructor(uint256 startTime,uint256 endTime){
        _self = msg.sender;
        _startTime = startTime;
        _endTime = endTime;
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


    modifier donatePeriod(){
        require(block.timestamp>=_startTime && block.timestamp<=_endTime,"BeggingContract: Donations are only allowed during the specific period");
        _;
    }

    function  _donate() public payable donatePeriod{
        uint256 amount = _donors[msg.sender];
        if(amount<1){//不重复记录
            _addrs.push(msg.sender);
        }
        _donors[msg.sender] += msg.value;
        
        _totalAmount += amount;
        emit Donation (msg.sender, amount);
    }

    modifier onlyOwnwe(){
        require(_self==msg.sender,"BeggingContract: Only owne");
        _;
    }

    function withdraw() external payable onlyOwnwe{
        uint256 len = _addrs.length;
        for(uint i=0;i<len;i++){
            _donors[_addrs[i]]=0;
        }   
        payable(_self).transfer(_totalAmount);
    }

    function getDonation(address addr) external view returns(uint256){
        return _donors[addr];
    }

    function top3Donation() private{
        _top3Donors = [address(0),address(0),address(0)];
        uint256 len = _addrs.length;
        for(uint i=0;i<len;i++){
            address addr = _addrs[i];
            uint256 amount = _donors[addr];

            for(uint j=0;j<3;j++){
                if (amount > _donors[_top3Donors[j]]){
                    
                    for(uint k=j+1;k<3;k++){
                      _top3Donors[k] = _top3Donors[k-1];
                    }
                     _top3Donors[j] = addr;
                    break;
                }
            }
        }   
    }

    function top3Donors() external view returns(address[3] memory){
        return _top3Donors;
    }
}