package main

import (
	"fmt"
)

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		s_i := string(i)
		fmt.Println(s_i)
		if i == 'Z' {
			for i := 'A'; i <= 'Z'; i++ {
				for j := 'A'; j <= 'Z'; j++ {
					s2_i := string(i)
					s2_j := string(j)
					ii := s2_i + s2_j
					fmt.Println(ii)
				}
			}
		}
	}
}
