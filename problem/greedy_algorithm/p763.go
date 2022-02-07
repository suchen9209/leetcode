package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	//fmt.Println(partitionLabels("a"))
}

func partitionLabels(s string) []int {
	EndPosMap := make(map[int32]int)
	for i, i2 := range s {
		EndPosMap[i2] = i
	}
	fmt.Println(EndPosMap)
	var LenArr []int
	lastCut := -1
	tmpPos := 0
	for i, i2 := range s {
		if tmpPos < EndPosMap[i2] {
			tmpPos = EndPosMap[i2]
		}
		if i == tmpPos {
			LenArr = append(LenArr, tmpPos-lastCut)
			lastCut = tmpPos
			tmpPos = 0
		}
	}
	fmt.Println(LenArr)
	return []int{0}
}
