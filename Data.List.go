package dataListGo

type List struct {
	head *ListElm
	last *ListElm
	size uint
}

type ListElm struct {
	data interface{}
	next *ListElm
	prev *ListElm
}

type FN func ()
type LR func (*ListElm) (*ListElm)
type LRFN func (*ListElm);
type LRFNBreak func (*ListElm, *bool)
