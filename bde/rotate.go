package bde

import (
	"fmt"

	"github.com/pkg/errors"
)

//RotateKey takes the first index of type RecoveryPassword and removes it.
//It will then generate a new key
func RotateKey() error {

	err := CreateProtectorsIfMissing()
	if err != nil {
		return errors.Wrap(err, "rotate: create key protectors if missing")
	}

	//First see if there is a key to rotate, if not, create one
	// key, err := GetActiveRecoveryPassword()
	// if err != nil {
	// 	return errors.Wrap(err, "rotate: cant get active recovery password")
	// }

	// if key == "" {
	// 	fmt.Println("No Recovery Passwords, Creating One...")
	// 	_, err := CreateRecoveryPasswordProtector()
	// 	if err != nil {
	// 		return errors.Wrap(err, "rotate: creating initial key protector")
	// 	}
	// 	return nil
	// }

	//If we have a key, rotate it
	status, err := GetEncryptionStatus()
	if err != nil {
		return errors.Wrap(err, "rotate: getting encryption status")
	}

	//This is here to protect us from unknown states, but we return early above so we should never get here
	if len(status.KeyProtector) == 0 {
		fmt.Println("no existing key protectors, need to create one")
		_, err := CreateRecoveryPasswordProtector()
		if err != nil {
			return errors.Wrap(err, "rotate: creating initial key protector")
		}
		return nil
	}

	fmt.Println("Active Key Found, Rotating...")

	k, err := CreateRecoveryPasswordProtector()
	if err != nil {
		return errors.Wrap(err, "rotate: creating another key protector")
	}

	var a []string

	//KeyProtector with Type 3 is RecoveryPassword
	for _, element := range status.KeyProtector {
		if element.KeyProtectorType == 3 {
			a = append(a, element.KeyProtectorID)
		}
	}

	//For debug
	if len(a) == 0 {
		return errors.New("rotate: no key found, but we expected one")
	}

	//Delete the recovery passwords on disk that we didn't just make
	//Windows seems to respect all keys, but only tells the users the first item is active
	//TODO This can error
	for _, element := range a {
		if element != k {
			err = DeleteKeyProtector(element)
		}
		if err != nil {
			return errors.Wrap(err, "rotate: failed to delete old key")
		}
	}

	//Finally, make sure we have a TPM protector or something equivalent
	//This tests to see if all key protectors are of type 3 (recovery password)
	//If all are of type 3 we assume the TPM (or something) is not active, and needs to be
	//TODO this entire function does to much, break it up!
	// for _, element := range status.KeyProtector {
	// 	t := 3
	// 	if element.KeyProtectorType != t {
	// 		fmt.Println("TPM or equivalent is present")
	// 		return nil
	// 	}
	// }

	// fmt.Println("No TPM or equivalent found, enabling the TPM protector")
	// err = CreateTpmProtector()
	// if err != nil {
	// 	return errors.Wrap(err, "roate: enabling the TPM")
	// }

	return nil
}
