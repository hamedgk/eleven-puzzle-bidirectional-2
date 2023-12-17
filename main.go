package main

import (
	ds "eleven-puzzle/data_structures"
	"eleven-puzzle/data_structures/puzzle"
	"eleven-puzzle/helpers"
	"fmt"
)

var examplePuzzle = puzzle.PuzzleBuffer{
	//{6, 4, 255, 8},
	//{5, 11, 1, 2},
	//{7, 9, 3, 10},

	//kazemini example
	{2, 7, 255, 6},
	{1, 7, 3, 8},
	{4, 5, 2, 4},

	//{1, 1, 1, 1},
	//{1, 0, 1, 1},
	//{255, 1, 1, 1},
}

func main() {
	sortedArray := puzzle.SortPuzzle(examplePuzzle)

	sourceExplored := map[puzzle.PuzzleBuffer]bool{}
	sourceFrontier := ds.NewQueue()
	sourceFrontier.Enqueue(
		ds.Node{
			Parent:    nil,
			Direction: puzzle.None,
			Puzzle:    puzzle.FromBuffer(examplePuzzle),
		},
	)
	targetExplored := map[puzzle.PuzzleBuffer]bool{}
	targetFrontier := ds.NewQueue()
	targetFrontier.Enqueue(
		ds.Node{
			Parent:    nil,
			Direction: puzzle.None,
			Puzzle:    puzzle.FromBuffer(sortedArray),
		},
	)

	for {

		if node, ok := sourceFrontier.Dequeue(); ok {
			if found := helpers.Intersect(node, targetExplored); found {
				if foundNode, found := helpers.SearchInNodes(node, targetFrontier); found {
					scounter := 0
					tcounter := 0
					ds.TraceBack(node, &scounter)
					fmt.Println("********************")
					ds.TraceForward(foundNode, &tcounter)
					fmt.Printf("\nSteps: %d\n", scounter+tcounter)
					fmt.Printf("\nNodes: %d\n", len(sourceExplored)+len(targetExplored))
					return
				} else {
					panic("found but not located !!!!")
				}
			}
			node.Expand(sourceFrontier, sourceExplored)
		}
		if node, ok := targetFrontier.Dequeue(); ok {
			if found := helpers.Intersect(node, sourceExplored); found {
				if foundNode, found := helpers.SearchInNodes(node, sourceFrontier); found {
					scounter := 0
					tcounter := 0
					ds.TraceBack(foundNode, &scounter)
					fmt.Println("********************")
					ds.TraceForward(node, &tcounter)
					fmt.Printf("\nSteps: %d\n", scounter+tcounter)
					fmt.Printf("\nNodes: %d\n", len(sourceExplored)+len(targetExplored))
					return
				} else {
					panic("found but not located !!!!")
				}
			}
			node.Expand(targetFrontier, targetExplored)
		}

	}
}
