package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go inputfile outputfile")
		return
	}

	inputfilename := os.Args[1]
	outputfilename := os.Args[2]

	inputfile, err := os.Open(inputfilename)
	if err != nil {
		fmt.Println("Error while opening input file:", err)
		return
	}
	defer inputfile.Close()

	outputfile, err := os.Create(outputfilename)
	if err != nil {
		fmt.Println("Error while creating output file:", err)
		return
	}
	defer outputfile.Close()

	scanner := bufio.NewScanner(inputfile)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		words = caps(words)
		words = capno(words)
		words = upno(words)
		words = lowno(words)

		words = punctuations(words)
		words = aps(words)

		finalstr := strings.Join(words, " ")
		outputfile.WriteString(finalstr + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while scanning input file:", err)
	}
}

func punctuations(words []string) []string {
	for i := 0; i < len(words); i++ {
		if i > 0 && (words[i][0] == '.' || words[i][0] == ',' || words[i][0] == '!' || words[i][0] == '?' || words[i][0] == ':' || words[i][0] == ';') {
			special := true
			for j := 0; j < len(words[i]); j++ {
				if words[i][j] == '.' || words[i][j] == ',' || words[i][j] == '!' || words[i][j] == '?' || words[i][j] == ':' || words[i][j] == ';' {
					words[i-1] += string(words[i][j])
				} else {
					special = false
					words[i] = words[i][j:]
					break
				}
			}
			if special {
				words = append(words[:i], words[i+1:]...)
				i-- // decrement i since we've removed an element
			}
		}
	}
	return words
}

func caps(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		if i > 0 {
			if word == "(cap)" {
				words[i-1] = strings.Title(words[i-1])
				words = append(words[:i], words[i+1:]...)

			} else if word == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
			} else if word == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
			} else if word == "(bin)" {
				prevWord := words[i-1]
				decimal, err := strconv.ParseInt(prevWord, 2, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decimal, 10)
					words = append(words[:i], words[i+1:]...)
				}
			} else if word == "(hex)" {
				prevWord := words[i-1]
				decimal, err := strconv.ParseInt(prevWord, 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decimal, 10)
					words = append(words[:i], words[i+1:]...)
				}
			} else if word == "a" {
				if words[i+1][0] == 'a' || words[i+1][0] == 'e' || words[i+1][0] == 'i' || words[i+1][0] == 'o' || words[i+1][0] == 'u' || words[i+1][0] == 'h' {
					words[i] = "an"
				}
			} else if word == "A" {
				if words[i+1][0] == 'a' || words[i+1][0] == 'e' || words[i+1][0] == 'i' || words[i+1][0] == 'o' || words[i+1][0] == 'u' || words[i+1][0] == 'h' {
					words[i] = "An"
				}
			}
		}
	}
	return words
}

func aps(words []string) []string {
	x := 0
	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "'" && i > 0 && i < len(words)-1 && x == 0 {
			fmt.Println("x", x)
			// plus :=
			temp := words[:i]
			words[i+1] = "'" + words[i+1]
			words[i] = ""
			temp = append(temp, words[i+1:]...)
			words = temp
			x = 1
		} else if word == "'" && i > 0 && x == 1 {
			fmt.Println("x", x)
			words[i-1] = words[i-1] + "'"
			words = append(words[:i], words[i+1:]...)
			x = 0
		}
	}
	return words
}

func capno(words []string) []string {
	for i, word := range words {
		if i > 0 {
			if word == "(cap," {
				newstr := words[i+1][:len(words[i+1])-1]
				new, err := strconv.Atoi(newstr)
				if err != nil {
					fmt.Println("Error while converting")
				}
				for j := 1; j <= new; j++ {
					words[i-j] = strings.Title(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	return words
}

func upno(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		if i > 0 {
			if word == "(up," {
				newstr := words[i+1][:len(words[i+1])-1]
				new, err := strconv.Atoi(newstr)
				if err != nil {
					fmt.Println("Error while converting")
				}
				for j := 1; j <= new; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	return words
}

func lowno(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		if i > 0 {
			if word == "(low," {
				newstr := words[i+1][:len(words[i+1])-1]
				new, err := strconv.Atoi(newstr)
				if err != nil {
					fmt.Println("Error while converting")
				}
				for j := 1; j <= new; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	return words
}
