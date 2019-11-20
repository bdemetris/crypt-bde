package bde

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/pkg/errors"
)

// CreateRecoveryPasswordProtector generates another key of type RecoveryPassword
// Windows seems to allow infinite keys to exist, each time this runs it will create a new key
// if the disk is encrypted already.  If the disk is not encrypted it will encrypt it.
func CreateRecoveryPasswordProtector() (string, error) {
	cmd := exec.Command("C:\\Windows\\System32\\manage-bde.exe", "-protectors", "-add", "-rp", "c:")

	o, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "exec Enable-Bitlocker: recovery password")
	}

	r, _ := regexp.Compile("{[A-H0-9-]{36}}")

	id := r.FindString(string(o))

	return id, nil
}

// CreateTpmProtector enables the TPM on a device
func CreateTpmProtector() error {
	cmd := exec.Command("C:\\Windows\\System32\\manage-bde.exe", "-protectors", "-add", "-tpm", "c:")

	_, err := cmd.Output()
	if err != nil {
		return errors.Wrap(err, "exec Enable-Bitlocker: tpm protector")
	}

	return nil
}

// CreateProtectorsIfMissing is meant to handle an out of the box situation
func CreateProtectorsIfMissing() error {
	// First see if there is a key to rotate, if not, create one
	key, err := GetActiveRecoveryPassword()
	if err != nil {
		return errors.Wrap(err, "create: cant get active recovery password")
	}

	if key == "" {
		fmt.Println("No Recovery Passwords, Creating One...")
		_, err := CreateRecoveryPasswordProtector()
		if err != nil {
			return errors.Wrap(err, "create: creating initial key protector")
		}
	}

	// Get the status to see if we have a TPM or better
	status, err := GetEncryptionStatus()
	if err != nil {
		return errors.Wrap(err, "create: getting encryption status")
	}

	for _, element := range status.KeyProtector {
		t := 3
		if element.KeyProtectorType != t {
			fmt.Println("TPM or Equivalent is Present")
			return nil
		}
	}

	fmt.Println("No TPM or Equivalent Found, Enabling the TPM Protector")
	err = CreateTpmProtector()
	if err != nil {
		return errors.Wrap(err, "create: enabling the TPM")
	}

	return nil
}
