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

// DaySteps хранит данные о шагах за день: количество шагов, длительность и личные данные
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку вида "шаги,длительность" и заполняет поля структуры
func (ds *DaySteps) Parse(datastring string) (err error) {
	var step int
	var dur time.Duration

	spl := strings.Split(datastring, ",")
	if len(spl) != 2 {
		return errors.New("invalid data format, should be: 678,0h50m")
	}

	step, err = strconv.Atoi(spl[0]) // Используем err вместо ers
	if err != nil {
		return err
	}
	if step <= 0 {
		return errors.New("number of steps must be greater than zero")
	}

	dur, err = time.ParseDuration(spl[1])
	if err != nil {
		return err
	}
	if dur <= 0 {
		return errors.New("duration must be greater than zero")
	}

	ds.Steps = step
	ds.Duration = dur
	return nil
}

// ActionInfo рассчитывает дистанцию и сожжённые калории, возвращает строку с результатами
func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cal), nil
}
