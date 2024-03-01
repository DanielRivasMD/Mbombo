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

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	out   string
	path  string
	files []string
	old   []string
	new   []string
)

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
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Where are the itmes to be forged?")
	rootCmd.MarkFlagRequired("path")
	rootCmd.PersistentFlags().StringArrayVarP(&files, "files", "f", []string{}, "These items will create...")
	rootCmd.MarkFlagRequired("files")
	rootCmd.PersistentFlags().StringVarP(&out, "out", "", "", "Where will the forge be delivered?")
	rootCmd.MarkFlagRequired("out")
	rootCmd.PersistentFlags().StringArrayVarP(&old, "old", "o", []string{}, "Value to replace")
	rootCmd.PersistentFlags().StringArrayVarP(&new, "new", "n", []string{}, "Value replacement")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
