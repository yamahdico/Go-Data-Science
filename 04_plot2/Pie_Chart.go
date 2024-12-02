package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"math"
)

func main() {
	// داده‌های دسته‌بندی و مقادیر
	data := []struct {
		label string
		value float64
		color draw.Color
	}{
		{"Category A", 40, draw.Red},
		{"Category B", 30, draw.Blue},
		{"Category C", 20, draw.Green},
		{"Category D", 10, draw.Yellow},
	}

	// مجموع کل داده‌ها
	total := 0.0
	for _, d := range data {
		total += d.value
	}

	// ایجاد نمودار و تنظیمات
	p := plot.New()
	p.Title.Text = "Pie Chart Example"
	p.HideAxes() // حذف محورها

	// زاویه‌ها برای هر بخش از دایره
	startAngle := 0.0
	for _, d := range data {
		// محاسبه زاویه بخش
		angle := (d.value / total) * 2 * math.Pi

		// رسم بخش دایره
		arc := plotterArc{
			start: startAngle,
			end:   startAngle + angle,
			color: d.color,
		}
		p.Add(&arc)

		// بروز رسانی زاویه شروع
		startAngle += angle
	}

	// ذخیره نمودار
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "pie_chart_gonum.png"); err != nil {
		panic(err)
	}
}

// تعریف یک نوع برای رسم قوس‌ها
type plotterArc struct {
	start, end float64
	color      draw.Color
}

// پیاده‌سازی روش Draw برای رسم قوس‌ها
func (arc *plotterArc) Plot(c draw.Canvas, plt *plot.Plot) {
	center := c.Center()
	radius := vg.Length(100)

	// ایجاد مسیر دایره‌ای
	path := c.Path()
	path.Move(center)
	path.Arc(center, radius, arc.start, arc.end)
	path.Close()

	// رنگ‌آمیزی بخش
	c.Fill(draw.FillStyle{Color: arc.color, Path: path})
}