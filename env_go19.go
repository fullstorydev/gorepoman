// +build go1.9

package gorepoman

func combineEnv(orig, override []string) []string {
	// Append our environment variables so that they overwrite any that are already there
	return append(orig, override...)
}