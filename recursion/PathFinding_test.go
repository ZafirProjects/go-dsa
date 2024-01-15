package recursion

import (
	"testing"
  "slices"
)

func TestPathFinding(t *testing.T)  {
  tests := []struct{
    maze []string
    wall string
    start, end Point
    solution []Point
  }{
    {
    	maze:     []string{
        "XXXXXXXXXX X",
        "X        X x",
        "x        x x",
        "x xxxxxxxx x",
        "x          x",
        "x xxxxxxxxxx",
      },
    	wall:     "X",
    	start:    Point{10, 0},
    	end:      Point{1, 5},
    	solution: []Point{
        Point{10, 0},
        Point{10, 1},
        Point{10, 2},
        Point{10, 3},
        Point{10, 4},
        Point{9, 4},
        Point{8, 4},
        Point{7, 4},
        Point{6, 4},
        Point{5, 4},
        Point{4, 4},
        Point{3, 4},
        Point{2, 4},
        Point{1, 4},
        Point{1, 5},
      },
    },
  }
  for i, test := range tests {
    if got := maze_solver(test.maze, test.wall, test.start, test.end); slices.Equal(got, test.solution){
      t.Fatalf("Failed maze %d, expected %d, got %d",i, test.solution, got)
    }
  }
}

