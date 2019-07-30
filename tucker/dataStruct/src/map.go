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
	key   string
	value string
}

type Map struct {
	KeyArray [3571][]KeyValue
}

func CreateMap() *Map {
	return &Map{}
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.KeyArray[h] = append(m.KeyArray[h], KeyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.KeyArray[h]); i++ {
		if m.KeyArray[h][i].key == key {
			return m.KeyArray[h][i].value
		}
	}

	return ""
}
