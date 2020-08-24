package auth

import "errors"

func Check(username string, password string) error {
	if username != "antares" || password != "vision" {
		return errors.New("username and/or password not valid")
	}

	return nil
}
