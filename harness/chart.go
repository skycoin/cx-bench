package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// chartRow holds the best CX, Lua, Go and C time per round of one
// benchmark over the C time. 0 means no value. notRun names the languages
// whose forms leave the benchmark out on purpose, "CX" for example.
type chartRow struct {
	bench  string
	ratios [4]float64
	notRun string
}

var (
	chartSurface   = color.RGBA{0xfc, 0xfc, 0xfb, 0xff}
	chartText      = color.RGBA{0x0b, 0x0b, 0x0b, 0xff}
	chartMuted     = color.RGBA{0x52, 0x51, 0x4e, 0xff}
	chartGrid      = color.RGBA{0xe4, 0xe3, 0xdf, 0xff}
	chartSeries    = [4]color.RGBA{{0x2a, 0x78, 0xd6, 0xff}, {0xeb, 0x68, 0x34, 0xff}, {0x1b, 0xaf, 0x7a, 0xff}, {0xed, 0xa1, 0x00, 0xff}}
	chartSeriesKey = [4]string{"best CX", "best Lua", "best Go", "C"}
)

// chartWrite draws a dot plot: a row per benchmark, the time per round of
// each language over C on a log scale.
func chartWrite(path string, rows []chartRow, title string) error {
	const (
		width, rowHeight, top, left, right = 1000, 34, 92, 150, 40
	)
	height := top + rowHeight*len(rows) + 60
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	fill(img, img.Bounds(), chartSurface)
	low, high := 1.0, 1.0
	for _, row := range rows {
		for _, ratio := range row.ratios {
			if ratio > 0 {
				low, high = math.Min(low, ratio), math.Max(high, ratio)
			}
		}
	}
	lowExp, highExp := math.Floor(math.Log2(low)), math.Ceil(math.Log2(high))
	if highExp == lowExp {
		highExp++
	}
	plotWidth := float64(width - left - right)
	x := func(ratio float64) int {
		return left + int((math.Log2(ratio)-lowExp)/(highExp-lowExp)*plotWidth)
	}
	text(img, 20, 28, title+": time per round relative to C, log2 scale (lower is faster)", chartText)
	legendX := 20
	for i, key := range chartSeriesKey {
		dot(img, legendX+6, 52, 5, chartSeries[i])
		text(img, legendX+16, 56, key, chartText)
		legendX += 30 + 8*len(key)
	}
	bottom := top + rowHeight*len(rows)
	labelStep := math.Ceil((highExp - lowExp + 1) / 16)
	for e := lowExp; e <= highExp; e++ {
		ratio := math.Pow(2, e)
		gx := x(ratio)
		shade := chartGrid
		if math.Mod(e-lowExp, labelStep) == 0 || e == 0 {
			shade = color.RGBA{0xc8, 0xc7, 0xc2, 0xff}
			label := fmt.Sprintf("%gx", ratio)
			text(img, gx-4*len(label), bottom+18, label, chartMuted)
		}
		fill(img, image.Rect(gx, top-6, gx+1, bottom), shade)
	}
	for i, row := range rows {
		cy := top + rowHeight*i + rowHeight/2
		if row.notRun != "" {
			text(img, 20, cy, row.bench, chartText)
			text(img, 20, cy+13, row.notRun+" not run", chartMuted)
		} else {
			text(img, 20, cy+4, row.bench, chartText)
		}
		fill(img, image.Rect(left, cy, width-right, cy+1), chartGrid)
		for s := 3; s >= 0; s-- {
			if row.ratios[s] > 0 {
				dot(img, x(row.ratios[s]), cy, 7, chartSurface)
				dot(img, x(row.ratios[s]), cy, 5, chartSeries[s])
			}
		}
	}
	text(img, left, bottom+40, "Each dot is the fastest form of that language for the benchmark, see results.md for the forms and the noise.", chartMuted)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}

func fill(img *image.RGBA, r image.Rectangle, c color.RGBA) {
	r = r.Intersect(img.Bounds())
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			img.SetRGBA(x, y, c)
		}
	}
}

func dot(img *image.RGBA, cx, cy, radius int, c color.RGBA) {
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				img.SetRGBA(cx+x, cy+y, c)
			}
		}
	}
}

func text(img *image.RGBA, x, y int, s string, c color.RGBA) {
	drawer := font.Drawer{Dst: img, Src: image.NewUniform(c), Face: basicfont.Face7x13, Dot: fixed.P(x, y)}
	drawer.DrawString(s)
}
