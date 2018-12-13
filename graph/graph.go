package graph

import (
	"fmt"
	"sync"
)

//узел
type node struct {
	id int
	name string
}


//узел
type Node interface {
	Id() int
	Name() string

}



type edge struct {
	from Node
	to Node
	weight float64
}

// это ребро
type Edge interface {
	From() Node
	To() Node
	Weight() float64
}


type DirectedGraph struct {
nodes 	map[int]Node
from  	map[int]map[int]Edge
to    	map[int]map[int]Edge
lock sync.RWMutex
}

// это основные методы графа
type Graph interface {


	// проверка существует ли такая нода
	Has(Node) bool

	// возвращаем просто все ноды без какого либа разбора-сортировки
	Nodes() []node

	// возвращаем список нод достижимых из заданной ноды
	From(Node) []Node

	// все ноды из которых можно достич указанную
	To(Node) []Node

	// проверка есть ли ребро из u в v
	HasEdgeFromTo(u, v Node) bool

	// возвращает ребро из u и v, если такое ребро существует, иначе nil
	Edge(u, v Node) Edge

	// добавить ноду (без ребер)
	AddNode(n Node)

	// удалить ноду
	RemoveNode(n Node)

	// добавить ребро
	SetEdge(e Edge)

	// удалить ребро
	RemoveEdge(e Edge)
}
// + проверка существует ли такая нода
func (g *DirectedGraph) Has(n Node) bool  {
	if _, ok := g.nodes[n.Id()]; ok {
		return true
	}
	return false
}
// + возвращаем просто все ноды без какого либа разбора-сортировки
func (g *DirectedGraph) Nodes() map[int]Node{
	return g.nodes
}

//+ возвращаем список нод достижимых из заданной ноды
func (g *DirectedGraph) From(n Node) []Node{
	g.lock.RLock()
	defer g.lock.RUnlock()
	var returnedNodes []Node
	if _, ok := g.nodes[n.Id()];ok{
		for _,value1:=range g.from[n.Id()]{
			returnedNodes=append(returnedNodes,value1.To())
		}
		return returnedNodes
	}

return nil
}

// все ноды из которых можно достич указанную
func (g *DirectedGraph) To(n Node) []Node{
	g.lock.RLock()
	defer g.lock.RUnlock()
	var returnedNodes []Node
	if _, ok := g.nodes[n.Id()];ok{
		for _,value:=range g.to[n.Id()]{
			returnedNodes=append(returnedNodes,value.From())
		}
		return returnedNodes
	}
	return nil
}
//+ проверка есть ли ребро из u в v
func (g *DirectedGraph) HasEdgeFromTo(u, v Node) bool{
	if _, ok := g.from[u.Id()][v.Id()]; ok {
		return true

	}

	return false
}

//+возвращает ребро из u и v, если такое ребро существует, иначе nil
func (g *DirectedGraph) Edge(u, v Node) Edge  {
	if _, ok := g.from[u.Id()][v.Id()]; ok {
	return g.from[u.Id()][v.Id()]
		}


	return nil
}

//+добавить ноду (без ребер)
func (g *DirectedGraph) AddNode(n Node) {
	//g.lock.Lock()
	//проверим, есть ли нода с таким же идентификатором в графе
		if g.Has(n){
			fmt.Println("Вершина с идентификатором '", n.Id(),"' уже существует. Вершина не была добавлена к графу")
			return
		}
	g.nodes[n.Id()]=n
//	g.lock.Unlock()
}
//+ удалить ноду
func (g *DirectedGraph) RemoveNode(n Node) {
	g.lock.Lock()
	//проверим, есть ли нода с таким же идентификатором в графе
	if g.Has(n)==false{
		fmt.Println("Вершина с идентификатором '", n.Id(),"' не существует. Невозможно удалить несуществующую вершину")
		return
	}

	//удаляем все рёбра связанные с вершиной

	for _,value:=range g.nodes {
		if g.HasEdgeFromTo(n,value){
			delete(g.from,value.Id())
		}

	}

	for _,value:=range g.nodes {
		if g.HasEdgeFromTo(value,n){
			delete(g.to,value.Id())
		}

	}




	if _, ok := g.to[n.Id()]; ok {
		delete(g.to, n.Id())
	}
	//удаляем вершину
	var indexDelete int
	for _,value:=range g.nodes{
		if value.Id() == n.Id(){
			indexDelete=n.Id()
			break
		}
	}
	delete(g.nodes, indexDelete)




	g.lock.Unlock()

}
// добавить ребро
func (g *DirectedGraph)SetEdge(e Edge)  {
	g.lock.Lock()
	if g.HasEdgeFromTo(e.From(),e.To()){
		fmt.Println("Дуга с началом в вершине '", e.From(),"' до вершины '", e.To(),"' " +
			"уже существует. Дуга не была добавлена к графу. Созданы вершины '", e.From()," и '", e.To(),"'")
		return
	}

		if !g.Has(e.From()){
			g.AddNode(e.From())
		}
		if !g.Has(e.To()){
			g.AddNode(e.To())
		}
	
	//Если первого ключа не существует, то создаём новый массив для вторых ключей
	//Добавление в список дуг ИЗ
	if _, ok := g.from[e.From().Id()]; !ok {
		g.from[e.From().Id()] = map[int]Edge{}
		g.from[e.From().Id()][e.To().Id()] = e
	} else{
		////Если первый ключ уже есть, то добавляем к уже существующему массиву для вторых ключей
		//Добавление в список дуг ИЗ
		g.from[e.From().Id()][e.To().Id()] = e
	}

	//Добавление в список дуг В
	if _, ok := g.to[e.From().Id()]; !ok {
		g.to[e.To().Id()] = map[int]Edge{}
		g.to[e.To().Id()][e.From().Id()] = e
	}else {
	//Добавление в список дуг В
	g.to[e.To().Id()][e.From().Id()] = e
	}


  		g.lock.Unlock()

}
//+ удалить ребро
func (g *DirectedGraph) RemoveEdge(e Edge) {
	g.lock.Lock()
	if !g.HasEdgeFromTo(e.From(),e.To()){
		fmt.Println("Дуга с началом в вершине '", e.From(),"' до вершины '", e.To(),"' " +
			"не существует. Невозможно удалить несуществующую дугу")
		return
	}
	//удаляём второй ключ
		delete(g.from[e.From().Id()], e.To().Id())
	//проверяем длину первой карты, если в ней больше нет элементов, то удаляем первичный ключ
	if len(g.from[e.From().Id()])==0 {
		delete(g.from, e.From().Id())
	}

	//удаляём второй ключ
	delete(g.to[e.To().Id()], e.From().Id())
	//проверяем длину первой карты, если в ней больше нет элементов, то удаляем первичный ключ
	if len(g.to[e.To().Id()])==0 {
		delete(g.to, e.To().Id())
	}

	g.lock.Unlock()
}

//* Получение значений инкапсулированных полей Ноды
func (n *node) Id() int {
 	return n.id
}
func (n *node) Name() string {
	return n.name
}

//* Получение значений инкапсулированных полей Дуги
func (e *edge) From() Node{
	return e.from
}
func (e *edge) To() Node{
	return e.to
}
func (e *edge) Weight() float64 {
	return e.weight
}


//создание нового графа
func NewGraph() *DirectedGraph{
	graphObject:=DirectedGraph{
		map[int]Node{},
		map[int]map[int]Edge{},
		map[int]map[int]Edge{},
		sync.RWMutex{},
	}
	return &graphObject
}
//создание нового узла
func NewNode(id int, name string) Node {
	nodeObject:=node{id:id,name:name}
	return &nodeObject
}
// создание новой дуги
func NewEdge(from Node, to Node, weight float64) Edge {
	edgeObject:=edge{from:from, to: to, weight:weight}
	return &edgeObject
}