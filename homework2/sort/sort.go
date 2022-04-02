// Пакет sort с функциями для сортировки числовых массивов различными методами
//
// Функция изменяет переданный слайс чисел, выполнив сортировку по возрастанию вставками
//
// SortByInserts(arr []int)
//
// Функция изменяет переданный слайс чисел, выполнив сортировку по возрастанию "пузырьком"
//
// BubbleSort(arr []int)
package sort

// Сортировка вставками
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

// Пузырьковая сортировка
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
