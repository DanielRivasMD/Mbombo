/*
Copyright © 2024 Daniel Rivas <danielrivasmd@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/atrox/homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////
// Mbombo, also called Bumba, is the creator god in the religion
// and mythology of the Kuba people of Central Africa in the area
// that is now known as Democratic Republic of the Congo.
// In the Mbombo creation myth, Mbombo was a giant in form and
// white in color. The myth describes the creation of the universe
// from nothing.
////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// rootCmd
var rootCmd = &cobra.Command{
	Use:   "mbombo",
	Short: "",
	Long: chalk.Green.Color("Daniel Rivas <danielrivasmd@gmail.com>") + `

` + chalk.Green.Color("Mbombo") + chalk.Blue.Color(` will forge a unifed product.


`) + chalk.Green.Color("Mbombo") + ` creates a convenient command line interphase
with built-in and accessible documentation.
`,

	Example: `
` + chalk.Cyan.Color("mbombo") + ` help`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute
func Execute() {
	ε := rootCmd.Execute()
	if ε != nil {
		log.Fatal(ε)
		os.Exit(1)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// initialize config
func initializeConfig(κ *cobra.Command, configPath string, configName string) error {

	// initialize viper
	ω := viper.New()

	// collect config path & file from persistent flags
	ω.AddConfigPath(configPath)
	ω.SetConfigName(configName)

	// read config file
	ε := ω.ReadInConfig()
	if ε != nil {
		// okay if no config file
		_, ϙ := ε.(viper.ConfigFileNotFoundError)
		if !ϙ {
			// error if not parse config file
			return ε
		}
	}

	// bind flags viper
	bindFlags(κ, ω)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bind each cobra flag viper configuration
func bindFlags(κ *cobra.Command, ω *viper.Viper) {

	κ.Flags().VisitAll(func(σ *pflag.Flag) {

		if !σ.Changed && ω.IsSet(σ.Name) {
			ν := ω.Get(σ.Name)
			κ.Flags().Set(σ.Name, fmt.Sprintf("%v", ν))
		// apply viper config value flag
		}
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {

	// persistent flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// fileExist checks if a file exists and is not a directory before try using it to prevent further errors
func fileExist(ƒ string) bool {
	info, ε := os.Stat(ƒ)
	if os.IsNotExist(ε) {
		return false
	}
	return !info.IsDir()
}

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
