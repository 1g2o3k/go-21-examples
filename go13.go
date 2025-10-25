package main 
import (
"fmt"
"strconv"
)
func main(){
i:=5
StrI:=strconv.Itoa(i)
fmt.Println(StrI)
m:="5"
IntM,err:=strconv.Atoi(m)
if err!=nil{
fmt.Println("error")
}
fmt.Println(IntM)
}

