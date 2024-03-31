package dict

type Dict[K comparable, V any] map[K]V

func New[K comparable, V any]() Dict[K, V] {
	return make(Dict[K, V])
}

func From[K comparable, V any](m map[K]V) Dict[K, V] {
	return m
}

func (m Dict[K, V]) Size() int {
	return len(m)
}

func (m Dict[K, V]) Add(key K, value V) Dict[K, V] {
	m[key] = value
	return m
}

func (m Dict[K, V]) Pop(key K) Dict[K, V] {
	delete(m, key)
	return m
}
