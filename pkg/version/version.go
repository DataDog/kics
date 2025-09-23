package version

import "runtime/debug"

// modulePath must match your go.mod module path exactly.
const modulePath = "github.com/you/yourlib" // <- change me

const fallback = "dev"

// Override lets you force a version at build time (optional).
//
//	go build -ldflags "-X github.com/you/yourlib/version.Override=v1.2.3"
var Override = ""

// String returns the best-known version of this module.
//
// Priority:
//  1. Override (via -ldflags)
//  2. Module version from build info (when used as a dep)
//  3. VCS info (devel-<commit>[+dirty])
//  4. "dev"
func String() string {
	if Override != "" {
		return Override
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return fallback
	}

	// If the library itself is the main module (e.g. `go install your/module@vX.Y.Z`)
	if bi.Main.Path == modulePath && bi.Main.Version != "" {
		return bi.Main.Version
	}

	// Otherwise, find this module among dependencies.
	for _, dep := range bi.Deps {
		if dep.Path == modulePath {
			// If the consumer has a replace with a version, prefer it.
			if dep.Replace != nil && dep.Replace.Version != "" {
				return dep.Replace.Version
			}
			if dep.Version != "" {
				return dep.Version // usually vX.Y.Z or a pseudo-version
			}
		}
	}

	// Fall back to VCS metadata (Go 1.18+; -buildvcs is on in module mode).
	var rev, dirty string
	for _, s := range bi.Settings {
		switch s.Key {
		case "vcs.revision":
			if len(s.Value) >= 12 {
				rev = s.Value[:12]
			} else {
				rev = s.Value
			}
		case "vcs.modified":
			if s.Value == "true" {
				dirty = "+dirty"
			}
		}
	}
	if rev != "" {
		return "devel-" + rev + dirty
	}

	return fallback
}
