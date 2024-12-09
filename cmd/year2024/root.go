package year2024

import (
	"aoc/cmd/year2024/day1"
	"aoc/cmd/year2024/day2"
	"aoc/cmd/year2024/day3"
	"aoc/cmd/year2024/day4"
	"aoc/cmd/year2024/day5"
	"aoc/cmd/year2024/day6"
	"aoc/cmd/year2024/day7"
	"aoc/cmd/year2024/day8"
	"aoc/cmd/year2024/day9"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2024",
	Short: "2024 - Pew Pew",
	Long:  `2024 is a command line utility to super charge your Advent of Code experience`,
}

func init() {
	Cmd.AddCommand(day1.Cmd)
	Cmd.AddCommand(day2.Cmd)
	Cmd.AddCommand(day3.Cmd)
	Cmd.AddCommand(day4.Cmd)
	Cmd.AddCommand(day5.Cmd)
	Cmd.AddCommand(day6.Cmd)
	Cmd.AddCommand(day7.Cmd)
	Cmd.AddCommand(day8.Cmd)
	Cmd.AddCommand(day9.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
