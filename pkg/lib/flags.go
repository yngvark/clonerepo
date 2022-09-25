// Package lib contains common functionality needed by multiple packages.
package lib

type Flags struct {
	CdToOutputDir bool
	ConfigFile    string
	DryRun        bool
	Verbose       bool
}
