package main

import "errors"

// Reset the router data connection.
func Reset(router Router) error {
	token, err := Token(router)
	if err != nil {
		return err
	}
	ok, err := Login(router, token)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("could not log in")
	}
	ok, err = SwitchData(router, token, false)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("could not switch data off")
	}
	ok, err = SwitchData(router, token, true)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("could not switch data on")
	}
	return nil
}
