package draw

import (
	"fmt"
	"github.com/fogleman/gg"
)

const (
	width  = 800
	height = 400
)

type Bucket struct {
	start int
	end   int
	color int
}

func DrawBars(dc *gg.Context, values []int, highlightIndex int, name string, frame int, secondHighlightIndex int) {
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	barWidth := width / len(values)

	for i, value := range values {
		barHeight := float64(value) / float64(height) * float64(height)
		x := float64(i) * float64(barWidth)
		y := float64(height) - barHeight

		if i == highlightIndex {
			dc.SetRGB(1, 0, 0)
		} else if i == secondHighlightIndex {
			dc.SetRGB(0, 0, 1)
		} else {
			dc.SetRGB(0, 1, 0)
		}

		dc.DrawRectangle(x, y, float64(barWidth), barHeight)
		dc.Fill()
		dc.Stroke()
	}

	dc.SavePNG(fmt.Sprintf("%s/frame_%06d.png", name, frame))
}
