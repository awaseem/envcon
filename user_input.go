package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
)

type userInput struct{}

// String prompt.
func (u *userInput) string(prompt string, args ...interface{}) string {
	fmt.Printf(prompt+": ", args...)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic("reading standard input:" + err.Error())
	}
	return ""
}

// String prompt (required).
func (u *userInput) stringRequired(prompt string, args ...interface{}) (s string) {
	for strings.Trim(s, " ") == "" {
		s = u.string(prompt, args...)
	}
	return s
}

// Confirm continues prompting until the input is boolean-ish.
func (u *userInput) confirm(prompt string, args ...interface{}) bool {
	for {
		switch u.string(prompt, args...) {
		case "Yes", "yes", "y", "Y":
			return true
		case "No", "no", "n", "N":
			return false
		}
	}
}

// Choose prompts for a single selection from `list`, returning in the index.
func (u *userInput) choose(prompt string, list []string) int {
	fmt.Println()
	for i, val := range list {
		fmt.Printf("  %d) %s\n", i+1, val)
	}

	fmt.Println()
	i := -1

	for {
		s := u.string(prompt)

		// index
		n, err := strconv.Atoi(s)
		if err == nil {
			if n > 0 && n <= len(list) {
				i = n - 1
				break
			} else {
				continue
			}
		}

		// value
		i = u.indexOf(s, list)
		if i != -1 {
			break
		}
	}

	return i
}

// Password prompt.
func (u *userInput) password(prompt string, args ...interface{}) string {
	fmt.Printf(prompt+": ", args...)
	password, _ := gopass.GetPasswd()
	s := string(password[0:])
	return s
}

// Password prompt with mask.
func (u *userInput) passwordMasked(prompt string, args ...interface{}) string {
	fmt.Printf(prompt+": ", args...)
	password, _ := gopass.GetPasswdMasked()
	s := string(password[0:])
	return s
}

// index of `s` in `list`.
func (u *userInput) indexOf(s string, list []string) int {
	for i, val := range list {
		if val == s {
			return i
		}
	}
	return -1
}
