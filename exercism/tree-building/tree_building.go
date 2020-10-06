package tree

import (
	"errors"
	"sort"
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

type byID []Record

func (records byID) Len() int {
	return len(records)
}

func (records byID) Swap(i, j int) {
	records[i], records[j] = records[j], records[i]
}

func (records byID) Less(i, j int) bool {
	return records[i].ID < records[j].ID
}

func add(record Record, node *Node) (bool, error) {
	if record.Parent == node.ID {
		if node.Children == nil {
			node.Children = []*Node{}
		}
		node.Children = append(node.Children, &Node{ID: record.ID})
		return true, nil
	}

	if node.Children == nil {
		return false, nil
	}

	i := 0
	isAdded := false
	err := errors.New("")
	for i < len(node.Children) && !isAdded {
		isAdded, err = add(record, node.Children[i])
		if err != nil {
			return false, err
		}
		i++
	}
	return isAdded, nil
}

// Build convert records array to the tree
func Build(records []Record) (*Node, error) {

	if records == nil {
		return nil, errors.New("records is nil")
	}
	if len(records) == 0 {
		return nil, nil
	}

	// sort by ID
	sort.Sort(byID(records))

	// check the root
	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("there is no root, or root has a parent")
	}

	// check not continous
	root := &Node{ID: 0}
	i := 1
	for i < len(records) {
		if i != records[i].ID {
			return nil, errors.New("not continuous")
		}
		isAdded, err := add(records[i], root)
		if err != nil || !isAdded {
			return nil, errors.New("cannot add")
		}
		i++
	}
	return root, nil
}
