package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i <= t; i++ {
		scanner.Scan()
		s := scanner.Text()
		var odd []byte
		var even []byte

		for j := 0; j < len(s); j++ {
			if j == 0 {
				even = append(even, s[j])
			} else {
				if j%2 != 0 { //if odd
					odd = append(odd, s[j])
				} else { //if even
					even = append(even, s[j])
				}
			}
		}

		fmt.Println(string(even) + " " + string(odd))
	}

}
