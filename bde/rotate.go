package bde

import (
	"fmt"

	"github.com/pkg/errors"
)

//RotateKey takes the first index of type RecoveryPassword and removes it.
//It will then generate a new key
func RotateKey() error {
	//TODO This can error
	status, err := GetBitlockerStatus()
	if err != nil {
		return errors.Wrap(err, "getting bitlocker status")
	}

	if status {
		fmt.Println("biltocker is enabled, need to rotate key")
		keys, _ := GetKeyProtectors()

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
		DeleteKeyProtector(a[0])

		CreateKeyProtector()

	} else {
		fmt.Println("bitlocker is not enabled, enabling for system volume")
		CreateKeyProtector()
	}

	return nil
}
