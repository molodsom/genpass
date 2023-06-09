package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

const (
	alpha   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	numbers = "0123456789"
	special = "#%&$~@.,:;?!_=*-+^|/\\\"'`{}[]<>()"
)

func generate(i string, length int) (output string) {
	r := []rune(i)
	for length > 0 {
		output += string(r[rand.Intn(len(r))])
		length--
	}
	return
}

func main() {
	hasAlpha := flag.Bool("a", false, "Include alphabet characters.")
	hasNumbers := flag.Bool("n", false, "Include numbers.")
	hasSpecial := flag.Bool("s", false, "Include special characters.")
	length := flag.Int("l", 20, "Length of each password.")
	results := flag.Int("r", 1, "Number of results (will be ignored if \"-u\" is specified).")
	chars := flag.String("c", "", "Custom character set.")
	users := flag.String("u", "", "Users (comma-separated).")

	flag.Parse()

	if *chars == "" {
		if *hasAlpha {
			*chars += alpha
		}
		if *hasNumbers {
			*chars += numbers
		}
		if *hasSpecial {
			*chars += special
		}
	}

	if *chars == "" {
		*chars = alpha + numbers + special
	}

	usersList := make([]string, *results)
	if *users != "" {
		usersList = strings.Split(*users, ",")
	}

	for i := range usersList {
		user := strings.TrimSpace(usersList[i])
		if user != "" {
			user += "\t"
		}
		fmt.Println(user + generate(*chars, *length))
	}
}
