package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("####")
	fmt.Println("Contains:", strings.Contains("INDIA", "IND"))
	fmt.Println("Contains:", strings.Contains("INDIA", "ind"))
	fmt.Println("Contains:", strings.Contains("INDIA", "DIA"))
	fmt.Println("####")
	fmt.Println("Count:", strings.Count("apple pinapple", "apple"))
	fmt.Println("Count:", strings.Count("apple pinapple", "APPLE"))
	fmt.Println("Count:", strings.Count("apple pinapple", "p"))
	fmt.Println("####")
	fmt.Println("ToLower:", strings.ToLower("CapItaL"))
	fmt.Println("####")
	fmt.Println("ToUpper:", strings.ToUpper("CapItaL"))

	fmt.Println("####")
	fmt.Println("Index:", strings.Index("Index finder", "find"))
	fmt.Println("Index:", strings.Index("Index finder", "found"))
	fmt.Println("####")

}
