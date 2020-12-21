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
	Number        int64
	Data          []string
	IDs           []int // 4 sides
	Flipped       bool
	UniqueEdgeIds []int
}

func main() {
	part1()
	//part2()
}

func part1() {
	lines := common.ReadFileString("day20.input")
	idMap := make(map[int][]Tile)
	tiles := make(map[int64]Tile)
	var tile Tile
	var flippedTile Tile
	for _, l := range lines {
		if l == "" {
			tile.IDs, flippedTile.IDs = getIDs(tile.Data)
			for _, t := range []Tile{tile, flippedTile} {
				for _, id := range t.IDs {
					_, ok := idMap[id]
					if !ok {
						idMap[id] = make([]Tile, 0)
					}
					idMap[id] = append(idMap[id], t)
				}
				tiles[t.Number] = t
			}
			continue
		}
		if strings.HasPrefix(l, "Tile") {
			tileNum, _ := strconv.ParseInt(l[5:9], 10, 64)
			tile = Tile{tileNum, make([]string, 0, 10), make([]int, 0, 4), false, make([]int, 0, 4)}
			flippedTile = Tile{tileNum * 10, make([]string, 0, 10), make([]int, 0, 4), true, make([]int, 0, 4)}
		} else {
			tile.Data = append(tile.Data, l)
		}
	}

	uniqueEdgeCount := make(map[int64]int)
	for k, v := range idMap {
		if len(v) == 1 {
			//fmt.Printf("Tile %d has unique edge (ID: %d)\n", v[0].Number, k)
			t := tiles[v[0].Number]
			t.UniqueEdgeIds = append(t.UniqueEdgeIds, k)
			tiles[v[0].Number] = t
			_, ok := uniqueEdgeCount[v[0].Number]
			if ok {
				uniqueEdgeCount[v[0].Number]++
			} else {
				uniqueEdgeCount[v[0].Number] = 1
			}
		}
	}

	// find tiles with two unique edges
	var cornerSum int64 = 1
	for k, v := range uniqueEdgeCount {
		if v == 2 {
			// of those, find tiles with two adjacent unique edges
			//fmt.Printf("Tile %d has two unique edges\n", k)
			ids := append(tiles[k].IDs, tiles[k].IDs...)
			for i := 1; i < len(ids)-1; i++ {
				if ids[i] == tiles[k].UniqueEdgeIds[0] {
					if ids[i-1] == tiles[k].UniqueEdgeIds[1] || ids[i+1] == tiles[k].UniqueEdgeIds[1] {
						if k < 10000 {
							fmt.Printf("Tile %d has two adjacent unique edges\n", k)
							cornerSum *= k
							break
						}
					}
				}
			}
		}
	}

	fmt.Printf("Result: %d\n", cornerSum)
}

func getIDs(data []string) ([]int, []int) {
	ids := make([]int, 0, 4)
	flippedIds := make([]int, 0, 4)
	left := ""
	right := ""
	for _, r := range data {
		left += string(r[0])
		right += string(r[9])
	}
	ids = append(ids, genID(left))
	ids = append(ids, genID(data[0])) //top
	ids = append(ids, genID(right))
	ids = append(ids, genID(data[9])) //bottom

	flippedIds = append(flippedIds, genID(reverse(left)))
	flippedIds = append(flippedIds, genID(reverse(data[0])))
	flippedIds = append(flippedIds, genID(reverse(right)))
	flippedIds = append(flippedIds, genID(reverse(data[9])))

	return ids, flippedIds
}

// treat the sequence as binary bits and produce an ID
// . is 0, # is 1
func genID(edge string) (id int) {
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
