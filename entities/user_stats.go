package entities

import (
	"math"
	"strconv"
	"time"

	"github.com/eure/si2018-second-half-1/models"
	"github.com/go-openapi/strfmt"
)

type UserStats struct {
	UserID          int64           `xorm:"user_id"`
	Birthday        float64         `xorm:"birthday"`
	HomeStateX      float64         `xorm:"home_state_x"`
	HomeStateY      float64         `xorm:"home_state_y"`
	ResidenceStateX float64         `xorm:"residence_state_x"`
	ResidenceStateY float64         `xorm:"residence_state_y"`
	Education       float64         `xorm:"education"`
	AnnualIncome    float64         `xorm:"annual_income"`
	Height          float64         `xorm:"height"`
	BodyBuild       float64         `xorm:"body_build"`
	Smoking         float64         `xorm:"Smoking"`
	Drinking        float64         `xorm:"drinking"`
	HolidayWeekday  float64         `xorm:"holiday_weekday"`
	HolidayWeekend  float64         `xorm:"holiday_weekend"`
	HolidayRandom   float64         `xorm:"holiday_random"`
	HolidayOthers   float64         `xorm:"holiday_others"`
	JobDoctor       float64         `xorm:"job_doctor"`
	JobOffice       float64         `xorm:"job_office"`
	JobClerk        float64         `xorm:"job_clerk"`
	CreatedAt       strfmt.DateTime `xorm:"created_at"`
	UpdatedAt       strfmt.DateTime `xorm:"updated_at"`
}

var DrinkingChoices = map[string]float64{"飲まない": 0, "ときどき飲む": 0.5, "飲む": 1}
var EducationChoices = map[string]float64{"飲まない": 0, "ときどき飲む": 0.5, "飲む": 1}
var BodyBuildChoices = map[string]float64{"飲まない": 0, "ときどき飲む": 0.5, "飲む": 1}
var SmokingChoices = map[string]float64{"飲まない": 0, "ときどき飲む": 0.5, "飲む": 1}

func getNearChoices(average float64, choices map[string]float64) []string {
	return []string{}
}

type Range struct {
	From int
	To   int
}

func getRoundedRange(average float64, lower, upper, unit int64) Range {
	return Range{From: 0, To: 0}
}

func getNearHeight(average float64) models.IdealTypeHeight {
	r := getRoundedRange(average, 135, 200, 5)
	return models.IdealTypeHeight{From: strconv.Itoa(r.From) + "cm", To: strconv.Itoa(r.To) + "cm"}
}

func getNearState(x, y float64) []string {
	return []string{"東京"}
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func getNearAnnualIncome(average float64) models.IdealTypeAnnualIncome {
	return models.IdealTypeAnnualIncome{From: "500万円", To: "500万円"}
}

func getNearAge(average float64) models.IdealTypeAge {
	now := time.Now()
	birth := time.Unix(int64(round(average)), 0)
	span := now.Sub(birth)
	years := span.Hours() / (24 * 365)
	r := getRoundedRange(years, 18, 65, 1)
	return models.IdealTypeAge{From: strconv.Itoa(r.From) + "歳", To: strconv.Itoa(r.To) + "歳"}
}

func getModeJob([]float64) []string {
	return []string{"会社員"}
}

func getModeHoliday([]float64) []string {
	return []string{"土日"}
}

func (u UserStats) Build() models.IdealType {
	age := getNearAge(u.Birthday)
	height := getNearHeight(u.Height)
	income := getNearAnnualIncome(u.AnnualIncome)
	return models.IdealType{
		Drinking:       getNearChoices(u.Drinking, DrinkingChoices),
		Education:      getNearChoices(u.Education, EducationChoices),
		BodyBuild:      getNearChoices(u.BodyBuild, BodyBuildChoices),
		Smoking:        getNearChoices(u.Smoking, SmokingChoices),
		Age:            &age,
		Height:         &height,
		AnnualIncome:   &income,
		HomeState:      getNearState(u.HomeStateX, u.HomeStateY),
		ResidenceState: getNearState(u.ResidenceStateX, u.ResidenceStateY),
		Job:            getModeJob([]float64{u.JobClerk, u.JobDoctor, u.JobOffice}),
		Holiday:        getModeHoliday([]float64{u.HolidayWeekend, u.HolidayWeekday, u.HolidayRandom, u.HolidayOthers})}
}
