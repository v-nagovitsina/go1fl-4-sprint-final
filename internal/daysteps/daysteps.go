package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, &strconv.NumError{Func: "parsePackage", Num: data, Err: strconv.ErrSyntax}
	}

	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, &strconv.NumError{Func: "parsePackage", Num: stepsStr, Err: strconv.ErrRange}
	}

	durationStr := parts[1]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, &strconv.NumError{Func: "parsePackage", Num: durationStr, Err: strconv.ErrRange}
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}

	// дистанция в метрах
	distanceMeters := float64(steps) * stepLength
	// дистанция в километрах
	distanceKm := distanceMeters / mInKm

	// вычисление калорий через WalkingSpentCalories из spentcalories
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}

	// форматирование строки
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceKm, calories)
}
