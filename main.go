//https://leetcode.com/problems/length-of-last-word/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"fmt"
	"strings"
)



func lengthOfLastWord(s string) int {

   s = strings.TrimSpace(s)
   words := strings.Split(s," ")
   

   if len (words) == 0 {
	return 0
   }
   lastWord := words[len(words) - 1]
   
    return len(lastWord)
} 

func main()  {
	fmt.Println(lengthOfLastWord("Hello World")) // Output: 5
    fmt.Println(lengthOfLastWord("    "))          // Output: 0

}