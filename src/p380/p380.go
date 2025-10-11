package p380

import (
	"math/rand"
)

type RandomizedSet struct {
	slice []int
	dict  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		slice: []int{},
		dict:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dict[val]
	if ok {
		return false
	}

	this.dict[val] = len(this.slice)
	this.slice = append(this.slice, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.dict[val]
	if !ok {
		return false
	}

	lastIdx := len(this.slice) - 1
	lastVal := this.slice[lastIdx]

	this.slice[idx] = lastVal
	this.dict[lastVal] = idx
	this.slice = this.slice[:lastIdx]

	delete(this.dict, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.slice))
	return this.slice[idx]
}
