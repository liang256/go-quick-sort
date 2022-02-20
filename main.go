package main

import "fmt"

func main() {
	sli := []int{4, 6, 2, 1, 5, 3, 8, 9}

	quickSort(sli)
	fmt.Println(sli)
}

func quickSort(sli []int) {
	if len(sli) <= 0 {
		return
	}
	quickSortRecur(sli, 0, len(sli)-1)
}

func quickSortRecur(sli []int, start_i int, end_i int) {
	// end condition
	if start_i >= end_i {
		return
	}

	// main logic
	guess_i := guessPV(start_i, end_i)
	guess_i = sort(sli, start_i, end_i, guess_i)

	// branch
	quickSortRecur(sli, start_i, guess_i-1)
	quickSortRecur(sli, guess_i+1, end_i)
}

func guessPV(start_i int, end_i int) int {
	return (start_i + end_i) / 2
}

func sort(sli []int, start_i int, end_i int, guess_i int) int {
	swap(sli, guess_i, end_i)
	pv_i := end_i
	end_i--

	for true {
		for true {
			if sli[start_i] > sli[pv_i] || start_i == end_i {
				break
			}
			start_i++
		}

		for true {
			if sli[end_i] < sli[pv_i] || start_i == end_i {
				break
			}
			end_i--
		}

		if start_i == end_i {
			break
		}
		swap(sli, start_i, end_i)
	}
	meet_i := start_i
	if sli[meet_i] >= sli[pv_i] {
		swap(sli, pv_i, meet_i)
		return meet_i
	}
	return pv_i
}

func swap(sli []int, l_i int, r_i int) {
	if l_i != r_i {
		sli[l_i] = sli[r_i] ^ sli[l_i]
		sli[r_i] = sli[r_i] ^ sli[l_i]
		sli[l_i] = sli[r_i] ^ sli[l_i]
	}
}
