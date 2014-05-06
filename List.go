package dataListGo

import (
	"errors"
	"fmt"
)

/* LIST & LISTELM DEFINITIONS */

func New(v ...interface{}) (*List) {
	l := new(List)
	for _, d := range v {
		el := NewElm(d)
		l.InsListElm(el)
	}
	return l
}

func (this *List) Size() uint {
	return this.size
}

func (this *List) Ins(data interface{}) (error) {
	return this.InsLast(data)
}

func (this *List) InsAfter() {
}

func (this *List) InsHead(data interface{}) (error) {
	el := NewElm(data)
	_, err := this.InsListElmHead(el)
	return err
}

func (this *List) InsLast(data interface{}) (error) {
	el := NewElm(data)
	_, err := this.InsListElm(el)
	return err
}

func (this *List) InsListElm(el *ListElm) (*ListElm, error) {
	if this.head == nil || this.last == nil {
		this.head = el
		this.last = el
	} else {
		last_prev := this.last
		this.last = el
		last_prev.SetNext(el)
		el.SetPrev(last_prev)
	}
	this.size += 1
	return el, nil

}

func (this *List) InsListElmHead(el *ListElm) (*ListElm, error) {
	if this.head == nil || this.last == nil {
		this.head = el
		this.last = el
	} else {
		head_prev := this.head
		this.head = el
		this.head.SetNext(head_prev)
	}
	this.size += 1
	return el, nil
}

func (this *List) Rem(el *ListElm) (error) {
	if this.head == nil || this.last == nil {
		return errors.New("List is empty")
	}
	el_prev := el.GetPrev()
	el_next := el.GetNext()
	if el_prev != nil {
		el_prev.SetNext(el_next)
	}
	if el_next != nil {
		el_next.SetPrev(el_prev)
	}

	if el == this.head {
		this.head = el_next
	} else if el == this.last {
		this.last = el_prev
	}

	return nil
}

func (this *List) Head() (*ListElm) {
	return this.head
}

func (this *List) Last() (*ListElm) {
	return this.last
}

func (this *List) Empty() (error) {
	this.head = nil
	this.last = nil
	return nil
}

func (this *List) DumpL() (error) {
	return this.DumpLR(this.head, goNext)
}

func (this *List) DumpR() (error) {
	return this.DumpLR(this.last, goPrev)
}

func (this *List) DumpLR(start *ListElm, lr LR) (error) {
	index := 0;
	banner := func() {
		println("------------------------");
	}
	printer := func(v *ListElm) {
		fmt.Printf("%d: %v\n", index, v.GetData());
		index += 1
	}
	banner()
	iterLR(start, nil, lr, printer);
	return nil
}


/* List Element */

func NewElm(data interface{}) (*ListElm) {
	l := new(ListElm)
	l.data = data
	return l
}

func (this *ListElm) GetData() (interface{}) {
	return this.data;
}

func (this *ListElm) SetData(data interface{}) {
	this.data = data;
}

func (this *ListElm) GetNext() (*ListElm) {
	return this.next;
}

func (this *ListElm) SetNext(next *ListElm) {
	this.next = next;
}

func (this *ListElm) GetPrev() (*ListElm) {
	return this.prev;
}

func (this *ListElm) SetPrev(prev *ListElm) {
	this.prev = prev;
}
