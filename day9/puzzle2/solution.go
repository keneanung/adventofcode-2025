package puzzle2

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Rectangle struct {
	Top int
	Bottom int
	Left int
	Right int
}

func (r Rectangle) Area() int {
	width := int(r.Right-r.Left) + 1
	height := int(r.Bottom-r.Top) + 1
	return width * height
}

func Solve(input []string) (int64, error) {
	points := make([]Point, 0, len(input))
	for _, line := range input {
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return 0, err
		}
		points = append(points, Point{X: x, Y: y})
	}
	rectangles := make([]Rectangle, 0)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			topLeft := points[i]
			bottomRight := points[j]
			if topLeft.X == bottomRight.X || topLeft.Y == bottomRight.Y {
				continue
			}
			minX := int(math.Min(float64(topLeft.X), float64(bottomRight.X)))
			maxX := int(math.Max(float64(topLeft.X), float64(bottomRight.X)))
			minY := int(math.Min(float64(topLeft.Y), float64(bottomRight.Y)))
			maxY := int(math.Max(float64(topLeft.Y), float64(bottomRight.Y)))
			
			rectangles = append(rectangles, Rectangle{
				Top:    minY,
				Bottom: maxY,
				Left:   minX,
				Right:  maxX,
			})
		}
	}
	slices.SortFunc(rectangles, func(r1 Rectangle, r2 Rectangle) int {
		return -1 * (r1.Area() - r2.Area())
	})
	return int64(rectangles[0].Area()), nil
}
