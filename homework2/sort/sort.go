// Пакет sort с функциями для сортировки числовых массивов различными методами
package sort

// Сортировка
func SortByInserts(arr []int) {
	for i := 1; i < len(arr); i++ {
		currentIndex := i

		for j := i - 1; j >= 0; j-- {
			if arr[currentIndex] < arr[j] {
				arr[currentIndex], arr[j] = arr[j], arr[currentIndex]

				currentIndex = j
			} else {
				continue
			}
		}
	}
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
