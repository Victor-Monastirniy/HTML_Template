package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Metric описує окремий показник (датчик)
type Metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type Device struct {
	Name        string   `json:"name"`
	IsActive    bool     `json:"is_active"`
	Metrics     []Metric `json:"metrics"`
	CurrentTime string   `json:"current_time"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Device{
		Name:     "Вентилятор",
		IsActive: true,
		Metrics: []Metric{
			{Name: "Напруга", Value: 220, Unit: "В"},
			{Name: "Потужність", Value: 50, Unit: "кВт*год"},
			{Name: "Час роботи", Value: 30, Unit: "хв"},
		},
		CurrentTime: time.Now().Format("02.01.2006 15:04:05"),
	}

	// Завантаження шаблону
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Помилка завантаження шаблону: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Rendering
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Помилка виконання шаблону: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Сервер запущено на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Помилка при запуску сервера: %v\n", err)
	}
}
