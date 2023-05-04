package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

type Input [][]int

func main() {
	var input Input
	file, err := ioutil.ReadFile("input/larg-input.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &input)
	if err != nil {
		log.Fatal(err)
	}

	// Solution with array
	sumArr := maxSumArr(input)
	fmt.Println("Result with array solution: ", sumArr)

	// Solution with tree
	// ** Low performance please use small-input.json before test this solution **//
	// root := buildTree(input, 0, 0)
	// sumTree := maxSumTree(root)
	// fmt.Println("Result with tree solution: ", sumTree)

}

func maxSumArr(input Input) int {
	size := len(input[len(input)-1])
	tempArr := []int{}
	for i := 0; i <= size; i++ {
		tempArr = append(tempArr, 0)
	}
	for i := len(input) - 1; i >= 0; i-- {
		row := input[i]
		for j := 0; j < len(row); j++ {
			left := float64(row[j] + tempArr[j])
			right := float64(row[j] + tempArr[j+1])
			tempArr[j] = int(math.Max(left, right))
		}
	}
	return tempArr[0]
}
