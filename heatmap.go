package main

import "errors"

type HeatMap struct {
	GridLevel    int     `json:"gridLevel"`
	Columns      int     `json:"columns"`
	Rows         int     `json:"rows"`
	MinX         float64 `json:"minX"`
	MinY         float64 `json:"minY"`
	MaxX         float64 `json:"maxX"`
	MaxY         float64 `json:"maxY"`
	CountsInts2D [][]int `json:"counts_ints2D"`
}

func (h *HeatMap) max() int {
	max := 0
	for _, row := range h.CountsInts2D {
		for _, col := range row {
			if max < col {
				max = col
			}
		}
	}
	return max
}

func (h *HeatMap) Add(other *HeatMap) error {
	if h.Columns != other.Columns || h.Rows != other.Columns {
		return errors.New("invalid size")
	}

	for y := 0; y < h.Rows; y++ {
		for x := 0; x < h.Columns; x++ {
			h.CountsInts2D[y][x] += other.CountsInts2D[y][x]
		}
	}

	return nil
}