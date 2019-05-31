package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// 画像ファイルを開く(元画像)
	src, _ := os.Open("./query.png")
	defer src.Close()

	// デコードしてイメージオブジェクトを準備
	srcImg, _, err := image.Decode(src)
	if err != nil {
		panic(err)
	}
	srcBounds := srcImg.Bounds()

	// 出力用イメージ
	dest := image.NewGray(srcBounds)

	// 色変更
	for v := srcBounds.Min.Y; v < srcBounds.Max.Y; v++ {
		for h := srcBounds.Min.X; h < srcBounds.Max.X; h++ {
			c := color.RGBAModel.Convert(srcImg.At(h, v))
			col := c.(color.RGBA)
			cr := uint8(int(col.R))
			cg := uint8(int(col.G))
			cb := uint8(int(col.B))
			dest.Set(h, v, color.RGBA{cr, cg, cb, col.A})
		}
	}

	// 書き出し用ファイル準備
	outfile, _ := os.Create("out.png")
	defer outfile.Close()
	// 書き出し
	png.Encode(outfile, dest)
}
