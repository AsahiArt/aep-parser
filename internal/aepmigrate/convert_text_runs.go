package aepmigrate

import (
	"github.com/AsahiArt/aep-parser/internal/aep"
	"github.com/AsahiArt/aep-parser/internal/profile"
)

func materializeTextRuns(layer *aep.Layer, runs []profile.TextStyleRun) error {
	for i, run := range runs {
		if err := materializeTextRun(layer, i, run); err != nil {
			return err
		}
	}
	return nil
}

func materializeTextRun(layer *aep.Layer, index int, run profile.TextStyleRun) error {
	if err := materializeTextRunMetrics(layer, index, run); err != nil {
		return err
	}
	if err := materializeTextRunOptions(layer, index, run); err != nil {
		return err
	}
	return materializeTextRunStroke(layer, index, run)
}
