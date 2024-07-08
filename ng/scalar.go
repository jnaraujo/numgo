package ng

type IntScalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UIntScalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type FloatScalar interface {
	~float32 | ~float64
}

type Scalar interface {
	IntScalar | UIntScalar | FloatScalar
}
