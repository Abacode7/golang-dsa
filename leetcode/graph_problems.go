package leetcode

/*
*
QUESTION: 997. Find the Town Judge
TAG: Easy

Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1

1: 3
2: 3
3: 1

- Build the graph adjacency list
- Find the vertex which has no edges
  - If more than one return -1

- Check if the vertex exists in the edges of all other vertices

*
*/
func FindJudge(n int, trust [][]int) int {
	graph := buildAdjacencyList(n, trust)

	potentialJudge := 0
	numOfPotentialJudges := 0

	// Count potential judges
	for key, values := range graph {
		if len(values) == 0 {
			potentialJudge = key
			numOfPotentialJudges++
			if numOfPotentialJudges > 1 {
				return -1
			}
		}
	}

	if numOfPotentialJudges == 0 {
		return -1
	}

	// Check if potential judge is trusted by all villagers
	for key, values := range graph {
		if key == potentialJudge {
			continue
		}

		if !values[potentialJudge] {
			return -1
		}
	}

	return potentialJudge
}

func buildAdjacencyList(n int, trusts [][]int) map[int]map[int]bool {
	graph := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		graph[i] = make(map[int]bool)
	}

	for _, trust := range trusts {
		truster, trusted := trust[0], trust[1]
		graph[truster][trusted] = true
	}
	return graph
}



/*
QUESTION: 1791. Find Center of Star Graph
TAG: Easy

Input: edges = [[1,2],[5,1],[1,3],[1,4]]
Output: 1

1: 2 5 3 4
2: 1
5: 1
3: 1
4: 1
*/
func FindCenter(edges [][]int) int {
	graph := buildAdjList(edges)

	for key, values := range graph {
		if len(values) == len(graph)-1 {
			return key
		}
	}
	return 0
}

func buildAdjList(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		vertexA, vertexB := edge[0], edge[1]
		graph[vertexA] = append(graph[vertexA], vertexB)
		graph[vertexB] = append(graph[vertexB], vertexA)
	}
	return graph
}




/*
QUESTION: 1971. Find if Path Exists in Graph
TAGS: Easy, BFS

Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
Output: false

0: 1 2
1: 0
2: 0
3: 5 4
4: 5 3
5: 3 4

Using BFS
queue [0]
- Mark source node as visited, add to queue to explore adjacents: [0]
- Pop from element the queue: 0
  - Iterate all adjacent element of popped element: 1, 2
  	- NB: Skip already visited elements: 0
	- Mark as adjacents as visited: 1, 2
  	- If destinatination is found, return true
  	- Else add adjacent element back to queue, for its adjacent to be explored: [1, 2]
- Repeat second step
*/
func ValidPathBFS(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	graph := buildAdjacencyLists(edges)
	visited := make(map[int]bool)
	queue := make([]int, 0)

	visited[source] = true
	queue = append(queue, source)

	// Explore breadth of visited nodes
	for len(queue) != 0 {
		poppedElement := queue[0]
		queue = queue[1:]

		for _, adjacent := range graph[poppedElement] {
			if adjacent == destination {
				return true
			}

			if !visited[adjacent] {
				visited[adjacent] = true
				queue = append(queue, adjacent)
			}
		}
	}
	return false
}


func ValidPathDFS(n int, edges [][]int, source int, destination int) bool {
    if len(edges) == 0 {
        return source == destination
    }
    graph := buildAdjacencyLists(edges)
    visited := make(map[int]bool)
    return dfs(graph, source, destination, visited)
}

func dfs(graph map[int][]int, source int, destination int, visited map[int]bool) bool {
    if source == destination {return true}
    visited[source] = true

    for _, adjacent := range graph[source] {
        if !visited[adjacent] {
            if dfs(graph, adjacent, destination, visited) {return true};
        }
    }
    return false;
}

func DfsIterative(graph map[int][]int, source int, destination int, visited map[int]bool) bool {
    if source == destination {return true}
    visited[source] = true

    for _, adjacent := range graph[source] {
        if !visited[adjacent] {
            if dfs(graph, adjacent, destination, visited) {
                return true
            }
        }
    }
    return false
}

func buildAdjacencyLists(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		vertexA, vertexB := edge[0], edge[1]
		graph[vertexA] = append(graph[vertexA], vertexB)
		graph[vertexB] = append(graph[vertexB], vertexA)
	}
	return graph
}
