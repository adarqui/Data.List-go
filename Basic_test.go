package dataListGo

import (
	"testing"
)

func Test_Index(t *testing.T) {
	l := New("hi","yo")
	if Index(0, l) != "hi" {
		t.Error("Test_Index: Index(0, l) != 'hi'")
	}
	if Index(1, l) != "yo" {
		t.Error("Test_Index: Index(1, l) != 'yo'")
	}
	if l.Last().GetData() != "yo" {
		t.Error("Test_Index: l.Last() != 'yo'")
	}
}

func Test_Rindex(t *testing.T) {
	l := New("hi","yo")
	if Rindex(0, l) != "yo" {
		t.Error("Test_Index: Index(0, l) != 'yo'")
	}
	if Rindex(1, l) != "hi" {
		t.Error("Test_Index: Index(1, l) != 'hi'")
	}
	if l.Last().GetData() != "yo" {
		t.Error("Test_Index: l.Last() != 'hi'")
	}
}


func Test_Head(t *testing.T) {
	l := New()
	if l.Head() != nil {
		t.Error("Test_Head: l.Head() != nil")
	}
	l = New("1","2","3")
	if l.Head().GetData() != "1" {
		t.Error("Test_Head: l.Head() != '1'")
	}
}


func Test_Last(t *testing.T) {
	l := New()
	if l.Last() != nil {
		t.Error("Test_Last: l.Last() != nil")
	}
	l = New("1","2","3")
	if l.Last().GetData() != "3" {
		t.Error("Test_Last: l.Last() != '3'")
	}
}


func Test_Size(t *testing.T) {
	l := New()
	if l.Size() != 0 {
		t.Error("Test_Size: l.Size() != 0")
	}
	if Size(l) != 0 {
		t.Error("Test_Size: Size(l) != 0")
	}
	l.Ins("1")
	if l.Size() != 1 {
		t.Error("Test_Size: l.Size() != 1")
	}
	if Size(l) != 1 {
		t.Error("Test_Size: Size(l) != 1")
	}
}

func Test_Length(t *testing.T) {
	l := New()
	if Length(l) != 0 {
		t.Error("Test_Length: Length(l) != 0")
	}
	l.Ins("1")
	if Length(l) != 1 {
		t.Error("Test_Length: l.Length(l) != 1")
	}
}


func Test_Nil(t *testing.T) {
	l := New()
	if Nil(l) != true {
		t.Error("Test_Nil: Nil(l) != true")
	}
	l.Ins("hi")
	if Nil(l) != false {
		t.Error("Test_Nil: Nil(l) != false")
	}
}


func Test_Empty(t *testing.T) {
	l := New()
	if Empty(l) != true {
		t.Error("Test_Empty: Empty(l) != true")
	}
	l.Ins("hi")
	if Empty(l) != false {
		t.Error("Test_Empty: Empty(l) != false")
	}
}
