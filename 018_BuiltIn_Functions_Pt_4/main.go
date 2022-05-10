package main

import (
	"log"
	"strconv"
)

func divide(a, b float32) float32 {
	if b == 0 {
		log.Fatal("Cannot divde by zero!\n")
	}
	return a / b
}

func main() {
	n, _ := strconv.ParseInt("87", 10, 64)
	log.Println("n =", n)
	st := "95f"
	m, err := strconv.ParseInt(st, 10, 64)
	if err != nil {
		log.Printf("Could not convert %v to number\n", st)
	}
	log.Println("m =", m)
	// string to int
	i, err := strconv.Atoi("-16")
	log.Println(i)
	// int to string
	s := strconv.Itoa(-17)
	log.Println(s)
	t, _ := strconv.ParseBool("1")
	log.Println("t =", t)
	log.Printf("7/5 = %v\n", divide(7, 5))
	log.Printf("7/0 = %v\n", divide(7, 0))
}
