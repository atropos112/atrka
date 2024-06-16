// Package helpers contains all the helper functions that can be used in the application.
package helpers

import "golang.org/x/sync/errgroup"

// ErrGo is a helpful function to run multiple functions in parallel and return the first error
func ErrGo(eg *errgroup.Group, funcs []func() error) {
	for _, fn := range funcs {
		fn := fn // capture range variable
		eg.Go(func() error {
			return fn()
		})
	}
}
