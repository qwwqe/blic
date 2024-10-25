package main

type Cloneable[T any] interface {
	Clone() T
}

func CloneSlice[T Cloneable[T]](slice []T) []T {
	newSlice := make([]T, 0, len(slice))
	for _, item := range slice {
		newSlice = append(newSlice, item.Clone())
	}
	return newSlice
}
