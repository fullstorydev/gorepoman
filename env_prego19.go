// +build !go1.9

package gorepoman

func combineEnv(orig, override []string) []string {
	// Prepend our environment variables so that they overwrite any that are already there
	return append(override, orig...)
	// This does not work as of Go 1.9 due to this change:
	// https://github.com/golang/go/commit/e73f4894949c4ced611881329ff8f37805152585
}