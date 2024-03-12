package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func test1() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	//rand.Seed(time.Now().UTC().UnixNano())
	//lissajous(os.Stdout)

	test3()
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func test2() {
	x := 2

	p := &x

	fmt.Println("p", p)
	fmt.Println(*p)

	*p = 3

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(x)

}

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func test3() {
	//flag.Parse()
	//fmt.Print(strings.Join(flag.Args(), *sep))
	//if !*n {
	//	fmt.Println()
	//
	//}
	//p := new(int) // p, *int 类型, 指向匿名的 int 变量
	////p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	//fmt.Println(*p) // "0"
	//fmt.Println(p)  // "0"
	//fmt.Println(&p) // "0"
	//*p = 2          // 设置 int 匿名变量的值为 2
	//fmt.Println(*p) // "2"
	//fmt.Println(p)  // "2"

	y := new(int)
	*y = 1

	p1 := new(bool)
	fmt.Println(*p1)
	*p1 = true
	fmt.Println(*p1)
}

func test4() {

}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
