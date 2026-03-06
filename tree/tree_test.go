package tree

import (
	"reflect"
	"slices"
	"testing"
)

func TestBinaryTree_BasicOperations(t *testing.T) {
	tree := New[int]()

	tests := []struct {
		name     string
		insert   []int
		expected []int
	}{
		{name: "пустое дерево", insert: []int{}, expected: []int{}},
		{name: "один элемент", insert: []int{42}, expected: []int{42}},
		{
			name:     "несколько элементов без дубликатов",
			insert:   []int{50, 30, 70, 20, 40, 60, 80},
			expected: []int{20, 30, 40, 50, 60, 70, 80},
		},
		{
			name:     "дубликаты игнорируются",
			insert:   []int{50, 30, 50, 70, 30, 20, 20, 40, 40},
			expected: []int{20, 30, 40, 50, 70},
		},
		{
			name:     "линейное дерево (все вправо)",
			insert:   []int{10, 20, 30, 40, 50},
			expected: []int{10, 20, 30, 40, 50},
		},
		{
			name:     "линейное дерево (все влево)",
			insert:   []int{50, 40, 30, 20, 10},
			expected: []int{10, 20, 30, 40, 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree.Clear()

			for _, v := range tt.insert {
				_ = tree.Insert(v)
			}

			var got []int
			tree.InOrder(func(v int) { got = append(got, v) })

			if !slices.Equal(got, tt.expected) {
				t.Errorf("InOrder = %v, ожидалось %v", got, tt.expected)
			}

			if tree.Size() != len(tt.expected) {
				t.Errorf("Size() = %d, ожидалось %d", tree.Size(), len(tt.expected))
			}

			if tree.IsEmpty() != (len(tt.expected) == 0) {
				t.Errorf("IsEmpty() = %v, ожидалось %v", tree.IsEmpty(), len(tt.expected) == 0)
			}
		})
	}
}

func TestBinaryTree_Search(t *testing.T) {
	tree := New[int]()
	for _, v := range []int{50, 30, 70, 20, 40, 60, 80} {
		tree.Insert(v)
	}

	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{"корень", 50, true},
		{"левый лист", 20, true},
		{"правый лист", 80, true},
		{"внутренний узел", 30, true},
		{"несуществующий больше максимума", 90, false},
		{"несуществующий меньше минимума", 10, false},
		{"несуществующий посередине", 55, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree.Search(tt.value); got != tt.want {
				t.Errorf("Search(%d) = %v, ожидалось %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Delete(t *testing.T) {
	baseValues := []int{50, 30, 70, 20, 40, 60, 80}

	tests := []struct {
		name     string
		delete   int
		wantOk   bool
		expected []int
	}{
		{"удаление листа (левый)", 20, true, []int{30, 40, 50, 60, 70, 80}},
		{"удаление листа (правый)", 80, true, []int{20, 30, 40, 50, 60, 70}},
		{"удаление узла с одним ребёнком (левый)", 30, true, []int{20, 40, 50, 60, 70, 80}},
		{"удаление узла с одним ребёнком (правый)", 70, true, []int{20, 30, 40, 50, 60, 80}},
		{"удаление корня с двумя детьми", 50, true, []int{20, 30, 40, 60, 70, 80}},
		{"удаление несуществующего", 999, false, []int{20, 30, 40, 50, 60, 70, 80}},
		{"удаление единственного узла", 42, true, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New[int]()
			for _, v := range baseValues {
				tree.Insert(v)
			}
			if tt.name == "удаление единственного узла" {
				tree.Clear()
				tree.Insert(42)
			}

			gotOk := tree.Delete(tt.delete)
			if gotOk != tt.wantOk {
				t.Errorf("Delete(%d) вернул %v, ожидалось %v", tt.delete, gotOk, tt.wantOk)
			}

			var got []int
			tree.InOrder(func(v int) { got = append(got, v) })

			if !slices.Equal(got, tt.expected) {
				t.Errorf("После удаления InOrder = %v, ожидалось %v", got, tt.expected)
			}
		})
	}
}

func TestBinaryTree_MinMax(t *testing.T) {
	tests := []struct {
		name    string
		insert  []int
		wantMin int
		wantMax int
		empty   bool
	}{
		{"пустое дерево", []int{}, 0, 0, true},
		{"один элемент", []int{42}, 42, 42, false},
		{"несколько элементов", []int{50, 30, 70, 20, 40}, 20, 70, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New[int]()
			for _, v := range tt.insert {
				tree.Insert(v)
			}

			min, minOk := tree.Min()
			max, maxOk := tree.Max()

			if tt.empty {
				if minOk || maxOk {
					t.Errorf("Ожидалось false для пустого дерева, получено minOk=%v, maxOk=%v", minOk, maxOk)
				}
			} else {
				if !minOk || min != tt.wantMin {
					t.Errorf("Min() = %d, %v; ожидалось %d, true", min, minOk, tt.wantMin)
				}
				if !maxOk || max != tt.wantMax {
					t.Errorf("Max() = %d, %v; ожидалось %d, true", max, maxOk, tt.wantMax)
				}
			}
		})
	}
}

func TestBinaryTree_Levels(t *testing.T) {
	tests := []struct {
		name   string
		insert []int
		want   int
	}{
		{"пустое дерево", []int{}, 0},
		{"один узел", []int{1}, 1},
		{"два уровня", []int{2, 1, 3}, 2},
		{"три уровня", []int{4, 2, 6, 1, 3, 5, 7}, 3},
		{"вырожденное в линию", []int{10, 20, 30, 40}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New[int]()
			for _, v := range tt.insert {
				tree.Insert(v)
			}
			if got := tree.Levels(); got != tt.want {
				t.Errorf("Levels() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_LevelValues(t *testing.T) {
	tree := New[int]()
	for _, v := range []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 55} {
		tree.Insert(v)
	}

	got := tree.LevelValues()
	want := [][]int{
		{50},
		{30, 70},
		{20, 40, 60, 80},
		{10, 25, 55},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LevelValues():\n got:  %v\nwant: %v", got, want)
	}
}

func TestBinaryTree_PrePostOrder(t *testing.T) {
	tree := New[int]()
	for _, v := range []int{50, 30, 70, 20, 40} {
		tree.Insert(v)
	}

	t.Run("PreOrder", func(t *testing.T) {
		var got []int
		tree.PreOrder(func(v int) { got = append(got, v) })
		want := []int{50, 30, 20, 40, 70}
		if !slices.Equal(got, want) {
			t.Errorf("PreOrder = %v, ожидалось %v", got, want)
		}
	})

	t.Run("PostOrder", func(t *testing.T) {
		var got []int
		tree.PostOrder(func(v int) { got = append(got, v) })
		want := []int{20, 40, 30, 70, 50}
		if !slices.Equal(got, want) {
			t.Errorf("PostOrder = %v, ожидалось %v", got, want)
		}
	})
}

// Дополнительный тест на разные типы (чтобы generics работал)
func TestBinaryTree_StringType(t *testing.T) {
	tree := New[string]()
	tree.Insert("banana")
	tree.Insert("apple")
	tree.Insert("cherry")
	tree.Insert("date")

	var got []string
	tree.InOrder(func(v string) { got = append(got, v) })

	want := []string{"apple", "banana", "cherry", "date"}
	if !slices.Equal(got, want) {
		t.Errorf("InOrder string = %v, ожидалось %v", got, want)
	}

	min, ok := tree.Min()
	if !ok || min != "apple" {
		t.Errorf("Min string = %q, %v; ожидалось \"apple\", true", min, ok)
	}
}
