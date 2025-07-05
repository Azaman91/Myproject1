package main

import (
	"slices"
	"sort"
	"strings"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Goals != b.Goals {
			return b.Goals - a.Goals
		}
		return strings.Compare(a.Name, b.Name)
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating != b.Rating {
			if a.Rating > b.Rating {
				return -1
			}
			return 1
		}
		return strings.Compare(a.Name, b.Name)
	})
	return players
}

func gmSort(players []Player) []Player {
	sorts := make([]Player, len(players))
	copy(sorts, players)

	sort.Slice(sorts, func(i, j int) bool {
		rI := calculateGMRatio(sorts[i])
		rJ := calculateGMRatio(sorts[j])
		if rJ != rI {
			return rI > rJ
		}
		return sorts[i].Name < sorts[j].Name
	})

	return sorts
}

func calculateGMRatio(p Player) float64 {
	if p.Misses == 0 {
		return float64(p.Goals) * 100
	}
	return float64(p.Goals) / float64(p.Misses)
}

func calculateRating(p Player) float64 {
	rating := float64(p.Goals) + float64(p.Assists)/2
	if p.Misses != 0 {
		rating /= float64(p.Misses)
	}
	return rating
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{Name: name, Goals: goals, Misses: misses, Assists: assists}
	p.Rating = calculateRating(p)
	return p
}
