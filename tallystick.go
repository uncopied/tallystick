package main

import (
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/eps"
	"github.com/tdewolff/canvas/pdf"
	"github.com/tdewolff/canvas/rasterizer"
	"github.com/tdewolff/canvas/svg"
	qrcode "github.com/uncopied/go-qrcode"
	"log"
	"math"
)

const (
	fontFamily              = "Montserrat"
	fontFileRegular         = "fonts/Montserrat-Regular.ttf"
	fontFileBold            = "fonts/Montserrat-Bold.ttf"
	fontSize                = 8.0
	pageWidth       float64 = 297 //210 // A4 = 297 x 210
	pageHeight      float64 = 210 //297 // A4 = 297 x 210

	// blocks
	hBlocks     = 9
	vBlocks     = 7

	// landscape
	tallyWidth  = pageHeight
	tallyHeight = pageWidth / 2
	tallyX      = (pageWidth - tallyWidth) / 2
	tallyY      = (pageHeight - tallyHeight) / 2
	blockWidth  = tallyWidth / hBlocks
	blockHeight = tallyHeight / vBlocks

	// portrait
	tallyXPortrait      = -pageHeight+tallyY
	tallyYPortrait      = tallyX
	blockWidthPortrait  = blockHeight
	blockHeightPortrait = blockWidth

	blockInnerMargin = 0.05
	// draw tally
	drawTallyYesNo = true
	// draw grids
	drawGrid = true
	// colorize
	colorize = true
)

func drawUncopiedLogo(fontFamily *canvas.FontFamily, ctx *canvas.Context, hBlock float64, vBlock float64, rotate bool ) {
	// Draw a comprehensive text box
	face := fontFamily.Face(fontSize, canvas.Black, canvas.FontBold, canvas.FontNormal)
	rich := canvas.NewRichText()
	rich.Add(face, "Mail to :\nun©opied\nPO XXX\n78XXX Versailles\n France")
	//metrics := face.Metrics()

	if colorize {
		ctx.SetFillColor(canvas.Lightpink)
	} else {
		ctx.SetFillColor(canvas.White)
	}
	if rotate {
		// portrait
		ctx.Rotate(-90)
		text := rich.ToText(blockWidthPortrait, blockHeightPortrait, canvas.Center, canvas.Center, 0.0, 0.0)
		ctx.DrawPath(tallyXPortrait+(vBlocks-vBlock-1)*blockWidthPortrait, tallyYPortrait+hBlock*blockHeightPortrait, canvas.RoundedRectangle(blockWidthPortrait, blockHeightPortrait, 1))
		ctx.DrawText(tallyXPortrait+(vBlocks-vBlock-1)*blockWidthPortrait, tallyYPortrait+(hBlock+1)*blockHeightPortrait, text)
		ctx.Rotate(90)
	} else {
		// landscape
		text := rich.ToText(blockWidth, blockHeight, canvas.Center, canvas.Center, 0.0, 0.0)
		ctx.DrawPath(tallyX+hBlock*blockWidth, tallyY+vBlock*blockHeight, canvas.RoundedRectangle(blockWidth, blockHeight, 1))
		ctx.DrawText(tallyX+hBlock*blockWidth, tallyY+(vBlock+1)*blockHeight, text)
	}
}

func drawText(fontFamily *canvas.FontFamily, ctx *canvas.Context, hBlock float64, vBlock float64, wBlock float64, textContent string, rotate bool ) {
	// Draw a comprehensive text box
	face := fontFamily.Face(fontSize, canvas.Black, canvas.FontBold, canvas.FontNormal)
	rich := canvas.NewRichText()
	rich.Add(face, textContent)
	//metrics := face.Metrics()

	if colorize {
		ctx.SetFillColor(canvas.Lightblue)
	} else {
		ctx.SetFillColor(canvas.White)
	}
	if rotate {
		ctx.Rotate(-90)
		text := rich.ToText(wBlock*blockWidthPortrait, blockHeightPortrait, canvas.Center, canvas.Center, 0.0, 0.0)
		ctx.DrawPath(tallyXPortrait+(vBlocks-vBlock-1)*blockWidthPortrait, tallyYPortrait+hBlock*blockHeightPortrait, canvas.RoundedRectangle(blockWidthPortrait*wBlock, blockHeightPortrait, 1))
		ctx.DrawText(tallyXPortrait+(vBlocks-vBlock-1)*blockWidthPortrait, tallyYPortrait+(hBlock+1)*blockHeightPortrait, text)
		ctx.Rotate(90)
	} else {
		text := rich.ToText(wBlock*blockWidth, blockHeight, canvas.Center, canvas.Center, 0.0, 0.0)
		ctx.DrawPath(tallyX+hBlock*blockWidth, tallyY+vBlock*blockHeight, canvas.RoundedRectangle(blockWidth*wBlock, blockHeight, 1))
		ctx.DrawText(tallyX+hBlock*blockWidth, tallyY+(vBlock+1)*blockHeight, text)
	}
}

func drawBarCode128(ctx *canvas.Context, hBlock float64, vBlock float64, wBlock float64, widthRatio float64,  textContent string, rotate bool ) {
	if colorize {
		ctx.SetFillColor(canvas.Darkred)
	} else {
		ctx.SetFillColor(canvas.White)
	}
	if rotate {
		ctx.Rotate(-90)
		ctx.DrawPath(tallyXPortrait+(vBlocks-vBlock-1)*blockWidthPortrait, tallyYPortrait+hBlock*blockHeightPortrait, canvas.Rectangle(blockWidthPortrait*wBlock, blockHeightPortrait*widthRatio))
		ctx.Rotate(90)
	} else {
		ctx.DrawPath(tallyX+hBlock*blockWidth, tallyY+vBlock*blockHeight, canvas.Rectangle(blockWidth*wBlock, blockHeight*widthRatio))
	}
}

func drawQRCode(ctx *canvas.Context, hBlock float64, vBlock float64, content string, rotate bool ) {
	q, err := qrcode.New(content, qrcode.Highest)
	if err != nil {
		log.Fatal(err)
	}
	if rotate {
		ctx.Rotate(-90)
		innerSquareWidth := math.Min(blockWidthPortrait, blockHeightPortrait) * (1 - blockInnerMargin)
		hBlockMargin := (blockWidthPortrait - innerSquareWidth) / 2
		vBlockMargin := (blockHeightPortrait - innerSquareWidth) / 2
		q.DrawQRCode(ctx, tallyXPortrait+(vBlocks-vBlock-1)*blockWidthPortrait+hBlockMargin, tallyYPortrait+hBlock*blockHeightPortrait+vBlockMargin, innerSquareWidth)
		ctx.Rotate(90)
	} else {
		innerSquareWidth := math.Min(blockWidth, blockHeight) * (1 - blockInnerMargin)
		hBlockMargin := (blockWidth - innerSquareWidth) / 2
		vBlockMargin := (blockHeight - innerSquareWidth) / 2
		q.DrawQRCode(ctx, tallyX+hBlock*blockWidth+hBlockMargin, tallyY+vBlock*blockHeight+vBlockMargin, innerSquareWidth)
	}
}

func drawTally(fontFamily *canvas.FontFamily, ctx *canvas.Context) {
	drawBarCode128(ctx,0,6,7, 1.8, "",true)
	drawBarCode128(ctx,7.2,6,7, 1.8, "",true)
	drawBarCode128(ctx,2,0,5, 1, "",false)
	drawBarCode128(ctx,2,2,5, 1, "",false)
	drawBarCode128(ctx,2,4,5, 1, "",false)
	drawBarCode128(ctx,2,6,5, 1, "",false)

	myTextContent := "Origin from Wikidata\nElian Carsenat, 11-2020 (1/15)"
	drawText(fontFamily, ctx, 3, 1, 3, myTextContent, false)
	drawText(fontFamily, ctx, 3, 3, 3,myTextContent, false)
	drawText(fontFamily, ctx, 3, 5, 3,myTextContent,false)

	drawText(fontFamily, ctx, 0, 4, 3,myTextContent, true)

	drawQRCode(ctx, 2, 3, "uncopied-B1", false)
	drawQRCode(ctx, 6, 3, "uncopied-B2", false)

	drawQRCode(ctx, 2, 5, "uncopied-A1", false)
	drawQRCode(ctx, 6, 5, "uncopied-A2", false)

	drawQRCode(ctx, 2, 1, "uncopied-C1", false)
	drawQRCode(ctx, 6, 1, "uncopied-C2", false)

	drawQRCode(ctx, 1, 5, "uncopied-D1A", false)
	drawQRCode(ctx, 1, 3, "uncopied-D1B", false)
	drawQRCode(ctx, 1, 1, "uncopied-D1C", false)

	drawQRCode(ctx, 7, 5, "uncopied-D2A", false)
	drawQRCode(ctx, 7, 3, "uncopied-D2B", false)
	drawQRCode(ctx, 7, 1, "uncopied-D2C", false)

	drawQRCode(ctx, 0, 6, "uncopied-D11", true)
	drawQRCode(ctx, 0, 0, "uncopied-D12", true)

	drawQRCode(ctx, 8, 6, "uncopied-D21", true)
	drawQRCode(ctx, 8, 0, "uncopied-D22", true)

	drawUncopiedLogo(fontFamily, ctx, 0, 1, false)
	drawUncopiedLogo(fontFamily, ctx, 0, 5, true)
	drawUncopiedLogo(fontFamily, ctx, 8, 1, false)
	drawUncopiedLogo(fontFamily, ctx, 8, 5, true)

}
func main() {
	c := canvas.New(pageWidth, pageHeight)
	ctx := canvas.NewContext(c)
	fontFamily := canvas.NewFontFamily(fontFamily)
	fontFamily.Use(canvas.CommonLigatures)
	if err := fontFamily.LoadFontFile(fontFileRegular, canvas.FontRegular); err != nil {
		panic(err)
	}
	if err := fontFamily.LoadFontFile(fontFileBold, canvas.FontBold); err != nil {
		panic(err)
	}

	// paint a grid on landscape
	if drawGrid {
		ctx.SetFillColor(canvas.Lightgray)
		for i := 0; i < hBlocks; i++ {
			for j := 0; j < vBlocks; j++ {
				if (i+j)%2 == 0 {
					ctx.DrawPath(tallyX+float64(i)*blockWidth, tallyY+float64(j)*blockHeight, canvas.RoundedRectangle(blockWidth, blockHeight, 1))
				}
			}
		}
		ctx.SetFillColor(canvas.Darkcyan)
		ctx.Rotate(-90)
		// paint a grid on portrait
		for i := 0; i < vBlocks; i++ {
			for j := 0; j < hBlocks; j++ {
				if (i+j)%2 == 0 {
					//ctx.DrawPath(tallyX+float64(i)*blockWidth, tallyY+float64(j)*blockHeight, canvas.RoundedRectangle(blockWidth, blockHeight, 1))
					ctx.DrawPath(tallyXPortrait+float64(i)*blockWidthPortrait+1, tallyYPortrait+1+float64(j)*blockHeightPortrait, canvas.RoundedRectangle(blockWidthPortrait-2, blockHeightPortrait-2, 1))
				}
			}
		}
		ctx.Rotate(90)
	}

	if drawTallyYesNo {
		drawTally(fontFamily,ctx)
	}

	//p:=canvas.MustParseSVG(qrSVG)
	//ctx.DrawPath(10,10,p)
	c.WriteFile("canvas_out.svg", svg.Writer)
	c.WriteFile("canvas_out.pdf", pdf.Writer)
	c.WriteFile("canvas_out.eps", eps.Writer)
	c.WriteFile("canvas_out.png", rasterizer.PNGWriter(3.2))
}