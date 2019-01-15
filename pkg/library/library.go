package library

import (
	"errors"
	"fmt"
	"regexp"
	"sort"

	"github.com/devinmcgloin/sail/pkg/sketch"
	"github.com/devinmcgloin/sail/pkg/sketch/accrew"
	"github.com/devinmcgloin/sail/pkg/sketch/delaunay"
	"github.com/devinmcgloin/sail/pkg/sketch/primitives"
	"github.com/devinmcgloin/sail/pkg/sketch/sampling"
)

// Sketches defines all the sketches that the system can render
var options = map[string]sketch.Renderable{
	"accrew/clouds":                accrew.Cloud{},
	"delaunay/ring":                delaunay.Ring{},
	"delaunay/mesh":                delaunay.Mesh{},
	"sampling/uniform-rectangle":   sampling.UniformRectangleDot{},
	"sampling/radial-rectangle":    sampling.RadialRectangleDot{},
	"sampling/dot-walk":            sampling.DotWalk{},
	"primitive/line-coloring":      primitives.LineColoring{},
	"primitive/bars":               primitives.Bars{},
	"primitive/rotated-lines":      primitives.RotatedLines{},
	"primitive/falling-rectangles": primitives.FallingRectangles{},
}

// Lookup finds a sketch based on the sketchID
func Lookup(sketchID string) (sketch.Renderable, error) {
	sketch, ok := options[sketchID]
	if !ok {
		return nil, errors.New("invalid sketch ID")
	}
	return sketch, nil
}

// List prints all avaliable sketches
func List(regex string) {
	var sketchIDs []string

	for sketchID := range options {
		sketchIDs = append(sketchIDs, sketchID)
	}
	sort.Strings(sketchIDs)
	for _, sketchID := range sketchIDs {
		matched, err := regexp.MatchString(regex, sketchID)
		if err != nil {
			fmt.Printf("%s -> %s\n", sketchID, err)
			continue
		}

		if matched && err == nil {
			fmt.Printf("%s\n", sketchID)
		}
	}
}
