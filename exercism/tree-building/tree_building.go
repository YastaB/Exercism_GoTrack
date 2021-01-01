package tree

import (
	"fmt"
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

// Build convert records array to the tree
func Build(records []Record) (*Node, error) {

	// sort the records with ascending ID numbers
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	recordsMap := map[int]*Node{}

	for i, r := range records {

		if i != r.ID || r.ID < r.Parent || r.ID != 0 && r.ID == r.Parent {
			return nil, fmt.Errorf("invalid node %v", r)
		}

		recordsMap[r.ID] = &Node{ID: r.ID}
		if r.ID != 0 {
			recordsMap[r.Parent].Children = append(recordsMap[r.Parent].Children, recordsMap[r.ID])
		}
	}

	// return the root
	return recordsMap[0], nil
}
