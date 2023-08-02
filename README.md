# Iterator library

A working progress attempt of an iterator library for go.
The objective of the library is also to support a set of functions for working with iterators, e.g., `map`, and `flatten`.

### Iterator interface

The iterator is composed of two methods `Next` that advance the iterator and `Get` for getting the element pointed by the iterator. Here is a usage example:

```golang
it := ToIter([]int{1, 2, 3, 4})
for it.Next() {
    fmt.Println(it.Get())
}
```