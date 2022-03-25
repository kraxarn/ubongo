package game

import (
	"container/list"
	"os"
)

type SceneManager struct {
	list *list.List
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		list: list.New(),
	}
}

func (s *SceneManager) Push(value Scene) Scene {
	s.list.PushBack(value)
	return value
}

func (s *SceneManager) Pop() Scene {
	if s.list.Len() == 0 {
		os.Exit(0)
	}

	front := s.list.Back()
	s.list.Remove(front)
	return front.Value.(Scene)
}

func (s *SceneManager) Peek() Scene {
	return s.list.Back().Value.(Scene)
}

func (s *SceneManager) Len() int {
	return s.list.Len()
}
