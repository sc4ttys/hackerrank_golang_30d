package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var names []string
	var numbers []int
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		names = append(names, line[0])
		num, _ := strconv.Atoi(line[1])
		numbers = append(numbers, num)
	}
	for {
		prev := scanner.Text()
		scanner.Scan()
		text := scanner.Text()
		if text == prev || len(text) == 0 {
			break
		}
		exists := false
		for i := range names {
			if text == names[i] {
				fmt.Println(names[i] + "=" + strconv.Itoa(numbers[i]))
				exists = true
			}
		}
		if !exists {
			fmt.Println("Not found")
		}
	}

}
