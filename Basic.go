package dataListGo

/* Basic functions */

/*
todo:
⇥   join : join,
⇥   tail : tail,
⇥   init : init,
*/

func Index(i int, L *List) (interface{}) {
	return LRindex(i, L.Head(), nil, goNext)
}

func Lindex(i int, L *List) (interface{}) {
	return Index(i, L)
}

func Rindex(i int, L *List) (interface{}) {
	return LRindex(i, L.Last(), nil, goPrev)
}

func LRindex(i int, start, last *ListElm, lr LR) (interface{}) {
	var data interface{}

	closure := func(el *ListElm, brk *bool) {
		if (i == 0) {
			data = el.GetData()
			*brk = true
			return
		}
		i -= 1;
	}

	iterLRbreak(start, last, lr, closure, func() { })
	return data
}

func Head(L *List) (*ListElm) {
	return L.Head()
}

func Last(L *List) (*ListElm) {
	return L.Last()
}

func Size(L *List) (uint) {
	return L.Size()
}

func Length(L *List) (uint) {
	return Size(L)
}

func Nil(L *List) (bool) {
	return L.Size() == 0
}

func Empty(L *List) (bool) {
	return Nil(L)
}
