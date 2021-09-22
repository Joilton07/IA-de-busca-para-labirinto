package main

import (
	"fmt"
)

var uses int

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

	tree := Node{initial, nil, []*Node{}}

	result := BFS(&tree, &initial)
	printWorld(result.value)
	print("-----------\n")
	for {
		if result.parent == nil {
			break
		}
		printWorld(result.value)
		result = result.parent
		print("\n")
	}

}

func initWorld(jogo *Game) {
	*jogo = Game{[][]int{{0, 0, 3, 0, 0}, {0, 1, 3, 0, 0}, {0, 3, 3, 0, 0}, {0, 0, 0, 0, 2}}, Point{1, 1}, Point{3, 4}}
}

func printWorld(world Game) {
	for i := 0; i < len(world.field); i++ {
		for j := 0; j < len(world.field[i]); j++ {
			switch world.field[i][j] {
			case 0:
				fmt.Print("-", "")
				break
			case 1:
				fmt.Print("☻", "")
				break
			case 2:
				fmt.Print("♥", "")
				break
			case 3:
				fmt.Print("█", "")
				break
			}
		}
		fmt.Print("\n")
	}
}

func isWin(jogo Game) bool {
	return (jogo.player.lin == jogo.obj.lin) && (jogo.player.col == jogo.obj.col)
}

func validMoviments(game Game) []Game {
	result := []Game{}

	defer recover()

	//verificando movimento para cima
	if game.player.lin-1 >= 0 {
		if game.field[game.player.lin-1][game.player.col] == 0 || game.field[game.player.lin-1][game.player.col] == 2 {
			aux := copy(game)
			aux.field[game.player.lin-1][game.player.col] = 1
			aux.field[game.player.lin][game.player.col] = 0
			aux.player = Point{aux.player.lin - 1, aux.player.col}
			result = append(result, aux)
		}
	}

	//verificando movimento para baixo
	if game.player.lin+1 <= 3 {
		if game.field[game.player.lin+1][game.player.col] == 0 || game.field[game.player.lin+1][game.player.col] == 2 {
			aux := copy(game)
			aux.field[game.player.lin+1][game.player.col] = 1
			aux.field[game.player.lin][game.player.col] = 0
			aux.player = Point{aux.player.lin + 1, aux.player.col}
			result = append(result, aux)
		}
	}

	//verificando movimento para esquerda
	if game.player.col-1 >= 0 {
		if game.field[game.player.lin][game.player.col-1] == 0 || game.field[game.player.lin][game.player.col-1] == 2 {
			aux := copy(game)
			aux.field[game.player.lin][game.player.col-1] = 1
			aux.field[game.player.lin][game.player.col] = 0
			aux.player = Point{aux.player.lin, aux.player.col - 1}
			result = append(result, aux)
		}
	}

	//verificando movimento para direita
	if game.player.col+1 <= 4 {
		if game.field[game.player.lin][game.player.col+1] == 0 || game.field[game.player.lin][game.player.col+1] == 2 {
			aux := copy(game)
			aux.field[game.player.lin][game.player.col+1] = 1
			aux.field[game.player.lin][game.player.col] = 0
			aux.player = Point{aux.player.lin, aux.player.col + 1}
			result = append(result, aux)
		}
	}

	return result
}

func copy(game Game) Game {
	novo := Game{}
	novo.player = game.player
	novo.obj = game.obj

	novo.field = [][]int{}

	for i := 0; i < len(game.field); i++ {

		novo.field = append(novo.field, []int{})

		for j := 0; j < len(game.field[i]); j++ {

			novo.field[i] = append(novo.field[i], game.field[i][j])
		}
	}
	return novo
}

func BFS(root *Node, initial *Game) *Node {
	uses = 0
	queue := []*Node{}
	queue = append(queue, root)
	return BFSconsume(queue)
}

func BFSconsume(queue []*Node) *Node {
	uses += 1
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

// 0 = espaço vazio
// 1 = player
// 2 = objetivo
// 3 = parede

/**
1 0 3 0 0
0 0 3 0 0
0 3 3 0 0
0 0 0 0 2

x - | - -
- - | - -
- | | - -
- - - - o

**/

// -Agente de busca

//----Fazer a arvore consumir
