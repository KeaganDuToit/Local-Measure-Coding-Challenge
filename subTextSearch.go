package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter string to search: ")
	readerText := bufio.NewReader(os.Stdin)
	searchAble, _ := readerText.ReadString('\n')

	fmt.Print("Enter string to search for: ")
	readerGiven := bufio.NewReader(os.Stdin)
	searchText, _ := readerGiven.ReadString('\n')

	// Remove the new line character from the input
	searchText = strings.TrimSuffix(searchText, "\n")
	searchAble = strings.TrimSuffix(searchAble, "\n")

	fmt.Println(simpleTextSearch(searchAble, searchText))
}

// simpleTextSearch is a simple text search algorithm that just iterates over the searchAble string
// and compares the first letter of the searchText with the current letter of the searchAble string,
// if the first letter of the searchText matches the current letter of the searchAble string we continue to compare
// the rest of the searchText with the rest of the searchAble string
func simpleTextSearch(searchAble string, searchText string) string {
	var fistLetterIndex []int = make([]int, 0)
	beginningChar := string(searchText[0])

	for pos := range searchAble {
		currentChar := string(searchAble[pos])

		if strings.EqualFold(currentChar, beginningChar) {
			for pos2 := range searchText {
				searchAbleNextChar := string(searchAble[pos+pos2])
				searchTextNextChar := string(searchText[pos2])
				if strings.EqualFold(searchAbleNextChar, searchTextNextChar) {
					if pos2 == len(searchText)-1 {
						fistLetterIndex = append(fistLetterIndex, pos+1)
					}
				} else {
					// This means the current letter of the searchAble string does not match the current letter of the searchText string,
					// so we break out of the loop and continue to the next letter of the searchAble string
					break
				}
			}
		} else {
			continue
		}
	}

	returnString := fmt.Sprintf("%v", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(fistLetterIndex)), ", "), "[]"))
	return returnString
}
