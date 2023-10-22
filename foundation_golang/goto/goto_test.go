package _goto

import (
	"fmt"
	"testing"
)

func TestGoto(t *testing.T) {
	i := 1

	for i <= 10 {
		if i == 5 {
			goto Skip
		}

		fmt.Println(i)
		i++
	}

Skip:
	fmt.Println("Skipped 5")
}

func TestGoto2(t *testing.T) {
	i := 1

Skip:
	if i == 5 {
		i++
		fmt.Println("Skipped 5")
	}
	for i <= 10 {
		if i == 5 {
			goto Skip
		}

		fmt.Println(i)
		i++
	}
}

type Task struct {
	Cid int64
}

func TestGotoSlice(t *testing.T) {
	//ConflictTaskSlice := make(map[int64]bool)
	ConflictTaskSlice := make([]int64, 0)
	TaskInfo := make([]*Task, 0)
	isConflict := false

	//ConflictTaskSlice[1] = true
	//ConflictTaskSlice[2] = true
	ConflictTaskSlice = append(ConflictTaskSlice, 1)
	ConflictTaskSlice = append(ConflictTaskSlice, 2)

	TaskInfo = append(TaskInfo, &Task{Cid: 2})
	TaskInfo = append(TaskInfo, &Task{Cid: 1})

loop:
	//break会跳出两个for
	for idx1, conflictCid := range ConflictTaskSlice {
		println(idx1, conflictCid)
		//if isConflict {
		//	break
		//}
		//loop:
		//loop:
		//break会跳出内层for，并将外层for加一
		for idx2, curTask := range TaskInfo {
			println(idx2, conflictCid)
			if conflictCid == curTask.Cid { // 和已有任务冲突
				isConflict = true
				//ss.Debugf("任务-新任务cid: %+v 和已有任务: %+v 冲突", newTaskCid, curTask)
				break loop // 【跳出外层的 for 循环】break loop是指定break到哪一层for循环。loop标签在内层for前一行，则break跳出内层循环；在外层for前一行，则break跳出外层for循环
			}
		}
	}

	//loop:
	println("isConflict: ", isConflict)
}
