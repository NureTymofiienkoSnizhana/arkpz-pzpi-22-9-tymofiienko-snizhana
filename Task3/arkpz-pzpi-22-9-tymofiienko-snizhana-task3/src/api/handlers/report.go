package handlers

import (
	"fmt"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func GetPetReport(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewPetReport(r)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if req.StartTime == 0 || req.EndTime == 0 {
		http.Error(w, "Missing start_time or end_time parameters", http.StatusBadRequest)
		return
	}

	startTime := primitive.Timestamp{T: uint32(req.StartTime), I: 0}
	endTime := primitive.Timestamp{T: uint32(req.EndTime), I: 0}

	petsDB := MongoDB(r).Pets()
	healthDB := MongoDB(r).HealthData()

	pet, err := petsDB.Get(req.PetID)
	if err != nil {
		http.Error(w, "Failed to retrieve pet information", http.StatusInternalServerError)
		return
	}

	filter := bson.M{
		"pet_id": pet.ID,
		"time": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	}

	healthData, err := healthDB.GetByFilter(filter)
	if err != nil {
		http.Error(w, "Failed to retrieve health data", http.StatusInternalServerError)
		return
	}

	// Ініціалізація PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("DejaVu", "", "fonts/DejaVuSans.ttf")

	pdf.SetFont("DejaVu", "", 14)
	pdf.AddPage()

	// Заголовок
	pdf.Cell(40, 10, "Звіт")
	pdf.Ln(10)

	pdf.SetFont("DejaVu", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Ім'я: %s", pet.Name))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Вид: %s", pet.Species))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Порода: %s", pet.Breed))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Вік: %d", pet.Age))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Власник: %s", requests.GetUserName(MongoDB(r), pet.OwnerID)))
	pdf.Ln(8)

	// Дані про здоров'я
	pdf.Ln(10)
	pdf.Cell(0, 10, "Дані стану здоров'я:")
	pdf.Ln(10)
	for _, health := range healthData {
		pdf.Cell(0, 10, fmt.Sprintf("Активність: %s", health.Activity))
		pdf.Ln(6)
		pdf.Cell(0, 10, fmt.Sprintf("Сон: %s", health.Sleep))
		pdf.Ln(6)
		pdf.Cell(0, 10, fmt.Sprintf("Прийом їжі: %s", health.Feeding))
		pdf.Ln(6)
		pdf.Cell(0, 10, fmt.Sprintf("Час: %s", time.Unix(int64(health.Time.T), 0).Format("2006-01-02 15:04:05")))
		pdf.Ln(10)
	}

	// Налаштування заголовків для завантаження файлу
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="pet_report.pdf"`)
	w.WriteHeader(http.StatusOK)

	// Виведення PDF у відповідь
	err = pdf.Output(w)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}
}
