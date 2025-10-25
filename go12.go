package main

import (
    "fmt"
    "strconv"
)

func main() {
    num1 := "1.65"

    // string -> float64
    flonum1, err := strconv.ParseFloat(num1, 64)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("float64:", flonum1)

    // string -> int (sadece tam sayı)
    num3 := "11"
    intnum3, err := strconv.Atoi(num3)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("int:", intnum3)

    // float64 -> string
    num2 := 1.64
    strnum1 := strconv.FormatFloat(num2, 'f', 2, 64) // 2 ondalık basamak
    fmt.Println("string:", strnum1)

    // Alternatif fmt.Sprintf ile
    strnum2 := fmt.Sprintf("%.2f", num2)
    fmt.Println("string (fmt.Sprintf):", strnum2)
}

