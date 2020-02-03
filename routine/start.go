package main

import (
	"fmt"
	"net/http"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%d ", i)
		resp, err := http.Get("http://2children.ru/news/index?NewsSearch%5Bword%5D=%D0%B2%D1%8B%D0%B0%D1%8B%D0%B2%D0%B0")
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c ", i)
		resp, err := http.Get("http://2children.ru/news/index?NewsSearch%5Bword%5D=%D0%B2%D1%8B%D0%B0%D1%8B%D0%B2%D0%B0")
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
	}
}
func main() {
	for index := 0; index < 1000; index++ {
		go numbers()
		go alphabets()
	}
	time.Sleep(30000 * time.Millisecond)
	fmt.Println("main terminated")
}
