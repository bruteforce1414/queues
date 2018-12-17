package graph

import (
	"fmt"
)

var (
	VertexsWithoutOutcoming = []Node{}
	VertexsRanged           = []Node{}
	//Значения среза Visited 1-непосещённые, 2 - посещённые, 3 - обработанная
	Visited = []int{}
	lastNode Node
	indexNode int
)

func GraphSortingTopologicalKan(g *DirectedGraph) {

	for range g.nodes {
		Visited = append(Visited, 1)
	}
	//Вначале ищем НОДУ (в единственном числе), из которой не исходят другие ноды. Просматриваем в структуре граф и его полю From карту карт по
	// первичному ключу, равному идентификатору вершины, если такого ключа нет, значит мы помещаем вершину в
	// список S
 	for _, value := range g.nodes {
		if _, ok := g.from[value.Id()]; !ok {
			VertexsWithoutOutcoming=append(	VertexsWithoutOutcoming,value)
			Visited[value.Id()-1]=3
		}
	}
	// Заполняем массив со статусом вершин значением 1 - непосещённые


	for !(len(VertexsWithoutOutcoming)==0){
		checkoutcoming:=false
		for _, value := range g.nodes {
			checkoutcoming=checkoutcoming|| g.HasEdgeFromTo(VertexsWithoutOutcoming[0],value)&&!(Visited[value.Id()-1]==3)
		}
		if !checkoutcoming{
			VertexsRanged=append(VertexsRanged,VertexsWithoutOutcoming[0])
			Visited[VertexsWithoutOutcoming[0].Id()-1]=3
			for _, value := range g.nodes {
				if g.HasEdgeFromTo(value,VertexsWithoutOutcoming[0])&&(Visited[value.Id()-1]==1){
					VertexsWithoutOutcoming=append(VertexsWithoutOutcoming,value)
					Visited[value.Id()-1]=2
				}

			}
			VertexsWithoutOutcoming=VertexsWithoutOutcoming[1:]
		} else{
		VertexsWithoutOutcoming=append(VertexsWithoutOutcoming,VertexsWithoutOutcoming[0])
		VertexsWithoutOutcoming=VertexsWithoutOutcoming[1:]
		}
	}
	for index, _:= range VertexsRanged{
	fmt.Println(index+1,"-й элемент для надевания", VertexsRanged[index].Id()," ", VertexsRanged[index].Name())
	}
	fmt.Scanln()
}
