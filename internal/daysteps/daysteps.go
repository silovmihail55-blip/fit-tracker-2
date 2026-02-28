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
	var ers error
	var step int
	var dur time.Duration

	spl := strings.Split(datastring, ",")
	if len(spl) != 2 {
		return errors.New("Данные неправильного формата, должно быть: 678,0h50m")
	}

	step, ers = strconv.Atoi(spl[0])
	if ers != nil {
		return ers
	}
	if step <= 0 {
		return errors.New("Количество шагов должно быть больше 0")
	}

	dur, ers = time.ParseDuration(spl[1])
	if ers != nil {
		return ers
	}
	if dur <= 0 {
		return errors.New("Длительность должна быть больше 0")
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
