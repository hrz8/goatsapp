package config

import (
	"fmt"
)

var (
	VersionNumber = [3]uint32{0, 1, 1}
	Version       = fmt.Sprintf("v%d.%d.%d", VersionNumber[0], VersionNumber[1], VersionNumber[2])
)
