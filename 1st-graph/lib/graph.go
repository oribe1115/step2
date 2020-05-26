package lib

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type XYAxis struct {
	X float64
	Y float64
}

func PlotGraph(data []XYAxis) error {
	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = "Relationship between Run Times and Size of Matrix"
	p.X.Label.Text = "Size of Matrix"
	p.Y.Label.Text = "Run Times (nanoseconds)"

	pts := make(plotter.XYs, len(data))
	for i, axis := range data {
		pts[i].X = axis.X
		pts[i].Y = axis.Y
	}

	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
		return err
	}

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "runTimes.png"); err != nil {
		panic(err)
	}

	return nil
}

func SetData(runTimes map[int][]int, start, end int) []XYAxis {
	data := make([]XYAxis, 0)

	for i := start; i <= end; i++ {
		values := runTimes[i]
		var average float64
		average = 0
		for _, v := range values {
			average += float64(v)
		}
		average /= float64(len(values))

		d := XYAxis{
			X: float64(i),
			Y: average,
		}
		data = append(data, d)
	}

	return data
}
