package kospell

import (
	"fmt"
	"sort"
)

func PrintDiff(orig string, checks []GrammarCheck) {

	runes := []rune(orig)

	for i := len(checks) - 1; 0 <= i; i-- {
		runes = getCorrectedWord(runes, checks[i])
	}

	fmt.Println(string(runes))
}

func getCorrectedWord(orig []rune, check GrammarCheck) []rune {
	sort.Slice(check.ErrInfo, func(i, j int) bool {
		return check.ErrInfo[i].ErrorIdx > check.ErrInfo[j].ErrorIdx
	})

	for _, errInfo := range check.ErrInfo {
		orig = replaceWordByErrInfo(orig, errInfo)
	}

	return orig
}

func replaceWordByErrInfo(orig []rune, errInfo GrammarCheckErrInfo) []rune {
	out := fmt.Sprintf("%s%s%s", string(orig[0:errInfo.Start]), errInfo.CandWord, string(orig[errInfo.End:]))
	return []rune(out)
}
