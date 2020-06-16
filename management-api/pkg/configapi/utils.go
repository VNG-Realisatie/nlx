package configapi

func contains(s []string, v string) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}

	return false
}
