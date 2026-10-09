package main

import (
    "strings"
    "unicode"
	"fmt"
)

// function to return Full Name initials in Caps
func getInitials(fullName string) string {
    parts := strings.Fields(fullName)
	fmt.Println("Parts:", parts) // Debugging line to print the parts of the name

    if len(parts) == 0 {
        return ""
    }

    initials := ""
    for _, part := range parts {
        if len(part) > 0 {
            initials += string(unicode.ToUpper(rune(part[0])))
        }
    }
    return initials
}
