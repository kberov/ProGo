package main

func Find(slice []string, vals ...string) (matches bool) {
	for _, sliceElement := range slice {
		for _, searchString := range vals {
			if sliceElement == searchString {
				matches = true
				return
			}
		}
	}
	return
}
