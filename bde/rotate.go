package bde

//RotateKey takes the first index of type RecoveryPassword and removes it.
//It will then generate a new key
func RotateKey() error {
	//TODO This can error
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

	//fmt.Println(a)
	CreateKeyProtector()

	return nil
}
