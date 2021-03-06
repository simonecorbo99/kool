// +build !windows

package environment

import (
	"fmt"
	"os"
)

func initAsuser() {
	if os.Getenv("KOOL_ASUSER") == "" {
		os.Setenv("KOOL_ASUSER", fmt.Sprintf("%d", os.Getuid()))
	}
}
