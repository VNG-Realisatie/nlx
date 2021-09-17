// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package certportal

import (
	"fmt"
	"time"
)

func GenerateDemoSerialNumber() (string, error) {
	return fmt.Sprintf("%020d", time.Now().UnixNano()), nil
}
