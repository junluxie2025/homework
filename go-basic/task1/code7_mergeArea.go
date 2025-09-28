package main

import "sort"

/*
合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 先按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片，放入第一个区间
	merged := [][]int{intervals[0]}

	for _, current := range intervals[1:] {
		// 获取结果中最后一个区间
		last := merged[len(merged)-1]
		// 如果当前区间的起始位置小于等于结果中最后区间的结束位置，说明有重叠
		if current[0] <= last[1] {
			// 合并区间：更新结果中最后区间的结束位置为两者较大值
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 没有重叠直接添加进新切片中
			merged = append(merged, current)
		}
	}

	return merged
}
