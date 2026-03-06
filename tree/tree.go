// Package tree представляет собой простую реализацию бинарного древа поиска
// с поддержкой обобщенныйх типов
// Поддерживает опции: вставка, поиск, удаление, обходы
// определение высоты, поиск минимального/максимального значения

// Package tree provides a generic binary search tree (BST) implementation
// using Go generics and cmp.Ordered constraint.
package tree

import (
	"cmp"
	"container/list"
	"fmt"
)

// Node - узел бинарного дерева

// Node represents a single node in the binary search tree.
type Node[T cmp.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// BinaryTree - структура бинарного дерева

// BinaryTree is a generic binary search tree.
// All operations maintain the BST property: left < node < right.
type BinaryTree[T cmp.Ordered] struct {
	Root *Node[T]
}

// New создает новое пустое дерево
func New[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{Root: nil}
}

// NewNode создает новый узел дерева с указанным значением

// NewNode creates a new tree node with the given value.
func NewNode[T cmp.Ordered](value T) *Node[T] {
	return &Node[T]{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

// Insert вставляет элемент с указанным значением в дерево
// Возвращает true если элемент был добавлен в дерево
// Если значение уже существует - вставка игнорируется и возвращается false

// Insert adds a value to the tree if it does not already exist.
// Returns true if the value was inserted, false if it was a duplicate.
func (bt *BinaryTree[T]) Insert(value T) bool {
	var found bool
	bt.Root, found = InsertNode(bt.Root, value)
	return found
}

func InsertNode[T cmp.Ordered](node *Node[T], value T) (*Node[T], bool) {
	if node == nil {
		return NewNode(value), true
	}

	if value < node.Value {
		var ok bool
		node.Left, ok = InsertNode(node.Left, value)
		return node, ok

	} else if value > node.Value {
		var ok bool
		node.Right, ok = InsertNode(node.Right, value)
		return node, ok
	}
	// Если значение уже есть в дереве вставка игнорируется и возвращаем false
	return node, false
}

// Delete удаляет элемент из дерева
// Возвращает true если элемент был найдени и удален

// Delete removes a value from the tree if it exists.
// Returns true if the value was found and removed, false otherwise.
func (bt *BinaryTree[T]) Delete(value T) bool {
	var changed bool
	bt.Root, changed = DeleteNode(bt.Root, value)
	return changed
}

func DeleteNode[T cmp.Ordered](node *Node[T], value T) (*Node[T], bool) {
	if node == nil {
		return nil, false
	}

	if value < node.Value {
		var delete bool
		node.Left, delete = DeleteNode(node.Left, value)
		return node, delete
	} else if value > node.Value {
		var delete bool
		node.Right, delete = DeleteNode(node.Right, value)
		return node, delete
	} else {
		if node.Left == nil {
			return node.Right, true
		}
		if node.Right == nil {
			return node.Left, true
		}

		//При наличии двух потомков через successor
		minNode := minValueNode(node.Right)
		node.Value = minNode.Value
		var delete bool
		node.Right, delete = DeleteNode(node.Right, minNode.Value)
		return node, delete
	}
}

func minValueNode[T cmp.Ordered](node *Node[T]) *Node[T] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Search ищет элемент в дереве

// Search checks whether a value exists in the tree.
func (bt *BinaryTree[T]) Search(Data T) bool {
	return SearchData(bt.Root, Data)
}

func SearchData[T cmp.Ordered](node *Node[T], value T) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	}

	if value > node.Value {
		return SearchData(node.Right, value)
	} else if value < node.Value {
		return SearchData(node.Left, value)
	}
	return false
}

// Levels возвращает количество уровней в дереве
// Пустое дерево -> 0, один узел -> 1

// Levels returns the number of levels in the tree (1 for a single node, 0 for empty).
func (bt *BinaryTree[T]) Levels() int {
	return levelsNode(bt.Root)
}

func levelsNode[T cmp.Ordered](node *Node[T]) int {
	if node == nil {
		return 0
	}
	left := levelsNode(node.Left)
	right := levelsNode(node.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// Min возвращает минимальное значение в дереве
// Если дерево пустое возвращает нулевое значение типа T и false

// Min returns the smallest value in the tree and true if the tree is not empty.
func (bt *BinaryTree[T]) Min() (T, bool) {
	if bt.Root == nil {
		var zero T
		return zero, false
	}
	return minNode(bt.Root)
}

func minNode[T cmp.Ordered](node *Node[T]) (T, bool) {
	if node.Left == nil {
		ok := true
		return node.Value, ok
	}
	return minNode(node.Left)
}

// Max возвращает максимальное значение в дереве
// Если дерево пустое возвращает нулевое значение типа T и false

// Max returns the largest value in the tree and true if the tree is not empty.
func (bt *BinaryTree[T]) Max() (T, bool) {
	if bt.Root == nil {
		var zero T
		return zero, false
	}
	return maxNode(bt.Root)
}

func maxNode[T cmp.Ordered](node *Node[T]) (T, bool) {
	if node.Right == nil {
		ok := true
		return node.Value, ok
	}
	return maxNode(node.Right)
}

// VisitFunc — тип функции, которую мы будем вызывать для каждого узла

// VisitFunc is the type of function called for each value during traversal.
type VisitFunc[T cmp.Ordered] func(value T)

// PreOrder обходит дерево в порядке: корень -> левое -> правое

// PreOrder traverses the tree in pre-order (root -> left -> right).
func (bt *BinaryTree[T]) PreOrder(visit VisitFunc[T]) {
	preOrderNode(bt.Root, visit)
}

func preOrderNode[T cmp.Ordered](node *Node[T], visit VisitFunc[T]) {
	if node == nil {
		return
	}
	visit(node.Value)
	preOrderNode(node.Left, visit)
	preOrderNode(node.Right, visit)
}

// InOrder обходит дерево в порядке левое -> корень -> правое

// InOrder traverses the tree in in-order (left -> root -> right).
func (bt *BinaryTree[T]) InOrder(visit VisitFunc[T]) {
	inOrderNode(bt.Root, visit)
}

func inOrderNode[T cmp.Ordered](node *Node[T], visit VisitFunc[T]) {
	if node == nil {
		return
	}
	inOrderNode(node.Left, visit)
	visit(node.Value)
	inOrderNode(node.Right, visit)

}

// PostOrder обходит дерево в порядке: левое -> правое -> корень

// PostOrder traverses the tree in post-order (left -> right -> root).
func (bt *BinaryTree[T]) PostOrder(visit VisitFunc[T]) {
	postOrderNode(bt.Root, visit)
}

func postOrderNode[T cmp.Ordered](node *Node[T], visit VisitFunc[T]) {
	if node == nil {
		return
	}
	postOrderNode(node.Left, visit)
	postOrderNode(node.Right, visit)
	visit(node.Value)
}

// LevelOrder обходит дерево по уровням (слева направо)

// LevelOrder traverses the tree level by level (breadth-first), left to right.
func (bt *BinaryTree[T]) LevelOrder(visit VisitFunc[T]) {
	if bt.Root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(bt.Root)

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)

		node := elem.Value.(*Node[T])
		visit(node.Value)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
}

// LevelValues обходит дерево по уровням (слева направо)
// Возвращает срезы по уровням

// LevelValues returns all values grouped by levels (breadth-first order).
// Each inner slice contains values from one level, left to right.
func (bt *BinaryTree[T]) LevelValues() [][]T {
	if bt.Root == nil {
		return nil
	}

	result := make([][]T, 0)
	queue := list.New()
	queue.PushBack(bt.Root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		level := make([]T, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			front := queue.Front()
			queue.Remove(front)
			node := front.Value.(*Node[T])

			level = append(level, node.Value)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}

// Size возвращает количество узлов в дереве

// Size returns the total number of nodes in the tree.
func (bt *BinaryTree[T]) Size() int {
	return sizeNode(bt.Root)
}

func sizeNode[T cmp.Ordered](node *Node[T]) int {
	if node == nil {
		return 0
	}
	return 1 + sizeNode(node.Left) + sizeNode(node.Right)
}

// IsEmpty проверяет пустое ли дерево

// IsEmpty reports whether the tree contains no nodes.
func (bt *BinaryTree[T]) IsEmpty() bool {
	return bt.Root == nil
}

// Clear очищает дерево (удаляет указатель на корень)

// Clear removes all nodes from the tree (sets root to nil).
func (bt *BinaryTree[T]) Clear() {
	bt.Root = nil
}

// PrettyPrint выводит дерево в читаемом виде(вертикально)
// Корень вверху, левое поддерево слева, правое справа

// PrettyPrint prints a human-readable representation of the tree.
func (bt *BinaryTree[T]) PrettyPrint() {
	if bt.Root == nil {
		fmt.Println("(пустое дерево)")
		return
	}

	printPretty(bt.Root, "", true)
}

func printPretty[T cmp.Ordered](node *Node[T], prefix string, isLeft bool) {
	if node == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Printf("%v\n", node.Value)

	nextPrefix := prefix
	if isLeft {
		nextPrefix += "│   "
	} else {
		nextPrefix += "    "
	}

	printPretty(node.Left, nextPrefix, true)
	printPretty(node.Right, nextPrefix, false)
}
