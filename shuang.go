// example.go
package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
)

//func bar(w http.ResponseWriter, _ *http.Request) *charts.Bar{
func bar() *charts.Bar{

	var nameItems = []string{"1999", "2009", "2019"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "三国GDP增长(万亿美元)"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("美国", []float64{9.63, 14.8, 20.41}).
		AddYAxis("俄罗斯", []float64{0.20, 1.23, 1.69}).
		AddYAxis("中国", []float64{1.09, 4.98, 14.09})
	//f, err := os.Create("3country-GDP.html")
	//if err != nil {
	//	log.Println(err)
	//}
	//bar.Render(w, f)

	return bar
}

//func gauge(w http.ResponseWriter, _ *http.Request) *charts.Gauge{
func gauge() *charts.Gauge{
	gauge := charts.NewGauge()
	m := make(map[string]interface{})
	m["内存使用率"] = 0
	gauge.Add("仪表盘", m)
	gauge.SetGlobalOptions(charts.TitleOpts{Title: "爽哥的服务器"})

	// 使用了 JS 函数
	fn := fmt.Sprintf(`setInterval(function () {
        option_%s.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
        myChart_%s.setOption(option_%s, true);
    }, 2000);`, gauge.ChartID, gauge.ChartID, gauge.ChartID)

	gauge.AddJSFuncs(fn)

	//f, err := os.Create("CVM-Memory-Usage.html")
	//if err != nil {
	//	log.Println(err)
	//}
	//gauge.Render(w, f)
	return gauge
}


//func mapChina(w http.ResponseWriter, _ *http.Request) *charts.Map{
func mapChina() *charts.Map{
	mapData := map[string]float32{
		"北京":   float32(rand.Intn(150)),
		"上海":   float32(rand.Intn(150)),
		"深圳":   float32(rand.Intn(150)),
		"辽宁":   float32(rand.Intn(150)),
		"青岛":   float32(rand.Intn(150)),
		"山西":   float32(rand.Intn(150)),
		"陕西":   float32(rand.Intn(150)),
		"乌鲁木齐": float32(rand.Intn(150)),
		"齐齐哈尔": float32(rand.Intn(150)),
	}

	mc := charts.NewMap("china")
	mc.SetGlobalOptions(charts.TitleOpts{Title: "爽哥-城市足迹"})
	mc.Add("the city that I have been to", mapData)

	//f, err := os.Create("map-china.html")
	//if err != nil {
	//	log.Println(err)
	//}
	//mc.Render(w, f)
	return mc
}


//func wordCloud(w http.ResponseWriter, _ *http.Request) *charts.WordCloud{
func wordCloud() *charts.WordCloud{

	var wcData = map[string]interface{}{
		"Sam S Club":               10000,
		"Macys":                    6181,
		"Amy Schumer":              4386,
		"Jurassic World":           4055,
		"Charter Communications":   2467,
		"Chick Fil A":              2244,
		"Planet Fitness":           1898,
		"Pitch Perfect":            1484,
		"Express":                  1689,
		"Home":                     1112,
		"Johnny Depp":              985,
		"Lena Dunham":              847,
		"Lewis Hamilton":           582,
		"KXAN":                     555,
		"Mary Ellen Mark":          550,
		"Farrah Abraham":           462,
		"Rita Ora":                 366,
		"Serena Williams":          282,
		"NCAA baseball tournament": 273,
		"Point Break":              265,
	}
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(charts.TitleOpts{Title: "WordCloud-示例图"})
	wc.Add("wordcloud", wcData, charts.WordCloudOpts{SizeRange: []float32{14, 80}})


	//f, err := os.Create("map-china.html")
	//if err != nil {
	//	log.Println(err)
	//}
	//wc.Render(w, f)
	return wc
}


var rangeColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genSurface3dData1() [][3]interface{} {
	data := make([][3]interface{}, 0)
	for i := -30; i < 30; i++ {
		y := float64(i) / 10
		for j := -30; j < 30; j++ {
			x := float64(j) / 10
			z := math.Sin(x*x+y*y) * x / math.Pi
			data = append(data, [3]interface{}{x, y, z})
		}
	}
	return data
}


// surface3d(w http.ResponseWriter, _ *http.Request) *charts.Surface3D {
func surface3d() *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.TitleOpts{Title: "surface3D-爽哥制图"},
		charts.VisualMapOpts{
			Calculable: true,
			InRange:    charts.VMInRange{Color: rangeColor},
			Max:        3,
			Min:        -3,
		},
	)
	surface3d.AddZAxis("surface3d", genSurface3dData1())

	//f, err := os.Create("surface3d.html")
	//if err != nil {
	//	log.Println(err)
	//}
	//surface3d.Render(w, f)

	return surface3d
}


func all(w http.ResponseWriter, _ *http.Request) {
	p := charts.NewPage()

	p.Add(
		bar(),
		gauge(),
		mapChina(),
		wordCloud(),
		surface3d(),
	)
	f, err := os.Create("page.html")
	if err != nil {
		log.Println(err)
	}
	p.Render(w, f)

}


func main() {
	//http.HandleFunc("/gdp", bar)
	//http.HandleFunc("/memory", gauge)
	//http.HandleFunc("/map", mapChina)
	//http.HandleFunc("/wordcloud", wordCloud)
	//http.HandleFunc("/surface3d", surface3d)
	http.HandleFunc("/chart", all)
	http.ListenAndServe(":8081", nil)
}
