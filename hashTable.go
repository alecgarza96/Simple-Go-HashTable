package main

import "fmt"

type Node struct {
	val  int
	Next *Node
}

func traverseNode(node *Node) *Node {
	tempNode := node
	for tempNode.Next != nil {
		tempNode = tempNode.Next
	}
	return tempNode
}

func addNode(node *Node, val int) {
	newNode := Node{val, nil}
	lastNode := traverseNode(node)
	lastNode.Next = &newNode
}

func findNode(node *Node, target int) bool {
	found := false
	for node != nil && found == false {
		if node.val == target {
			found = true
		}
		node = node.Next
	}
	return found
}

func createHashTable(size int) []Node {
	hashTable := make([]Node, size)
	return hashTable
}

func hash(hashTable []Node, num int) bool {
	num = findHashValue(num, len(hashTable))
	//need to be able to traverse
	return hashTable[num%len(hashTable)].val != 0
}

func findHashValue(num int, tableSize int) int {
	for num > tableSize {
		num /= 10
	}
	return num
}

func main() {
	hashTable := createHashTable(9)
	if hash(hashTable, 800) == false {
		location := findHashValue(800, len(hashTable))
		node := hashTable[location]
		lastNode := traverseNode(&node)
		addNode(lastNode, 1)
		//newLastNode := traverseNode(&node)
		fmt.Println(hash(hashTable, 800))
	}
}
