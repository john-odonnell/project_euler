package output

import (
	"fmt"
	"time"
)

func PrintAnswer(description, answer string, pDuration, sDuration time.Duration) {
	answerStr := fmt.Sprintf("Answer:           %s", answer)
	pDurationStr := fmt.Sprintf("Process Duration: %s", pDuration)
	sDurationStr := fmt.Sprintf("Solve Duration:   %s", sDuration)
	tDurationStr := fmt.Sprintf("Total Duration:   %s", pDuration+sDuration)

	longest := len(description)
	if len(answerStr) > longest {
		longest = len(answerStr)
	}
	if len(pDurationStr) > longest {
		longest = len(pDurationStr)
	}
	if len(sDurationStr) > longest {
		longest = len(sDurationStr)
	}
	if len(tDurationStr) > longest {
		longest = len(tDurationStr)
	}

	barrierStr := ""
	for i := 0; i < longest; i++ {
		barrierStr += "="
	}

	fmt.Println(barrierStr)
	fmt.Println(description)
	fmt.Println(answerStr)
	fmt.Println(pDurationStr)
	fmt.Println(sDurationStr)
	fmt.Println(tDurationStr)
	fmt.Println(barrierStr)
}
