package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getStateFilepath(executablePath string) string {
	filename := fmt.Sprintf(
		"%s.json",
		strings.Join(
			strings.Split(
				executablePath,
				string(os.PathSeparator),
			)[1:],
			"_",
		),
	)
	return string(os.PathSeparator) + filepath.Join("tmp", filename)
}
