//lissajous产生随机利萨茹图形的GIF动画
//执行 go run lissajous.go web ,执行后监听本地8080端口，从浏览器打开 http://127.0.0.1:8080/ 可以看到生成的图像。
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//var palette = []color.Color{color.White, color.Black}  //切片里存黑白两种颜色，切片是color类型
var palette = []color.Color{color.RGBA{0x00, 0x22, 0x33, 0xff}, color.Black}

const (
	whiteIndex = 0 //画板中的第一种颜色
	blackIndex = 1 //画板中的下一种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette) //使用颜色和尺寸创建图
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex) //每次循环设置一些像素为黑色
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img) //将图像追加到GIF帧列表中
	}
	gif.EncodeAll(out, &anim)
}
