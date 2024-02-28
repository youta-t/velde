package env

import (
	"os"
	"path/filepath"
)

func VeldePath() (string, error) {
	dh := os.Getenv("VELDE_PATH")
	if dh != "" {
		return dh, nil
	}

	userhome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(userhome, ".velde"), nil
}
