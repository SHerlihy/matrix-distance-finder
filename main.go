package main

import (
    "fmt"
    "math"
	"sort"
)

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	var distOrderCells [][]int

	centerCo := []int{rCenter, cCenter}

	lowLimit := cCenter
	topLimit := cols - cCenter -1

	rightLimit := rows - rCenter - 1
	leftLimit := rCenter

	allLimits := []int{rightLimit, leftLimit, lowLimit, topLimit}

	var maxToEdge int


	for idx, limit := range allLimits {
		if idx == 0 || maxToEdge < limit {
			maxToEdge = limit
		}
	}

	coordinatesByHypotenuse := map[float64][][]int{
		0: [][]int{centerCo},
	}

	//loop for y=0
	for x := 1; x <= maxToEdge; x++ {
		hyp := simplifiedDistance(centerCo, []int{rCenter - x, cCenter})

            fmt.Printf("\nFrom 1,1 loop: %v", coordinatesByHypotenuse)
		equidistantCells := [][]int{}

		l := []int{rCenter - x, cCenter}
		r := []int{rCenter + x, cCenter}

		if x <= rightLimit {
			equidistantCells = append(equidistantCells, r)
		}

		if x <= leftLimit {
			equidistantCells = append(equidistantCells, l)
		}

		_, ok := coordinatesByHypotenuse[hyp]

		if !ok {
			coordinatesByHypotenuse[hyp] = equidistantCells
		} else {
			updatedCos := append(coordinatesByHypotenuse[hyp], equidistantCells...)
			coordinatesByHypotenuse[hyp] = updatedCos
		}
            fmt.Printf("\nFrom 1,1 loop: %v", coordinatesByHypotenuse)
	}

	//loop for x =0
	for y := 1; y <= maxToEdge; y++ {
		hyp := simplifiedDistance(centerCo, []int{rCenter, cCenter - y})

		equidistantCells := [][]int{}

		t := []int{rCenter, cCenter + y}
		b := []int{rCenter, cCenter - y}

		if y <= lowLimit {
			equidistantCells = append(equidistantCells, b)
		}

		if y <= topLimit {
			equidistantCells = append(equidistantCells, t)
		}

		_, ok := coordinatesByHypotenuse[hyp]

		if !ok {
			coordinatesByHypotenuse[hyp] = equidistantCells
		} else {
			updatedCos := append(coordinatesByHypotenuse[hyp], equidistantCells...)
			coordinatesByHypotenuse[hyp] = updatedCos
		}
	}

	for i := 1; i <= maxToEdge; i++ {
		for j := 1; j <= maxToEdge; j++ {
			hyp := simplifiedDistance(centerCo, []int{rCenter - i, cCenter - j})


			equidistantCells := [][]int{}

			tl := []int{rCenter - i, cCenter + j}
			tr := []int{rCenter + i, cCenter + j}
			bl := []int{rCenter - i, cCenter - j}
			br := []int{rCenter + i, cCenter - j}

			if i <= rightLimit {
				if j <= lowLimit {
					equidistantCells = append(equidistantCells, br)
				}
				if j <= topLimit {
					equidistantCells = append(equidistantCells, tr)
				}
			}

			if i <= leftLimit {
				if j <= lowLimit {
					equidistantCells = append(equidistantCells, bl)
				}
				if j <= topLimit {
					equidistantCells = append(equidistantCells, tl)
				}
			}
			_, ok := coordinatesByHypotenuse[hyp]

			if !ok {
				coordinatesByHypotenuse[hyp] = equidistantCells
			} else {
				updatedCos := append(coordinatesByHypotenuse[hyp], equidistantCells...)
				coordinatesByHypotenuse[hyp] = updatedCos
			}
		}
	}

	orderedHypotenus := make([]float64, len(coordinatesByHypotenuse))

	i := 0
	for k := range coordinatesByHypotenuse {
		orderedHypotenus[i] = k
		i++
	}

	sort.Float64s(orderedHypotenus)

    fmt.Printf("\n orderedDistances: %v", orderedHypotenus)
    fmt.Printf("\n dist by coos: %v", coordinatesByHypotenuse)

	for _, hyp := range orderedHypotenus {
		currentCoos := coordinatesByHypotenuse[float64(hyp)]

        //ordering for ease of debugging
        sort.Slice(currentCoos, func(p, n int) bool { return currentCoos[p][0] < currentCoos[n][0] })

        fmt.Printf("\noreded dist %v coos: %v", hyp, currentCoos)

		if hyp == 0 {
			distOrderCells = currentCoos
		} else {
			distOrderCells = append(distOrderCells, currentCoos...)
		}
	}

	return distOrderCells
}

func simplifiedDistance(coA []int, coB []int) float64 {
    r1 := float64(coA[0])
    r2 := float64(coB[0])

    c1 := float64(coA[1])
    c2 := float64(coB[1])

    distance :=  math.Abs(r1-r2) + math.Abs(c1-c2)

    fmt.Printf("\nFrom %v to %v dist is %v", coA, coB, distance)
     return distance
}
