package cmd

import (
	"fmt"
	"github.com/prongbang/analyticsgen/internal/analyticsgen"
	"github.com/prongbang/analyticsgen/pkg/parameter"
)

func Run(params *parameter.Parameter) error {
	fmt.Println("--> START")
	fmt.Println("# Platform:", params.Platform)

	analyticsGen := analyticsgen.New()
	analyticsGen.Process(params)

	fmt.Println("<-- END")
	return nil
}
