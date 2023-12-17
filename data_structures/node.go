package data_structures

import (
	"eleven-puzzle/data_structures/puzzle"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

type Node struct {
	Parent    *Node
	Direction puzzle.Direction
	Puzzle    puzzle.Puzzle
}

func (node *Node) Expand(queue Queue, explored map[puzzle.PuzzleBuffer]bool) {
	possibleMoves := node.Puzzle.PossibleBlankMoves()
	for _, direction := range possibleMoves {
		copyNode := *node
		copyNode.Direction = direction
		copyNode.Parent = node
		copyNode.Puzzle.MoveBlank(direction)
		if explored[copyNode.Puzzle.Buffer] {
			continue
		}
		queue.Enqueue(copyNode)
		explored[copyNode.Puzzle.Buffer] = true
	}
}

func (node *Node) IsGoal(buffer puzzle.PuzzleBuffer) bool {
	return node.Puzzle.Buffer == buffer
}

func TraceBack(node Node, counter *int) {
	if node.Parent == nil {
		fmt.Println()
		node.Print()
		return
	}

	TraceBack(*node.Parent, counter)
	*counter++
	switch node.Direction {
	case puzzle.Up:
		fmt.Println()
		node.Print()
	case puzzle.Down:
		fmt.Println()
		node.Print()
	case puzzle.Right:
		fmt.Println()
		node.Print()
	case puzzle.Left:
		fmt.Println()
		node.Print()
	default:
		return
	}
}

func TraceForward(node Node, counter *int) {
	if node.Parent == nil {
		fmt.Println()
		node.Print()
		return
	}

	*counter++
	switch node.Direction {
	case puzzle.Up:
		fmt.Println()
		node.Print()
	case puzzle.Down:
		fmt.Println()
		node.Print()
	case puzzle.Right:
		fmt.Println()
		node.Print()
	case puzzle.Left:
		fmt.Println()
		node.Print()
	default:
		return
	}
	TraceForward(*node.Parent, counter)
}

func (node Node) Print() {
	greenBold := color.New(color.FgGreen, color.Bold)
	red := color.New(color.FgHiYellow, color.Bold)
	w := tabwriter.NewWriter(os.Stdout, 4, 1, 2, ' ', 0)
	for i := 0; i < puzzle.Rows; i++ {
		for j := 0; j < puzzle.Cols; j++ {
			if node.Puzzle.Buffer[i][j] == puzzle.Blank {
				red.Fprintf(w, "%v\t", puzzle.BlankStr)
			} else {
				greenBold.Fprintf(w, "%v\t", node.Puzzle.Buffer[i][j])
			}
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	color.Cyan("----------------------------")
}
