package security

import (
	"strings"
	"fmt"
	"go-crawler/http"
	"regexp"
	"math/rand"
)

var DBMS_ERRORS = map[string][]string{
	"MySQL": []string{ `SQL syntax.*MySQL`, `Warning.*mysql_.*`, `valid MySQL result`, `MySqlClient\.`},
	// TODO ADD MORE DATABASES
}

var BOOLEAN_TESTS = []string{" AND %d=%d", " OR NOT (%d=%d)"}

func CheckSqlInjection(url string) string {
	matched := strings.HasSuffix(url,".html")
	// TODO use regexp later
	if matched {
		return url + " not match URL format"
	}
	body := http.Get(url + "%29%28%22%27")

	for dbName, regs := range DBMS_ERRORS {
		for _, reg := range regs {
			res, err := regexp.MatchString(reg, body)

			if err != nil {
				panic(err.Error())
			}

			if res {
				return fmt.Sprintf("SQLInjection Found: %s database: %s", url, dbName)
			}
		}
	}

	for _, payload := range BOOLEAN_TESTS {
		origin := http.Get(url)
		randNum := rand.Intn(100)
		testURL := url + payload
		boolEqualURL := fmt.Sprintf(testURL, randNum, randNum)
		boolEqual := http.Get(boolEqualURL)
		boolNotEqualURL :=	fmt.Sprintf(testURL, randNum, randNum+1)
		boolNotEqual := http.Get(boolNotEqualURL)

		if (origin == boolEqual) && (origin != boolNotEqual) {
			return fmt.Sprintf("SQLInjection Found:  %s", url)
		}

	}

	return ""
}