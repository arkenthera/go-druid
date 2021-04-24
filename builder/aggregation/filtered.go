package aggregation

import (
	"github.com/grafadruid/go-druid/builder"
)

type Filtered struct {
	Base
	Aggregator builder.Aggregator
	Filter     builder.Filter
}

func NewFiltered() *Filtered {
	f := &Filtered{}
	f.SetType("filtered")
	return f
}

func (f *Filtered) SetName(name string) *Filtered {
	f.Base.SetName(name)
	return f
}
