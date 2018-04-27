package main

import (
	"bufio"
	"io"
)

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		sc := bufio.NewScanner(r)
		for sc.Scan() {
			ch <- sc.Text()
		}
		close(ch)
	}()
	return ch
}
