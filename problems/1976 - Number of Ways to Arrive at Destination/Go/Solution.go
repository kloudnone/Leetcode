// You are in a city that consists of n intersections numbered from 0 to n - 1 with bi-directional roads between some intersections. The inputs are generated such that you can reach any intersection from any other intersection and that there is at most one road between any two intersections.

// You are given an integer n and a 2D integer array roads where roads[i] = [ui, vi, timei] means that there is a road between intersections ui and vi that takes timei minutes to travel. You want to know in how many ways you can travel from intersection 0 to intersection n - 1 in the shortest amount of time.

// Return the number of ways you can arrive at your destination in the shortest amount of time. Since the answer may be large, return it modulo 109 + 7.

package main

import (
	"container/heap"
)

type pair struct {
	node int
	dist int64
}

type minHeap []pair

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func countPaths(n int, roads [][]int) int {
	// Build adjacency list
	adj := make([][]pair, n)
	for _, road := range roads {
		u, v, time := road[0], road[1], road[2]
		adj[u] = append(adj[u], pair{v, int64(time)})
		adj[v] = append(adj[v], pair{u, int64(time)})
	}

	// Initialize distances and ways arrays
	dist := make([]int64, n)
	ways := make([]int64, n)
	for i := range dist {
		dist[i] = 1e18
	}

	// Start from node 0
	dist[0] = 0
	ways[0] = 1

	// Use Dijkstra's algorithm with min-heap
	pq := &minHeap{}
	heap.Init(pq)
	heap.Push(pq, pair{0, 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(pair)
		if curr.dist > dist[curr.node] {
			continue
		}

		// Explore neighbors
		for _, next := range adj[curr.node] {
			if dist[curr.node]+next.dist < dist[next.node] {
				dist[next.node] = dist[curr.node] + next.dist
				ways[next.node] = ways[curr.node]
				heap.Push(pq, pair{next.node, dist[next.node]})
			} else if dist[curr.node]+next.dist == dist[next.node] {
				ways[next.node] = (ways[next.node] + ways[curr.node]) % 1000000007
			}
		}
	}

	return int(ways[n-1])
}
