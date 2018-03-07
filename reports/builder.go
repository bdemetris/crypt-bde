package reports

import (
	"fmt"

	"github.com/bdemetris/crypt-bde/bde"
	"github.com/bdemetris/crypt-bde/config"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

//BuildCheckin builds the checkin object
func BuildCheckin(conf *config.Config) (*Checkin, error) {

	win32Bios, err := GetWin32Bios()
	if err != nil {
		return nil, errors.Wrap(err, "get win32Bios")
	}

	win32ComputerSystem, err := GetWin32ComputerSystem()
	if err != nil {
		return nil, errors.Wrap(err, "get win32ComputerSystem")
	}

	test := bde.CreateKeyProtector()
	fmt.Println(test)

	// Sending this in place of the recovery key for now
	u1 := uuid.Must(uuid.NewV4(), err).String()

	checkin := &Checkin{
		Serial:      win32Bios.SerialNumber,
		RecoveryKey: u1,
		UserName:    win32ComputerSystem.UserName,
		MacName:     win32Bios.PSComputerName,
	}

	fmt.Printf("%+v\n", checkin)

	return checkin, nil
}

// Checkin is what Crypt-Server expects us to POST
type Checkin struct {
	Serial      string
	RecoveryKey string
	UserName    string
	MacName     string
}
