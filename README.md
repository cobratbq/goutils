# GoUtils

Common utilities for Go.

The GoUtils module provides common operations that typically operate on a single type or within a single package. These are all simple sequences of logical operations that are context-independent and can therefore be captured in an independent module.

Utilities are packaged per representative package, for example `std/crypto/rand` for utilities operating on the `crypto/rand` package in the Go standard library.

## Findability patterns

- The utility is located in a package with the same name as the type on which it operates. For example, utilities for working with errors and `panic`ing can be found in `std/builtin`.

## Naming patterns

A number of patterns are applied, so utilities with certain characteristics are easy to find. Assume `<original>` is the name of the original function for which the utility is written.

- `Must<original>` (prefixed with `Must`) - eliminates the need for error checking. In case an error does occur, the function will `panic`, as the developer chooses this function to declare that these error cases do not happen. `Must-` functions basically strip off the `error` type from the return type.
- `Require<requirement>` (prefixed with `Require`) - functions to check that (obvious) requirement conditions hold for the provided instance. Trivial cases are `Require(bool)` checking if a bool is true, `RequireSuccess(error, string)` checking if `error == nil`. If these cases do not hold, the function will `panic`. These functions provide guards to ensure minimum input requirements are met, such that subsequent logic can safely make the necessary assumptions.
- `<original>Logged`, `<original>Panicked` - execute function `<original>` and if an error occurs, perform appropriate action according to name: `Logged` logs the error and continues as normal, `Panicked` calls `panic`.
- `<original><addition>` - execute a function `<original>` but this utility provides additional functionality `<addition>`. For example, `rand.Uint64NonZero()` provides a random `Uint64` as per the original function, but also ensures the returned value is `non-zero`, as per the `<addition>`.
- `NewName` (name is loose derivation from `<original>`) - new logic operating on existing type/using an existing function.

## Implementation considerations

- Avoid using goroutines unless completely internalized, for internal implementation performance improvements. For the use of a utility, the decision to execute in a goroutine should be left to the user.
