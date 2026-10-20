package aepmigrate

import "github.com/AsahiArt/aep-parser/internal/profile"

func isExplicitMatteRef(layer profile.Layer) bool {
	return layer.MatteRef != nil && layer.MatteRefKind == "explicit"
}

func supportsExplicitMatteTarget(target VersionLabel) bool {
	rule, ok := DefaultCapabilityLedger().RuleByID("layer-explicit-matte-source")
	return ok && rule.SupportForTarget(target) == CapabilityPreserved
}
