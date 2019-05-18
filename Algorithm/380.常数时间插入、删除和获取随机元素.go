import "math/rand"

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 */
type RandomizedSet struct {
	m        map[int]int
	elements []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		m:        make(map[int]int, 0),
		elements: make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.elements = append(this.elements, val)
	this.m[val] = len(this.elements) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.m[val]; !ok {
		return false
	} else {
		lastValue := this.elements[len(this.elements)-1]
		lastIndex := this.m[lastValue]
		this.elements[index], this.elements[lastIndex] = this.elements[lastIndex], this.elements[index]
		this.m[lastValue] = index
		this.elements = this.elements[:len(this.elements)-1]
		return true
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	selectIndex := rand.Intn(len(this.elements))
	return this.elements[selectIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

