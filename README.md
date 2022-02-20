# go-quick-sort
Implement quick sort in Go

## Concept
- a DFS, left-side-first implementation
- pre-order traversal: the main sorting logic happens before going down the next layer
- time complexity: 
  - best: <img src="https://render.githubusercontent.com/render/math?math=O(n\log n)">
  - worst: <img src="https://render.githubusercontent.com/render/math?math=O(n^{2})">

## Usage
```
go run main.go
```
Should print out this sorted slice on Terminal:
```
[1 2 3 4 5 6 8 9]
```
