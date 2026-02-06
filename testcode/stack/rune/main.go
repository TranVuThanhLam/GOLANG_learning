package main

type stackRune struct {
	elements []rune
}

func (s *stackRune) Push(r rune) {
	s.elements = append(s.elements, r)
}

func (s *stackRune) Pop() {
	if s.isEmpty() {
		return
	}
	s.elements = s.elements[:(len(s.elements) - 1)]
}

func (s *stackRune) Peak() rune {
	var r rune
	if s.isEmpty() {
		return r
	}

	return s.elements[len(s.elements)-1]
}

func (s *stackRune) isEmpty() bool {
	if len(s.elements) == 0 {
		return true
	} else {
		return false
	}
}

func isValid(s string) bool {
	var sr stackRune
	for _, value := range s {
		if sr.Peak() == '(' {
			if value == ')' {
				sr.Pop()
				continue
			}
		} else if sr.Peak() == '[' {
			if value == ']' {
				sr.Pop()
				continue
			}
		} else if sr.Peak() == '{' {
			if value == '}' {
				sr.Pop()
				continue
			}
		}
		sr.Push(value)
	}

	if sr.isEmpty() {
		return true
	} else {
		return false
	}
}

func main() {
	s := "({[)"
	// s := ""
	if isValid(s) {
	} else {
	}

}
