// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

//二分查找 (Binary Search)
//题目描述：在一个有序数组中查找目标值。
contract Search {
    function find(int[] memory a,int num) external pure returns (int) {

        int left=0;
        int right = int(a.length-1);
        while(left<=right){
            int mid = (left+right)/2;
            if(a[uint256(mid)]>num){
                right = mid-1;
            }else if(a[uint256(mid)]<num){
                left = mid+1;
            }else{
                return mid;
            }
        }
        return -1;
    }
}
