// Stringers
import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old)", p.Name, p.Age)
}

func main() {
	a := Person{"Master programming", 23}
	z := Person{"Dev programming", 28}
	fmt.Println(a, z) // Master programming (23 years old) Dev programming (28 years old)
}

// Example IP generator

type IPAddr [4]byte

func (ip IPAddr) String() string {
	str := ""
	for i, ipValue := range ip {
		srt += fmt.Sprint(ipValue)
		if i < len(ip) -1 {
			srt += "."
		}
	}
	return fmt.Sprintf("%v", str)
}


func main() {
	hosts := map[string]IPAddr {
		"loopback": {127,0,0,1},
		"googleDNS": {8,8,8,8}
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip) // loopback: 127.0.0.1 googleDNS: 8.8.8.8
	}
}