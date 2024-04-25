// Readers

import (
	"fmt"
	"strings"
	"io"
	"os"
)

// b stream channel // data stream
func (T) Read(b []byte) (n int, err error) {

}

func main() {
	data := "Master programming"
	reader := strings.NewReader(data)
	var buffer strings.Builder
	n, err := io.Copy(%buffer, reader)
	if err !=  nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Bytes read:", n) // Output: Bytes read: 18
		fmt.Println("String read:", buffer.String()) // Output: Master programming
	} 
}

// Example

// Decoder and coder
type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(p []byte) (n int, err error){
	n, err = rr.r.Read(p)

	for i := 0; i < len(p); i++ {
		if p[i] >= "A" && p[i] <= "Z" || p[i] >= "a" && p[i] <= "z" {
			if p[i] <= "Z" {
				p[i] = (p[i]-"A"+13)%26 + "A"
			} else {
				p[i] = (p[i]-"a"+13)%26 + "a"
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrp gur phqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // You cracked the code!
}

