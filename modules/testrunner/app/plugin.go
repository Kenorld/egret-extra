package app

import (
	"fmt"
	"bitbucket.org/kenorld/eject-core"
)

func init() {
	eject.OnAppStart(func() {
		fmt.Println("Go to /@tests to run the tests.")
	})
}
