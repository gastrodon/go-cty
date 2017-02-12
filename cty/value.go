package cty

// Value represents a value of a particular type, and is the interface by
// which operations are executed on typed values.
//
// The philosophy for the operations API is that it's the caller's
// responsibility to ensure that the given types and values satisfy the
// specified invariants during a separate type check, so that the caller is
// able to return errors to its user from the application's own perspective.
//
// Consequently the design of these methods assumes such checks have already
// been done and panics if any invariants turn out not to be satisfied. These
// panic errors are not intended to be handled, but rather indicate a bug in
// the calling application that should be fixed with more checks prior to
// executing operations.
//
// A related consequence of this philosophy is that no automatic type
// conversions are done. If a method specifies that its argument must be
// number then it's the caller's responsibility to do that conversion before
// the call, thus allowing the application to have more constrained conversion
// rules than are offered by the built-in converter where necessary.
type Value struct {
	ty Type
	v  interface{}
}

// Type returns the type of the value.
func (val Value) Type() Type {
	return val.ty
}

// IsKnown returns true if the value is known. That is, if it is not
// the result of the unknown value constructor Unknown(...), and is not
// the result of an operation on another unknown value.
//
// Unknown values are only produced either directly or as a result of
// operating on other unknown values, and so an application that never
// introduces Unknown values can be guaranteed to never receive any either.
func (val Value) IsKnown() bool {
	return val.v != unknown
}

// IsNull returns true if the value is null. Values of any type can be
// null, but any operations on a null value will panic. No operation ever
// produces null, so an application that never introduces Null values can
// be guaranteed to never receive any either.
func (val Value) IsNull() bool {
	return val.v == nil
}
