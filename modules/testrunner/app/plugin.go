package app

import (
	"fmt"
	"github.com/eject/eject"
)

func init() {
	eject.OnAppStart(func() {
		fmt.Println("Go to /@tests to run the tests.")
	})
}
