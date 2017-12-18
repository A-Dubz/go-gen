package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	stringSet string = "0123456789ABCDEFGHJKLMNPQRSTWXYZabcdefghijkmnopqrstuvwxyz"
	symSet    string = "!@#$%^&*()+=?{}~_"
)

func main() {
	var s = flag.Bool("s", false, "Add symbols to generated string")
	var r = flag.Int("r", 1, "Number of generated strings to output. Default: 1")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Please specify length of generated string (1-100)")
		return
	}

	len, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Println("Length must be a integer")
		return
	}

	if len <= 0 || len > 100 {
		fmt.Println("Generate length value out of range")
		return
	}

	var buf bytes.Buffer
	set := stringSet

	if *s {
		set += symSet
	}

	for i := 0; i < *r; i++ {
		genString(&buf, set, len)
		fmt.Println(buf.String())
		buf.Reset()
	}
}

// genString generates a random alphanumric string from subset
func genString(b *bytes.Buffer, s string, l int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < l; i++ {
		b.WriteString(string(s[rand.Intn(len(s)-1)]))
	}
}
