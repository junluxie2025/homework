// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

//题目描述：将两个有序数组合并为一个有序数组。
contract MergeArray {
    function merge(
        uint[] memory a,
        uint[] memory b
    ) external pure returns (uint[] memory) {
        uint aLen = a.length;
        uint bLen = b.length;
        uint len = aLen + bLen;
        uint[] memory res = new uint[](len);

        uint aIndex = 0;
        uint bIndex = 0;
        for (uint i = 0; i < len; i++) {
            if (aIndex < aLen && bIndex < bLen) {
                if (a[aIndex] < b[bIndex]) {
                    res[i] = a[aIndex];
                    aIndex++;
                }else if (a[aIndex] > b[bIndex]) {
                    res[i] = b[bIndex];
                    bIndex++;
                }else{
                    res[i] = a[aIndex];
                    aIndex++;
                }
            } else if (aIndex >= aLen) {
                res[i] = b[bIndex];
                bIndex++;
            } else {
                res[i] = a[aIndex];
                aIndex++;
            }

        }
        return res;
    }
}
