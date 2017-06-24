package bubblesort

import (
	"fmt"
)

func BubbleSort(value []int) {

	for i := 0; i < len(value)-1; i++ {
		for j := 0; j < len(value)-1; j++{
			if value[j] > value[j+1]{
				value[j],value[j+1] = value[j+1],value[j]
			} 
		}
	}
}
1