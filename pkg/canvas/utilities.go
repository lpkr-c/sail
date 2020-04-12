package canvas

import "github.com/fogleman/gg"

func H(c *gg.Context, rel float64) float64 {
	return rel * float64(c.Height())
}

func W(c *gg.Context, rel float64) float64 {
	return rel * float64(c.Width())
}
