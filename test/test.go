package main

import (
	"fmt"
	"openvpn-manage/lib"
)

func main() {
	s, _ := lib.ReadLine("C:/easy-rsa/pki/ca.crt")
	fmt.Printf(s)
}
