package core

// GetShortestPath returns the shortest path (of stations) and cost
func GetShortestPath(from, to *Location, mode Mode) ([]*Station, int) {
	destinations := make(map[*Station]bool)
	for _, s := range to.Stations {
		destinations[s] = true
	}
	heap := NewNodeHeap()
	visited := make(map[*Station]bool)
	//trace store the parent of each node
	trace := make(map[*Station]*Station)
	for _, s := range from.Stations {
		node := &Node{Station: s, Cost: 0}
		heap.Push(node)
		visited[s] = true
	}

	node := BreadthFirstSearch(mode, heap, destinations, visited, trace)
	if node == nil {
		return nil, NotReachable
	}
	// build reverse path from trace map
	last := node.Station
	path := []*Station{last}
	for trace[last] != nil {
		path = append(path, trace[last])
		last = trace[last]
	}
	// reverse
	for i, j := 0, len(path)-1; i < j; {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}
	return path, node.Cost
}

// BreadthFirstSearch finds the shortest path to arrive location
// returns the reached Station (any Station in destination Location)
// if not reachable return nil
func BreadthFirstSearch(mode Mode, heap *NodeHeap, destinations map[*Station]bool, visited map[*Station]bool, trace map[*Station]*Station) *Node {
	for heap.Len() > 0 {
		nearestNode := heap.Pop()
		nearestStation := nearestNode.Station
		if _, contains := destinations[nearestStation]; contains {
			return nearestNode
		}
		for _, neighborStation := range nearestStation.Neighbors {
			if visited[neighborStation] {
				continue
			}
			moveCost := getMoveCost(mode, nearestStation, neighborStation)
			// cannot reach to this neighbor
			if moveCost < 0 {
				continue
			}
			neighborNode := &Node{Station: neighborStation, Cost: nearestNode.Cost + moveCost}
			trace[neighborStation] = nearestStation
			visited[neighborStation] = true
			heap.Push(neighborNode)
		}
	}
	return nil
}

func getMoveCost(mode Mode, station, neighbor *Station) int {
	moveCost := NotReachable
	if station.Line == neighbor.Line {
		moveCost = trafficCost.Train[neighbor.Line][mode]
	} else {
		moveCost = trafficCost.Wait[mode]
	}
	return moveCost
}
