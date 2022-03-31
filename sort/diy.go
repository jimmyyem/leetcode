package sort

import (
	"fmt"
	"sort"
)

type Student struct {
	ID    int
	Name  string
	Score int
}

// 排序
func DiySort(students []Student) {
	fmt.Println("原始数据")
	for _, v := range students {
		fmt.Printf("%d => %d\n", v.ID, v.Score)
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].Score > students[j].Score
	})

	fmt.Println("排序后数据")
	for _, v := range students {
		fmt.Printf("%d => %d\n", v.ID, v.Score)
	}
}
