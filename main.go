package main

import (
	"fmt"

)

func main() {

sample_array := [30]string{"dog","cat"}
fmt.Printf("sample_array: %v\n", sample_array)

sample_slice := append(sample_array[:])
fmt.Printf("sample_slice: %v\n", sample_slice)

var sample_map = map [string] int { "hey":78, "hello":56,}
fmt.Printf("codes: %v\n", sample_map)



}
