package core

type request interface {
}

func ch[T request](x T, cb func(*T)) T {
	cb(&x)
	return x
}
