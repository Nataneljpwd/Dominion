package utils

type Predicate[T any] func(T) bool

type PredicateMap[T any, V any] struct {
	pairs []PredicatePair[T, V]
}

func (pm *PredicateMap[T, V]) Add(predicate Predicate[T], value V) bool {
	pm.pairs = append(pm.pairs, *&PredicatePair[T, V]{predicate, value})
	return true
}

func (pm *PredicateMap[T, V]) Get(value T) (V, bool) {
	for _, pair := range pm.pairs {
		if pair.predicate(value) {
			return pair.value, true
		}
	}

	var result V
	return result, false
}

func NewPredicateMap[T any, V any]() *PredicateMap[T, V] {
	return &PredicateMap[T, V]{[]PredicatePair[T, V]{}}
}

type PredicatePair[T any, V any] struct {
	predicate Predicate[T]
	value     V
}
