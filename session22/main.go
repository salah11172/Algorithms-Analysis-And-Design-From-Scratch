package main

import (
	"fmt"
)

type Item struct {
	name   string
	value  float64
	weight float64
	ratio  float64
}

type Knapsack struct {
	maxCapacity     float64
	currentCapacity float64
	totalValue      float64
	items           []*Item
}

func NewItem(value, weight float64, name string) *Item {
	return &Item{
		name:   name,
		value:  value,
		weight: weight,
		ratio:  value / weight,
	}
}

func NewKnapsack(maxCapacity float64) *Knapsack {
	return &Knapsack{
		maxCapacity:     maxCapacity,
		currentCapacity: 0,
		totalValue:      0,
		items:           make([]*Item, 0),
	}
}

func (k *Knapsack) addItem(newItem *Item) {
	if newItem.weight > k.maxCapacity-k.currentCapacity {
		diff := k.maxCapacity - k.currentCapacity
		newItem.weight = diff
		newItem.value = diff * newItem.ratio
	}

	k.items = append(k.items, newItem)
	k.currentCapacity += newItem.weight
	k.totalValue += newItem.value
}

type ByRatio []*Item

func (a ByRatio) Len() int           { return len(a) }
func (a ByRatio) Less(i, j int) bool { return a[i].ratio > a[j].ratio }
func (a ByRatio) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func mergeSort(myArray []*Item) []*Item {
	if len(myArray) == 1 {
		return myArray
	}
	mid := len(myArray) / 2
	left := mergeSort(myArray[:mid])
	right := mergeSort(myArray[mid:])
	return merge(left, right)
}

func merge(left, right []*Item) []*Item {
	res := make([]*Item, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0].ratio < right[0].ratio {
			res = append(res, right[0])
			right = right[1:]
		} else {
			res = append(res, left[0])
			left = left[1:]
		}
	}
	res = append(res, left...)
	res = append(res, right...)
	return res
}

func printItems(bag *Knapsack) {
	fmt.Println("----------------------------")
	fmt.Println("total value:", bag.totalValue)
	fmt.Println("current capacity:", bag.currentCapacity)
	fmt.Println("items:")
	fmt.Println("n v w")
	for _, item := range bag.items {
		fmt.Printf("%s %f %f\n", item.name, item.value, item.weight)
	}
}

func printArray(items []*Item) {
	fmt.Println("n v w r")
	for _, item := range items {
		fmt.Printf("%s %f %f %f\n", item.name, item.value, item.weight, item.ratio)
	}
}

func main() {
	values := []float64{4, 9, 12, 11, 6, 5}
	weights := []float64{1, 2, 10, 4, 3, 5}

	items := make([]*Item, 0)
	for i := 0; i < len(values); i++ {
		newItem := NewItem(values[i], weights[i], fmt.Sprintf("#%d", i))
		items = append(items, newItem)
	}

	items = mergeSort(items)

	printArray(items)

	maxCapacity := 12.0
	bag := NewKnapsack(maxCapacity)

	j := 0
	for bag.currentCapacity < bag.maxCapacity && j < len(items) {
		bag.addItem(items[j])
		j++
	}

	printItems(bag)
}
