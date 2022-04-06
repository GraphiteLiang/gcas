package main

import (
	"crypto/md5"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"strconv"
	"time"
)
func main() {
	crutime := time.Now().Unix()
	fmt.Println("crutime-->", crutime)

	h := md5.New()
	fmt.Println("h-->", h)

	fmt.Println("strconv.FormatInt(crutime, 10)-->", strconv.FormatInt(crutime, 10))
	io.WriteString(h, strconv.FormatInt(crutime, 10))

	fmt.Println("h-->", h)

	token := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("token--->", token)

	fmt.Println(len("8e1a188743c6077110da3c9778183031"))
}
