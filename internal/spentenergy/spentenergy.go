package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	//Валидируем входные параметры
	if steps <= 0 {
		return 0, errors.New("walkingSpentCalories: steps must be greater than 0")
	}
	if weight <= 0 {
		return 0, errors.New("walkingSpentCalories: weight must be greater than 0")
	}
	if height <= 0 {
		return 0, errors.New("walkingSpentCalories: height must be greater than 0")
	}
	if duration <= 0 {
		return 0, errors.New("walkingSpentCalories: duration must be greater than 0")
	}

	//Получаем среднюю скорость
	speed := MeanSpeed(steps, height, duration)

	//Переводим в минуты
	minutes := duration.Minutes()

	//Считаем калории
	calories := (weight * speed * minutes) / minInH

	//Коэф хотьбы
	calories *= walkingCaloriesCoefficient

	//Возвращаем каолрии
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	//Валидируем входные параметры
	if steps <= 0 {
		return 0, errors.New("runningSpentCalories: steps must be greater than 0")
	}
	if weight <= 0 {
		return 0, errors.New("runningSpentCalories: weight must be greater than 0")
	}
	if height <= 0 {
		return 0, errors.New("runningSpentCalories: height must be greater than 0")
	}
	if duration <= 0 {
		return 0, errors.New("runningSpentCalories: duration must be greater than 0")
	}

	//Получаем среднюю скорость
	speed := MeanSpeed(steps, height, duration)

	//Получаем длительность в минутах
	minutes := duration.Minutes()

	//Считаем калории
	calories := (weight * speed * minutes) / minInH

	//Возвращаем каллории
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	//Проверяем на отрицательные шаги и длительность
	if steps < 0 || duration <= 0 {
		return 0
	}

	//Дистанция в км
	distance := Distance(steps, height)

	//Переводи длительность в часы
	hours := duration.Hours()

	//Средняя скорость км/ч
	speed := distance / hours

	//Возвращаем среднюю скорость
	return speed

}

func Distance(steps int, height float64) float64 {

	//Длина одного шага
	stepLength := height * stepLengthCoefficient

	//Дистанция в метрах
	distanceInMeters := float64(steps) * stepLength

	//Переводим в км
	distanceInKm := distanceInMeters / mInKm

	//Возвращаем дистанцию в км
	return distanceInKm
}
