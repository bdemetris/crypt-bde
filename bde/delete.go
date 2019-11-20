package bde

import (
	"os/exec"

	"github.com/pkg/errors"
)

// DeleteKeyProtector Removes the Protector ID Passed in as type Int
func DeleteKeyProtector(id string) error {
	cmd := exec.Command("powershell", "Remove-BitlockerKeyProtector", "-MountPoint", "$env:SystemDrive", "-KeyProtectorId", (`"` + id + `"`))

	// cmd.Stderr = os.Stderr
	_, err := cmd.Output()
	if err != nil {
		return errors.Wrap(err, "exec Remove-BitlockerKeyProtector")
	}

	return nil
}
