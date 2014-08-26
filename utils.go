package main

import (
	// "fmt"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/fatih/color"
	"os"
	"text/tabwriter"
)

func loadLogger() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")

	if err != nil {
		// log.Errorf("Erreur de loading de seelog.xml : %s", err)
		log.Info("Fallback to default config.")
		defaultConfig := `
			<seelog>
				<outputs formatid="common">
					<console/>
					<filter levels="error, critical">
						<file path="errors.log" formatid="critical"/>
					</filter>
				</outputs>
				<formats>
					<format id="common" format="[%LEV] %Msg"/>
					<format id="critical" format="%Time %Date %RelFile %Func %Msg %n"/>
					<format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
				</formats>
			</seelog>
		`
		logger, _ = log.LoggerFromConfigAsBytes([]byte(defaultConfig))
	}

	log.ReplaceLogger(logger)
}

// func writeColoredHeader() {
// func writeColoredHeader(firstWord string, a ...interface{}) {
// 	// fmt.Println(a...)
// 	fmt.Println(w, infoBold(firstWord), "\t", info("Path with namespace"), "\t", info("Web URL"))
// }

// func infoBold(texte string) (int, error) {
// 	return color.New(color.FgYellow, color.BgBlack, color.Bold).Println(texte)
// }

func getColors() (
	info func(a ...interface{}) string,
	infoBold func(a ...interface{}) string,
	w *tabwriter.Writer,
) {
	info = color.New(color.FgYellow, color.BgBlack).SprintFunc()
	infoBold = color.New(color.FgYellow, color.BgBlack, color.Bold).SprintFunc()

	w = new(tabwriter.Writer)
	// Format in space-separated columns of minimal width 8
	w.Init(os.Stdout, 8, 0, 3, ' ', 0)

	return info, infoBold, w
}

func getTabWriter() (w *tabwriter.Writer) {
	w = new(tabwriter.Writer)
	// Format in space-separated columns of minimal width 8
	w.Init(os.Stdout, 8, 0, 3, ' ', 0)

	return w
}

func printLine(w *tabwriter.Writer, cells ...interface{}) {
	normal := color.New(color.FgWhite, color.BgBlack).SprintFunc()
	normalBold := color.New(color.FgWhite, color.BgBlack, color.Bold).SprintFunc()

	line := ""
	for i, v := range cells {
		if i == 0 {
			line += normalBold(v) + "\t"
		} else {
			line += normal(v) + "\t"
		}
	}
	fmt.Fprintln(w, line)
}

func printHeader(w *tabwriter.Writer, cells ...interface{}) {
	info := color.New(color.FgYellow, color.BgBlack).SprintFunc()
	infoBold := color.New(color.FgYellow, color.BgBlack, color.Bold).SprintFunc()

	header := ""
	for i, v := range cells {
		if i == 0 {
			header += infoBold(v) + "\t"
		} else {
			header += info(v) + "\t"
		}
	}
	fmt.Fprintln(w, header)
}
