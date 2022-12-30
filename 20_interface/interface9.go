package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student age: %d", s.Age)
}

func PrintAge(stringer Stringer) { // 1. Student 인스턴스가 Stringer 타입으로 사용될 수 있음

	fmt.Printf("stringer type: %T\n", stringer) // 인자로 넘어온 stringer가 *Student 타입 인스턴스를 가리킴
	fmt.Println(stringer.String())              // Stringer 타입의 메서드 사용 가능

	// Student 타입의 Age 속성에 접근하기 위해 Student 포인터 타입으로 변환하고자 함
	// 2. Student 타입이 Stringer 인터페이스를 구현했기 때문에 문법 상 타입 변환할 수 있어 컴파일타임 에러가 발생하지 않음
	// 3. stringer가 *Student 인스턴스를 가리키고 있기 때문에 구체 타입으로 변환할 수 있어 런타임 에러가 발생하지 않음
	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)

}

func main() {
	s := &Student{19} // Student는 Stringer 타입으로 사용될 수 있음
	PrintAge(s)       // Stringer 타입으로 PrintAge 함수에 인자로 입력
}
