package duplicate

import (
	"bufio"
	"errors"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"fmt"
	"strings"
)

var (
	filepath []string
	StartCmd  = &cobra.Command{
		Use:          "duplicate",
		Short:        "Start Duplicate Tool",
		Example:      "mogo duplicate -f=filepath,filepath2",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup() // 初始化配置一类的
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringSliceVar(&filepath, "f", nil, "ex: -f=f1,f2")
}

// os
// ioutil
// strings
// bufio
// cobra 的 slice
func run() error {
	counts := make(map[string]int)
	if (len(filepath) < 1) {
		return errors.New("filepath is not empty")
	}

	for _, f := range filepath {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			return err
		}
		// open file
		f, err := os.Open(f)
		defer f.Close()

		if err != nil {
			return err
		}
		//read file
		countFileLine(f, counts)
	}

	// 一次性读入内存的方式
	for _, ff := range filepath {
		data, err := ioutil.ReadFile(ff)
		if err != nil {
			return err
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	return nil
}

func setup() {

}

func countFileLine(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

