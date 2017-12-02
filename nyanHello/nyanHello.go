package nyanHello

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func Nyan() {
	fmt.Printf("nyan\n")
}

func TryReadAndCopy() {
	aga, err := ioutil.ReadFile(`Makefile`)
	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReader(string(aga))

	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
}
