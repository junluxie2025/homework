// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Vote{

    address[] addrrs;

    //一个mapping来存储候选人的得票数
    mapping(address=>uint) public voters;

    //一个vote函数，允许用户投票给某个候选人
    function vote(address addr) external {
        voters[addr]++;
        // addrrs.push(addr);
    }
    //一个getVotes函数，返回某个候选人的得票数
    function getVotes(address addr) external view returns(uint){
        return voters[addr];
    }
    //一个resetVotes函数，重置所有候选人的得票数
    function resetVotes() external {
        for(uint i=0;i<addrrs.length;i++){
            delete voters[addrrs[i]];
        }
        delete addrrs;
    }

    //反转字符串 (Reverse String)
    function reverse(string memory str0) external pure returns(string memory){
        bytes memory chars = bytes(str0);
        uint  len = chars.length;
        bytes memory reserChars = new bytes(len);
        for(uint i=0;i<len;i++){
            reserChars[len-i-1]=chars[i];
        }
        return string(reserChars);
    }

}