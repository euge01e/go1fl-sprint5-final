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

var (
	errInvalidSteps    = errors.New("количество шагов должно быть больше нуля")
	errInvalidWeight   = errors.New("вес должен быть больше нуля")
	errInvalidHeight   = errors.New("рост должен быть больше нуля")
	errInvalidDuration = errors.New("продолжительность должна быть больше нуля")
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validate(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / minInH

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validate(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps < 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}

func validate(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errInvalidSteps
	}
	if weight <= 0 {
		return errInvalidWeight
	}
	if height <= 0 {
		return errInvalidHeight
	}
	if duration <= 0 {
		return errInvalidDuration
	}

	return nil
}
