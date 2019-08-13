package src

func Hash(s string) int {
	h := 0
	A := 256
	B := 3571

	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}
	return h
}

type KeyValue struct {
	Key   string
	Value string
}

type Map struct {
	KeyArray [3571][]KeyValue
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.KeyArray[h] = append(m.KeyArray[h], KeyValue{key})
}
