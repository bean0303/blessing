package main

import (
	"fmt"

	"github.com/bean0303/blessing/wechat"
	"github.com/fogleman/gg"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 0, 0)
	dc.Clear()
	dc.SetRGB(1, 1, 0)
	if err := dc.LoadFontFace("../bytedance/千图小兔体.ttf", 120); err != nil {
		panic(err)
	}
	var name, s, bless string
	fmt.Println("请输入姓名，考试科目和祝福语：")
	fmt.Scanf("%s %s %s", &name, &s, &bless)
	//name, s, bless = "豆豆", "语文", "旗开得胜！"
	s = fmt.Sprintf("%v考试%v", s, bless)

	ellipsisWidth, _ := dc.MeasureString("...")

	maxTextWidth := S * 0.8
	lineSpace := 80.0
	maxLine := int(S / (dc.FontHeight() + lineSpace))

	line := 0
	lineTexts := make([]string, 0)
	for len(s) > 0 {
		line++
		if line > maxLine {
			break
		}
		if line == maxLine {
			sw, _ := dc.MeasureString(s)
			if sw > maxTextWidth {
				maxTextWidth -= ellipsisWidth
			}
		}
		lineText := TruncateText(dc, s, maxTextWidth)
		if line == maxLine && len(lineText) < len(s) {
			lineText += "..."
		}
		lineTexts = append(lineTexts, lineText)
		if len(lineText) >= len(s) {
			break
		}
		s = s[len(lineText):]
	}

	lineY := (S - dc.FontHeight()*float64(len(lineTexts)) - lineSpace*float64(len(lineTexts)-1)) / 2
	lineY -= dc.FontHeight()

	blessHead := "预祝" + name + "同学"
	var start float64 = float64(100)
	dc.DrawString(blessHead, start, lineY)
	lineY += dc.FontHeight() + lineSpace*2

	for _, text := range lineTexts {
		sWidth, _ := dc.MeasureString(text)
		lineX := (S - sWidth) / 2
		dc.DrawString(text, lineX, lineY)
		lineY += dc.FontHeight() + lineSpace
	}

	sign := "--琦飞"
	sWidth, _ := dc.MeasureString(sign)
	lineX := S - sWidth - 100
	lineY += lineSpace
	dc.DrawString(sign, lineX, lineY)
	_ = dc.SavePNG("out.png")
	fmt.Printf("Send massage to %v\n", name)
	wechat := wechat.Wechat{}
	wechat.Login()
	friend := wechat.Search(name)
	fmt.Printf("friend：%v send succeed", friend)

	wechat.SendImageMessage(friend, "out.png")
}

func TruncateText(dc *gg.Context, originalText string, maxTextWidth float64) string {
	tmpStr := ""
	result := make([]rune, 0)
	for _, r := range originalText {
		tmpStr = tmpStr + string(r)
		w, _ := dc.MeasureString(tmpStr)
		if w > maxTextWidth {
			if len(tmpStr) <= 1 {
				return ""
			} else {
				break
			}
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
