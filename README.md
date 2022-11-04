# go-graph-algs

## Base on this tutorial series

<https://youtube.com/playlist?list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P>

## Common Graph Theory Problems

1. Shortest path problem
2. Connectivity
3. Negative cycles
4. Strongly connected components
5. Traveling salesman problem
6. Bridge / cut edge
7. Articulation point / cut vertex
8. Minimum spanning tree
9. Network flow: max flow

## Common Proramatically Representation

1. Adjacency Matrix
2. Adjacency List
3. Edge List

## Notes on Breadth First Search on a Matrix

```text
Suppose we have a matrix of size MxN = 4x7
col		00 01 02 03 04 05 06
row 00: 00 01 02 03 04 05 06
row 01: 07 08 09 10 11 12 13
row 02: 14 15 16 17 18 19 20
row 03: 21 22 23 24 25 26 27
Then we can derive the cordinate (x, y) via a single cell number K like this: x = K/N and y = K-x*N

Suppose we have this matrix of size 5x7 that include obstacles
1 1 1 0 1 1 1
1 0 1 1 1 0 1
1 0 1 1 1 1 1
1 1 0 0 1 1 1
0 1 0 1 1 0 1
With starting vertex (0, 0) and ending vertex (4, 3), the shortest path is [0, 1, 2, 9, 10, 17, 18, 25, 32, 31]
Our adjacency list is composed of all K vertexes, each vertex contains a list of its adjacent vertexes, and each element of that list is
an array consists of 2 elements: Kth and W, K is its vertex number and W is its weight, in this case W is either 1 or 0 (obstacle)

To reconstruct K from any (x, y): K = x*N + y

Coordinates of 8 directions:
(x-1, y-1) | (x-1, y+0) | (x-1, y+1)
(x+0, y-1) |   (x, y)   | (x+0, y+1)
(x+1, y-1) | (x+1, y+0) | (x+1, y+1)
For only 4 directions:
kX: [-1, 0, 1, 0]
kY: [0, 1, 0, -1]
currX = x + kX
currY = y + kY
```
