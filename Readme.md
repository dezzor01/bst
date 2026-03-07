# go-bst — Generic Binary Search Tree in Go

Простая и типобезопасная реализация **бинарного дерева поиска** (BST) с использованием generics (Go 1.18+).

## Особенности

- Типобезопасно благодаря `cmp.Ordered`
- Поддержка дубликатов: вставка игнорируется (возвращает `false`)
- Методы: Insert, Delete, Search, Min/Max, Size, Levels, IsEmpty, Clear
- Обходы: PreOrder, InOrder, PostOrder, LevelOrder (BFS), LevelValues (по уровням)
- Графическая печать дерева (`PrettyPrint`)
- Полные тесты

## Установка

```bash
go get github.com/dezzor01/bst
