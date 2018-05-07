package main

import "fmt"

type QData struct {
	Value string
	QData *QData
}

type Queue struct {
	Tail *QData
}

func (q *Queue) Add(value string) {
	_data := &QData{Value: value}
	if q.Tail != nil {
		q.Tail.QData = _data
	}
	q.Tail = _data
}

func (q *Queue) PrintQueue() {
	d := q.Tail
	for d != nil {
		fmt.Println(d.Value)
		d = d.QData
	}
}
