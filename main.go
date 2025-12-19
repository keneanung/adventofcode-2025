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
	day4Puzzle1 "github.com/keneanung/adventofcode-2025/day4/puzzle1"
	day4Puzzle2 "github.com/keneanung/adventofcode-2025/day4/puzzle2"
	day5Puzzle1 "github.com/keneanung/adventofcode-2025/day5/puzzle1"
	day5Puzzle2 "github.com/keneanung/adventofcode-2025/day5/puzzle2"
	day6Puzzle1 "github.com/keneanung/adventofcode-2025/day6/puzzle1"
	day6Puzzle2 "github.com/keneanung/adventofcode-2025/day6/puzzle2"
	day7Puzzle1 "github.com/keneanung/adventofcode-2025/day7/puzzle1"
	day7Puzzle2 "github.com/keneanung/adventofcode-2025/day7/puzzle2"
	day8Puzzle1 "github.com/keneanung/adventofcode-2025/day8/puzzle1"
	day8Puzzle2 "github.com/keneanung/adventofcode-2025/day8/puzzle2"
	day9Puzzle1 "github.com/keneanung/adventofcode-2025/day9/puzzle1"
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
			{
				Name: "day4",
				Commands: []*cli.Command{
					{
						Name: "puzzle1",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day4", "puzzle1", day4Puzzle1.Solve)
						},
					},
					{
						Name: "puzzle2",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day4", "puzzle2", day4Puzzle2.Solve)
						},
					},
				},
			},
			{
				Name: "day5",
				Commands: []*cli.Command{
					{
						Name: "puzzle1",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day5", "puzzle1", day5Puzzle1.Solve)
						},
					},
					{
						Name: "puzzle2",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day5", "puzzle2", day5Puzzle2.Solve)
						},
					},
				},
			},
			{
				Name: "day6",
				Commands: []*cli.Command{
					{
						Name: "puzzle1",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day6", "puzzle1", day6Puzzle1.Solve)
						},
					},
					{
						Name: "puzzle2",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day6", "puzzle2", day6Puzzle2.Solve)
						},
					},
				},
			},
			{
				Name: "day7",
				Commands: []*cli.Command{
					{
						Name: "puzzle1",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day7", "puzzle1", day7Puzzle1.Solve)
						},
					},
					{
						Name: "puzzle2",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day7", "puzzle2", day7Puzzle2.Solve)
						},
					},
				},
			},
			{
				Name: "day8",
				Commands: []*cli.Command{
					{
						Name: "puzzle1",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day8", "puzzle1", day8Puzzle1.Solve)
						},
					},
					{
						Name: "puzzle2",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day8", "puzzle2", day8Puzzle2.Solve)
						},
					},
				},
			},
			{
				Name: "day9",
				Commands: []*cli.Command{
					{
						Name: "puzzle1",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day9", "puzzle1", day9Puzzle1.Solve)
						},
					},
					{
						Name: "puzzle2",
						Action: func(ctx context.Context, c *cli.Command) error {
							return runSolver("day8", "puzzle2", day8Puzzle2.Solve)
						},
					},
				},
			},
		},
	}).Run(context.Background(), os.Args)
}
