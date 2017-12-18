package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const (
	stringSet string = "0123456789ABCDEFGHJKLMNPQRSTWXYZabcdefghijkmnopqrstuvwxyz"
	symSet    string = "!@#$%^&*()+=?{}~_"
)

type genStr struct {
	id  int
	str bytes.Buffer
}

type genWork struct {
	genStr genStr
}

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

	set := stringSet
	if *s {
		set += symSet
	}

	workers := make(chan genWork, *r)
	var wg sync.WaitGroup

	for i := 0; i < *r; i++ {
		wg.Add(1)
		go func() {
			for work := range workers {
				var gen = work.genStr
				genString(&gen.str, set, len)
				fmt.Println(gen.str.String())
			}
			wg.Done()
		}()
	}

	for i := 1; i <= *r; i++ {
		workers <- genWork{
			genStr{i, bytes.Buffer{}},
		}
	}

	close(workers)
	wg.Wait()
}

// genString generates a random alphanumric string
func genString(str *bytes.Buffer, set string, l int) {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < l; i++ {
		str.WriteString(string(set[rand.Intn(len(set)-1)]))
	}
}
