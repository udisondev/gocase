package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Accept either a single argument or stdin for Helix :pipe
	if len(os.Args) > 2 {
		fmt.Printf("Usage: %s [word]\n", os.Args[0])
		os.Exit(1)
	}

	var input string
	if len(os.Args) == 2 {
		input = os.Args[1]
	} else {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		input = strings.TrimSpace(string(data))
	}

	if input == "" {
		os.Exit(0)
	}

	switch detectCase(input) {
	case "snake":
		words := strings.Split(input, "_")
		// snake_case → kebab-case
		fmt.Println(strings.ToLower(strings.Join(words, "-")))

	case "kebab":
		words := strings.Split(input, "-")
		// kebab-case → PascalCase
		for i, w := range words {
			words[i] = capitalize(w)
		}
		fmt.Println(strings.Join(words, ""))

	case "pascal":
		words := splitCamel(input)
		// PascalCase → camelCase
		if len(words) > 0 {
			words[0] = strings.ToLower(words[0])
			for i := 1; i < len(words); i++ {
				words[i] = capitalize(words[i])
			}
		}
		fmt.Println(strings.Join(words, ""))

	case "camel":
		words := splitCamel(input)
		// camelCase → snake_case
		for i, w := range words {
			words[i] = strings.ToLower(w)
		}
		fmt.Println(strings.Join(words, "_"))

	default:
		fmt.Fprintln(os.Stderr, "Unsupported format")
		os.Exit(1)
	}
}

// detectCase identifies the naming convention of s
func detectCase(s string) string {
	if strings.Contains(s, "_") {
		return "snake"
	}
	if strings.Contains(s, "-") {
		return "kebab"
	}
	runes := []rune(s)
	if len(runes) > 0 && unicode.IsUpper(runes[0]) {
		return "pascal"
	}
	if len(runes) > 0 && unicode.IsLower(runes[0]) {
		return "camel"
	}
	return ""
}

// capitalize makes the first rune uppercase and the rest lowercase
func capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	return string(unicode.ToUpper(r[0])) + strings.ToLower(string(r[1:]))
}

// splitCamel splits CamelCase or PascalCase into individual words
func splitCamel(s string) []string {
	var words []string
	last := 0
	runes := []rune(s)
	for i := 1; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) {
			words = append(words, string(runes[last:i]))
			last = i
		}
	}
	words = append(words, string(runes[last:]))
	return words
}
