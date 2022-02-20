package banner

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
	converting the banners to a map
*/
func AsciiMap(banner string) map[int][]string {
	file, err := ioutil.ReadFile(banner)
	if err != nil {
		fmt.Println("Invalid file")
	}
	splitFile := strings.Split(string(file), "\n")
	asciiArtMap := make(map[int][]string)
	art := []string{} // slices for every art
	posCount := 0     // position in the map
	ascii := 32       // ascii value counter
	for ascii <= 126 {
		num := (9 * posCount) // equation to get the position in the map
		for i := num; i < num+9; i++ {
			art = append(art, splitFile[i])
		}
		asciiArtMap[ascii] = art
		art = []string{}
		posCount++
		ascii++
	}
	return asciiArtMap
}

/*
	transform the string to art by calling the key of the inputStr in the map
*/
func AsciiToArt(inputStr, inputBanner string) string {
	art := ""
	p := AsciiMap(inputBanner)
	for i := 1; i <= 8; i++ {
		for _, value := range inputStr {
			art += p[int((value))][i]
		}
		art += "\n"
	}
	return art
}

/*
	this function catch the \n. It will split the input string with newline and print the word line by line.
*/
func PrintAsciiArt(arg1, ban string) string {
	art := ""
	switch arg1 {
	case "":
		return ""
	case "\\n":
		return "\n"
	default:
		argSplit := strings.Split(arg1, "\\n")
		for _, word := range argSplit {
			if word != "" {
				art += AsciiToArt(word, ban)
			} else {
				art += "\n"
			}
		}
	}
	return art
}
