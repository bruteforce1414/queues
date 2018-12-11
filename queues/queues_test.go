package queues

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q:=NewQueue();
	a:=assert.New(t)
	q.Enqueue("1")
	q.PrintQueue()
	a.Equal(q.Size(),1)
	q.Enqueue("2")
	q.PrintQueue()
	a.Equal(q.Size(),2)
	q.Enqueue("3")
	q.PrintQueue()
	a.Equal(q.Size(),3)
	fmt.Println("Первый элемент в очереди",q.Front())
	a.Equal(q.Front(),"1")
	q.Dequeue()
	q.PrintQueue()
	a.Equal(q.Size(),2)
	q.Dequeue()
	q.PrintQueue()
	a.Equal(q.Size(),1)
	q.Dequeue()
	q.PrintQueue()
	a.Equal(q.Size(),0)
	q.Dequeue()

}