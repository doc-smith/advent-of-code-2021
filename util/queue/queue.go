package queue

type Queue struct {
	elements []interface{}
	read     int
	write    int
}

func NewQueue() *Queue {
	return &Queue{
		elements: make([]interface{}, 1, 1),
	}
}

func (q *Queue) Empty() bool {
	return q.write == q.read
}

func (q *Queue) Full() bool {
	/* it's a simple trick: we declare the queue full when there is one element
		 left in the buffer (oh no, such a waste)
	   otherwise there will be an ambiguity:
	   	 read == write when the queue is full and
		 read == write when the queue is empty
	*/
	return q.advance(q.write) == q.read
}

func (q *Queue) Enqueue(x interface{}) {
	q.growIfNecessary()
	q.elements[q.write] = x
	q.write = q.advance(q.write)
}

func (q *Queue) Dequeue() interface{} {
	if q.Empty() {
		panic("dequeueing from an empty queue")
	}
	front := q.elements[q.read]
	q.read = q.advance(q.read)
	return front
}

func (q *Queue) advance(p int) int {
	return (p + 1) % len(q.elements)
}

func (q *Queue) growIfNecessary() {
	if !q.Full() {
		return
	}

	oldSize := len(q.elements)
	newSize := oldSize * 2
	newElements := make([]interface{}, newSize, newSize)

	if q.read < q.write {
		copy(newElements, q.elements[q.read:q.write])
	} else {
		n := copy(newElements, q.elements[q.read:])
		copy(newElements[n:], q.elements[:q.write])
	}

	q.elements = newElements
	q.read = 0
	q.write = oldSize - 1
}
