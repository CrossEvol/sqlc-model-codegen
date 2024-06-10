package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
)

func main() {
	s1 := strcase.ToLowerCamel("EmailVerified")
	fmt.Println(s1)
	s2 := strcase.ToLowerCamel("Emailverified")
	fmt.Println(s2)
}
