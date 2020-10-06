package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// EffectScatter represents an effect scatter chart.
type EffectScatter struct {
	RectChart
}

func (EffectScatter) Type() string { return types.ChartEffectScatter }

// NewEffectScatter creates a new effect scatter chart.
func NewEffectScatter() *EffectScatter {
	c := &EffectScatter{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.HasXYAxis = true
	return c
}

// SetXAxis adds the X axis.
func (c *EffectScatter) SetXAxis(x interface{}) *EffectScatter {
	c.xAxisData = x
	return c
}

// AddSeries adds the Y axis.
func (c *EffectScatter) AddSeries(name string, data []opts.EffectScatterData, options ...SeriesOpts) *EffectScatter {
	series := SingleSeries{Name: name, Type: types.ChartEffectScatter, Data: data}
	series.configureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// Validate validates the given configuration.
// TODO: add more EffectScatter validate cases
func (c *EffectScatter) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	c.Assets.Validate(c.AssetsHost)
}
