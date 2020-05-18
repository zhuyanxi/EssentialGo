package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Right *Tree

	Value int
}

func NewTree(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = t.insert((1 + v) * k)
	}
	return t
}

func (t *Tree) insert(v int) *Tree {
	if t == nil {
		return &Tree{
			Left:  nil,
			Right: nil,
			Value: v,
		}
	}
	if v < t.Value {
		t.Left = t.Left.insert(v)
	} else {
		t.Right = t.Right.insert(v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
