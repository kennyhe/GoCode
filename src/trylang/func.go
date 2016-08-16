/**
 * Author: Kenny He (he.scu2013@gmail.com)
 * Try some functional programming features of golang
 */

package main

import (
    "fmt"
)


func TryFunc() {
    f := func()(func()) {
        fmt.Println("Call the func 1")
        f1 := func() {
            fmt.Println("Call the func 2")
        }
        return f1
    } ()

    f()
    
    f()
    
    f()
}

func TryAppend() {
    s1 := make([]int, 5, 10)
    for i:=0; i<5; i++ {
        s1[i] = i
    }
    var s2 []int
    for i:=5; i<=12; i++ {
        s1 = append(s1, i)
    }
    fmt.Println(len(s1))
    fmt.Println(len(s2))
}

func main() {
    TryAppend()
}