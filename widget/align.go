package widget

type Alignment uint8

const (
	AlignTop    Alignment = 1 << iota
	AlignRight  Alignment = 1 << iota
	AlignBottom Alignment = 1 << iota
	AlignLeft   Alignment = 1 << iota
)
