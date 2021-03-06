package utils

import (
	"bufio"
	"go/build"
	"os"
	"path/filepath"
	"strconv"
)

func GetArrayOfSize(n int) []int {
	p, err := build.Default.Import("go-algorithm/sort/utils", "", build.FindOnly)
	if err != nil {
		//
	}
	fName := filepath.Join(p.Dir, "IntegerArray.txt")
	f, _ := os.Open(fName)
	defer f.Close()

	numbers := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, s)
	}
	return numbers[:n]
}