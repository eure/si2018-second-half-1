package user

import (
	"fmt"
	"testing"
)

func TestReHeight(t *testing.T) {
	test := ReHeight("170cm", "180cm")
	fmt.Println(test)
}

func TestReAnnualIncome(t *testing.T) {
	test := ReAnnualIncome("200万円", "600万円")
	fmt.Println(test)
}

func TestReAge(t *testing.T) {
	test := ReAge("20歳")
	fmt.Println(test)
}
