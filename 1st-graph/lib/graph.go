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

func PlotTest() error {
	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = "title"
	p.X.Label.Text = "size of matrix"
	p.Y.Label.Text = "run times(nanoseconds)"

	nums := []XYAxis{
		{3, 384},
		{4, 436},
		{5, 605},
	}

	pts := make(plotter.XYs, len(nums))
	for i, axis := range nums {
		pts[i].X = axis.X
		pts[i].Y = axis.Y
	}

	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
		return err
	}

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "sample1.png"); err != nil {
		panic(err)
	}

	return nil
}
