package utils

import "fmt"

func AppendError(existErr error, newErr error) error {
	if existErr == nil {
		return newErr
	}
	return fmt.Errorf("%v\n%v", existErr, newErr)
}
