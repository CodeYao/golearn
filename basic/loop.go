package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	s := ""
	for ; n > 0; n = n / 2 {
		s = strconv.Itoa(n%2) + s
	}
	return s
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1101
		convertToBin(7238775),
		convertToBin(0),
	)

	printFile("abc.txt")
}
