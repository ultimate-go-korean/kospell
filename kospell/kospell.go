package kospell

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const pusanUrl = "https://speller.cs.pusan.ac.kr/results"

var noErrorRegexp = regexp.MustCompile("맞춤법과 문법 오류를 찾지\\s+못했습니다")

type GrammarCheckErrInfo struct {
	Help          string `json:"help"`
	ErrorIdx      int    `json:"errorIdx"`
	CorrectMethod int    `json:"correctMethod"`
	Start         int    `json:"start"`
	End           int    `json:"end"`
	OrgStr        string `json:"orgStr"`
	CandWord      string `json:"candWord"`
}

type GrammarCheck struct {
	Str     string                `json:"str"`
	ErrInfo []GrammarCheckErrInfo `json:"errInfo"`
	Idx     int                   `json:"idx"`
}

func Check(text string) ([]GrammarCheck, error) {
	formData := url.Values{
		"text1": []string{text},
	}
	resp, err := http.PostForm(pusanUrl, formData)

	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	out, err := ExtractGrammarChecks(bs)
	return out, err
}

func ExtractGrammarChecks(data []byte) ([]GrammarCheck, error) {

	if noErrorRegexp.Match(data) {
		return nil, nil
	}

	dataStr := string(data)

	firstNeedle := "data = [{"
	secondNeedle := "}];"

	startIndex := strings.Index(dataStr, firstNeedle)
	if startIndex == -1 {
		return nil, fmt.Errorf("`%s` is not found. data = %s", firstNeedle, dataStr)
	}
	endIndex := strings.LastIndex(dataStr, secondNeedle)
	if endIndex == -1 {
		return nil, fmt.Errorf("`%s` is not found. data = %s", secondNeedle, dataStr)
	}

	dataJson := dataStr[startIndex+len(firstNeedle)-2 : endIndex+len(secondNeedle)-1]
	var grammarChecks []GrammarCheck
	err := json.Unmarshal([]byte(dataJson), &grammarChecks)

	return grammarChecks, err
}
