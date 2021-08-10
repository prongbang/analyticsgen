package analyticsgen

import (
	"fmt"
	"github.com/prongbang/analyticsgen/internal/analyticsgen/generate"
	"github.com/prongbang/analyticsgen/internal/analyticsgen/sheets"
	"github.com/prongbang/analyticsgen/pkg/parameter"
	"github.com/prongbang/analyticsgen/pkg/parameter/asset"
	"github.com/prongbang/analyticsgen/pkg/parameter/platform"
	"github.com/prongbang/filex"
	"strings"
)

const (
	AndroidAnalyticsKey = "analytics_key.xml"
	AndroidAnalytics    = "Analytics.kt"
	FlutterAnalyticsKey = "analytics_key.dart"
	FlutterAnalytics    = "analytics.dart"
	IOSAnalyticsKey     = "AnalyticsKey.swift"
	IOSAnalytics        = "Analytics.swift"
)

type AnalyticsGen interface {
	Process(params *parameter.Parameter)
}

type analyticsGen struct {
	FileX     filex.FileX
	AndroidUc generate.AndroidUseCase
	FlutterUc generate.FlutterUseCase
	IOSUc     generate.IOSUseCase
}

func (a *analyticsGen) Process(params *parameter.Parameter) {
	sheet := sheets.Sheets{
		Id:         params.Sheet,
		DocumentId: params.Document,
		Package:    params.Package,
	}
	if params.Platform == platform.Android {
		if strings.Contains(params.Asset, asset.Key) {

		}
		if strings.Contains(params.Asset, asset.Code) {

		}
		if strings.Contains(params.Asset, asset.Test) {

		}
	} else if params.Platform == platform.Flutter {
		if strings.Contains(params.Asset, asset.Key) {
			code, err := a.FlutterUc.GenKey(sheet)
			a.createFile(err, params.Target, FlutterAnalyticsKey, code)
		}
		if strings.Contains(params.Asset, asset.Code) {
			code, err := a.FlutterUc.GenCode(sheet)
			a.createFile(err, params.Target, FlutterAnalytics, code)
		}
		if strings.Contains(params.Asset, asset.Test) {

		}
	} else if params.Platform == platform.IOS {
		if strings.Contains(params.Asset, asset.Key) {
			code, err := a.IOSUc.GenKey(sheet)
			a.createFile(err, params.Target, IOSAnalyticsKey, code)
		}
		if strings.Contains(params.Asset, asset.Code) {
			code, err := a.IOSUc.GenCode(sheet)
			a.createFile(err, params.Target, IOSAnalytics, code)
		}
		if strings.Contains(params.Asset, asset.Test) {

		}
	} else {
		fmt.Println("# Platform not supported")
	}
}

func (a *analyticsGen) createFile(err error, path string, filename string, data string) {
	if err != nil {
		fmt.Println("# Generate:", err)
	} else {
		_, err = a.FileX.CreateFile(path, filename, data)
		if err != nil {
			fmt.Println("# Generate:", err)
		} else {
			fmt.Println("# Generate file " + path + "/" + filename + ": Success")
		}
	}
}

func NewAnalyticsGen(
	fileX filex.FileX,
	androidUc generate.AndroidUseCase,
	flutterUc generate.FlutterUseCase,
	iOSUc generate.IOSUseCase,
) AnalyticsGen {
	return &analyticsGen{
		FileX:     fileX,
		AndroidUc: androidUc,
		FlutterUc: flutterUc,
		IOSUc:     iOSUc,
	}
}
