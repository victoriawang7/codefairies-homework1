package main

import "fmt"

var animalList = []string{"fly", "spider", "bird", "cat", "dog", "cow"}
var contentList = []string{"", "That wriggled and wiggled and tickled inside her.", "How absurd to swallow a bird.", "Fancy that to swallow a cat!", "What a hog, to swallow a dog!", "I don't know how she swallowed a cow!"}

func main() {

	getPrintSong()

}

func getPrintSong() {
	for i := 0; i < len(animalList); i++ {
		if i == 0 {
			printContent2(animalList[i], i)
			printContent1()
			continue
		} else {
			printContent2(animalList[i], i)
			fmt.Printf(contentList[i])
			if i >= 1 {
				for j := i; j >= 1; j-- {
					printContent3(animalList[j], animalList[j-1], j)
				}
			}
			printContent1()
		}
	}
	printEnd()
}

func printContent1() {
	fmt.Printf("I don't know why she swallowed a fly - perhaps she'll die!")
}
func printContent2(str string, index int) {
	if index == 0 {
		fmt.Printf("There was an old lady who swallowed a " + str + ".")

	} else {
		fmt.Printf("There was an old lady who swallowed a " + str + ";")
	}
}
func printContent3(str string, str1 string, index int) {
	if index == 1 {
		fmt.Printf("She swallowed the " + str + " to catch the " + str1 + ";")
	} else {
		fmt.Printf("She swallowed the " + str + " to catch the " + str1 + ",")

	}
}

func printEnd() {
	fmt.Printf("There was an old lady who swallowed a horse...")
	fmt.Printf("...She's dead, of course!")
	fmt.Printf("")
}
