package main

import "fmt"

func main() {
	programmers := map[string]string{"Katie": "Android"}
	programmers["Magda"] = "Facebook"
	fmt.Println(programmers)
	for key, value := range programmers {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}
