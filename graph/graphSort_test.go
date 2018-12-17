package graph_test

import (
	"fmt"
	"github.com/bruteforce1414/queues/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	a := assert.New(t)
	graphClothes := graph.NewGraph()
	underpantsNode1 := graph.NewNode(1, "underpants")
	graphClothes.AddNode(underpantsNode1)
	trousersNode2 := graph.NewNode(2, "trousers")
	graphClothes.AddNode(trousersNode2)
	beltNode3 := graph.NewNode(3, "belt")
	graphClothes.AddNode(beltNode3)
	shirtNode4 := graph.NewNode(4, "shirt")
	graphClothes.AddNode(shirtNode4)
	tieNode5 := graph.NewNode(5, "tie")
	graphClothes.AddNode(tieNode5)
	blazerNode6 := graph.NewNode(6, "blazer")
	graphClothes.AddNode(blazerNode6)
	socksNode7 := graph.NewNode(7, "socks")
	graphClothes.AddNode(socksNode7)
	shoesNode8 := graph.NewNode(8, "shoes")
	graphClothes.AddNode(shoesNode8)
	watchNode9 := graph.NewNode(9, "watch")
	graphClothes.AddNode(watchNode9)

	underpantsNode1_trousersNode2 := graph.NewEdge(underpantsNode1, trousersNode2, 0)
	graphClothes.SetEdge(underpantsNode1_trousersNode2)
	trousersNode2_beltNode3 := graph.NewEdge(trousersNode2, beltNode3, 0)
	graphClothes.SetEdge(trousersNode2_beltNode3)
	beltNode3_blazerNode6 := graph.NewEdge(beltNode3, blazerNode6, 0)
	graphClothes.SetEdge(beltNode3_blazerNode6)
	shirtNode4_beltNode3 := graph.NewEdge(shirtNode4, beltNode3, 0)
	graphClothes.SetEdge(shirtNode4_beltNode3)
	shirtNode4_tieNode5 := graph.NewEdge(shirtNode4, tieNode5, 0)
	graphClothes.SetEdge(shirtNode4_tieNode5)
	tieNode5_blazerNode6 := graph.NewEdge(tieNode5, blazerNode6, 0)
	graphClothes.SetEdge(tieNode5_blazerNode6)
	underpantsNode1_shoesNode8 := graph.NewEdge(underpantsNode1, shoesNode8, 0)
	graphClothes.SetEdge(underpantsNode1_shoesNode8)
	trousersNode2_shoesNode8 := graph.NewEdge(trousersNode2, shoesNode8, 0)
	graphClothes.SetEdge(trousersNode2_shoesNode8)
	socksNode7_shoesNode8 := graph.NewEdge(socksNode7, shoesNode8, 0)
	graphClothes.SetEdge(socksNode7_shoesNode8)

	t.Log("Порядок надевания одежды до упорядочивания", graphClothes)

	VertexRanged := graph.GraphSortingTopologicalKan(graphClothes)
	for index, _ := range VertexRanged {
		t.Log(index+1, "-й элемент для надевания", VertexRanged[index].Id(), " ", VertexRanged[index].Name())
	}
	fmt.Scanln()
	fmt.Scanln()
	var (
		index1Underpants int
		index4Shirt      int
		index2Trousers   int
		index7Socks      int
		index3Belt       int
		index5Tie        int
		index9Watch      int
		index8Shoes      int
		index6Blazer     int
	)
	for index, value := range VertexRanged {
		if value.Name() == "underpants" {
			index1Underpants = index
		}

		if value.Name() == "shirt" {
			index4Shirt = index
		}
		if value.Name() == "trousers" {
			index2Trousers = index
		}
		if value.Name() == "socks" {
			index7Socks = index
		}
		if value.Name() == "belt" {
			index3Belt = index
		}

		if value.Name() == "tie" {
			index5Tie = index
		}
		if value.Name() == "watch" {
			index9Watch = index
		}
		if value.Name() == "shoes" {
			index8Shoes = index
		}

		if value.Name() == "blazer" {
			index6Blazer = index
		}

	}

	a.Equal(true, index6Blazer == 0 || index6Blazer == 1 || index6Blazer == 2)
	a.Equal(true, index8Shoes == 0 || index8Shoes == 1 || index8Shoes == 2)
	a.Equal(true, index9Watch == 0 || index9Watch == 1 || index9Watch == 2)
	a.Equal(true, index7Socks > index8Shoes)
	a.Equal(true, index5Tie > index6Blazer)
	a.Equal(true, index3Belt > index6Blazer)
	a.Equal(true, index4Shirt > index5Tie)
	a.Equal(true, index4Shirt > index3Belt)
	a.Equal(true, index2Trousers > index3Belt)
	a.Equal(true, index1Underpants > index2Trousers)
}
