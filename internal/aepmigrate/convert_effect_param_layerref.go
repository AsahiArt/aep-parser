package aepmigrate

import "github.com/AsahiArt/aep-parser/internal/profile"

func effectParamLayerRefIsSelf(ref *profile.LayerRef, layer profile.Layer) bool {
	if ref == nil {
		return false
	}
	if ref.Name != "" && ref.Name == layer.Name {
		return true
	}
	if ref.ID != 0 && ref.ID == layer.ID {
		return true
	}
	if ref.Index != 0 && ref.Index == layer.Index {
		return true
	}
	return false
}
