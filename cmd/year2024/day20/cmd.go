package day20

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"fmt"
	"math"
	"slices"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day20",
	Long:  `day20`,
	Use:   "day20",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, func(s []byte) int { return part1(s, 2, 100) }, "part 1")
	common.Run(parent, command, 1, func(s []byte) int { return part1(s, 20, 100) }, "part 2")
}

func part1(s []byte, cheatLength, minimum int) int {
	g := grid.New(s, ``)

	start := g.FindCell(`S`)
	end := g.FindCell(`E`)

	lastStep := g.BFS(start[0], start[1], end[0], end[1], 0)
	track := map[string]int{}

	for r, row := range g {
		for c, v := range row {
			if v == "." {
				track[grid.Key(r, c)] = -1
			}
		}
	}

	scores := map[int]int{0: len(lastStep.Path)}

	// queue := []cheat{}
	// seen := map[string]bool{}
	queue := []cheat{}
	for i, step := range lastStep.Path {
		r := step[0]
		c := step[1]
		track[grid.Key(r, c)] = i
		queue = append(queue, cheat{R: r, C: c, Path: [][]int{{r, c}}, depth: 1, visited: map[string]bool{}})
	}

	// scores
	// logrus.Infof("%#v", queue)
	var p cheat
	seen := map[string]int{}
	for len(queue) > 0 {
		p, queue = queue[0], queue[1:]
		if p.visited[grid.Key(p.R, p.C)] {
			continue
		}
		p.visited[grid.Key(p.R, p.C)] = true

		for _, dir := range grid.DIR_CROSS {
			nr := p.R + dir[0]
			nc := p.C + dir[1]

			if p.depth > cheatLength {
				continue
			}

			if !inBound(g, nr, nc) {
				continue
			}

			//ensure first step is into a wall
			// if p.depth == 1 && g[nr][nc] != `#` {
			// 	continue
			// }
			// if _, ok := track[grid.Key(nr, nc)]; ok && p.depth > 0 && g[p.R][p.C] == `#` {
			if _, ok := track[grid.Key(nr, nc)]; ok {
				score := (max(track[grid.Key(nr, nc)], track[grid.Key(p.Path[0][0], p.Path[0][1])]) - min(track[grid.Key(nr, nc)], track[grid.Key(p.Path[0][0], p.Path[0][1])])) - distance(p.Path[0][0], p.Path[0][1], nr, nc)
				key := sort([][]int{{nr, nc}, {p.Path[0][0], p.Path[0][1]}})

				if seen[fmt.Sprintf("%#v", key)] == 0 || score < seen[fmt.Sprintf("%#v", key)] {
					seen[fmt.Sprintf("%#v", key)] = score
				}
			}

			queue = append(queue, cheat{
				R:       nr,
				C:       nc,
				depth:   p.depth + 1,
				Path:    append(slices.Clone(p.Path), []int{nr, nc}),
				visited: p.visited,
			})
		}

	}

	score := 0

	for _, v := range seen {
		scores[v]++
		continue
	}

	for k, v := range scores {
		if k >= minimum {
			score += v
		}
	}

	logrus.Infof("%#v", scores)
	return score
}

func part2(s []byte, cheatLength, minimum int) int {

	return part1(s, cheatLength, minimum)
}

func sort(in [][]int) [][]int {
	if in[1][0] < in[0][0] || (in[1][0] == in[0][0] && in[1][1] < in[0][1]) {
		in[0], in[1] = in[1], in[0]
	}

	return in
}

func exists(points [][]int, point []int) bool {
	for _, p := range points {
		if point[0] == p[0] && point[1] == p[1] {
			return true
		}
	}
	return false
}

func inBound(g grid.Strings, r, c int) bool {
	if !g.InBound(r, c) {
		return false
	}

	//TOP row and bottom row are not valid
	// if r == 0 || r == len(g)-1 {
	// 	return false
	// }

	// // outer columns are not valid
	// if c == 0 || c == len(g[0])-1 {
	// 	return false
	// }

	return true
}

type cheat struct {
	R, C    int
	Path    [][]int
	depth   int
	visited map[string]bool
}

func distance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}
