package delaunay

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/fogleman/delaunay"
	"github.com/fogleman/gg"
)

// Ring defines the type that other satisfies the sketch interface
type Ring struct{}

// Dimensions returns the dimensions of the sketch
func (c Ring) Dimensions() (int, int) {
	return 1400, 1400
}

// Draw handles all the fundamental drawing logic
func (c Ring) Draw(context *gg.Context, rand *rand.Rand) {
	points := int(rand.Float64() * 600)
	hue := uint16(rand.Intn(365))
	center := gg.Point{X: 700, Y: 700}
	radius := rand.Float64() * 300
	minDistance := rand.Float64() * 500

	fmt.Printf("\tpoints: %d\n", points)
	fmt.Printf("\thue: %d\n", hue)
	fmt.Printf("\tcenter: %f\n", center)
	fmt.Printf("\tradius: %f\n", radius)

	pointLocations := make([]delaunay.Point, points)
	for i := range pointLocations {

		r := rand.Float64()*radius + minDistance
		theta := rand.NormFloat64() * radius
		xDiff := r * math.Cos(theta)
		yDiff := r * math.Sin(theta)
		pointLocations[i] = delaunay.Point{X: center.X + xDiff, Y: center.Y + yDiff}
	}

	r, g, b := clr.HSV{H: hue, S: 78, V: 70}.RGB()
	context.SetRGB(float64(r), float64(g), float64(b))

	triangulation, err := delaunay.Triangulate(pointLocations)
	if err != nil {
		log.Fatal(err)
	}
	renderDelaunay(context, triangulation, pointLocations)
}

func renderDelaunay(context *gg.Context, triangulation *delaunay.Triangulation, points []delaunay.Point) {

	ts := triangulation.Triangles
	hs := triangulation.Halfedges
	for i, h := range hs {
		if i > h {
			p := points[ts[i]]
			q := points[ts[nextHalfEdge(i)]]
			context.DrawLine(p.X, p.Y, q.X, q.Y)
		}
	}
	context.Stroke()
}

func nextHalfEdge(e int) int {
	if e%3 == 2 {
		return e - 2
	}
	return e + 1
}
