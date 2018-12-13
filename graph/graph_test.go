package graph_test

import (
	"github.com/bruteforce1414/queues/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGraph(t *testing.T) {

	a:=assert.New(t)
	//create object graph without nodes and edges
	testGraph := graph.NewGraph()
	t.Log("Создан тестовый граф")

	//create node and add to graph
	testNode1 := graph.NewNode(1, "Вершина1");testGraph.AddNode(testNode1)
	a.Equal(true,testGraph.Has(testNode1))
	testNode2 := graph.NewNode(2, "Вершина2");testGraph.AddNode(testNode2)
	a.Equal(true,testGraph.Has(testNode2))
	testNode3 := graph.NewNode(3, "Вершина3");testGraph.AddNode(testNode3)
	a.Equal(true,testGraph.Has(testNode3));
	testNode4 := graph.NewNode(4, "Вершина4");testGraph.AddNode(testNode4)
	a.Equal(true,testGraph.Has(testNode4))

	//testNode0:=graph.NewNode(0,"Несуществующая нода 0")
	//testNode10:=graph.NewNode(10,"Несуществующая нода 10")
	//testEdge01:=graph.NewEdge(testNode0,testNode1,0)
	//testGraph.SetEdge(testEdge01)
	//t.Log(testNode10)
	t.Log("Список вершин", testGraph.Nodes())
	testEdge12 := graph.NewEdge(testNode1, testNode2, 12);testGraph.SetEdge(testEdge12)
	a.Equal(true,testGraph.HasEdgeFromTo(testNode1,testNode2))
	testEdge13 := graph.NewEdge(testNode1, testNode3, 13);testGraph.SetEdge(testEdge13)
	testEdge14 := graph.NewEdge(testNode1, testNode4, 14);testGraph.SetEdge(testEdge14)
	testEdge34 := graph.NewEdge(testNode3, testNode4, 34);testGraph.SetEdge(testEdge34)
	t.Log("Вершины, достижимые из вершины 1", testGraph.From(testNode1))
  	 a.Equal(len(testGraph.From(testNode1)),3)
	t.Log("Вершины, идущие к вершине 4.", testGraph.To(testNode4))
	a.Equal(len(testGraph.To(testNode4)),2)
  //  testGraph.RemoveEdge(testEdge12); t.Log("Список ребёр после удаления ребра 12", testGraph)
  //	testGraph.RemoveEdge(testEdge13); t.Log("Список ребёр после удаления ребра 13", testGraph)

	t.Log("Проверка наличия пути 13", testGraph.Edge(testNode1, testNode3))
	t.Log("Проверка наличия пути 14", testGraph.Edge(testNode1, testNode4))
	t.Log("Проверка наличия пути 24", testGraph.Edge(testNode2, testNode4))
	t.Log("Список вершин до удаления вершины 2   ", testGraph)
	//Удаление вершины из списка вершин
	testGraph.RemoveNode(testNode2);
	t.Log("Список вершин после удаления вершины 2", testGraph)
	//Попытка добавить уже существующую вершину
	//testNode1double:=graph.NewNode(1,"Вершина 1");testGraph.AddNode(testNode1double)

}
