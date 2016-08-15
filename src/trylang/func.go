/**
 * Author: Kenny He (he.scu2013@gmail.com)
 * Try some functional programming features of golang
 */

package main

import (
    "fmt"
)


func main() {
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

