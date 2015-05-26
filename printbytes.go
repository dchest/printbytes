// Command printbytes reads bytes from standard input and prints them in Go
// source format.
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)
	lines := 1
	for i, v := range in {
		if i > 0 && i%10 == 0 {
			fmt.Fprintf(buf, "\n\t")
			lines++
		}
		fmt.Fprintf(buf, "0x%02x", v)
		if i != len(in)-1 {
			fmt.Fprintf(buf, ", ")
		}
	}
	fmt.Printf("[]byte{")
	if lines > 1 {
		fmt.Printf("\n\t%s,\n", buf.String())
	} else {
		fmt.Printf("%s", buf.String())
	}
	fmt.Println("}")
}
