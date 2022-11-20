package structure

type Stack[T any] struct {
    d []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{}
}

func (s *Stack[T]) Length() int {
    return len(s.d)
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.d) == 0
}

func (s *Stack[T]) Push(val T) {
    s.d = append(s.d, val)
}

func (s *Stack[T]) Pop() (val T, ok bool) {
    if s.Length() == 0 {
        ok = false
        return
    }
    ok = true
    val = s.d[s.Length() - 1]
    s.d = s.d[0:s.Length() - 1]
    return
}

