package main

import "fmt"

func main() {
	Array_ori := [11]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	Slice_a := Array_ori[2:5]
	Slice_b := Array_ori[3:5]
	Slice_c := []string{"x", "y", "z"}
	Slice_c = append(Slice_c, string(`喂`)) // runes vs bytes - UTF8 encoding
	fmt.Println(Array_ori, Slice_a, Slice_b, Slice_c)
}
