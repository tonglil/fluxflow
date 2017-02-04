package submenu

type subMenu string

const (
	// Darkroom preserves night vision.
	Darkroom subMenu = "Darkroom"
	// Movie sets movie mode for 2.5 hours.
	Movie subMenu = "Movie mode"

	// Hour disables Flux for 1 hour.
	Hour subMenu = "for an hour"
	// Sunrise disables Flux until the sunrise.
	Sunrise subMenu = "until sunrise"
)
