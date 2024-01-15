package recursion

type Point struct{
  x, y int
}

var dir = [4][2]int{
  {-1, 0},
  {1, 0},
  {0, -1},
  {0, 1},
}

func walk(maze []string, wall string, start Point, end Point, seen [][]bool, path []Point) bool {
  curr := start
  if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
    return false
  }
  if string(maze[curr.y][curr.x]) == wall {
    return false
  }
  if curr.x == end.x && curr.y == end.y {
    path = append(path, end)
    return true
  }
  if seen[curr.y][curr.x] {
    return false
  }

  seen[curr.y][curr.x] = true
  path = append(path, curr)

  for i := 0; i < len(dir); i++ {
    x, y := dir[i][0], dir[i][1]
    if walk(maze, wall, Point{curr.x + x, curr.y + y}, end, seen, path) {
      return true;
    }
  }

  path = path[:len(path)-1]
  return false
}

func maze_solver(maze []string, wall string, start, end Point) []Point { 
  seen := make([][]bool, len(maze))
  for i := range seen {
    seen[i] = make([]bool, len(maze[0]))
  }
  path := []Point{}

  walk(maze, wall, start, end, seen, path)

  return path
}
