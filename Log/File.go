package Log

import (
	"bufio"
	"os"
)

func readFile(fileName string) ([]string, []string) {
	return revertLine(fileName)
}
func revertLine(fileName string) ([]string, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var origin []string
	var revert []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		origin = append(origin, scanner.Text())
	}
	for i := len(origin) - 1; i >= 0; i-- {
		revert = append(revert, origin[i])
	}
	return origin, revert
}
