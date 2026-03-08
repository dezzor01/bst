# go-bst — Generic Binary Search Tree in Go

Простая и типобезопасная реализация **бинарного дерева поиска** (BST) с использованием generics (Go 1.18+).

[![Go Reference](https://pkg.go.dev/badge/github.com/dezzor01/bst/tree.svg)](https://pkg.go.dev/github.com/dezzor01/bst/tree)

## Особенности

- Типобезопасно благодаря `cmp.Ordered`
- Поддержка дубликатов: вставка игнорируется (возвращает `false`)
- Методы: Insert, Delete, Search, Min/Max, Size, Levels, IsEmpty, Clear
- Обходы: PreOrder, InOrder, PostOrder, LevelOrder (BFS), LevelValues (по уровням)
- Графическая печать дерева (`PrettyPrint`)
- Полные тесты

## Установка

```bash
go get github.com/dezzor01/bst/tree

```
## Примеры использования

### Базовый пример (вставка, обход, печать, удаление)

```go
package main

import (
	"fmt"

	"github.com/dezzor01/bst/tree"
)

func main() {
	bst := tree.New[int]()

	// Вставка (дубликаты игнорируются)
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)
	bst.Insert(50) // ← не добавится

	// Отсортированный вывод (InOrder)
	fmt.Print("In-order (отсортировано): ")
	bst.InOrder(func(v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Println()
	// Output: 20 30 40 50 60 70 80

	// Графическое представления дерева
	fmt.Println("\nСтруктура дерева:")
	bst.PrettyPrint()
	/*
	Binary Search Tree:
	50
	├── 30
	│   ├── 20
	│   └── 40
	└── 70
	    ├── 60
	    └── 80
	*/

	// Дополнительные методы
	min, _ := bst.Min()
	fmt.Printf("\nМинимальное: %d\n", min)     // 20

	fmt.Printf("Узлов: %d\n", bst.Size())      // 7
	fmt.Printf("Уровней: %d\n", bst.Levels())  // 3

	// Удаление узла
	bst.Delete(30)
	fmt.Println("\nПосле удаления 30:")
	bst.PrettyPrint()
}
```

### Обход по уровням
```go
bst.LevelOrder(func(v int) {
	fmt.Printf("%d ", v)
})
// Output: 50 30 70 20 40 60 80
```

### Значения по уровням
```go
levels := bst.LevelValues()
for i, level := range levels {
	fmt.Printf("Level %d: %v\n", i, level)
}
// Output:
// Level 0: [50]
// Level 1: [30 70]
// Level 2: [20 40 60 80]
```

