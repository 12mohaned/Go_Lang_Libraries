package main
import "fmt"
func main(){
fmt.Println()
}
func doubel_Matrix_Transpoe (matrix[2][2] int) [2][2]int{
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