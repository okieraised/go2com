package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/okieraised/go2com/pkg/dicom/tag"
)

func main() {
	fileName := "../../pkg/dicom/tag/tag_definitions.go"
	var file *os.File

	_, err := os.Stat(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, err = os.Create(fileName)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println(err)
			return
		}
	} else {
		err = os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	pkgDef := []byte("package tag\n\n")
	_, err = file.Write(pkgDef)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate tag variables
	r := regexp.MustCompile(`^\(.*,.*\)\t[a-zA-Z]{2}\t\w+\t.*\t\w+(?:\/retired)?`)
	scanner := bufio.NewScanner(strings.NewReader(tag.PublicTagDict))
	for scanner.Scan() {
		line := scanner.Text()
		if r.MatchString(line) {
			lineArr := regexp.MustCompile("\t").Split(line, -1)
			tagStr := strings.ReplaceAll(lineArr[0], "(", "")
			tagStr = strings.ReplaceAll(tagStr, ")", "")
			if strings.Contains(tagStr, "-") {
				continue
			}
			tagArr := strings.Split(tagStr, ",")
			tagGroup := tagArr[0]
			tagElem := tagArr[1]
			tagName := lineArr[2]
			tagName = strings.TrimPrefix(tagName, "RETIRED_")
			tagVarLine := fmt.Sprintf("var %s = DicomTag{0x%s, 0x%s}\n", tagName, tagGroup, tagElem)
			tagDef := []byte(tagVarLine)
			_, err = file.Write(tagDef)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}

	_, err = file.Write([]byte("\n\nvar TagDict map[DicomTag]TagInfo\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Write([]byte("\n\nfunc initTag() {\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Write([]byte("\n\tTagDict = make(map[DicomTag]TagInfo)\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate Tag dictionary
	scanner = bufio.NewScanner(strings.NewReader(tag.PublicTagDict))
	for scanner.Scan() {
		line := scanner.Text()
		if r.MatchString(line) {
			lineArr := regexp.MustCompile("\t").Split(line, -1)
			tagStr := strings.ReplaceAll(lineArr[0], "(", "")
			tagStr = strings.ReplaceAll(tagStr, ")", "")
			if strings.Contains(tagStr, "-u-") || strings.Contains(tagStr, "-o-") {
				continue
			}

			tagArr := strings.Split(tagStr, ",")
			tagElem := tagArr[1]
			tagVR := lineArr[1]
			tagName := lineArr[2]
			tagName = strings.TrimPrefix(tagName, "RETIRED_")
			tagVM := strings.ToUpper(lineArr[3])
			tagRetired := ""
			if strings.HasSuffix(lineArr[4], "retired") {
				tagRetired = "retired"
			}

			tagGroup := tagArr[0]
			if strings.Contains(tagGroup, "-") {
				tagGroupBegin := tagGroup[0:2]
				for i := 0x00; i <= 0xFF; i++ {
					tagGroup = fmt.Sprintf("%s%02X", tagGroupBegin, i)
					mapTagLine := fmt.Sprintf("\tTagDict[DicomTag{0x%s, 0x%s}] = TagInfo{\"%s\", \"%s\", \"%s\", \"%s\"}\n",
						tagGroup, tagElem, tagVR, tagName, tagVM, tagRetired)
					_, err = file.Write([]byte(mapTagLine))
					if err != nil {
						fmt.Println(err)
						return
					}

				}
				continue
			}

			if strings.Contains(tagElem, "-") {
				tagElemBegin := tagElem[0:2]
				for i := 0x00; i <= 0xFF; i++ {
					tagElem = fmt.Sprintf("%s%02X", tagElemBegin, i)
					mapTagLine := fmt.Sprintf("\tTagDict[DicomTag{0x%s, 0x%s}] = TagInfo{\"%s\", \"%s\", \"%s\", \"%s\"}\n",
						tagGroup, tagElem, tagVR, tagName, tagVM, tagRetired)
					_, err = file.Write([]byte(mapTagLine))
					if err != nil {
						fmt.Println(err)
						return
					}
				}
				continue
			}

			mapTagLine := fmt.Sprintf("\tTagDict[DicomTag{0x%s, 0x%s}] = TagInfo{\"%s\", \"%s\", \"%s\", \"%s\"}\n",
				tagGroup, tagElem, tagVR, tagName, tagVM, tagRetired)

			_, err = file.Write([]byte(mapTagLine))
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}

	_, err = file.Write([]byte("\n\n} \n"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
