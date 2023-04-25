package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func GetMainMenu() {
	fmt.Println("                                      ")
	fmt.Println("1. Add Distributor with Permission")
	fmt.Println("2. List all Distributors")
	fmt.Println("3. Check Permission for a Distributor")
	fmt.Println("4. Create a Network of Distributors")
	fmt.Println("5. Back to the Main Menu")
	fmt.Println("6. Exit")
	fmt.Println("                                      ")
}
func RemoveSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}

	return strings.ToLower(string(rr))
}