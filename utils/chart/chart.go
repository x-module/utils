/**
 * Created by PhpStorm.
 * @file   chart.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/18 12:46
 * @desc   统计组件服务
 */

package chart

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"github.com/go-xmodule/utils/utils/datetime"
	"github.com/go-xmodule/utils/utils/xlog"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/util"
	"os"
	"time"
)

type Chart struct {
	height int
	width  int
	path   string
}

func NewChart() *Chart {
	return new(Chart)
}

func (c *Chart) Init(height int, width int, savePath string) *Chart {
	c.width = width
	c.height = height
	c.height = height
	c.path = savePath
	return c
}

func (c *Chart) Generate(xValues []time.Time, yValues []float64, dirPath string) (string, error) {
	mainSeries := chart.TimeSeries{
		Style: chart.Style{
			Show:        true,
			StrokeWidth: 0.8,
			StrokeColor: chart.ColorBlue,
			FillColor:   chart.ColorBlue.WithAlpha(100),
		},
		XValues: xValues,
		YValues: yValues,
	}

	minSeries := &chart.MinSeries{
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateGreen,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}

	maxSeries := &chart.MaxSeries{
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}

	graph := chart.Chart{
		Width:  c.width,
		Height: c.height,
		Background: chart.Style{
			FillColor: chart.ColorWhite,
			Padding: chart.Box{
				Top: 10,
			},
		},
		YAxis: chart.YAxis{
			Name:      "person",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			TickStyle: chart.Style{
				TextRotationDegrees: 45.0,
			},
			ValueFormatter: func(v any) string {
				return fmt.Sprintf("%d P", int(v.(float64)))
			},
		},
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
			ValueFormatter: chart.TimeValueFormatterWithFormat(datetime.OnlyTimeTemplate),
			GridMajorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorAlternateGray,
				StrokeWidth: 1.0,
			},
			GridLines: []chart.GridLine{
				{Value: util.Time.ToFloat64(time.Date(2016, 8, 1, 9, 30, 0, 0, time.UTC))},
			},
		},
		Series: []chart.Series{
			mainSeries,
			minSeries,
			maxSeries,
			chart.LastValueAnnotation(minSeries),
			chart.LastValueAnnotation(maxSeries),
		},
	}

	if !gotool.FileUtils.Exists(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			xlog.Logger.Errorf("create temp dir err:%s. path:%s:", err.Error(), dirPath)
		}
	}
	path := fmt.Sprintf("%s/%s", dirPath, c.path)
	buffer, err := os.Create(path)
	if err != nil {
		return "", err
	}
	err = graph.Render(chart.PNG, buffer)
	if err != nil {
		return "", err
	}
	return path, nil
}
