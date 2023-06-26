package main


/*
Queues
A queue consists of elements to be processed in a particular order or based on priority. A priotory-based queue of orders is shown in the following code, structured as a heap. Operations such as enqueue, dequeue, and peak can be performed on queue. Elements are added to the end and are removed from the start of the collection.
*/

type Queue []*Order

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// New method
func (order *Order) New(priority int, quantity int, product string, customerName string) {
	order = &Order{
		priority:     priority,
		quantity:     quantity,
		product:      product,
		customerName: customerName,
	}
}

// Add method adds the order to the queue
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false
		for i, addedOrder := range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

// main method

