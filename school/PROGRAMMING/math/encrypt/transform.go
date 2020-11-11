package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	str "strings"
)

const chars = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwqxyz0123456789 !?#@$%^&*_+/-=[]{}().,<>`

func fromMessage(msg string) (s string) {
	for _, c := range msg {
		s += strconv.Itoa(charToInt(c))
	}
	return
}

func toMessage(s string) (msg string) {
	for i := 0; i < len(s); i += 2 {
		is := string(s[i]) + string(s[i+1])
		n, err := strconv.Atoi(is)
		if err != nil {
			panic(fmt.Sprintf("err parse %s: %v", is, err))
		}
		msg += string(intToChar(n))
	}
	return
}

func intToChar(i int) rune {
	return rune(chars[i-10])
}

func charToInt(c rune) int {
	for i, e := range chars {
		if e == c {
			return i + 10
		}
	}
	panic(fmt.Sprintf("char %c not supported", c))
}

//
//
//
//
//
//
//
//

func transformMain() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		fmt.Println(handle(scan.Text()))
	}
}

func handle(in string) (out string) {
	str.TrimSpace(in)
	if str.HasPrefix(in, ":") {
		in = str.TrimSpace(str.TrimPrefix(in, ":"))
		return fromMessage(in)
	}
	return toMessage(in)
}
