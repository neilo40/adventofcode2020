package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2020/internal/common"
)

//Tile is a tile
type Tile struct {
	Number int64
	Data   []string
	IDs    []int // 4 sides + 4 sides reversed
}

func main() {
	part1()
	//part2()
}

func part1() {
	lines := common.ReadFileString("day20.input")
	tiles := make(map[int64]Tile)
	idMap := make(map[int][]Tile)
	var tile Tile
	for _, l := range lines {
		if l == "" {
			tile.IDs = getIDs(tile.Data)
			for _, id := range tile.IDs {
				_, ok := idMap[id]
				if !ok {
					idMap[id] = make([]Tile, 0)
				}
				idMap[id] = append(idMap[id], tile)
			}
			tiles[tile.Number] = tile
			continue
		}
		if strings.HasPrefix(l, "Tile") {
			tileNum, _ := strconv.ParseInt(l[5:9], 10, 64)
			tile = Tile{tileNum, make([]string, 0, 10), make([]int, 0, 8)}
		} else {
			tile.Data = append(tile.Data, l)
		}
	}

	uniqueEdgeCount := make(map[int64]int)
	for k, v := range idMap {
		if len(v) == 1 {
			fmt.Printf("Tile %d has unique edge (ID: %d)\n", v[0].Number, k)
			_, ok := uniqueEdgeCount[v[0].Number]
			if ok {
				uniqueEdgeCount[v[0].Number]++
			} else {
				uniqueEdgeCount[v[0].Number] = 1
			}
		}
	}

	for k, v := range uniqueEdgeCount {
		if v == 2 {
			fmt.Printf("Tile %d has two unique edges\n", k)
		}
	}

	// need to find two unique edges that are adjacent and on the same flip

	fmt.Printf("Result: \n")
}

func getIDs(data []string) []int {
	ids := make([]int, 0, 8)
	ids = append(ids, genId(data[0]))
	ids = append(ids, genId(reverse(data[0])))
	ids = append(ids, genId(data[9]))
	ids = append(ids, genId(reverse(data[0])))

	left := ""
	right := ""
	for _, r := range data {
		left += string(r[0])
		right += string(r[9])
	}
	ids = append(ids, genId(left))
	ids = append(ids, genId(reverse(left)))
	ids = append(ids, genId(right))
	ids = append(ids, genId(reverse(right)))

	return ids
}

// treat the sequence as binary bits and produce an ID
// . is 0, # is 1
func genId(edge string) (id int) {
	id = 0
	for i, d := range edge {
		if d == '#' {
			id += int(math.Pow(2, float64(i)))
		}
	}
	return
}

func reverse(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}

func part2() {
	//lines := common.ReadFileString("dayX.input")
	fmt.Printf("Result: \n")
}
