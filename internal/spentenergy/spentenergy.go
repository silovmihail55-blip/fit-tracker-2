package spentenergy

import (
	"errors"
	"time"
)

// Константы для расчётов: метры в км, минуты в часе, коэффициенты длины шага и расхода калорий
const (
	mInKm                      = 1000 // метров в километре
	minInH                     = 60   // минут в часе
	stepLengthCoefficient      = 0.45 // коэффициент длины шага от роста
	walkingCaloriesCoefficient = 0.5  // коэффициент расхода калорий при ходьбе
)

// WalkingSpentCalories рассчитывает калории, сожжённые при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("Количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("Вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("Рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("Длительность должна быть больше 0")
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	return (walkingCaloriesCoefficient * weight * avgSpeed * duration.Minutes()) / minInH, nil
}

// RunningSpentCalories рассчитывает калории, сожжённые при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("Количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("Вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("Рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("Длительность должна быть больше 0")
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	return (weight * avgSpeed * duration.Minutes()) / minInH, nil
}

// MeanSpeed вычисляет среднюю скорость движения (км/ч)
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// Distance вычисляет пройденное расстояние в километрах
func Distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	return float64(steps) * stepLen / mInKm
}
