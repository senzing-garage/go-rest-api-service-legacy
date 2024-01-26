/*
 */
package cmd

import (
	"context"
	"os"

	"github.com/senzing-garage/go-cmdhelping/cmdhelper"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-cmdhelping/option/optiontype"
	"github.com/senzing-garage/go-rest-api-service-legacy/examplepackage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Short string = "go-rest-api-service-legacy short description"
	Use   string = "go-rest-api-service-legacy"
	Long  string = `
go-rest-api-service-legacy long description.
    `
)

// ----------------------------------------------------------------------------
// Context variables
// ----------------------------------------------------------------------------

var SomethingToSay = option.ContextVariable{
	Arg:     "something-to-say",
	Default: option.OsLookupEnvString("SENZING_TOOLS_SOMETHING_TO_SAY", "Main says 'Hi!'"),
	Envar:   "SENZING_TOOLS_SOMETHING_TO_SAY",
	Help:    "Just a phrase to say [%s]",
	Type:    optiontype.String,
}

var ContextVariablesForMultiPlatform = []option.ContextVariable{
	option.Configuration,
	option.EngineConfigurationJson,
	option.LogLevel,
	SomethingToSay,
}

var ContextVariables = append(ContextVariablesForMultiPlatform, ContextVariablesForOsArch...)

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define command line parameters.
func init() {
	cmdhelper.Init(RootCmd, ContextVariables)
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Used in construction of cobra.Command
func PreRun(cobraCommand *cobra.Command, args []string) {
	cmdhelper.PreRun(cobraCommand, args, Use, ContextVariables)
}

// Used in construction of cobra.Command
func RunE(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	examplePackage := &examplepackage.ExamplePackageImpl{
		Something: viper.GetString(SomethingToSay.Arg),
	}
	return examplePackage.SaySomething(ctx)
}

// Used in construction of cobra.Command
func Version() string {
	return cmdhelper.Version(githubVersion, githubIteration)
}

// ----------------------------------------------------------------------------
// Command
// ----------------------------------------------------------------------------

// RootCmd represents the command.
var RootCmd = &cobra.Command{
	Use:     Use,
	Short:   Short,
	Long:    Long,
	PreRun:  PreRun,
	RunE:    RunE,
	Version: Version(),
}
