package main

import "fmt"

const Pig = 0
const Cow = 1
const Chicken = 2

const (
    C1 = iota
    C2 = iota
    C3 = iota
)

const (
    Blue = 0xff << (iota * 8)
    Green
    Red
)

func saySomething(animal int){
    switch animal {
    case Pig:
        fmt.Println("꿀꿀")
    case Cow:
        fmt.Println("음메")
    case Chicken:
        fmt.Println("꼬끼오")
    default:
        fmt.Println("ERROR")
    }
}

func main() {
    const myConstant = 123
    fmt.Println(myConstant)
    saySomething(Pig)
    saySomething(Cow)
    saySomething(Chicken)
    fmt.Println(C1)
    fmt.Println(C2)
    fmt.Println(C3)
    fmt.Println(Red)
    fmt.Println(Green)
    fmt.Println(Blue)
}
