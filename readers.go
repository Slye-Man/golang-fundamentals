package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, inputErr := read.ReadString(' ')
		nex := strings.TrimSpace(input)
		if nex == "" {
			continue
		}
		num , convErr := strconv.Atoi(nex)
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}
	fmt.Printf("Sum: %v\n", sum)
}
