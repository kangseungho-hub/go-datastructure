package datastructure

type Map struct {
	keyArray [3571][]KeyValue
}

type KeyValue struct {
	key   string
	value string
}

func RollingHash(str string) int {
	h := 0
	A := 256
	B := 3571
	for _, character := range str {
		h = (h*A + int(character)) % B
	}
	return h
}

func (m *Map) Add(key, value string) {
	h := RollingHash(key)
	m.keyArray[h] = append(m.keyArray[h], KeyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := RollingHash(key)

	for _, value := range m.keyArray[h] {
		if value.key == key {
			return value.value
		}
	}

	return "Empty"
}
