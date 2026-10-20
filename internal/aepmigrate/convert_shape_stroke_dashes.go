package aepmigrate

import (
	"github.com/AsahiArt/aep-parser/internal/aep"
	"github.com/AsahiArt/aep-parser/internal/profile"
)

func materializeShapeStrokeDashes(stroke *aep.StrokeNode, source profile.Layer) error {
	if err := applyShapeFloatProperty(source.Properties, "ADBE Vector Stroke Dash 1", stroke.Dashes().SetDash); err != nil {
		return err
	}
	if err := applyShapeFloatProperty(source.Properties, "ADBE Vector Stroke Gap 1", stroke.Dashes().SetGap); err != nil {
		return err
	}
	return nil
}
