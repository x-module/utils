/**
 * Created by goland.
 * @file   slice.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/1 19:40
 * @desc   slice.go
 */

package algorithm

import (
	"github.com/go-xmodule/utils/utils/lancetconstraints"
)

type AscComparator struct{}

func (c *AscComparator) Compare(v1 any, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	// ascending order
	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

type DescComparator struct{}

func (c *DescComparator) Compare(v1 any, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	// ascending order
	if val1 < val2 {
		return 1
	} else if val1 > val2 {
		return -1
	}
	return 0
}

// AscSort 切片正序排序
func AscSort[T any](params []T) {
	QuickSort(params, new(AscComparator))
}

// DescSort 切片倒叙排序
func DescSort[T any](params []T) {
	QuickSort(params, new(DescComparator))
}

// Sort 自定义排序
func Sort[T any](params []T, comparator lancetconstraints.Comparator) {
	QuickSort(params, comparator)
}
