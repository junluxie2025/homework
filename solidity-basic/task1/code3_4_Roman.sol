// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Roman{
    mapping(bytes1=>int) public romanMap;

    constructor(){
        romanMap['I']=1;
        romanMap['V']=5;
        romanMap['X']=10;
        romanMap['L']=50;
        romanMap['C']=100;
        romanMap['D']=500;
        romanMap['M']=1000;
    }

    //实现罗马数字转整数
    function intToRoman(uint num) external pure returns(string memory){
        uint[13] memory values = [uint(1000), uint(900), uint(500), uint(400),uint(100), uint(90), uint(50), uint(40), uint(10), uint(9), uint(5), uint(4), uint(1)];
        string[13] memory symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        string memory roman = "";
        for (uint i = 0; i < values.length && num > 0; i++) {
            while (num >= values[i]) {
                num -= values[i];
                roman = string(abi.encodePacked(roman, symbols[i]));
            }
        }
        return roman;
    }

    //实现整数转罗马数字
    function romanToInt(string memory s) external view returns(int256){
        bytes memory charArray = bytes(s);
        int256 rs=0;
        for(uint i=0;i<charArray.length;i++){
            bool existGreater = false;
            for(uint j=i+1;j<charArray.length;j++){
                if (romanMap[charArray[i]]<romanMap[charArray[j]]) {
                    existGreater=true;
                    break;
                }
            }
            if(existGreater){
                rs -=romanMap[charArray[i]];
            }else{
                rs +=romanMap[charArray[i]];
            }
        }
        return rs;
    }
}