package dataListGo

func iterL(L *List, last *ListElm, fn LRFN) {
	iterLR(L.Head(), last, goNext, fn)
}

func iterR(L *List, last *ListElm, fn LRFN) {
	iterLR(L.Last(), last, goPrev, fn)
}

func iterLR(start, last *ListElm, lr LR, fn LRFN) {
	for v := start ; v != last; v = lr(v) {
		fn(v)
	}
}


/*
 * iteration with the ability to break out
 */
func iterLbreak(L *List, last *ListElm, fn LRFNBreak, defaultcb FN) {
	iterLRbreak(L.Head(), last, goNext, fn, defaultcb)
}

func iterRbreak(L *List, last *ListElm, fn LRFNBreak, defaultcb FN) {
	iterLRbreak(L.Last(), last, goPrev, fn, defaultcb)
}

func iterLRbreak(start, last *ListElm, lr LR, fn LRFNBreak, defaultcb FN) {
	brk := false
	for v := start ; v != last ; v = lr(v) {
		fn(v, &brk)
		if brk == true {
			return
		}
	}
	defaultcb()
}

func goNext(el *ListElm) (*ListElm) {
	return el.GetNext()
}

func goPrev(el *ListElm) (*ListElm) {
	return el.GetPrev()
}
