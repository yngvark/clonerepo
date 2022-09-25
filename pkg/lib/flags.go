// Package lib contains common functionality needed by multiple packages.
package lib

type Flags struct {
	PrintOutputDirFlag bool
	ConfigFile         string
	DryRun             bool
	Verbose            bool
}
