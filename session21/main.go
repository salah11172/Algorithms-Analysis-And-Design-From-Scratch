package main

import (
	"container/heap"
	"fmt"
)

type HeapNode struct {
	data  rune
	freq  int
	left  *HeapNode
	right *HeapNode
}

type Huffman struct {
	codes map[rune]string
}

type Item struct {
	value    interface{}
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func NewHeapNode(data rune, freq int) *HeapNode {
	return &HeapNode{data: data, freq: freq}
}

func NewHuffman(message string) *Huffman {
	freqHash := make(map[rune]int)
	for _, char := range message {
		freqHash[char]++
	}

	minHeap := make(PriorityQueue, len(freqHash))
	i := 0
	for k, v := range freqHash {
		minHeap[i] = &Item{value: NewHeapNode(k, v), priority: v}
		i++
	}
	heap.Init(&minHeap)

	internalChar := rune(0)
	for len(minHeap) > 1 {
		left := heap.Pop(&minHeap).(*Item).value.(*HeapNode)
		right := heap.Pop(&minHeap).(*Item).value.(*HeapNode)
		newFreq := left.freq + right.freq
		top := NewHeapNode(internalChar, newFreq)
		top.left = left
		top.right = right
		heap.Push(&minHeap, &Item{value: top, priority: newFreq})
	}

	h := &Huffman{codes: make(map[rune]string)}
	h.generateCodes(minHeap[0].value.(*HeapNode), "")
	return h
}

func (h *Huffman) generateCodes(node *HeapNode, str string) {
	if node == nil {
		return
	}
	if node.data != rune(0) {
		h.codes[node.data] = str
	}
	h.generateCodes(node.left, str+"0")
	h.generateCodes(node.right, str+"1")
}

func main() {
	msg := "The output from Huffman's algorithm can be viewed as a variable length code table for encoding a source symbol. The algorithm derives this table from the estimated probability or frequency of occurrence for each possible value of the source symbol. as in other entropy encoding methods, more common symbols are generally represented using fewer bits than less common symbols. Huffman's method can be efficiently implemented, finding a code in time linear to the number of input weights if these weights are sorted. However, although optimal among methods encoding symbols separately, Huffman coding is not always optimal among all compression methods it is replaced with arithmetic coding or asymmetric numeral systems if better compression ratio is required."

	huff := NewHuffman(msg)

	for k, v := range huff.codes {
		fmt.Printf("%c %s\n", k, v)
	}
}
