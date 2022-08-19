package engine

// TODO: ensure all values come from here
type Key string

func (k Key) Value() string {
	return string(k)
}
