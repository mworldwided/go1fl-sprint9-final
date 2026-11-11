package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "нулевой размер", size: 0, wantLen: 0},
		{name: "отрицательный размер", size: -5, wantLen: 0},
		{name: "один элемент", size: 1, wantLen: 1},
		{name: "обычный размер", size: 1000, wantLen: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if data == nil {
				t.Fatal("ожидался слайс, получен nil")
			}
			if len(data) != tt.wantLen {
				t.Fatalf("длина слайса: получено %d, ожидалось %d", len(data), tt.wantLen)
			}
			for i, value := range data {
				if value <= 0 {
					t.Fatalf("элемент %d = %d, ожидалось положительное число", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil-слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{9, 1, 5, 3}, want: 9},
		{name: "максимум в середине", data: []int{1, 5, 9, 3}, want: 9},
		{name: "максимум в конце", data: []int{1, 5, 3, 9}, want: 9},
		{name: "все элементы одинаковые", data: []int{7, 7, 7, 7}, want: 7},
		{name: "повторяющийся максимум", data: []int{3, 9, 2, 9}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil-слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "элементов меньше, чем частей", data: []int{4, 8, 1, 6, 2}, want: 8},
		{name: "элементов ровно CHUNKS", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "максимум в остатке после деления", data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 99}, want: 99},
		{name: "максимум в первой части", data: []int{50, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, want: 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

// TestMaxChunksMatchesMaximum проверяет, что многопоточный и однопоточный
// поиск дают одинаковый результат на случайных данных разного размера.
func TestMaxChunksMatchesMaximum(t *testing.T) {
	for _, size := range []int{1, 7, 8, 9, 100, 1003, 100_000} {
		data := generateRandomElements(size)

		if got, want := maxChunks(data), maximum(data); got != want {
			t.Errorf("размер %d: maxChunks = %d, maximum = %d", size, got, want)
		}
	}
}
