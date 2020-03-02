/*
Method(s) of String Library Implemented and supporting Go-Lang Language which support String Method(s)
*/
package main
import "strings"
func Contain_Char(String string,Char rune) bool {
result := strings.IndexRune(String,Char)
if(result != -1){
	return true
}
return false
}
