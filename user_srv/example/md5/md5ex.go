package main

import (
	"fmt"
	"talkon_srvs/user_srv/utils"
)

func main() {

	s, e, _ := utils.NewEncodedPwd("534534534534534fghgfhgfhsdsa")

	fmt.Println(s)
	fmt.Println(e)
	str := utils.GetDefaultPwdSep("300WybXfgSkWkWyx",
		"87fbdd4d1a86149b56814f1dba9c345f0d5adf66773471b7efeb9b1400c23cfc")
	fmt.Println(len(str))
	ok, _ := utils.VerifyPwdSep(utils.GetDefaultPwdSep("300WybXfgSkWkWyx",
		"87fbdd4d1a86149b56814f1dba9c345f0d5adf66773471b7efeb9b1400c23cfc"),
		"534534534534534fghgfhgfhsdsa")
	fmt.Println(ok)
}
