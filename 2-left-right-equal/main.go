package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print("Input: ")

		input, err := consoleReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		trimInput := strings.TrimSuffix(input, "\n")
		if strings.HasPrefix(trimInput, "exit") {
			fmt.Println("Good bye!")
			os.Exit(0)
		}

		_err := validate(trimInput)
		if _err != nil {
			fmt.Println(_err.Error())
			continue
		}

		result := solution(trimInput)
		fmt.Println("Result: ", result)
	}
}

func solution(str string) string {
	arr := []int{}
	last := 0
	for i, s := range str {
		if i == 0 {
			if string(s) == "L" {
				arr = append(arr, 2)
				arr = append(arr, 1)
				last = 1
			} else if string(s) == "R" {
				arr = append(arr, 1)
				arr = append(arr, 2)
				last = 2
			} else {
				arr = append(arr, 0)
				arr = append(arr, 0)
				last = 0
			}
		} else {
			if string(s) == "L" {
				arr = append(arr, last-1)
				last = last - 1
			} else if string(s) == "R" {
				arr = append(arr, last+1)
				last = last + 1
			} else {
				arr = append(arr, last)
			}
		}
	}
	min := min(arr)
	if min > 0 {
		min = 0
	} else {
		min = min * (-1)
	}

	pre := ""
	for arr[0] == 0 {
		arr = arr[1:]
		pre += "0"

	}

	chrArr := []string{}
	for i := range arr {
		number := arr[i] + min
		text := strconv.Itoa(number)
		chrArr = append(chrArr, text)
	}
	return pre + strings.Join(chrArr, "")
}

func min(arr []int) int {
	min := 0
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func validate(str string) error {
	for _, s := range str {
		if string(s) != "L" && string(s) != "R" && string(s) != "=" {
			return errors.New("Invalid input character.")
		}
	}
	return nil
}
