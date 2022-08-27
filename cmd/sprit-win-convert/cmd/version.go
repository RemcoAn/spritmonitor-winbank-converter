package cmd

import (
	"fmt"
	"runtime"
	"win-sprit-converter/cmd/sprit-win-convert/build"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "prints version",
		Run:   version,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func version(_ *cobra.Command, args []string) {
	currOs := runtime.GOOS
	currArch := runtime.GOARCH
	fmt.Printf("%s %s %s/%s\n", build.Program, build.Version, currOs, currArch)
}
