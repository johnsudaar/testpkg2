package testpkg2

import "github.com/johnsudaar/testpkg1"

func Version() string {
	return testpkg1.Version() + " - 1.2.0"
}
