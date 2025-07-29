package main

import (
    "fmt"
    "reflect"
    )

func main(){
    
    fmt.Println(reflect.TypeOf("A normal string which can have 123, #$%^&, true and false, 34.5"))
    
    fmt.Println(reflect.TypeOf('K'))
    
    fmt.Println(reflect.TypeOf(false))
    fmt.Println(reflect.TypeOf(34))
    fmt.Println(reflect.TypeOf(116.5))
}
