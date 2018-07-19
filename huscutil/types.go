package huscutil

// "enum" for the type of each attribute
// s = string, n = number, b = boolean o = object, a = array,
// x = null,
// v = value (In an attribute. Uses value field)
// when prefixed with "a", means part of array. Use value only
const (
	s  = iota
	n  = iota
	b  = iota
	o  = iota
	a  = iota
	x  = iota
	v  = iota
	as = iota
	an = iota
	ab = iota
	ax = iota
)

// huscObject is used to represent the entire husc file parsed.
// Since any particular husc object may contain simple data, an
// array, or even another huscObject, these types need to be
// simple, with complex definitions to use in parsing to JSON.
//
// name		Array/attribute/object name.
// 			If in an array, name is only used for objects/arrays
//
// t		type, defined in const above. if o/a, treated specially
//
// value	Array of huscObjects.
type huscObject struct {
	name  string       // attribute/array/object name
	t     int          // type
	value []huscObject // values it has
}
