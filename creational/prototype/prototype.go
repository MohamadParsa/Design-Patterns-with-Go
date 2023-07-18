package prototype

type Prototype interface {
	Clone() Prototype

	String() string
}
