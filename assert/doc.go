// SPDX-License-Identifier: LGPL-3.0-only

// assert provides various assertion functions that can be used to confirm certain conditions such
// that these conditions are guaranteed true afterwards. These functions are particularly useful to
// catch unexpected and unsupported use cases, without having to litter the code with if-statements.
// Assertions may be placeholders for use cases that will later be supported, or they may indicate
// failure conditions that will not or cannot ever be supported, or cannot even occur. Some use
// cases or possible error conditions are illusions created by the type-system, for example when a
// function implements an interface but will never fail the operation.
// Regardless, assertions allow you to (subtly) handle cases and failure conditions that you do not
// handle otherwise.
//
// For use in testing, check `/std/testing` for assertions that accept `t` following Go's
// testing/benchmarking practices.
package assert
