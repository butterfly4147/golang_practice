package main

// EConstJudgeRange the judgement range
type EConstJudgeRange int

const (
	EConstLeft  EConstJudgeRange = 1
	EConstRight EConstJudgeRange = 2
)

func quickSort(salaries []int, lIdx int, rIdx int) {
	if lIdx >= rIdx {
		return
	}

	benchmark := salaries[0] //benchmark
	left := lIdx             //left idx
	right := rIdx            // right idx
	mark := EConstLeft       //a flag to decide the judgement range is left or right, default left

	for left < right {
		if mark == EConstLeft {
			if salaries[left] < benchmark {
				left++
				continue
			}
			// ==, do nothing
			if salaries[left] > benchmark {
				benchmark, salaries[left] = salaries[left], benchmark
			}

			mark = EConstRight
		}

		if mark == EConstRight {
			if benchmark < salaries[right] {
				right--
				continue
			}
			// ==, do nothing
			if benchmark > salaries[right] {
				salaries[right], benchmark = benchmark, salaries[right]
			}

			mark = EConstLeft
		}
	}

	//Actually, at the time, left == right
	quickSort(salaries, lIdx, left-1)
	quickSort(salaries, right+1, rIdx)
}

//func main() {
//	salaries := []int{10, 30, 80, 90, 40, 50, 70}
//	quickSort(salaries, 0, len(salaries)-1)
//	log.Printf("%+v\n", salaries)
//}
