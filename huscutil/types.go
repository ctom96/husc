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
	toString() string
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

func (h huscObject) toString() string {
	retStr := h.name + ":{\n"
	for _, val := range h.values {
		retStr += "\t" + val.toString() + "\n"
	}
	retStr += "}"

	return retStr
}

type huscArray struct {
	name   string
	values []huscCompliant
}

// Make huscArrays huscComplaint
func (h huscArray) dataType() int {
	return a
}

func (h huscArray) toString() string {
	retStr := h.name + ":[\n"
	for _, val := range h.values {
		retStr += "\t" + val.toString() + "\n"
	}
	retStr += "]"

	return retStr
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

func (h huscSingle) toString() string {
	return h.name + ": " + h.value
}
