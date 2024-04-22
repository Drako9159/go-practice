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
