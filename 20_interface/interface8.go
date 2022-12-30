package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student(%s, %d)", s.Name, s.Age)
}

type Actor struct {
	Name  string
	Age   int
	Genre string
}

func (a *Actor) String() string {
	return fmt.Sprintf("Actor(%s, %d, %s)", a.Name, a.Age, a.Genre)
}

func convertFromStudentToActor(stringer Stringer) { // 1. Student 인스턴스가 Stringer 타입으로 사용될 수 있음
	fmt.Printf("stringer: %T\n", stringer)

	// 2. Actor는 Stringer 인터페이스를 구현하므로 컴파일타임 에러 발생하지 않음
	// 3. stringer 변수가 가리키는 값인 student 인스턴스가 *Student 타입이므로, *Actor 구체 타입으로 변환할 수 없어 런타임 에러 발생함
	actor := stringer.(*Actor)
	fmt.Println(actor)
}

func main() {
	student := &Student{"Eraser", 28}
	convertFromStudentToActor(student) // student 인스턴스를 Stringer 타입으로 사용할 수 있음 -> 인자로 넘김

}
