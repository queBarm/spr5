package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("Parse: wrong format !=3")
	}
	//Парсим шаги
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return errors.New("Parsing: wrong steps count: " + err.Error())
	}
	if steps <= 0 {
		return errors.New("Parsing: steps <= 0")
	}
	t.Steps = steps

	//Парсим тип
	t.TrainingType = strings.TrimSpace(parts[1])
	if t.TrainingType == "" {
		return errors.New("Parsing: empty training type")
	}

	//Парсим длительность
	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return errors.New("Parsing: wrong duration: " + err.Error())
	}
	if duration <= 0 {
		return errors.New("Parsing: duration <= 0")
	}
	t.Duration = duration

	return nil

}

func (t Training) ActionInfo() (string, error) {
	//Валидация
	if t.Duration <= 0 {
		return "", errors.New("Parse: duration must be greates than 0")
	}
	if t.Steps <= 0 {
		return "", errors.New("Parse: steps count must be greatest then 0")
	}

	//Получаем дистанцию и скорость
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки: " + t.TrainingType)
	}

	if err != nil {
		return "", err
	}

	//Формируем результат
	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		speed,
		calories,
	)

	return info, nil
}
