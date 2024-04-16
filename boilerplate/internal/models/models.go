package 'models'

// public
func MyModel() {

}

// private
func myModel() {

}

import 'fmt'

var number int = 2


func Main() {
	i, j := 42, 2701
	fmt.Println(i ,j) - 42 2701
}

type Person struct {
	Name string
	LastName string
	Age int
}

var person = Person {
	Name: "Master",
	LastName: "Master Last",
	Age: 26
}

fmt.Println(person.Name)

var array = [2]int{1,2}

var values = [6]int{2,3,5,7,11,13}
var s []int = values[1:4]
fmt.Println(s) // [3 5 7]
s = append(s, 17)
fmt.Println(s) // [3 5 7 17]
fmt.Println(values) // [2 3 5 7 17 11 13]

// its same
var a[10] int
a[0:10]
a[:10]
a[0:]
a[:]

// make
a := make([]int, 0, 5)

// punters
type ElementType struct {
	Name string
}

var exampleElement = ElementType {
	Name: "Master"
}

func MyFunc(element *ElementType){
	element.Name = "NewMaster"
}

fmt.Println(element.Name) // Master

MyFunc(&exampleElement) // Apply

fmt.Println(element.Name) // NewMaster

var a = 1
var p *int
i := 42
p = &i

// No primitives types
v := Person("Master")
p := &v
fmt.Println(p.Name)

// Primitives types
fmt.Println(*p) // 42
*p = 21
fmt.Println(*p) // 21


