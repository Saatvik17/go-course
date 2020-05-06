package main

import "fmt"

func main() {

	fmt.Println("voila")
	var S string
	var T string

	fmt.Scan(&S)
	fmt.Scan(&T)

	n := len(S)
	q :=S

	S = []string{S}

	word := lex(S, n )

	if word !=T && q !=T {
		fmt.Println(word)
	}
	else {
		fmt.Println(-1)
	}
}

func lex(s string , n int) {

	for i := n - 1 ; i < -1 ; i-- {
		if s[i] != 'z'
			k := int(s[i])
			s[i] = string(k+1)
			return append(s)
		s[i] = 'a'
	}
	
}
