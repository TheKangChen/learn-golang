package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "John: 123-456-7890 William: 111-222-3333 Steve: 444-555-6666 Thomas: 777-888-9999"

	updatedText := RemoveHyphensFromPhoneNumbers(text)

	fmt.Println("Text with hyphens removed from phone numbers:")
	fmt.Println(updatedText)
}

func RemoveHyphensFromPhoneNumbers(text string) string {
	re := regexp.MustCompile(`\b(\d{3})-(\d{3})-(\d{4})\b`)
	text = re.ReplaceAllString(text, "$1$2$3")
	return text
}
