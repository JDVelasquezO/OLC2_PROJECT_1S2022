package AbstractOptimizer

type Expression interface {
	GetValue() interface{}
}

type Instruction interface {
	Execute() interface{}
}
