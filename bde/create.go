package bde

import (
	"os/exec"

	"github.com/pkg/errors"
)

//CreateKeyProtector generates another key of type RecoveryPassword
//Windows seems to allow infinite keys to exist, each time this runs it will create a new key
//If the disk is encrypted already.  If the disk is not encrypted it will encrypt it.
func CreateKeyProtector() error {
	cmd := exec.Command("powershell", "Enable-Bitlocker", "-MountPoint", "$env:SystemDrive", "-EncryptionMethod", "Aes256", "-RecoveryPasswordProtector")

	// cmd.Stderr = os.Stderr
	_, err := cmd.Output()
	if err != nil {
		return errors.Wrap(err, "exec Enable-Bitlocker")
	}

	return nil
}
