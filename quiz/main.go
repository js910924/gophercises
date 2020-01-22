package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("problem.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var count int = 0
	for _, line := range lines {
		ans, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		if eval(line[0]) != ans {
			fmt.Println("Wrong:", eval(line[0]), ans)
			count++
		}
	}

	fmt.Println(count)
}

func eval(str string) int {
	operationalIndex := findOperationalSymbol(str)
	num1, err := strconv.Atoi(str[:operationalIndex])
	if err != nil {
		log.Fatal(err)
	}

	num2, err := strconv.Atoi(str[operationalIndex+1:])
	if err != nil {
		log.Fatal(err)
	}

	var result int
	switch string(str[operationalIndex]) {
	case "+":
		result = num1 + num2

	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	return result
}

func findOperationalSymbol(str string) int {
	var index int
	for i, s := range str {
		if _, err := strconv.Atoi(string(s)); err != nil {
			index = i
			break
		}
	}

	return index
}
