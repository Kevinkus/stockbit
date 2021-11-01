package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getInput(reader *bufio.Reader) []string {
	fmt.Print("$ ")
	text, _, _ := reader.ReadLine()

	return strings.Split(string(text), " ")
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// input by space
	input := getInput(reader)

	MapResult := make(map[string][]string)

	for _, val := range input {
		// sort string
		str := sortString(val)
		MapResult[str] = append(MapResult[str], val)
	}

	for _, val := range MapResult {
		fmt.Println(val)
	}

}
