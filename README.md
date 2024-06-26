__⚠ warning__ This library is still in development. There will be changes in structure and distribution. __⚠__

# GoUtils

Common utilities for Go. The boundary to inclusion is that these utilities (primarily) work on a single type and rely only on the Go standard library. These utilities aim to be useful when a specialized solution is not (yet) necessary.

The _goutils_ module provides common operations that operate on a single type or within a single package. These are all simple sequences of logical operations that are context-independent. The module does not rely on any external dependencies. Therefore, it provides a minimal threshold to inclusion in any project. That is, these functions are not geared towards a specific functional domain.

- `assert` - assertions for quickly testing and guarding conditions.
- `std` - additional utilities for the standard library.  
  - ...
  - `errors` - provides error handling utils: basic functions, `Context(..)`, `Stacktrace(..)`.
  - `testing` - additional functions for `testing`, useful specifically for writing tests.
  - ...
- `structure` - composite data-structures that represent basic concepts.
- `codec` - encoding/decoding, various codings for representing a concept as a basic data-type, grouped by coding-type.
- `types` - (syntax) type-based constants and definitions.

The goal for these utils is _not_ to provide advanced, best-of-breed implementations or scientifically perfected algorithms. In the first place, it is to provide workable solutions, that may be fine-tuned if reasonable/feasible. That way, we can select external dependencies based on highly optimized or highly specialized function.

## Conventions

- `N` to indicate that the function takes a parameter specifying a _count_, _number_ or similar value. This is typically a variant that allows the caller to specify a value rather than the default.
- `Keyed` to indicate that the function is a variant with a _key_ (e.g. map key) parameter included.
- `Indexed` to indicate that the function is a variant with an _index_ (e.g. array/slice index) parameter included.

## Don't Repeat Yourself (DRY)

An often-heard argument is that _don't repeat yourself_ is overrated. This utility library is composed with the idea that these utils are an extension to the standard library provided by Go. It contains functions that capture common, domain-agnostic logic. The logic should serve a single purpose, or be split up in multiple functions.

Some functions may be trivial, but that is not the point. These functions exist so that they do not have to be reinvented and instead are readily available. These utils provide higher-level functions ready to use. The utilities' significance is in keeping the application's code-base small and readable.

_Don't repeat yourself_ may be explained differently for different use cases: the first is in utility functions. The second in the repeated use of literals that may be captured as constants or variables, such that a change of value need only be updated at a single location. The utilities provided here do not, in any way, reduce the DRY burden for the repeated use of literal values, e.g. originating in a specification.

## Package notes

### `std/builtin`

Utilities that improve working with the built-in primitives of the language.

### `std/builtin/set`

Utility functions for using any `map[K]struct{}`, with `K` being any `comparable` type, for set-like uses. This mechanism works for handling the data, but does not have any optimizations.

### `std/builtin/multiset`

Utility functions for using any `map[K]uint`, with `K` being any `comparable` type, for multiset- or bag-like uses. This mechanism works for handling the data, but does not have any optimizations. When the provided utility functions are used consistently, the invariant is preserved.

Use `len(multiset)` to count distinct entries. Use `multiset.Count` function to summize all entry counts.

_INVARIANT_ any entry that decreases to 0 occurrences is removed.

_NOTE_ for clarification: this package does not provide a _multiset_. This package provides the necessary functions (logic) to use a map as a "poor man's multiset" for whenever a full-blown specialized multiset implementation is not necessary.

### `std/errors`

The core notion is that any error in its root should be a fixed value, accessible as a variable. This is also in the root of the Go standard library. Many error situations require additional information. This information should be added, i.e. wrapped, as context-information. The root-cause will be the single variable that can be processed easily in code, while the context-information can be printed to provide additional information to developers/power-users.

Let the ability to apply above pattern also serve as an indicator of quality of the code: if error reporting requires a complex structure in order to communicate what is wrong, it likely means that a piece of code takes on too many concerns simultaneously.

## Implementation considerations

- Avoid using goroutines unless completely internalized, for internal implementation performance improvements. For the use of a utility, the decision to execute in a goroutine should be left to the user.

## To-Do

- TODO is there standard utils for converting bytes 0 and 1 to actual binary numbers?
- TODO what to do with something like an interval (Advent of Code 2022, Day 15)? It's not an encoding. It's a data structure with a specific purpose in the domain of ranges/intervals. Allowed or banned?

## Build/Usage

__Build-tags__

- `disable_assert` to replace assertions with empty functions to negate overhead, although overhead should be minimal regardless
- `enable_trace` to enable trace-logging

## Notes

Random notes. One day I might sort these out. :-)

- _unexported method on interface_ makes it possible to define interfaces that can only be implemented within the same package. Otherwise the requirement, that is the unexported method, cannot be satisfied.

## References

- [Go language specification]
- [Writing Go code]
- [Effective Go]
- [Go Fuzzing]
- [Go-Perfbook]
- [LICENSE]: [LGPL-3.0-only](<https://spdx.org/licenses/LGPL-3.0-only.html> "GNU Lesser General Public License v3.0 only")

[Go language specification]: <https://go.dev/ref/spec> "The Go Programming Language Specification"
[Writing Go code]: <https://go.dev/doc/code> "How to Write Go Code"
[Effective Go]: <https://go.dev/doc/effective_go> "Effective Go"
[Go Fuzzing]: <https://go.dev/security/fuzz/> "Go Fuzzing"
[Go-Perfbook]: <https://github.com/dgryski/go-perfbook> "Go-Perfbook: best-practices for writing high-performance Go code"
[LICENSE]: <LICENSE>
