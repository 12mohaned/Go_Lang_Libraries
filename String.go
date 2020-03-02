package main
import "strings"
func Contain_Char(String string,Char rune) bool {
result := strings.IndexRune(String,Char)
if(result != -1){
	return true
}
return false
}
