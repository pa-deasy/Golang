package bfs

type Person struct {
	Name        string
	MangoSeller bool
}

type Node struct {
	Contacts []Node
	Value    Person
}

const NotFound string = "No one"

func ClosestMangoSeller(n Node) string {
	var visitedNodes []Node
	queue := []Node(n.Contacts)

	for len(queue) > 0 {
		q, n := dequeue(queue)

		queue = q

		if contains(visitedNodes, n) {
			continue
		}

		if n.Value.MangoSeller {
			return n.Value.Name
		}

		for _, c := range n.Contacts {
			queue = enqueue(queue, c)
		}

		visitedNodes = append(visitedNodes, n)
	}

	return NotFound
}

func enqueue(q []Node, n Node) []Node {
	q = append(q, n)
	return q
}

func dequeue(q []Node) ([]Node, Node) {
	element := q[0]
	q = q[1:]

	return q, element
}

func contains(nodes []Node, target Node) bool {
	for _, node := range nodes {
		if node.Value == target.Value {
			return true
		}
	}

	return false
}
