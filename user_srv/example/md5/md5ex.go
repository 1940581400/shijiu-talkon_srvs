package main

import (
	"fmt"
	"talkon_srvs/user_srv/global/pwd"
)

func main() {

	s, e, _ := pwd.NewEncodedPwd("534534534534534fghgfhgfhsdsa")

	fmt.Println(s)
	fmt.Println(e)
	str := pwd.GetDefaultPwdSep("300WybXfgSkWkWyx",
		"87fbdd4d1a86149b56814f1dba9c345f0d5adf66773471b7efeb9b1400c23cfc")
	fmt.Println(len(str))
	ok, _ := pwd.VerifyPwdSep(pwd.GetDefaultPwdSep("300WybXfgSkWkWyx",
		"87fbdd4d1a86149b56814f1dba9c345f0d5adf66773471b7efeb9b1400c23cfc"),
		"534534534534534fghgfhgfhsdsa")
	fmt.Println(ok)
}
