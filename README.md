# 🚀 CH001 변수와 자료형

> 기본 문법 - 변수(Variables)

## 변수 선언

- 기본 형태
  ```go
  ...
  func main() {
      var myValue int = 0
      var myMessage string = "hello!!!"

      fmt.Println(myValue)
      fmt.Println(myMessage)
      fmt.Println(myValue, myMessage)
  }
  --- Results
  0
  hello!!!
  0 hello!!!
  ```

- 다른 형태
  ```go
  ...
  func main() {
      var myValue1 int // 초기값 생략: 각 타입별 기본값으로 초기화
      var myValue2 = "hello!" // 타입 선언 생략: 우변의 값에 대한 타입 지정
      myValue3 := 123 // 선업 대입문: :=로 var 키워드와 타입 생략
      // 선업 대입문으로 변수를 선언할 때 정수형 등 비트 수를 지정가능한 타입이라면 비트 수는 운영체제 아키텍처 비트수로 지정됨
      fmt.Println(myValue1, myValue2, myValue3)
  }
  --- Results
  0 hello! 123
  ```

## 타입 캐스팅

  ```go
  ...
  func main() {
      a := 309       // int64로 초기화: 0b100110101
      fmt.Println(a) // 309 그대로 출력
      
      b := int8(a)   // int8로 타입 캐스팅: 상위 비트 제거 유도
      fmt.Println(b) // 53출력: 0b110101
    }
  --- Results
  309
  53
  ```

## 병렬 할당과 변수값 교체

```go
  ...
  func main() {
      a, b := 1, 2 // 병렬 할당
      fmt.Println(a, b)
      a, b = b, a // 변수값 교체
      fmt.Println(a, b)
  }
```

## 변수 스코프

  ```go
  ...
  func main() {
      a := 123
      fmt.Println(a)
      {
      b := 234
      fmt.Println(b)
      }
      //fmt.Println(b) // 컴파일 에서
      c := 345
      fmt.Println(c)
  }
  ```
