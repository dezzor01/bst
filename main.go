package main

import (
	tree "bst/tree"
	"fmt"
)

func main() {
	tree := tree.New[int]()

	data := []int{7, 2, 5, 9, 11, 4, 8, 3}

	for _, v := range data {
		tree.Insert(v)
	}

	fmt.Println("In-order :")
	tree.InOrder(func(v int) {
		fmt.Printf("%d ", v)
	})

	fmt.Println("\n\nДерево:")
	tree.PrettyPrint()

	min, _ := tree.Min()
	fmt.Println("Минимальный элемент: ", min)
	fmt.Printf("Высота: %d\n", tree.Levels())

}
