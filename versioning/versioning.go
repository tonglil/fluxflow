package versioning

// Deliberately uninitialized so it can be set during the build process.
var v string

// Set sets the version.
func Set(version string) {
	v = version
}

// Get returns the version.
func Get() string {
	if v != "" {
		return v
	}

	return "unknown"
}
