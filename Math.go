/*
Method(s) of Math Library Implemented and supporting Go-Lang Language
*/
package main
import ("fmt"
		"math"	
	)

func main(){
fmt.Println("Welcome to the library of Go this is a personal project where i implement methods for Go library feel free to add ",
"any idea of yours or to use any method :=) ")

}
/* Convert a 2D matrix to a transpose
	Complexity : O(n**2)
	Idea : transpose Algorithm
*/func doubel_Matrix_Transpoe (matrix[2][2] int) [2][2]int{
var i,j,temp int;
for i=0; i < 2;i++ {
for j=j+1;j < 2;j++ {
temp = matrix[j][i]
matrix[j][i] = matrix[i][j]
matrix[i][j] = temp
}
}
return matrix
}

/* Count the number of primenumbers between 1 and n 
	Complexity : O(Log(n)*N)
	Idea : Sieve of EratosthenesCal
*/
func Calculate_prime(n int)int {
	var count int
	var i,j int
	var arr[1000000] bool
	var number float64 
	number = float64(n)
	number = math.Sqrt(number)
	number1 := int(number)
	if(n > 1000000){
	fmt.Println("The Number has to be less than 1000000")
	return -1;	
	}
for i = 2; i < number1 ; i++ {
for j = i * i; j < n; j = i + j {
arr[j] = true
}
}
for i = 0; i < n ; i++ {
if(!arr[i]){
count++
}
}
return count
}
/* Sorting Arrays of numbers
 complexity of O(n+m)
 Idea : Counting Sort 
*/
func sort(Array[1000000]int)[] int{
	var indices[1000000]int
	max := Array[0]
	var i,j int
for i=0; i <  1000000; i++ {
	if(max < Array[i]){
	max = Array[i]	
	}
	}
new_Array := Array[0:max]
for i =0 ; i < max; i++ {
if(new_Array[i] != 0){
indices[new_Array[i]]++
}
}
for i=0; i <= max; i++ {
	for x := indices[i]; x > 0; x-- {
	new_Array[j] = i;
	indices[i]--
	j++
}
}
return new_Array
}