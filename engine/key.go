package engine

// Key is used to access values.
// TODO: ensure all values come from here
type Key string

// Value is the stringified value of the Key type.
func (k Key) Value() string {
	return string(k)
}
