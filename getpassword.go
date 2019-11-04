package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type KeyType struct {
	BasicKey []string
	HardKey  []string
}

func main() {
	var key KeyType
	key.BasicKey = append(key.BasicKey, "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNPQRSTUVWXYZ")
	key.HardKey = append(key.HardKey, "abcdefghijklmnopqrstuvwxyz0123456789.?/#*-;,!+%&{(=)}:><$ABCDEFGHIJKLMNPQRSTUVWXYZ")
	choice, length := Input()
	pass := Pass(choice, length, key)

	fmt.Println("Password ->", pass)
}

func Input() (string, int) {
	var length int
	var choice string
	flag.StringVar(&choice, "difficulty", "basic", "password difficulty")
	flag.IntVar(&length, "length", 7, "password length")
	flag.Parse()

	if length > 25 {
		log.Println("Max length -> 25")
		length = 25
	}

	return choice, length
}

func Pass(choice string, length int, key KeyType) string {
	var pass string
	r := rand.New(rand.NewSource(time.Now().Unix()))
	if choice == "basic" {
		for i := 0; i < length; i++ {
			randIndex := r.Intn(len(key.BasicKey[0]))
			pass += string(key.BasicKey[0][randIndex])
		}
	} else if choice == "hard" {
		for i := 0; i < length; i++ {
			randIndex := r.Intn(len(key.HardKey[0]))
			pass += string(key.HardKey[0][randIndex])
		}
	} else {
		fmt.Println("wrong difficulty level")
		main()
	}

	return pass
}
