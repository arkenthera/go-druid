package postaggregation

type ThetaSketch struct {
	Base
	Field FieldAccess `json:"field,omitempty"`
}


func NewThetaSketch() *ThetaSketch  {
	f := &ThetaSketch{}
	f.SetType("thetaSketchEstimate")
	return f
}

func (f *ThetaSketch) SetName(name string) *ThetaSketch {
	f.Base.SetName(name)
	return f
}

func (f *ThetaSketch) SetField(fa *FieldAccess) *ThetaSketch {
	f.Field = *fa 
	return f
}