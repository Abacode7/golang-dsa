package heap

import (
	"fmt"
	"testing"
)

func TestNewPersonList(t *testing.T) {
	// // [{Eve 2} {Alice 23} {Bob 25}]
	// []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	family := PersonList{
		{"Person", 9},
		{"Person", 31},
		{"Person", 40},
		{"Person", 22},
		{"Person", 10},
		{"Person", 15},
		{"Person", 1},
		{"Person", 25},
		{"Person", 91},
	}
	fmt.Println("Family: ", family)

	family2 := NewPersonList(
		Person{"Person", 9}, Person{"Person", 31}, Person{"Person", 40},
		Person{"Person", 22}, Person{"Person", 10}, Person{"Person", 15},
		Person{"Person", 1}, Person{"Person", 25}, Person{"Person", 91})

	fmt.Println("Family 2: ", family2)

	fmt.Println("Family 2 pop heap: ", family2.PopHeap())
	// &[{Person 10} {Person 15} {Person 22} {Person 31} {Person 91} {Person 40} {Person 25} {Person 1}]

	fmt.Println("Family 2 after: ", family2)
}
