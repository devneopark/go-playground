package main

import "fmt"

type Data struct {
    value string
    data [10]uint8
}

func changeData(data Data)  {
    data.value = "hi"
    data.data[0] = 255
}

func changeDataWithPointer(data *Data)  {
    data.value = "hi"
    data.data[0] = 255
}

func main() {
    var i *int
    fmt.Printf("변수 i의 값: %v\n", i)

    var value = 100
    fmt.Printf("변수 value의 값: %v\n", value)
    fmt.Printf("변수 value가 가리키는 주소: %v\n", &value)

    i = &value
    fmt.Printf("변수 i의 값(가리키는 주소 == value 가 할당된 주소): %v\n", i)
    fmt.Printf("변수 i가 할당된 주소: %v\n", &i)
    fmt.Printf("변수 i가 가리키는 주소에 할당된 값: %v\n", *i)

    var data Data
    fmt.Println(data)
    changeData(data)
    fmt.Println(data)
    // 왜 data의 필드를 변경할 수 없나?
    // 이전에 알아봤듯 golang은 함수 매개변수에 값을 복사해서 전달하기 때문
    // 만약 함수를 거쳐서 입력값을 조작하려면 포인터를 전달해야함
    changeDataWithPointer(&data)
    fmt.Println(data)

    newData := new(Data) // 내장 기능 new 연산자로 초기화 할 수 있음
    fmt.Println(newData)
    fmt.Println(&newData)

    pointerData := &Data{} // 이 방식으로도 초기화 가능
    fmt.Println(pointerData)
    fmt.Println(&pointerData)

    // golang에서는 객체(인스턴스)를 메모리 공간에 할당하는 방식이 다른 언어들과 조금 다름
    // 컴파일 타임에 컴파일러가 탈출분석* 으로 메모리 할당 위치를 미리 지정
    // *탈출 분석: 특정 함수 외부로 인스턴스가 나가는 경우 등
}


