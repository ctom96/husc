package huscutil

// "enum" for the type of each attribute
// s = string, n = number, b = boolean o = object, a = array,
// N = null
const (
	s = iota
	n = iota
	b = iota
	o = iota
	a = iota
	N = iota
)

type huscCompliant interface {
	dataType() int
}

// huscObject is used to represent a single object, like a single
// JSON object. This specifically represents a huscObject of type o
// so it must contain at least 1 other huscObject
type huscObject struct {
	name   string          // attribute/array/object name
	values []huscCompliant // values it has
}

// Make huscObjects huscComplaint
func (h huscObject) dataType() int {
	return o
}

type huscArray struct {
	name   string
	values []huscCompliant
}

// huscSingle is a one-line huscObject, the simplest type
type huscSingle struct {
	name  string // Name of this single
	dType int    // type of the data inside this object, defined in format.txt
	value string // actual value
}

// Make huscSingles huscCompliant
func (h huscSingle) dataType() int {
	return h.dType
}
