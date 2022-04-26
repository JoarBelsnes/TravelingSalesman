package salesman

import (
	"fmt"
	"log"
	"math"
)

//Shortest calculates the shortest path from src to dest
func (g *Graph) Shortest(src, dest int) (BestPath, error) {
	return g.evaluate(src, dest, true)
}

//Longest calculates the longest path from src to dest
func (g *Graph) Longest(src, dest int) (BestPath, error) {
	return g.evaluate(src, dest, false)
}

func (g *Graph) setup(shortest bool, src int, list int) {
	//-1 auto list
	//Get a new list regardless
	if list >= 0 {
		g.forceList(list)
	} else if shortest {
		g.forceList(-1)
	} else {
		g.forceList(-2)
	}
	//Reset state
	g.visitedDest = false
	//Reset the best current value (worst so it gets overwritten)
	// and set the defaults *almost* as bad
	// set all best verticies to -1 (unused)
	if shortest {
		g.setDefaults(int64(math.MaxInt64)-2, -1)
		g.best = int64(math.MaxInt64)
	} else {
		g.setDefaults(int64(math.MinInt64)+2, -1)
		g.best = int64(math.MinInt64)
	}
	//Set the distance of initial vertex 0
	g.Verticies[src].distance = 0
	//Add the source vertex to the list
	g.visiting.PushOrdered(&g.Verticies[src])
}

func (g *Graph) forceList(i int) {
	//-2 long auto
	//-1 short auto
	//0 short pq
	//1 long pq
	//2 short ll
	//3 long ll
	switch i {
	case -2:
		if len(g.Verticies) < 800 {
			g.forceList(2)
		} else {
			g.forceList(0)
		}
		break
	case -1:
		if len(g.Verticies) < 800 {
			g.forceList(3)
		} else {
			g.forceList(1)
		}
		break
	case 0:
		g.visiting = priorityQueueNewShort()
		break
	case 1:
		g.visiting = priorityQueueNewLong()
		break
	case 2:
		g.visiting = linkedListNewShort()
		break
	case 3:
		g.visiting = linkedListNewLong()
		break
	default:
		panic(i)
	}
}

func (g *Graph) bestPath(src, dest int) BestPath {
	var path []int
	for c := g.Verticies[dest]; c.ID != src; c = g.Verticies[c.bestVerticies[0]] {
		path = append(path, c.ID)
	}
	path = append(path, src)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return BestPath{g.Verticies[dest].distance, path}
}

func (g *Graph) evaluate(src, dest int, shortest bool) (BestPath, error) {
	//Setup graph
	g.setup(shortest, src, -1)
	return g.postSetupEvaluate(src, dest, shortest)
}

func (g *Graph) postSetupEvaluate(src, dest int, shortest bool) (BestPath, error) {
	var current *Vertex
	oldCurrent := -1
	for g.visiting.Len() > 0 {
		//Visit the current lowest distanced Vertex
		//TODO WTF
		current = g.visiting.PopOrdered()
		if oldCurrent == current.ID {
			continue
		}
		oldCurrent = current.ID
		//If the current distance is already worse than the best try another Vertex
		if shortest && current.distance >= g.best {
			continue
		}
		for v, dist := range current.arcs {
			//If the arc has better access, than the current best, update the Vertex being touched
			if (shortest && current.distance+dist < g.Verticies[v].distance) ||
				(!shortest && current.distance+dist > g.Verticies[v].distance) {
				if current.bestVerticies[0] == v && g.Verticies[v].ID != dest {
					//also only do this if we aren't checkout out the best distance again
					//This seems familiar 8^)
					return BestPath{}, newErrLoop(current.ID, v)
				}
				g.Verticies[v].distance = current.distance + dist
				g.Verticies[v].bestVerticies[0] = current.ID
				if v == dest {
					//If this is the destination update best, so we can stop looking at
					// useless Verticies
					g.best = current.distance + dist
					g.visitedDest = true
					continue // Do not push if dest
				}
				//Push this updated Vertex into the list to be evaluated, pushes in
				// sorted form
				g.visiting.PushOrdered(&g.Verticies[v])
			}
		}
	}
	return g.finally(src, dest)
}

func (g *Graph) finally(src, dest int) (BestPath, error) {
	if !g.visitedDest {
		return BestPath{}, ErrNoPath
	}
	return g.bestPath(src, dest), nil
}

//BestPath contains the solution of the most optimal path
type BestPath struct {
	Distance int64
	Path     []int
}

//BestPaths contains the list of best solutions
type BestPaths []BestPath
//creating graph of Romanian cities and their edges
func Salesman() {
	graph := NewGraph()
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)
	graph.AddVertex(7)
	graph.AddVertex(8)
	graph.AddVertex(9)
	graph.AddVertex(10)
	graph.AddVertex(11)
	graph.AddVertex(12)
	graph.AddVertex(13)
	graph.AddVertex(14)
	graph.AddVertex(15)
	graph.AddVertex(16)
	graph.AddVertex(17)
	graph.AddVertex(18)
	graph.AddVertex(19)

	graph.AddArc(0, 1, 71)
	graph.AddArc(0, 7, 151)
	graph.AddArc(1, 0, 71)
	graph.AddArc(1, 2, 75)
	graph.AddArc(2, 1, 75)
	graph.AddArc(2, 3, 118)
	graph.AddArc(2, 7, 140)
	graph.AddArc(3, 2, 118)
	graph.AddArc(3, 4, 111)
	graph.AddArc(4, 3, 111)
	graph.AddArc(4, 5, 70)
	graph.AddArc(5, 4, 70)
	graph.AddArc(5, 6, 75)
	graph.AddArc(6, 6, 75)
	graph.AddArc(6, 7, 120)
	graph.AddArc(7, 0, 151)
	graph.AddArc(7, 2, 140)
	graph.AddArc(7, 8, 80)
	graph.AddArc(7, 10, 99)
	graph.AddArc(8, 7, 80)
	graph.AddArc(8, 11, 97)
	graph.AddArc(8, 9, 146)
	graph.AddArc(9, 6, 120)
	graph.AddArc(9, 8, 146)
	graph.AddArc(9, 11, 138)
	graph.AddArc(10, 8, 99)
	graph.AddArc(10, 13, 211)
	graph.AddArc(11, 8, 97)
	graph.AddArc(11, 9, 138)
	graph.AddArc(11, 13, 101)
	graph.AddArc(12, 13, 90)
	graph.AddArc(13, 10, 211)
	graph.AddArc(13, 11, 101)
	graph.AddArc(13, 12, 90)
	graph.AddArc(14, 16, 87)
	graph.AddArc(15, 17, 142)
	graph.AddArc(15, 18, 98)
	graph.AddArc(16, 14, 87)
	graph.AddArc(16, 17, 92)
	graph.AddArc(17, 16, 92)
	graph.AddArc(17, 14, 142)
	graph.AddArc(18, 14, 98)
	graph.AddArc(18, 19, 86)
	graph.AddArc(19, 18, 86)
	findPath(*graph)
}
//finding shortest path from node 3 to 13
func findPath(graph Graph) {
	best, err := graph.Shortest(3, 13)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cheapest travel from Timisoara to Bucharest is costs", best.Distance, "Rumensk leu following path ", best.Path)

	// cannot find longest distance because of it would be infinite
	/*
		best, err = graph.Longest(3, 13)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Longest distance ", best.Distance, " following path ", best.Path)
	*/
}
