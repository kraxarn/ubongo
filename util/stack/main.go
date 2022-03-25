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
	return stack.list.PushBack(value).Value
}

func Pop[V any](stack *Stack[V]) V {
	if stack.list.Len() == 0 {
		return nil
	}

	front := stack.list.Back()
	stack.list.Remove(front)
	return front.Value
}

func Peek[V any](stack *Stack[V]) V {
	return stack.list.Back().Value
}
