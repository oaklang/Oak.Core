package core

import "github.com/oaklang/runtime"

type Basics_Bool runtime.TypeUnion

func Basics_True() Basics_Bool { panic("not implemented") }

func Basics_False() Basics_Bool { panic("not implemented") }

type Basics_Order runtime.TypeUnion

func Basics_LT() Basics_Order { panic("not implemented") }

func Basics_GT() Basics_Order { panic("not implemented") }

func Basics_EQ() Basics_Order { panic("not implemented") }

type Maybe_Maybe[T any] runtime.TypeUnion

func Maybe_Just[T any](T) Maybe_Maybe[T] { panic("not implemented") }

func Maybe_Nothing[T any]() Maybe_Maybe[T] { panic("not implemented") }
