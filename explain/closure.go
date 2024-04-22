// Closure
var x = 10
var y = 20

func sum(x,y int) int {
	x := 1
	y := 2
	return x + y
}

func split() int {
	return x / y
}

// Methods
type Person struct {
	Name, LastName, string
}
// func is from Person
func (p Person) Hello() string {
	return "Hello "+ p.Name
}
// Go not have Class
func main() {
	p := Person{"Master", "Last Master"}
	fmt.Println(p.Hello())
}


func ChangeName(p *Person) {
	p.Name = "Master Pop"
}

func (p *Person) ChangeNameOwn() {
	p.Name = "Master Pop"
}

func main2() {
	p := Person{"Master Pep", "Last Master"}
	punter := &Person{"Master pep", "Last Master"}
	ChangeName(punter)
	p.ChangeNameOwn()
	fmt.Println(p.Hello())
}

// Interface

type IPerson interface {
	Hello() string
	Move() string
}

type Student struct {
	Name string
}

func (a Student) Hello() string {
	return "Hello "+ a.Name
}

func (a Student) Move() string {
	return "Moving"
}

func main3() {
	var person IPerson = Student{"Master Student"}

	fmt.Println(person.Hello())
	fmt.Println(person.Move())
}