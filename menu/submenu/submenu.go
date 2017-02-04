package submenu

// Item represents the nested menu item.
type Item string

const (
	// Darkroom preserves night vision.
	Darkroom Item = "Darkroom"
	// Movie sets movie mode for 2.5 hours.
	Movie Item = "Movie mode"

	// Hour disables Flux for 1 hour.
	Hour Item = "for an hour"
	// Sunrise disables Flux until the sunrise.
	Sunrise Item = "until sunrise"
)
