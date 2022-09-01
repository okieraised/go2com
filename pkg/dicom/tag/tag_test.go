package tag

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestGenTag(t *testing.T) {
	// fileName := "./tag_definitions.go"

	// _, err := os.Stat(fileName)
	// if err != nil {
	// 	if errors.Is(err, os.ErrNotExist) {
	// 		_, err := os.Create(fileName)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 	} else {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// } else {
	// 	err = os.Remove(fileName)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	_, err := os.Create(fileName)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }

	r := regexp.MustCompile(`^\(.*,.*\)\t[A-Z]{2}\t\w+\t.*\t\w+(?:\/retired)?`)
	scanner := bufio.NewScanner(strings.NewReader(PublicTagDict))
	for scanner.Scan() {
		line := scanner.Text()
		if r.MatchString(line) {
			lineArr := regexp.MustCompile("\t").Split(line, -1)
			// fmt.Println(lineArr)
			tagStr := strings.ReplaceAll(lineArr[0], "(", "")
			tagStr = strings.ReplaceAll(tagStr, ")", "")
			fmt.Println(tagStr)
			if strings.Contains(tagStr, "-u-") || strings.Contains(tagStr, "-o-") {
				continue
			}

		}

	}
}
