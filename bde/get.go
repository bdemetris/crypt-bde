package bde

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"
)

// GetActiveKeyProtector returns the primary activation key (not the TPM)
func GetActiveKeyProtector() (string, error) {
	status, err := GetEncryptionStatus()
	if err != nil {
		return "", errors.Wrap(err, "get active: getting encryption status")
	}

	keys := status.KeyProtector

	var kp []string

	for _, key := range keys {
		if key.KeyProtectorType == 3 {
			kp = append(kp, key.RecoveryPassword)
		}
	}

	ak := kp[0]

	return ak, nil
}

// GetEncryptionStatus does some checks to see whats going on with the disk
func GetEncryptionStatus() (EncryptionStatus, error) {
	cmd := exec.Command("powershell", "Get-BitlockerVolume", "-MountPoint", "$env:SystemDrive", "|", "ConvertTo-Json")

	o, err := cmd.Output()
	if err != nil {
		return EncryptionStatus{}, errors.Wrap(err, "get status: exec Get-BitlockerVolume")
	}

	var es EncryptionStatus

	if err := json.Unmarshal(o, &es); err != nil {
		return EncryptionStatus{}, errors.Wrap(err, "failed unmarshalling Encryption Status")
	}

	return es, nil
}

// KeyProtectors represents each item that can unlock the disk
type KeyProtectors struct {
	KeyProtectorID   string `json:"KeyProtectorId"`
	KeyProtectorType int    `json:"KeyProtectorType"`
	RecoveryPassword string `json:"RecoveryPassword"`
}

// EncryptionStatus returns the disk encryption status
type EncryptionStatus struct {
	MountPoint       string `json:"MountPoint"`
	EncryptionMethod int    `json:"EncryptionMethod"`
	KeyProtector     []KeyProtectors
}
