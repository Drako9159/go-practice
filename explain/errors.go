// Errors

import "fmt"

func main() {
	// Atoi is equals ParseInt
	i, err := strconv.Atoi("45")

	if err != nil {
		fmt.Printf(err) 
		return
	} else {
		fmt.Printf(i)
	}
}

