package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("wrong format: parts != 2")
	}

	stepStr := parts[0]
	durationStr := parts[1]

	// Проверяем наличие пробелов
	if stepStr != strings.TrimSpace(stepStr) || durationStr != strings.TrimSpace(durationStr) {
		return errors.New("wrong format: have space")
	}

	steps, err := strconv.Atoi(stepStr)
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("steps <= 0")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration <= 0")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Вычисляем калории
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Формируем строку
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
