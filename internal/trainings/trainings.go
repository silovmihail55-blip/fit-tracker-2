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

// Training хранит данные о тренировке: шаги, тип, длительность и личные данные
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку вида "шаги,тип,длительность" и заполняет поля структуры
func (t *Training) Parse(datastring string) (err error) {
	spl := strings.Split(datastring, ",")
	if len(spl) != 3 {
		return errors.New("Данные неправильного формата, должно быть: 3456,Ходьба,3h00m")
	}

	step, err := strconv.Atoi(spl[0])
	if err != nil {
		return err
	}
	if step <= 0 {
		return errors.New("Количество шагов должно быть больше 0")
	}

	t.Steps = step
	t.TrainingType = spl[1]

	dur, err := time.ParseDuration(spl[2])
	if err != nil {
		return err
	}
	if dur <= 0 {
		return errors.New("Длительность должна быть больше 0")
	}

	t.Duration = dur
	return nil
}

// ActionInfo рассчитывает параметры тренировки и возвращает строку с результатами
func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var cal float64
	var err error

	switch t.TrainingType {
	case "Бег":
		cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	durationHours := fmt.Sprintf("%.2f", t.Duration.Hours())
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %s ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, durationHours, dist, avgSpeed, cal,
	), nil
}
