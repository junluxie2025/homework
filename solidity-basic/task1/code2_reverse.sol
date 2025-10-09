// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract code2_reverse {

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
