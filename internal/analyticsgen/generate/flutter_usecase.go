package generate

import (
	"fmt"
	"github.com/prongbang/analyticsgen/internal/analyticsgen/sheets"
	"github.com/prongbang/analyticsgen/pkg/core"
	"github.com/prongbang/analyticsgen/pkg/csvx"
	"strings"
)

type FlutterUseCase interface {
	UseCase
}

type flutterUc struct {
	SheetUc sheets.UseCase
}

func (f *flutterUc) GenKey(sheet sheets.Sheets) (string, error) {
	values := f.SheetUc.Get(sheet)
	contents := ""
	extensions := ""
	mapKey := map[string]string{}
	mapTopic := map[string]string{}
	for i := 0; i < len(Topics); i++ {
		if Topics[i] == InformationValue {
			className := "Analytics" + Topics[i]

			extensions += "extension " + className + "Extension on " + className + " {\n"
			extensions += "\tString get value {\n"
			extensions += "\t\tswitch (this) {\n"

			header := "enum Analytics" + Topics[i] + " {\n"
			if mapTopic[Topics[i]] == "" {
				mapTopic[Topics[i]] = "1"
				mapKey = map[string]string{}
			}
			end := "}\n\n"
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
						variableName := core.VariableCamel(key)
						typeCate := "\t" + variableName + ",\n"
						body += typeCate
						extensions += "\t\t\tcase " + className + "." + variableName + ":\n"
						extensions += "\t\t\t\treturn '" + key + "';\n"
					}
				}
			}
			contents += header + body + end
			extensions += "\t\t\tdefault:\n"
			extensions += "\t\t\t\treturn '';\n"
			extensions += "\t\t}\n"
			extensions += "\t}\n"
			extensions += "}"
		} else {
			header := "class Analytics" + Topics[i] + " {\n"
			if mapTopic[Topics[i]] == "" {
				mapTopic[Topics[i]] = "1"
				mapKey = map[string]string{}
			}
			end := "}\n\n"
			body := ""
			for n := 1; n < len(values); n++ {
				cell := values[n][i]
				key := strings.ToLower(cell)
				typeCate := "\tstatic final String " + core.VariableCamel(key) + " = '" + key + "';\n"
				if mapKey[key] == "" && key != "" {
					mapKey[key] = "1"
					body += typeCate
				}
			}
			contents += header + body + end
		}
	}

	contents += extensions

	return contents, nil
}

func (f *flutterUc) GenCode(sheet sheets.Sheets) (string, error) {
	values := f.SheetUc.Get(sheet)
	functions := f.PrepareFunction(values)
	var keys = core.GetMapKeys(functions)

	var contents = ""
	if sheet.Package != "" {
		contents += "import 'package:" + sheet.Package + "/analytics_key.dart';\n"
		contents += "import 'package:" + sheet.Package + "/analytics_utility.dart';\n"
		contents += "\n"
	}
	contents += "class Analytics {\n"
	contents += "\n"
	contents += "\tfinal AnalyticsUtility _analyticsUtility;\n"
	contents += "\n"
	contents += "\tAnalytics(this._analyticsUtility);\n"
	contents += "\n"
	contents += "\tFuture<void> logScreen(String screenName, {String screenClassOverride}) {\n"
	contents += "\t\treturn _analyticsUtility.logScreen(screenName,\n"
	contents += "\t\t\tscreenClassOverride: screenClassOverride);\n"
	contents += "\t}"

	var body = "\n"
	for k := range keys {
		var screen = fmt.Sprintf("%s", keys[k])
		var funcName = functions[screen]
		for i := 0; i < len(funcName); i++ {
			var statement = funcName[i].(map[string]interface{})
			body += f.BuildFunction(statement)
		}
	}
	contents += body
	contents += "\n}"
	return contents, nil
}

func (f *flutterUc) PrepareFunction(values csvx.CsvList) map[string][]interface{} {
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
				params += "AnalyticsInformationValue " + core.VariableCamel(infoKey) + ", "
			} else {
				informationKeys[infoKeyArgs] = map[string]string{
					"value": infoKey,
					"type":  "String",
				}
				params += "String " + core.VariableCamel(infoKey) + ", "
			}
		} else {
			if infoKey != "" {
				paramsSize := fmt.Sprintf("%d", len(strings.Split(params, ",")))
				if infoValue != "" {
					informationKeys[infoKeyArgs] = map[string]string{
						"value": infoKey,
						"type":  "InformationValue",
					}
					params = "WithParams(" + params + "AnalyticsInformationValue " + core.VariableCamel(infoKey) + ")"
				} else {
					informationKeys[infoKeyArgs] = map[string]string{
						"value": infoKey,
						"type":  "String",
					}
					params = "With" + paramsSize + "Params(" + params + "String " + core.VariableCamel(infoKey) + ")"
				}
			} else {
				informationKeys = map[string]map[string]string{}
				params = "()"
			}

			funcName += "Future<void> " + core.VariableCamel("event_"+screen+"_"+functionName) + params

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

func (f *flutterUc) BuildFunction(statement map[string]interface{}) string {
	var body = ""
	body += "\n\t" + core.ToString(statement["name"]) + " {\n"
	body += "\t\treturn _analyticsUtility.logEvent("
	body += "AnalyticsLogEvent." + core.VariableCamel(core.ToString(statement["event"])) + ", {\n"
	body += ""
	body += "\t\t\t'category': AnalyticsCategory." + core.VariableCamel(core.ToString(statement["category"])) + ",\n"
	if core.ToString(statement["action"]) != "" {
		body += "\t\t\t'action': AnalyticsAction." + core.VariableCamel(core.ToString(statement["action"])) + ", \n"
	} else {
		body += "\t\t\t'action': '', \n"
	}
	body += "\t\t\t'label': AnalyticsLabel." + core.VariableCamel(core.ToString(statement["label"])) + ",\n"
	body += "\t\t\t'screen': AnalyticsScreenName." + core.VariableCamel(core.ToString(statement["screen"])) + ",\n"

	var info = ""
	var infoKeys = core.GetMapKeys(statement["informationKey"])
	if len(infoKeys) > 0 {
		for i := range infoKeys {
			var key = fmt.Sprintf("%s", infoKeys[i])
			information := statement["informationKey"].(map[string]map[string]string)
			var value = information[key]["value"]
			var types = information[key]["type"]
			if types == "String" {
				info += "\t\t\t'" + value + "': " + core.VariableCamel(value)
			} else {
				info += "\t\t\t'" + value + "': " + core.VariableCamel(value) + ".value"
			}
			info += ",\n"
		}
		body += info
		body += "\t\t});\n"
	} else {
		body += "\t\t});\n"
	}
	body += "\t}\n"
	return body
}

func NewFlutterUseCase(sheetUc sheets.UseCase) FlutterUseCase {
	return &flutterUc{
		SheetUc: sheetUc,
	}
}
