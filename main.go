package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func gravatar(mail string, size int) string {
	hash := md5.Sum([]byte(strings.TrimSpace(strings.ToLower(mail))))
	return fmt.Sprintf("http://www.gravatar.com/avatar/%s?s=%d&d=identicon", hex.EncodeToString(hash[:]), size)
}

func main() {
	email := os.Args[1]
	size := os.Args[2]
	s, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := gravatar(email, s)
	fmt.Println(url)
}
