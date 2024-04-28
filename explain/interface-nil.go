// Interface values with nil

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("< nil >")
		return
	}
	fmt.Println(t.S)
}

func main(){
	var i I
	var t *T
	
	i = t
	
	describe(i)
	i.M() // < nil >

	i = &T{"Hello"}

	describe(i) // Hello
	i.M()
}

// Empty interface

var i interface{}


// Assertion type

func main2() {
	var i interface {} = "Hello World"

	// first try
	s := i.(string)
	fmt.Println(s) // Hello World

	// second try
	s, ok := i.(string)
	fmt.Println(s, ok) // Hello World true

	// third try
	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	// fourth try
	f = i.(float64) // Panic
	fmt.Println(f) // Error stopping by panic


	// panic
	t, error := i.(int)

	if error != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(t)
	}

	// Switch types
	switch v := i.(type) {
		case int: fmt.Println("String")
		case string: fmt.Println("String")
		default: fmt.Println("Unknown")
	}

}


type IHello interface {
	Hello()
}

type Person struct {
	Name string 
}

func (p Person) Hello() {
	fmt.Println("Hello, I am, "p.Name)
}

type Number int 

func (n Number) Hello() {
	if n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "not is even")
	}
}

func main3() {
	greeters := []IHello{
		Person{"Master Programming"},
		Person{"Dev Programming"},
		Number(3),
		Number(2)
	}
	for _, greeter := range greeters {
		switch value := greeter.(type) {
			case Person: fmt.Println(value.Name)
			case Number: fmt.Println(value)
			default: fmt.Println("I don´ known")
		}
		greeter.Hello()
	}
}