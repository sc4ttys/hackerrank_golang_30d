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
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		name := line[0]
		num, _ := strconv.Atoi(line[1])
		m[name] = num
	}
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		_, exists := m[text]
		if !exists {
			fmt.Println("Not found")
		} else {
			fmt.Println(text + "=" + strconv.Itoa(m[text]))
		}

	}

}
