package stack

import "container/list"

type Stack[V any] struct {
	list *list.List
}

func New[V any]() *Stack[V] {
	return &Stack[V]{
		list: list.New(),
	}
}

// Currently, methods can't contain generics, so functions are required

func Push[V any](stack *Stack[V], value V) V {
	return stack.list.PushBack(value).Value.(V)
}

func Pop[V any](stack *Stack[V]) V {
	if stack.list.Len() == 0 {
		// TODO: We can't return V's address as it doesn't have to be a pointer
		panic("Stack empty")
	}

	front := stack.list.Back()
	stack.list.Remove(front)
	return front.Value.(V)
}

func Peek[V any](stack *Stack[V]) V {
	return stack.list.Front().Value.(V)
}

func Len[V any](stack *Stack[V]) int {
	return stack.list.Len()
}
