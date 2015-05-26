// Command printrandombytes prints random bytes for use in Go programs.
package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
)

var fCount = flag.Int("c", 16, "number of bytes")

func main() {
	flag.Parse()
	b := make([]byte, *fCount)
	_, err := rand.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)
	lines := 1
	for i, v := range b {
		if i > 0 && i%10 == 0 {
			fmt.Fprintf(buf, "\n\t")
			lines++
		}
		fmt.Fprintf(buf, "0x%02x", v)
		if i != len(b)-1 {
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
