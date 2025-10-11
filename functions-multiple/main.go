package main

import (
	"fmt"
	"strings"
)


func getInitials(n string) (string, string) {
   s := strings.ToUpper(n)
   name := strings.Split(s, " ")

   var initials []string
   for _,v := range name{
      initials = append(initials, v[:1])
   }

   if(len(initials) > 1){
      return initials[0], initials[1]
   }

   return initials[0], "_"
}

func main() {
   n1,n2 := getInitials("Samuel Jr")
   fmt.Println(n1, n2)
}


