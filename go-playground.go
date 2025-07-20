package main

import (
    "fmt"
)

func main() {
    var myArray1 []int // 배열 선언, 따로 초기화 하지 않아도 정해진 타입의 요소를 담는 빈 배열 생성
    fmt.Println(myArray1)

    // fmt.Println(myArray1[0]) // 패닉: 런타임 에러: 배열에 존재하지 않는 인덱스

    myArray2 := [3]int{1, 2, 3} // 배열 선언 + 초기화
    fmt.Println(myArray2)

    myArray3 := [...]int{4, 5, 6} // ...으로 배열 길이 생략
    fmt.Println(myArray3)

    myArray4 := [5]int{1, 1} // 배열 길이와 초기화할 일부 요소만 초기화, 순서대로 0, 1 인덱스 요소만 초기화 -> 나머지 요소는 기본값으로 초기화
    fmt.Println(myArray4)

    myArray5 := [5]int{1: 10, 4: 99} // 배열 길이를 정하고 초기화 할 요소의 인덱스를 지정
    fmt.Println(myArray5)

    const length int = 5
    myArray6 := [length]int{}
    //var variableLength int = 5
    //myArray7 := [variableLength]int{} // 배열 길이는 상수만 입력 가능, 컴파일 불가능

    // ---

    for i := 0; i < length; i++ {
        fmt.Print(myArray6[i]) // 배열 요소 접근
        fmt.Print(", ")
    }

    for index, value := range myArray6 { // range 순회(iterator 반복자)
        fmt.Println(index, value) // 인덱스, 요소 출력
    }

    for _, value := range myArray6 { // range 순회(iterator 반복자): range 리턴값 중 인덱스 무시
        fmt.Println(value) // 인덱스, 요소 출력
    }

    // ---

    a := [...]int{1, 2, 3}
    b := [...]int{4, 5, 6}

    fmt.Println(b)
    fmt.Printf("%p\n", &b)
    for i, v := range b {
        fmt.Println(i, v, &v)
    }

    // 배열 a 의 요소들을 b에 모두 복사
    // b에 a에 할당된 배열 객체를 그대로 할당하는것이 아닌 요소의 값만 복사해서 입력
    // 복사한 이후에도 b의 포인터(메모리 주소)값은 바뀌지 않고 값들만 새로 입력됨
    // 따라서 a에 할당한 배열 객체와 b에 할당된 객체는 서로 다른 메모리 공간에 위치한 다른 객체
    b = a

    fmt.Println(b)
    fmt.Printf("%p\n", &b)
    for i, v := range b {
        fmt.Println(i, v, &v)
    }

    fmt.Println(a)
    fmt.Printf("%p\n", &a)
    for i, v := range a {
        fmt.Println(i, v, &v)
    }

    c := []int{1, 2, 3, 4, 5}
    fmt.Println(c)
    // b = c // 만약 배열 길이가 다르다면 복사할 수 없음. 컴파일 불가

    d := [...]string{"hi,", "i'm", "Neo."}
    fmt.Println(d)
    // b = d // 마찬가지로 타입이 달라도 복사할 수 없음. 컴파일 불가

    // ---

    //aa := [...][...]int{{11, 12, 13}, {21, 22, 23}, {31, 32, 33}} // 컴파일 불가
    aa := [...][3]int{{11, 12, 13}, {21, 22, 23}, {31, 32, 33}}
    fmt.Println(aa)

    bb := [...][3]int{{11, 12, 13}, {21, 22, 23}, {31, 32, 33}}
    fmt.Println(bb)

    cc := [...][3]int{
        {11, 12, 13},
        {21, 22, 23},
        {31, 32, 33},
    } // 이렇게 닫는 중괄호가 마지막 요소와 다른줄에 있다면 반드시 마지막 요소 뒤에 쉼표를 입력해야함
    fmt.Println(cc)

    // ---

    quizArr := [3]int{0, 1, 2}

    Quiz := func (arr [3]int) { // 익명 함수 선언
        arr[1] = 200
    }
    Quiz(quizArr)

    // 값이 다른 이유는 함수를 호출할 때 함수의 매개변수로 배열의 포인터가 전달되는것이 아닌 매개변수에 배열을 복사하기 때문
    // 이 때, 값을 변경하기위해 필요한 것이 포인터, 포인터는 다른 챕터에서 확인
    fmt.Println(quizArr[1] == 0)
}
