package learn_go_with_tests

func sum(slice1 []int) (sum int) {
	for _, element := range slice1 {
		sum += element
	}
	return
}

func sumAll(slices ...[]int) (final []int) {

	for _, slice := range slices {
		final = append(final, sum(slice))
	}
	return
}

func sumAllTails(slices ...[]int) (sumTail []int) {
	for _, sliceElement := range slices {
		if len(sliceElement) == 0 {
			sumTail = append(sumTail, 0)
		} else {
			sumTail = append(sumTail, sum(sliceElement[1:]))
		}
	}
	return
}
