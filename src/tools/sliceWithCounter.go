package tools

type SliceWithCounter[T any] struct {
	Slice []T
	Len   int
}

func NewSliceWithCounter[T any](size int) SliceWithCounter[T] {
	return SliceWithCounter[T]{
		Slice: make([]T, size),
		Len:   0,
	}
}

func (slice *SliceWithCounter[T]) AddToEnd(value T) {
	slice.Slice[slice.Len] = value
	slice.Len++
}

func (slice *SliceWithCounter[T]) Append(value T) {
	slice.Slice = append(slice.Slice, value)
	slice.Len++
}

func (slice *SliceWithCounter[T]) Get(index int) T {
	return slice.Slice[index]
}

func (slice *SliceWithCounter[T]) Set(index int, value T) {
	slice.Slice[index] = value
	if index >= slice.Len {
		slice.Len = index + 1
	}
}

func (slice *SliceWithCounter[T]) RemoveLast() {
	slice.Len--
}

// Shifts elements to the left
func (slice *SliceWithCounter[T]) Remove(index int) {
	/*for i := index; i < slice.Len-1; i++ {
		slice.Slice[i] = slice.Slice[i+1]
	}*/
	if index == slice.Len-1 {
		slice.RemoveLast()
	} else {
		slice.Slice[index] = slice.Slice[slice.Len-1]
		slice.Len--
	}
}

func (slice *SliceWithCounter[T]) RemoveAndGet(index int) (removed T) {
	removed = slice.Slice[index]
	/*for i := index; i < slice.Len-1; i++ {
		slice.Slice[i] = slice.Slice[i+1]
	}
	slice.Len--*/
	slice.Remove(index)
	return
}

// Sadly can't define a generic function whose type is more restrictive directly on the struct... so have to do like this.
func RemoveByValueSlice[V comparable](slice *SliceWithCounter[V], value V) {
	for i := 0; i < slice.Len; i++ {
		if slice.Slice[i] == value {
			slice.Remove(i)
			return
		}
	}
}

func (slice *SliceWithCounter[T]) Clear() {
	slice.Len = 0
}

func (slice *SliceWithCounter[T]) DeepClear() {
	for i := 0; i < len(slice.Slice); i++ { //Deep clears even positions that were previously removed.
		var zeroValue T
		slice.Slice[i] = zeroValue
	}
	slice.Len = 0
}

func (slice *SliceWithCounter[T]) IsEmpty() bool {
	return slice.Len == 0
}

func (slice *SliceWithCounter[T]) ToSlice() []T {
	return slice.Slice[:slice.Len]
}

func (slice *SliceWithCounter[T]) Cap() int {
	return cap(slice.Slice)
}

func (slice *SliceWithCounter[T]) Copy() SliceWithCounter[T] {
	newSlice := SliceWithCounter[T]{Slice: make([]T, cap(slice.Slice)), Len: slice.Len}
	copy(newSlice.Slice, slice.Slice[:slice.Len])
	return newSlice
}
