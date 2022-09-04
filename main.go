package main

import (
	"fmt"

)

func main() {

//fixed size not a reference type
sample_array := [30]string{"dog","cat"}
fmt.Printf("sample_array: %v\n", sample_array)

//refreences underlying array , size can change
sample_slice := append(sample_array[:])
fmt.Printf("sample_slice: %v\n", sample_slice)

// k.v pair
var sample_map = map [string] int { "hey":78, "hello":56,}
fmt.Printf("codes: %v\n", sample_map)

//like a clss , accepts multiple types
type Doctor struct{
   actorName string
   age int
}

sampleStruct := Doctor{actorName:"guy",age:57}
fmt.Printf("sampleStruct: %v\n", sampleStruct)



type Animal struct{
	name string
	speices string
}

//embedding animal struct into Mammal
type Mammal struct{
	Animal
	breed string
	legs int
}

var lua = Mammal{
legs:4,	
breed:"fresian",
// name:"cow",
// speices:"meat",
}
//assigning Anmal Struct values to Mammal using embedded fields
lua.name="cow"
lua.speices="meat"

fmt.Printf("lua: %v\n", lua)

//Tagging a field to add extar information
type Goon struct{
	name string `required max:"100"`
	speices string
}






}




