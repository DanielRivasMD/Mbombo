////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"log"
	"os"

	"github.com/atrox/homedir"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// find home directory
func findHome() string {
	Λ, ε := homedir.Dir()
	if ε != nil {
		log.Fatal(ε)
		os.Exit(1)
	}
	return Λ
}

////////////////////////////////////////////////////////////////////////////////////////////////////