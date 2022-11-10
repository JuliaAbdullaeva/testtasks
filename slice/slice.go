package slice

//На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение
func intersectionSlices(a, b []int) (intersectSlice []int) {
	mapIntersect := make(map[int]int)
	for _, itemA := range a {
		for _, itemB := range b {
			if itemA == itemB {
				_, ok := mapIntersect[itemB]
				if !ok {
					mapIntersect[itemB] = itemB
					intersectSlice = append(intersectSlice, itemB)
				}
			}
		}
	}
	return
}
