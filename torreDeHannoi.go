package main

import "fmt"

var users int

type Point struct {
	lin int
	col int
}

type Game struct {
	field  [][]int
	player Point
	obj    Point
}

type Node struct {
	value     Game
	parent    *Node
	childrens []*Node
}

func main() {
	initial := Game{}
	initWorld(&initial)
	printWorld(initial)
}

func initWorld(jogo *Game) {
	*jogo = Game{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, Point{0, 2}, Point{2, 2}}
}

func printWorld(world Game) {
	for i := 0; i < len(world.field); i++ {
		for j := 0; j < len(world.field[i]); j++ {
			switch world.field[i][j] {
			case 0:
				fmt.Print("-", "", "")
			case 1:
				fmt.Print("_", "-", "")
			case 2:
				fmt.Print("|", "_", "-")
			case 3:
				fmt.Print("|", "-", "")
			case 4:
				fmt.Print("|", "_", "")
			case 5:
				fmt.Print("|", "", "")
			case 6:
				fmt.Print("_", "", "")
			}
		}
	}
}

func isWin(jogo Game) bool {
	return (jogo.player.lin == jogo.obj.lin) && (jogo.player.col == jogo.obj.col)
}

func validMoviments(game Game) []Game {
	result := []Game{}

}
func BFS(root *Node, initial *Game) *Node {
	users = 0
	queue := []*Node{}
	queue = append(queue, root)
	return BFSconsume(queue)
}

func BFSconsume(queue []*Node) *Node {
	users += 1
	//time.Sleep(time.Second * 3)

	if len(queue) == 0 {
		fmt.Println("Não existe resultado")
		empty := Node{}
		return &empty
	}

	printWorld(queue[0].value)

	if isWin(queue[0].value) {
		fmt.Println("Resultado encontrado em ", uses, " passos")
		result := queue[0]
		return result
	}

	for _, move := range validMoviments(queue[0].value) {
		novoNo := Node{move, queue[0], []*Node{}}
		queue[0].childrens = append(queue[0].childrens, &novoNo)
		queue = append(queue, &novoNo)
	}

	return BFSconsume(queue[1:])

}

/**
 * |
 * _-
 *
 */
