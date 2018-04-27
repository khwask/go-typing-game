package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	r := os.Stdin
	fmt.Println("######################################")
	fmt.Println("###  Golang Typing Game (ver.beta) ###")
	fmt.Println("### * Type more words in 10 second ###")
	fmt.Println("### * To start, press any button.. ###")
	fmt.Println("######################################")

	// スタートボタンの入力をまつ
	s := bufio.NewScanner(r)
	s.Scan()
	s.Text()

	// 入力待ち関数
	ch := input(r)

	// 問題をだして、正誤判定する関数
	ret := question(ch)

	// 正解数をカウントする関数
	count := counter(ret)

	// 10秒待つ
	<-time.After(10 * time.Second)
	fmt.Println("### TIME UP!! Your score is", *count, ":) ###")
}
