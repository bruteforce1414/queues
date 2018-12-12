package graph_test

import (
	"fmt"
	"github.com/bruteforce1414/queues/graph"
	"testing"
)

func TestNewGraph(t *testing.T) {

	//create object graph without nodes and edges
	testGraph:=graph.NewGraph()
	fmt.Println("Создан тестовый граф:",testGraph)
	//create node and add to graph
	testNode1:=graph.NewNode(1,"Вершина1");testGraph.AddNode(testNode1)
	testNode2:=graph.NewNode(2,"Вершина2");testGraph.AddNode(testNode2)
	testNode3:=graph.NewNode(3,"Вершина3");testGraph.AddNode(testNode3)
	testNode4:=graph.NewNode(4,"Вершина4");testGraph.AddNode(testNode4)
	fmt.Println("Список вершин", testGraph.Nodes())

  //  testEdge12:=graph.NewEdge(testNode1,testNode2,12);testGraph.SetEdge(testEdge12)
//	testEdge13:=graph.NewEdge(testNode1,testNode3,13);testGraph.SetEdge(testEdge13)
//	testEdge14:=graph.NewEdge(testNode1,testNode4,14);testGraph.SetEdge(testEdge14)
//	testEdge34:=graph.NewEdge(testNode3,testNode4,34);testGraph.SetEdge(testEdge34)
/*	fmt.Println("Вершины, ведущие к вершине 4", testGraph.To(testNode4))
	fmt.Println("Вершины, ведущие от вершины 1", testGraph.From(testNode1))
*/
//	testGraph.RemoveEdge(testEdge12); fmt.Println("Список рёбер после удаления дуги 12", testGraph)

	testGraph.RemoveNode(testNode2); fmt.Println("Список вершин после удаления вершины 2", testGraph)
	//Попытка добавить уже существующую вершину
//	testNode1double:=graph.NewNode(1,"Вершина 1");testGraph.AddNode(testNode1double)

}