package format

type Field struct {
	ValueType  string
	Optional   bool
	Required   bool
	Default    interface{}
	Computed   bool
	ForceNew   bool
	Deprecated string
	Removed    string
	Elem       interface{}
}
