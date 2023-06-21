package utils

func CompareMaps(map1, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, ok := map2[key]; ok {
			if value1 != value2 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}