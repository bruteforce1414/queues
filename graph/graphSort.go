package graph

import (
	"fmt"
)

var (
	VertexsWithoutOutcoming = []Node{}
	VertexsRanged           = []Node{}
	//Значения среза Visited 1-непосещённые, 2 - посещённые, 3 - обработанная
	Visited = []int{}
)

func GraphSortingTopologicalKan(g *DirectedGraph) {

	fmt.Println(g)
	//Вначале ищем ноды, из которой не исходят другие ноды. Просматриваем в структуре граф и его полю From карту карт по
	// первичному ключу, равному идентификатору вершины, если такого ключа нет, значит мы помещаем вершину в
	// список S

	for _, value := range g.nodes {
		if _, ok := g.from[value.Id()]; !ok {
			VertexsWithoutOutcoming = append(VertexsWithoutOutcoming,value)
		//	Visited[value.Id()-1] = 2
		}
	}
	// Заполняем массив со статусом вершин значением 1 - непосещённые
	for range g.nodes {
		Visited = append(Visited, 1)
	}


			dfs(g, VertexsWithoutOutcoming[0])


	VertexsRanged:=VertexsRanged[:len(VertexsRanged)-1]
	fmt.Println("Порядок надевания одежды", VertexsRanged)
	fmt.Scanln()
}

func dfs(g *DirectedGraph, v Node) {
	Visited[(v).Id()-1] = 2
	for _, value := range g.nodes {
		if g.HasEdgeFromTo(value, v) && (Visited[value.Id()-1] == 1) {
			checking:=false
			for index2, value2 := range VertexsWithoutOutcoming {
					checking=checking||(value == value2)
					}
				if checking==false{
				VertexsWithoutOutcoming = append(VertexsWithoutOutcoming, value)
				Visited[value.Id()-1] = 2
					if checking==true{
						VertexsWithoutOutcoming = append(VertexsWithoutOutcoming, value)
						Visited[value.Id()-1] = 2
			}
		}
	}
	if len(VertexsWithoutOutcoming) > 1 {
		VertexsWithoutOutcoming = append(VertexsWithoutOutcoming[1:])
	}


	if len(VertexsWithoutOutcoming) == 1 {
		x:=VertexsWithoutOutcoming[0]
		VertexsWithoutOutcoming = append(VertexsWithoutOutcoming[:0])
		VertexsRanged = append(VertexsRanged, v)
		dfs(g, x)
	}

	if len(VertexsWithoutOutcoming) == 0 {
		VertexsRanged = append(VertexsRanged, v)
		return
	}

	VertexsRanged = append(VertexsRanged, v)
	Visited[v.Id()-1] = 3

	for i := 1; i < len(Visited); i++{
		if !(Visited[i]==3){
			Visited[i]=1
		}

	}

	dfs(g, VertexsWithoutOutcoming[0])
}
