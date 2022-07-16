# GoUtils

Common utilities for Go.

The _goutils_ module provides common operations that typically operate on a single type or within a single package. These are all simple sequences of logical operations that are context-independent and can therefore be captured in an independent module.

- `assert` - assertions for quickly testing and guarding conditions.
- `std` - additional utilities for the standard library.  
- `encoding` - various encodings, grouped by encoded type.

## Don't Repeat Yourself (DRY)

An often-heard argument is that _don't repeat yourself_ is overrated. This utility library is composed with the idea that these utils are an extension to the standard library provided by Go. It contains functions that capture common, domain-agnostic logic. The logic should serve a single purpose, or be split up in multiple functions.

Some functions may be trivial, but that is not the point. These functions exist so that they do not have to be reinvented and instead are readily available. These utils provide higher-level functions ready to use. The utilities' significance is in keeping the application's code-base small and readable.

_Don't repeat yourself_ may be explained differently for different use cases: the first is in utility functions. The second in the repeated use of literals that may be captured as constants or variables, such that a change of value need only be updated at a single location. The utilities provided here do not, in any way, reduce the DRY burden for the repeated use of literal values, e.g. originating in a specification.

## Naming patterns

A number of patterns are applied, so utilities with certain characteristics are easy to find. Assume `<original>` is the name of the original function for which the utility is written.

- `Must<original>` (prefixed with `Must`) - eliminates the need for error checking. In case an error does occur, the function will _panic_, as the developer chooses this function to declare that these error cases do not happen, or are an unsupported case. `Must-` functions strip off the `error` type from the return type.
- `Match<name>` (prefixed with `Match`) - test for a certain condition and returns boolean value indicating whether or not there is a match.
- `<original>Logged`, `<original>Panicked` - execute function `<original>` and if an error occurs, perform appropriate action according to name: `Logged` logs the error and continues as normal, `Panicked` calls `panic`.
- `<original><addition>` - execute a function `<original>` but this utility provides additional functionality `<addition>`. For example, `rand.Uint64NonZero()` provides a random `Uint64` as per the original function, but also ensures the returned value is `non-zero`, as per the `<addition>`.
- `NewName` (name is loose derivation from `<original>`) - new logic operating on existing type/using an existing function.

## Implementation considerations

- Avoid using goroutines unless completely internalized, for internal implementation performance improvements. For the use of a utility, the decision to execute in a goroutine should be left to the user.
