// Exercise from educative.io "The Way To Go" course
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IdentifyPrefixPostfix(userID, email string) bool {
	return strings.HasPrefix(email, userID) || strings.HasSuffix(email, userID)
}

func ContainsEducative(email string) bool {
	return strings.Contains(strings.ToLower(email), "educative")
}

func MaskUserName(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) < 2 {
		return email
	}
	userName := parts[0]
	nameLen := len(userName)
	parts[0] = string(userName[0]) + strings.Repeat("*", nameLen-2) + string(userName[nameLen-1])
	return strings.Join(parts, "@")
}

func IndexOfAtSymbol(email string) int {
	return strings.Index(email, "@")
}

func TrimAndSplitUserID(userID string) string {
	slicedUID := strings.Split(strings.TrimSpace(userID), "-")
	if len(slicedUID) < 2 {
		panic("Error: Incorrect user ID format")
	}
	return slicedUID[1]
}

func ConvertStringToInt(str string) int {
	intStr, err := strconv.Atoi(str)
	if err != nil {
		panic("Error: Failed to convert string to integer")
	}
	return intStr
}

func main() {
	fmt.Println(IdentifyPrefixPostfix(".io", "evangeline@educative.io")) // true
	fmt.Println(IdentifyPrefixPostfix("UID", "UID-0123"))                // true
	fmt.Println(IdentifyPrefixPostfix("UID", "evangeline@educative.io")) // false
	fmt.Println(ContainsEducative("evangeline@educative.io"))            // true
	fmt.Println(MaskUserName("evangeline@educative.io"))                 // e******e@educative.io
	fmt.Println(IndexOfAtSymbol("evangeline@educative.io"))              // 10
	fmt.Println(TrimAndSplitUserID("UID-0123"))                          // 0123
	fmt.Println(ConvertStringToInt("123"))                               // 123
}
