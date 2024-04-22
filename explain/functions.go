// Functions

func Hello(name string) string{
	fmt.Println("Hello "+ name)
}

func main2() {
	message := Hello("Lucy")
	fmt.Println(message)
}

func GetCallback(callback func(float64, float64) float64) float64 {
	return callback(2,3)
}

func hypotenuse(x,y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	fmt.Println(hypotenuse(3,4))
	fmt.Println(GetCallback(hypotenuse))
}

