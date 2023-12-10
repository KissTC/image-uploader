package main

import "errors"

func main() {

}

func connect() error {
	return errors.New("connection failed")
}

func createUser() error {
	err := connect()
	if err != nil {
		return err
	}

	return nil
}
