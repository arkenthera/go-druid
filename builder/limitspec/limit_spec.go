package limitspec

import (
	"encoding/json"
	"errors"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Type builder.ComponentType
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Type = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Type
}

func Load(data []byte) (builder.LimitSpec, error) {
	var l builder.LimitSpec
	if string(data) == "null" {
		return l, nil
	}
	var t struct {
		Type builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Type {
	case "default":
		l = NewDefault()
	default:
		return nil, errors.New("unsupported limitspec type")
	}
	return l, json.Unmarshal(data, &l)
}
