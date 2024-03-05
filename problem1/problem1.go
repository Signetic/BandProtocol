package main

import (
	"fmt"
	"strings"
)

func bossBaby(shots string) string {
	strBad := "Bad boy"
	strGood := "Good boy"
	splitString := strings.Split(shots, "")

	if splitString[0] == "R" {
		return strBad
	} else if splitString[len(splitString)-1] == "S" {
		return strBad
	}
	lenS := 0
	lenR := 0
	for x := 0; x < len(splitString); x++ {
		if splitString[x] == "R" {
			lenR++
		} else if splitString[x] == "S" {
			lenS++
		}
	}
	if lenR < lenS {
		return strBad
	}
	return strGood
}
func main() {
	str1 := bossBaby("SRSSRRR")
	fmt.Println(str1)
	str2 := bossBaby("RSSRR")
	fmt.Println(str2)
	str3 := bossBaby("SSSRRRRS")
	fmt.Println(str3)
	str4 := bossBaby("SR")
	fmt.Println(str4)
}
