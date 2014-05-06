package dataListGo

import (
	"testing"
)

func Test_New_Empty(t *testing.T) {
	l := New()
	if l.Head() != nil {
		t.Error("Test_New_Empty: l.Head() != nil")
	}
	if l.Last() != nil {
		t.Error("Test_New_Empty: l.Last() != nil")
	}
}

func Test_New_List(t *testing.T) {
	l := New("hi")
	if l == nil {
		t.Error("Test_New_List: l == nil")
	}
	n := l.Size()
	if n != 1 {
		t.Error("Test_New_List: l.Size() != 1")
	}
	el := l.Head()
	if el == nil || el.GetData() != "hi" {
		t.Error("Test_New_List: l.Head() != 'hi'", el.GetData())
	}
}

func Test_New_ListElm(t *testing.T) {
	el := NewElm("el")
	if el == nil {
		t.Error("Test_New_ListElm: el == nil")
	}
	if el.GetNext() != nil {
		t.Error("Test_New_ListElm: next != nil")
	}
	if el.GetPrev() != nil {
		t.Error("Test_New_ListElm: prev != nil")
	}
	if s := el.GetData(); s != "el" {
		t.Error("Test_New_ListElm: data != 'el'")
	}
}
