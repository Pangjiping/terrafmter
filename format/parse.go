package format

import (
	"bufio"
	"github.com/Pangjiping/terrafmtter/util"
	"github.com/waigani/diffparser"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	NAME        = "Name"
	ARGUMENT    = "Argument"
	ATTRIBUTE   = "Attribute"
	OPTIONAL    = "Optional"
	FORCENEW    = "ForceNew"
	REQUIRED    = "Required"
	TYPE        = "Type"
	DESCRIPTION = "Description"
	PARAMS      = "params"
)

type parsed struct {
	name      string
	arguments map[string]map[string]string
}

// parseResource parse .md file to parsed struct.
// resourceName - cs_kubernetes_version
// todo: cancel attr regex
func parseResource(resourceName string) (*parsed, error) {
	filePath := strings.Join([]string{util.FILE_LOC_PREFIX, resourceName, ".md"}, "")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	argsRegex := regexp.MustCompile(`## Argument Reference`)
	secondLevelRegex := regexp.MustCompile(`^#+`)
	argumentsFieldRegex := regexp.MustCompile("^\\- `([a-zA-Z_0-9]*)`[ \\\\]*-? ?(\\(.*\\)) ?(.*)")
	// defaultSetRegex := regexp.MustCompile("Default [a-z]?(.*)")

	result := &parsed{
		name:      resourceName,
		arguments: map[string]map[string]string{},
	}

	scanner := bufio.NewScanner(file)
	phase := ARGUMENT
	record := false
	for scanner.Scan() {
		line := scanner.Text()
		if argsRegex.MatchString(line) {
			record = true
			phase = ARGUMENT
			continue
		}
		if secondLevelRegex.MatchString(line) && strings.HasSuffix(line, PARAMS) {
			record = true
			continue
		}

		if record {
			if secondLevelRegex.MatchString(line) && !strings.HasSuffix(line, PARAMS) {
				record = false
				continue
			}
			var matched [][]string
			if phase == ARGUMENT {
				matched = argumentsFieldRegex.FindAllStringSubmatch(line, 1)
			}

			for _, m := range matched {
				field := parseMatchLine(m, phase)
				field[TYPE] = phase
				if v, exist := field[NAME]; exist {
					result.arguments[v] = field
				}
			}
		}
	}
	return result, nil
}

func parseMatchLine(words []string, phase string) map[string]string {
	res := make(map[string]string, 0)
	if phase == ARGUMENT && len(words) >= 4 {
		res[NAME] = words[1]
		res[DESCRIPTION] = words[3]
		if strings.Contains(words[2], OPTIONAL) {
			res[OPTIONAL] = "true"
		}
		if strings.Contains(words[2], REQUIRED) {
			res[REQUIRED] = "true"
		}
		if strings.Contains(words[2], FORCENEW) {
			res[FORCENEW] = "true"
		}

		// todo: get default

		return res
	}
	return nil
}

func parseField(hunk diffparser.DiffRange, length int) map[string]map[string]interface{} {
	schemaRegex := regexp.MustCompile("^\\t*\"([a-zA-Z_]*)\"")
	typeRegex := regexp.MustCompile("^\\t*Type:\\s+schema.([a-zA-Z]*)")
	optionRegex := regexp.MustCompile("^\\t*Optional:\\s+([a-z]*),")
	forceNewRegex := regexp.MustCompile("^\\t*ForceNew:\\s+([a-z]*),")
	requiredRegex := regexp.MustCompile("^\\t*Required:\\s+([a-z]*),")
	validateStringRegex := regexp.MustCompile("^\\t*ValidateFunc: ?validation.StringInSlice\\(\\[\\]string\\{([a-z\\-A-Z_,\"\\s]*)")

	temp := map[string]interface{}{}
	schemaName := ""
	raw := make(map[string]map[string]interface{}, 0)
	for i := 0; i < length; i++ {
		currentLine := hunk.Lines[i]
		content := currentLine.Content
		fieldNameMatched := schemaRegex.FindAllStringSubmatch(content, -1)
		if fieldNameMatched != nil && fieldNameMatched[0] != nil {
			if len(schemaName) != 0 && schemaName != fieldNameMatched[0][1] {
				temp["Name"] = schemaName
				raw[schemaName] = temp
				temp = map[string]interface{}{}
			}
			schemaName = fieldNameMatched[0][1]
		}

		if !schemaRegex.MatchString(currentLine.Content) && currentLine.Mode == diffparser.UNCHANGED {
			continue
		}

		typeMatched := typeRegex.FindAllStringSubmatch(content, -1)
		typeValue := ""
		if typeMatched != nil && typeMatched[0] != nil {
			typeValue = typeMatched[0][1]
			temp["Type"] = typeValue
		}

		optionalMatched := optionRegex.FindAllStringSubmatch(content, -1)
		optionValue := ""
		if optionalMatched != nil && optionalMatched[0] != nil {
			optionValue = optionalMatched[0][1]
			op, _ := strconv.ParseBool(optionValue)
			temp["Optional"] = op
		}

		forceNewMatched := forceNewRegex.FindAllStringSubmatch(content, -1)
		forceNewValue := ""
		if forceNewMatched != nil && forceNewMatched[0] != nil {
			forceNewValue = forceNewMatched[0][1]
			fc, _ := strconv.ParseBool(forceNewValue)
			temp["ForceNew"] = fc
		}

		requiredMatched := requiredRegex.FindAllStringSubmatch(content, -1)
		requiredValue := ""
		if requiredMatched != nil && requiredMatched[0] != nil {
			requiredValue = requiredMatched[0][1]
			rq, _ := strconv.ParseBool(requiredValue)
			temp["Required"] = rq
		}

		validateStringMatched := validateStringRegex.FindAllStringSubmatch(content, -1)
		validateStringValue := ""
		if validateStringMatched != nil && validateStringMatched[0] != nil {
			validateStringValue = validateStringMatched[0][1]
			temp["ValidateFuncString"] = strings.Split(validateStringValue, ",")
		}

	}
	if _, exist := raw[schemaName]; !exist && len(temp) >= 1 {
		temp["Name"] = schemaName
		raw[schemaName] = temp
	}
	return raw
}
