package main

import (
	"fmt"
)

type MyStruct struct {
    Id     string
    Count  int
    Name   string
}

type Student struct {
	Id string
	Name string
	Class Class // 구조체를 포함하는 구조체
}

type Class struct {
	Title string
	Point int
}

type User struct {
	Name string
}

type AdminUser struct { // 다른 구조체의 필드를 포함하는 구조체, 마치 다른 OOP언어의 상속처럼 보일 수 있음
	User
	Role string
}

type ManagerUser struct {
	User
	Name string // 필드 이름이 겹치는 경우
}

func main() {
    var myStruct1 MyStruct // 구조체를 선언하면 초기화된다.
    fmt.Println(myStruct1) // 모든 필드가 기본값으로 초기화된 구조체(zero value 구조체) 출력

    // 구조체의 각 필드 사용
    myStruct1.Id = "aaa123"
    myStruct1.Count += 1
    myStruct1.Name = "MymyStruct"
    fmt.Println(myStruct1)

    var myStruct2 *MyStruct = nil // 만약, zero value 구조체가 아닌 nil을 할당하려면 null pointer를 할당해야함
    // var myStruct2 *MyStruct // nil을 명시적으로 할당하지 않아도 nil로 초기화됨
    fmt.Println(myStruct2)

	myStruct3 := MyStruct{"bbb222", 100, "anotherStruct"}
	fmt.Println(myStruct3)

	myStruct4 := MyStruct{ // 각 필드이름으로 일부만 원하는 값으로 초기화할 수 있음
		Count: 255,
		Name:  "someStruct",
	}
	fmt.Println(myStruct4)

	var std1 Student // 다른 구조체를 포함하는 구조체
	fmt.Println(std1)

	fmt.Println(std1.Class.Title) // 구조체 안의 구조체에 접근하는 방식

	myUser := AdminUser{User{"관리자이름"}, "관리자"}
	fmt.Println(myUser)

	fmt.Println(myUser.Name) // 구조체의 필드를 포함하는 방식이라면 '.'을 한번만 표현해도 접근 가능
	fmt.Println(myUser.User.Name) // 물론 이런 방식도 가능

	myManager := ManagerUser{User{"매니저1"}, "매니저11"}
	fmt.Println(myManager) // 필드 이름이 겹치는 경우

	fmt.Println(myManager.Name) // 이 경우 myManager가 선언된 타입 안에 있는 'Name'필드에 우선 접근
}
