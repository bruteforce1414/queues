package graph_test

import (
	"fmt"
	"github.com/bruteforce1414/queues/graph"
	"testing"
)
func TestSort(t *testing.T) {
//	a:=assert.New(t)
	graphClothes:=graph.NewGraph()
	underpantsNode1:=graph.NewNode(1,"underpants"); graphClothes.AddNode(underpantsNode1)
	trousersNode2:=graph.NewNode(2,"trousers"); graphClothes.AddNode(trousersNode2)
	beltNode3:=graph.NewNode(3,"belt"); graphClothes.AddNode(beltNode3)
	shirtNode4:=graph.NewNode(4,"shirt"); graphClothes.AddNode(shirtNode4)
	tieNode5:=graph.NewNode(5,"tie"); graphClothes.AddNode(tieNode5)
	blazerNode6:=graph.NewNode(6,"blazer"); graphClothes.AddNode(blazerNode6)
	socksNode7:=graph.NewNode(7,"socks"); graphClothes.AddNode(socksNode7)
	shoesNode8:=graph.NewNode(8,"shoes");graphClothes.AddNode(shoesNode8)
	watchNode9:=graph.NewNode(9,"watch"); graphClothes.AddNode(watchNode9)

	underpantsNode1_trousersNode2:=graph.NewEdge(underpantsNode1,trousersNode2,0); graphClothes.SetEdge(underpantsNode1_trousersNode2)
	trousersNode2_beltNode3:=graph.NewEdge(trousersNode2,beltNode3,0);graphClothes.SetEdge(trousersNode2_beltNode3)
	beltNode3_blazerNode6:=graph.NewEdge(beltNode3,blazerNode6,0);graphClothes.SetEdge(beltNode3_blazerNode6)
	shirtNode4_beltNode3:=graph.NewEdge(shirtNode4,beltNode3,0);graphClothes.SetEdge(shirtNode4_beltNode3)
	shirtNode4_tieNode5:=graph.NewEdge(shirtNode4,tieNode5,0);graphClothes.SetEdge(shirtNode4_tieNode5)
	tieNode5_blazerNode6:=graph.NewEdge(tieNode5,blazerNode6,0); graphClothes.SetEdge(tieNode5_blazerNode6)
	underpantsNode1_shoesNode8:=graph.NewEdge(underpantsNode1,shoesNode8,0);graphClothes.SetEdge(underpantsNode1_shoesNode8)
	trousersNode2_shoesNode8:=graph.NewEdge(trousersNode2,shoesNode8,0);graphClothes.SetEdge(trousersNode2_shoesNode8)
	socksNode7_shoesNode8:=graph.NewEdge(socksNode7,shoesNode8,0);graphClothes.SetEdge(socksNode7_shoesNode8)

	fmt.Println("Порядок надевания одежды до упорядочивания", graphClothes)
	graph.GraphSortingTopologicalKan(graphClothes)
	fmt.Scanln()

}
