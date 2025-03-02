package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"permission-cat/internal/datetime"
	"strconv"
	"strings"
	"time"
)

var inputTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time calculation",
	Long:  "time calculation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "the current time",
	Long:  "the current time",
	Run: func(cmd *cobra.Command, args []string) {
		currentTime := datetime.GetNowTime()
		log.Printf(
			"The current time: %s, %d\n",
			currentTime.Format(time.DateTime),
			currentTime.Unix())
	},
}

var calcTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "calculate the time span",
	Long:  "calculate the time span",
	Run: func(cmd *cobra.Command, args []string) {
		var baseTime time.Time
		var layout = time.DateTime

		if inputTime == "" {
			baseTime = datetime.GetNowTime()
		} else {
			var err error

			if !strings.Contains(inputTime, " ") {
				layout = time.DateOnly
			}

			var location, _ = time.LoadLocation("Asia/Shanghai")

			baseTime, err = time.ParseInLocation(layout, inputTime, location)
			if err != nil {
				timeOnly, _ := strconv.Atoi(inputTime)
				baseTime = time.Unix(int64(timeOnly), 0)
			}

			result, err := datetime.GetCalculatedTime(baseTime, duration)
			if err != nil {
				log.Fatalf("calculate the time error: %v\n", err)
			}

			log.Printf("Result: %s\n", result)
		}

	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calcTimeCmd)

	calcTimeCmd.Flags().StringVarP(&inputTime, "base", "b", "", "input base time ")
	calcTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "input duration")
}
