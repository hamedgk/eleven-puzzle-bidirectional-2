package helpers

import (
	ds "eleven-puzzle/data_structures"
	"eleven-puzzle/data_structures/puzzle"
)

func Intersect(node ds.Node, explored map[puzzle.PuzzleBuffer]bool) bool {
	return explored[node.Puzzle.Buffer]
}

func SearchInNodes(node ds.Node, frontier ds.Queue) (ds.Node, bool){
	first := frontier.Front()
	for current := first; current != nil; current = current.Next(){
		exploredNode := current.Value.(ds.Node)
		
		for current := &exploredNode; current != nil; current = current.Parent{
			if current.Puzzle == node.Puzzle{
				return *current, true
			}
		}

	}
	return ds.Node{}, false
}