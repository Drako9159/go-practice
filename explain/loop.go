// Loop Range

var arr = []int{5,4,3,2,1}

for i, v := range arr {
	fmt.Println("index: %d, value: %d", i, v)
}

// Range is for slices, arrays, maps and strings

// Maps

type Person struct {
	DNI, Name, string
}

var m = map[string]Person{
	"12345": {"12345", "Master"},
	"12346": {"12346", "Dev"}
}

func main() {
	fmt.Println(m["12345"])
	m["12345"] = Person["12345", "New Master"]
}



