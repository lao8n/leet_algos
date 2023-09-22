package data_structures

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		iIsLetterLog := isLetterLog(logs[i])
		jIsLetterLog := isLetterLog(logs[j])

		if iIsLetterLog && !jIsLetterLog {
			return true
		}
		if !iIsLetterLog && jIsLetterLog {
			return false
		}

		if iIsLetterLog && jIsLetterLog {
			return compareLetterLog(logs[i], logs[j])
		}

		return false
	})
	return logs
}

func isLetterLog(log string) bool {
	for i := 0; i < len(log); i++ {
		if log[i] != ' ' {
			continue
		}
		if '0' <= log[i+1] && log[i+1] <= '9' {
			return false
		}
		break
	}
	return true
}

func compareLetterLog(a, b string) bool {
	indexA := strings.Index(a, " ")
	indexB := strings.Index(b, " ")

	idA := a[:indexA]
	idB := b[:indexB]

	restA := a[indexA:]
	restB := b[indexB:]

	cmp := strings.Compare(restA, restB)
	if cmp != 0 {
		return cmp == -1
	}

	return idA < idB
}
