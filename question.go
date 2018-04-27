package main

import (
	"fmt"
	"math/rand"
	"time"
)

var q = [...]string{"golang", "java", "c++", "javascript", "lisp", "kotlin", "swift", "python", "c", "ruby"}

func question(ch <-chan string) <-chan bool {
	ret := make(chan bool)
	go func() {
		for {
			t := time.Now().UnixNano()
			rand.Seed(t)
			s := rand.Intn(len(q))
			// 問題を出す
			fmt.Println("[", q[s], "]")
			// 入力を受け取り判定
			if q[s] == <-ch {
				ret <- true
			}
		}
	}()
	return ret
}
