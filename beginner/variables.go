package main

import (
    "fmt"
    "reflect"
    )

func main(){
    
    // variable declaration and value assign
    var rollNumber int
    rollNumber = 27
    
    // same thing at a single shot
    var fname, sname string
    fname,sname="KP","S"
    
    // no need to mention type of data explicitly
    var height, weight = 5.7, 80
    
    fmt.Println(rollNumber)
    fmt.Print(fname)
    fmt.Println(sname)
    fmt.Println(reflect.TypeOf(weight))
    fmt.Println(reflect.TypeOf(height))
}
