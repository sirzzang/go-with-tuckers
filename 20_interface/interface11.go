package main

import "fmt"

type Attacker interface {
	Attack()
}

type Monster struct {
	Name string
}

func (m *Monster) Attack() {
	fmt.Println("monster attack")
}

func DoAttack(att Attacker) {
	if att != nil {
		att.Attack()
	}
}

func main() {

	var att Attacker
	// att.Attack()
	DoAttack(att)

}
