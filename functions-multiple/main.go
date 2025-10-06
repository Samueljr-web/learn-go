package main

import (
	"fmt"
	"strings"
)


func getInitials(n string) (string, string) {
   s := strings.ToUpper(n) //capitalize strings
   names := strings.Split(s, " ") //split into a slice

   var initials []string
   
   for _,v := range names{
    initials = append(initials, v[:1])
   }

   if len(initials) > 1 {
    return  initials[0], initials[1]
   }

   return initials[0], "_"
}

func main() {
   fn1, sn1 := getInitials("Samuel Adeyemi")
   fmt.Println(fn1, sn1)
   fn2, sn2 := getInitials("Frontend Eng")
   fmt.Println(fn2, sn2)


}