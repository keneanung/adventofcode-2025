package main

import (
	"bufio"
	"context"
	"os"
	"strconv"

	day1Puzzle1 "github.com/keneanung/adventofcode-2025/day1/puzzle1"
	day1Puzzle2 "github.com/keneanung/adventofcode-2025/day1/puzzle2"
	day2Puzzle1 "github.com/keneanung/adventofcode-2025/day2/puzzle1"
	day2Puzzle2 "github.com/keneanung/adventofcode-2025/day2/puzzle2"
	day3Puzzle1 "github.com/keneanung/adventofcode-2025/day3/puzzle1"
	day3Puzzle2 "github.com/keneanung/adventofcode-2025/day3/puzzle2"
	"github.com/urfave/cli/v3"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func runSolver(day string, puzzle string, solver func([]string) (int64, error)) error {
	lines, err := readLines(day + "/input.txt")
	if err != nil {
		return err
	}
	result, err := solver(lines)
	if err != nil {
		return err
	}
	println("Solution for " + day + " " + puzzle + ": " + strconv.FormatInt(result, 10))
	return nil
}

func main() {
	(&cli.Command{
		Commands: []*cli.Command{
			{
				Name: "day1",
				Commands: []*cli.Command{
					{
						Name: "puzzle1",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day1", "puzzle1", day1Puzzle1.Solve)
						},
					},
					{
						Name: "puzzle2",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day1", "puzzle2", day1Puzzle2.Solve)
						},
					},
				},
			},
			{
				Name: "day2",
				Commands: []*cli.Command{
					{
						Name: "puzzle1",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day2", "puzzle1", day2Puzzle1.Solve)
						},
					},
					{
						Name: "puzzle2",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day2", "puzzle2", day2Puzzle2.Solve)
						},
					},
				},
			},
			{
				Name: "day3",
				Commands: []*cli.Command{
					{
						Name: "puzzle1",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day3", "puzzle1", day3Puzzle1.Solve)
						},
					},
					{
						Name: "puzzle2",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day3", "puzzle2", day3Puzzle2.Solve)
						},
					},
				},
			},
		},
	}).Run(context.Background(), os.Args)
}
