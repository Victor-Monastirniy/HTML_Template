package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// Metric описує показник
type Metric struct {
	Name  string
	Value float64
	Unit  string
}

// Device описує окремий пристрій
type Device struct {
	Type    string
	Model   string
	IsOn    bool
	Metrics []Metric
}

// PowerUnit — основна структура (те, що йде в шаблон)
type PowerUnit struct {
	UnitName   string
	Location   string
	Devices    []Device // Тепер тут слайс пристроїв
	LastUpdate string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Формуємо дані з декількома пристроями
	data := PowerUnit{
		UnitName:   "Дім",
		Location:   "Запорізька область",
		LastUpdate: time.Now().Format("15:04:05"),
		Devices: []Device{
			{
				Type:  "Плита",
				Model: "АС-2",
				IsOn:  true,
				Metrics: []Metric{
					{Name: "Напруга", Value: 220.5, Unit: "kV"},
					{Name: "Навантаження", Value: 78.0, Unit: "%"},
				},
			},
			{
				Type:  "Холодильник",
				Model: "СН56",
				IsOn:  false,
				Metrics: []Metric{
					{Name: "Температура", Value: 5, Unit: "°C"},
				},
			},
			{
				Type:  "Кондиціонер",
				Model: "НН12",
				IsOn:  true,
				Metrics: []Metric{
					{Name: "Температура", Value: 20.0, Unit: "°C"},
					{Name: "Потужність", Value: 120.0, Unit: "кВт"},
				},
			},
		},
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Сервер стартував на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
