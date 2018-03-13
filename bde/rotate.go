package bde

import (
	"fmt"

	"github.com/pkg/errors"
)

//RotateKey takes the first index of type RecoveryPassword and removes it.
//It will then generate a new key
func RotateKey() error {

	status, err := GetEncryptionStatus()
	if err != nil {
		return errors.Wrap(err, "rotate: getting encryption status")
	}

	if len(status.KeyProtector) == 0 {
		fmt.Println("no existing key protectors, need to create one")
		err := CreateKeyProtector()
		if err != nil {
			return errors.Wrap(err, "rotate: creating initial key protector")
		}
		return nil
	}

	keys, err := GetKeyProtectors()
	if err != nil {
		return errors.Wrap(err, "rotate: getting active key protector")
	}

	fmt.Println("active keys found, rotating")
	var a []string

	//KeyProtector with Type 3 is RecoveryPassword
	for _, element := range keys {
		if element.KeyProtectorType == 3 {
			a = append(a, element.KeyProtectorID)
		}
	}

	//Delete the first element in the ID slice
	//Windows seems to respect all keys, but only tells the users the first item is active
	//TODO This can error
	err = DeleteKeyProtector(a[0])
	if err != nil {
		return errors.Wrap(err, "rotate: failed to delete old key")
	}

	err = CreateKeyProtector()
	if err != nil {
		return errors.Wrap(err, "rotate: creating another key protector")
	}

	return nil
}
