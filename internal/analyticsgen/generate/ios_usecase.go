package generate

import (
	"fmt"
	"github.com/prongbang/analyticsgen/internal/analyticsgen/sheets"
	"github.com/prongbang/analyticsgen/pkg/core"
	"github.com/prongbang/analyticsgen/pkg/csvx"
	"strings"
)

type IOSUseCase interface {
	UseCase
}

type iosUc struct {
	SheetUc sheets.UseCase
}

func (i *iosUc) GenKey(sheet sheets.Sheets) (string, error) {
	values := i.SheetUc.Get(sheet)
	contents := "import Foundation\n\npublic struct AnalyticsKey {\n"
	mapKey := map[string]string{}
	mapTopic := map[string]string{}

	for i := 0; i < len(Topics); i++ {
		if Topics[i] == "InformationValue" {
			header := "\n\tpublic enum " + Topics[i] + ": String {\n"
			if mapTopic[Topics[i]] == "" {
				mapTopic[Topics[i]] = "1"
				mapKey = map[string]string{}
			}
			end := "\t}\n"
			body := ""
			allValues := map[string]string{}
			for n := 1; n < len(values); n++ {
				cell := values[n][i]
				keys := strings.Split(strings.ToLower(cell), ",")
				for m := 0; m < len(keys); m++ {
					var key = keys[m]
					allValues[key] = key
					if mapKey[key] == "" && key != "" {
						mapKey[key] = "1"
						var typeCate = "\t\tcase " + key + "\n"
						body += typeCate
					}
				}
			}
			contents += header + body + end
		} else {
			header := "\n\tenum " + Topics[i] + ": String {\n"
			if mapTopic[Topics[i]] == "" {
				mapTopic[Topics[i]] = "1"
				mapKey = map[string]string{}
			}
			end := "\t}\n"
			body := ""
			for n := 1; n < len(values); n++ {
				cell := values[n][i]
				key := strings.ToLower(cell)
				typeCate := "\t\tcase " + key + "\n"
				if mapKey[key] == "" && key != "" {
					mapKey[key] = "1"
					body += typeCate
				}
			}
			contents += header + body + end
		}
	}
	contents += "}\n"

	return contents, nil
}

func (i *iosUc) GenCode(sheet sheets.Sheets) (string, error) {
	values := i.SheetUc.Get(sheet)
	functions := i.PrepareFunction(values)
	var keys = core.GetMapKeys(functions)

	var extension = "import Foundation\n\n extension Analytic {\n"
	var body = ""
	for k := range keys {
		var screen = fmt.Sprintf("%s", keys[k])
		var funcName = functions[screen]
		for j := 0; j < len(funcName); j++ {
			var statement = funcName[j].(map[string]interface{})
			body += i.BuildFunction(statement)
		}
	}
	extension += body
	extension += "\n}"
	return extension, nil
}

func (i *iosUc) PrepareFunction(values csvx.CsvList) map[string][]interface{} {
	functions := map[string][]interface{}{}
	informationKeys := map[string]map[string]string{}
	params := ""

	for r := 1; r < len(values)-1; r++ {
		var row = values[r]
		var next = values[r+1]

		var category = row[0]
		var screen = row[1]
		var logEvent = row[2]
		var label = row[3]
		var action = row[4]
		var infoKey = row[5]
		var infoValue = strings.Trim(row[6], " ")
		var functionName = row[7]

		var screenNext = next[1]
		var logEventNext = next[2]
		var labelNext = next[3]
		var actionNext = next[4]

		var funcName = ""
		var infoKeyArgs = category + "_" + infoKey

		if screen == screenNext && logEvent == logEventNext && label == labelNext && action == actionNext {
			if infoValue != "" {
				informationKeys[infoKeyArgs] = map[string]string{
					"value": infoKey,
					"type":  "InformationValue",
				}
				params += core.VariableCamel(infoKey) + ": AnalyticsKey.InformationValue?, "
			} else {
				informationKeys[infoKeyArgs] = map[string]string{
					"value": infoKey,
					"type":  "String",
				}
				params += core.VariableCamel(infoKey) + ": String?, "
			}
		} else {
			if infoKey != "" {
				if infoValue != "" {
					informationKeys[infoKeyArgs] = map[string]string{
						"value": infoKey,
						"type":  "InformationValue",
					}
					params = "(" + params + core.VariableCamel(infoKey) + ": AnalyticsKey.InformationValue?)"
				} else {
					informationKeys[infoKeyArgs] = map[string]string{
						"value": infoKey,
						"type":  "String",
					}
					params = "(" + params + core.VariableCamel(infoKey) + ": String?)"
				}
			} else {
				informationKeys = map[string]map[string]string{}
				params = "()"
			}

			funcName += "public func " + core.VariableCamel("event_"+screen+"_"+functionName) + params

			var fileName = screen
			if functions[screen] == nil {
				functions[fileName] = []interface{}{}
			}

			functions[fileName] = append(functions[fileName], map[string]interface{}{
				"package":        category,
				"event":          logEvent,
				"category":       category,
				"label":          label,
				"action":         action,
				"name":           funcName,
				"screen":         fileName,
				"informationKey": informationKeys,
			})

			// clear
			informationKeys = map[string]map[string]string{}
			params = ""
		}
	}
	return functions
}

func (i *iosUc) BuildFunction(statement map[string]interface{}) string {
	var body = ""
	body += "\n\t" + core.ToString(statement["name"]) + " {\n"
	body += "\t\tevent("
	body += "." + core.ToString(statement["event"]) + ", "
	body += "category: ." + core.ToString(statement["category"]) + ",\n"
	if core.ToString(statement["action"]) != "" {
		body += "\t\t\taction: ." + core.ToString(statement["action"]) + ", "
	} else {
		body += "\t\t\taction: nil, "
	}
	body += "\n\t\t\tlabel: ." + core.ToString(statement["label"]) + ",\n"
	body += "\n\t\t\tscreen: ." + core.ToString(statement["screen"]) + ",\n"
	var info = ""
	var infoKeys = core.GetMapKeys(statement["informationKey"])
	if len(infoKeys) > 0 {
		body += "\t\t\tinformation: [\n"
		for i := range infoKeys {
			var key = fmt.Sprintf("%s", infoKeys[i])
			information := statement["informationKey"].(map[string]map[string]string)
			var value = information[key]["value"]
			var types = information[key]["type"]
			if types == "String" {
				info += `\t\t\t\t"` + value + `": ` + core.VariableCamel(value)
			} else {
				info += `\t\t\t\t"` + value + `": ` + core.VariableCamel(value) + "?.rawValue"
			}
			if i < len(infoKeys)-1 {
				info += ",\n"
			} else {
				info += "\n"
			}
		}
		body += info
		body += "\t\t\t])\n"
	} else {
		body += "\t\t\tinformation: nil)\n"
	}
	body += "\t}\n"
	return body
}

func NewIOSUseCase(sheetUc sheets.UseCase) IOSUseCase {
	return &iosUc{
		SheetUc: sheetUc,
	}
}
