package main

import (
	"fmt"
	"math/rand"
	"time"
)

type listInt []int
type node struct {
	lchild *node
	rchild *node
	data   int
}

var sorted_list = make(listInt, 0, 10)

func main() {
	tree := new(node)
	list := make(listInt, 10)
	var temp int

	rand.Seed(time.Now().UTC().UnixNano())
	for i, _ := range list {
		for {
			temp = rand.Intn(100)
			if !list.exists(temp) {
				list[i] = temp
				break
			}
		}
	}

	fmt.Printf("%v\n", list)

	for _, v := range list {
		tree.MakeNode(v)
	}
	tree.PreOderTraverse()
	tree.sort()
	fmt.Printf("%v\n", sorted_list)

}
func (self listInt) exists(data int) bool {
	for _, v := range self {
		if data == v {
			return true
		}
	}

	return false
}
func (this *node) MakeNode(data int) {
	if this.data == 0 {
		this.data = data
	}
	if data < this.data {
		if this.lchild == nil {
			n := node{
				lchild: nil,
				rchild: nil,
				data:   data,
			}
			this.lchild = &n
		} else {
			this.lchild.MakeNode(data)
		}
	} else if data > this.data {
		if this.rchild == nil {
			n := node{
				lchild: nil,
				rchild: nil,
				data:   data,
			}
			this.rchild = &n
		} else {
			this.rchild.MakeNode(data)
		}
	}
}
func (this *node) sort() { //排序
	if this.lchild != nil {
		this.lchild.sort() //找到最小的分支
	}
	sorted_list = append(sorted_list, this.data)
	if this.rchild != nil {
		this.rchild.sort()
	}
}
func (this *node) PreOderTraverse() { //先根遍历
	if this == nil {
		return
	}
	fmt.Println(this.data)
	this.lchild.PreOderTraverse()
	this.rchild.PreOderTraverse()

}
func (this *node) MidOderTraverse() { //中根遍历
	if this == nil {
		return
	}
	this.lchild.MidOderTraverse()
	fmt.Println(this.data)
	this.rchild.MidOderTraverse()
}
func (this *node) AftOderTraverse() { //后根遍历
	if this == nil {
		return
	}
	this.lchild.AftOderTraverse()
	this.rchild.AftOderTraverse()
	fmt.Println(this.data)
}
