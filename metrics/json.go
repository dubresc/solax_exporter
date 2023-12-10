package metrics

func GetField[K any](m map[string]interface{}, f string) (K, bool) {
	var res K

	iface, ok := m[f]

	if !ok {
		return res, false
	}

	res, ok = iface.(K)

	return res, ok
}

func GetNullableField[K any](m map[string]interface{}, f string) *K {
	var res K

	iface, ok := m[f]

	if !ok {
		return nil
	}

	res, ok = iface.(K)

	if !ok {
		return nil
	}

	return &res
}
