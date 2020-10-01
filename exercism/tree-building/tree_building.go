package tree

import (
	"errors"
)

// Record is the items in array that will be converted to the tree
type Record struct {
	ID     int
	Parent int
}

// Node tree element
type Node struct {
	ID       int
	Children []*Node
}

// findTheParent finds the parent of the record in the tree
func findTheParent(root *Node, record Record) *Node {
	curNode := root
	if root.ID == record.Parent {
		return curNode
	}
	for _, node := range root.Children {
		found := findTheParent(node, record)
		if found != nil {
			return found
		}
	}
	return nil
}

// removeFromRecords remove efficiently without keeping the order
func removeFromRecords(index int, slice []Record) ([]Record, error) {
	if index >= len(slice) {
		return nil, errors.New("cannot remove")
	}
	slice[index] = slice[len(slice)-1]
	slice[len(slice)-1] = Record{}
	return slice[:len(slice)-1], nil
}

// Build convert records array to the tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	var root *Node = nil
	var err error = nil
	// find the root
	for i, record := range records {
		if record.ID == 0 {
			if record.Parent != 0 {
				return nil, errors.New("root cannot have parent")
			}
			root = &Node{ID: 0}
			records, err = removeFromRecords(i, records)
			if err != nil {
				return nil, err
			}
			break
		}
	}
	if root == nil {
		return nil, errors.New("root cannot be founded")
	}

	currentIndex := 1
	found := true
	for found && len(records) > 0 {
		found = false
		for i, record := range records {
			if record.ID == currentIndex {
				node := findTheParent(root, record)
				if node == nil {
					return nil, errors.New("not found")
				}
				if node.Children == nil {
					node.Children = []*Node{}
				} else {
					// check for duplicate
					for _, node := range node.Children {
						if node.ID == record.ID {
							return nil, errors.New("duplicate is found")
						}
					}
				}
				node.Children = append(node.Children, &Node{ID: record.ID})

				// remove the founded one
				records, err = removeFromRecords(i, records)
				if err != nil {
					return nil, err
				}
				currentIndex++
				found = true
			}
		}
	}
	if len(records) > 0 {
		return nil, errors.New("cannot add")
	}
	return root, nil
}