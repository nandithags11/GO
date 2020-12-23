package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var hash string

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		to_hash := string(i)
		hash = hasher(to_hash)
		fmt.Println(hash)
		if i == 'Z' {
			for i := 'A'; i <= 'Z'; i++ {
				for j := 'A'; j <= 'Z'; j++ {
					s2_i := string(i)
					s2_j := string(j)
					to_hash := s2_i + s2_j
					t_hash := to_hash
					hash = hasher(to_hash)
					fmt.Println(hash)

					if t_hash == "ZZ" {
						for i := 'A'; i <= 'Z'; i++ {
							for j := 'A'; j <= 'Z'; j++ {
								for k := 'A'; k <= 'Z'; k++ {
									s3_i := string(i)
									s3_j := string(j)
									s3_k := string(k)
									to_hash := s3_i + s3_j + s3_k
									hash = hasher(to_hash)
									fmt.Printf("%s    %s \n", to_hash, hash)
								}
							}
						}
					}
				}
			}
		}
	}
}
func hasher(to_hash string) string {
	h := md5.New()
	h.Write([]byte(string(to_hash)))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash

}
