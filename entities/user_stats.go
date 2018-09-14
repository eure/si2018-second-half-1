package entities

import (
	"fmt"
	"math"
	"sort"
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
	JobEmployee     float64         `xorm:"job_employee"`
	JobStudent      float64         `xorm:"job_student"`
	JobCreator      float64         `xorm:"job_creator"`
	CreatedAt       strfmt.DateTime `xorm:"created_at"`
	UpdatedAt       strfmt.DateTime `xorm:"updated_at"`
}

func (s UserStats) Multiply(ratio float64) UserStats {
	return UserStats{
		UserID:          s.UserID,
		Birthday:        s.Birthday * ratio,
		HomeStateX:      s.HomeStateX * ratio,
		HomeStateY:      s.HomeStateY * ratio,
		ResidenceStateX: s.ResidenceStateX * ratio,
		ResidenceStateY: s.ResidenceStateY * ratio,
		Education:       s.Education * ratio,
		AnnualIncome:    s.AnnualIncome * ratio,
		Height:          s.Height * ratio,
		BodyBuild:       s.BodyBuild * ratio,
		Smoking:         s.Smoking * ratio,
		Drinking:        s.Drinking * ratio,
		HolidayWeekday:  s.HolidayWeekday * ratio,
		HolidayWeekend:  s.HolidayWeekend * ratio,
		HolidayRandom:   s.HolidayRandom * ratio,
		HolidayOthers:   s.HolidayOthers * ratio,
		JobEmployee:     s.JobEmployee * ratio,
		JobStudent:      s.JobStudent * ratio,
		JobCreator:      s.JobCreator * ratio,
		CreatedAt:       s.CreatedAt,
		UpdatedAt:       s.UpdatedAt,
	}
}

func (s UserStats) Add(u UserStats, ratio float64) UserStats {
	return UserStats{
		UserID:          s.UserID,
		Birthday:        s.Birthday + u.Birthday*ratio,
		HomeStateX:      s.HomeStateX + u.HomeStateX*ratio,
		HomeStateY:      s.HomeStateY + u.HomeStateY*ratio,
		ResidenceStateX: s.ResidenceStateX + u.ResidenceStateX*ratio,
		ResidenceStateY: s.ResidenceStateY + u.ResidenceStateY*ratio,
		Education:       s.Education + u.Education*ratio,
		AnnualIncome:    s.AnnualIncome + u.AnnualIncome*ratio,
		Height:          s.Height + u.Height*ratio,
		BodyBuild:       s.BodyBuild + u.BodyBuild*ratio,
		Smoking:         s.Smoking + u.Smoking*ratio,
		Drinking:        s.Drinking + u.Drinking*ratio,
		HolidayWeekday:  s.HolidayWeekday + u.HolidayWeekday*ratio,
		HolidayWeekend:  s.HolidayWeekend + u.HolidayWeekend*ratio,
		HolidayRandom:   s.HolidayRandom + u.HolidayRandom*ratio,
		HolidayOthers:   s.HolidayOthers + u.HolidayOthers*ratio,
		JobEmployee:     s.JobEmployee + u.JobEmployee*ratio,
		JobStudent:      s.JobStudent + u.JobStudent*ratio,
		JobCreator:      s.JobCreator + u.JobCreator*ratio,
		CreatedAt:       s.CreatedAt,
		UpdatedAt:       s.UpdatedAt,
	}
}

type Coordinate struct {
	Latitude  float64 // 緯度
	Longitude float64 // 経度
}

var DrinkingChoices = map[string]float64{"飲まない": 0, "ときどき飲む": 1.0 / 2, "飲む": 2.0 / 2}
var EducationChoices = map[string]float64{"その他": 0, "高校卒": 1.0 / 4, "短大/専門学校": 2.0 / 4, "大学卒": 3.0 / 4, "大学院卒": 4.0 / 4}
var BodyBuildChoices = map[string]float64{"スリム": 0, "やや細め": 1.0 / 6, "普通": 2.0 / 6, "グラマー": 3.0 / 6, "筋肉質": 4.0 / 6, "ややぽっちゃり": 5.0 / 6, "ぽっちゃり": 6.0 / 6}
var SmokingChoices = map[string]float64{"吸わない": 0, "非喫煙者の前では吸わない": 1.0 / 5, "相手が嫌ならやめる": 2.0 / 5, "ときどき吸う": 3.0 / 5, "吸う（電子タバコ）": 4.0 / 5, "吸う": 5.0 / 5}
var CoordinateMap = map[string]Coordinate{
	"北海道": {43.06417, 141.34694},
	"青森":  {40.82444, 140.74},
	"岩手":  {39.70361, 141.1525},
	"宮城":  {38.26889, 140.87194},
	"秋田":  {39.71861, 140.1025},
	"山形":  {38.24056, 140.36333},
	"福島":  {37.75, 140.46778},
	"茨城":  {36.34139, 140.44667},
	"栃木":  {36.56583, 139.88361},
	"群馬":  {36.39111, 139.06083},
	"埼玉":  {35.85694, 139.64889},
	"千葉":  {35.60472, 140.12333},
	"東京":  {35.68944, 139.69167},
	"神奈川": {35.44778, 139.6425},
	"新潟":  {37.90222, 139.02361},
	"富山":  {36.69528, 137.21139},
	"石川":  {36.59444, 136.62556},
	"福井":  {36.06528, 136.22194},
	"山梨":  {35.66389, 138.56833},
	"長野":  {36.65139, 138.18111},
	"岐阜":  {35.39111, 136.72222},
	"静岡":  {34.97694, 138.38306},
	"愛知":  {35.18028, 136.90667},
	"三重":  {34.73028, 136.50861},
	"滋賀":  {35.00444, 135.86833},
	"京都":  {35.02139, 135.75556},
	"大阪":  {34.68639, 135.52},
	"兵庫":  {34.69139, 135.18306},
	"奈良":  {34.68528, 135.83278},
	"和歌山": {34.22611, 135.1675},
	"鳥取":  {35.50361, 134.23833},
	"島根":  {35.47222, 133.05056},
	"岡山":  {34.66167, 133.935},
	"広島":  {34.39639, 132.45944},
	"山口":  {34.18583, 131.47139},
	"徳島":  {34.06583, 134.55944},
	"香川":  {34.34028, 134.04333},
	"愛媛":  {33.84167, 132.76611},
	"高知":  {33.55972, 133.53111},
	"福岡":  {33.60639, 130.41806},
	"佐賀":  {33.24944, 130.29889},
	"長崎":  {32.74472, 129.87361},
	"熊本":  {32.78972, 130.74167},
	"大分":  {33.23806, 131.6125},
	"宮崎":  {31.91111, 131.42389},
	"鹿児島": {31.56028, 130.55806},
	"沖縄":  {26.2125, 127.68111},
}
var States = []string{
	"北海道", "青森", "岩手", "宮城", "秋田", "山形", "福島", "茨城", "栃木", "群馬", "埼玉", "千葉",
	"東京", "神奈川", "新潟", "富山", "石川", "福井", "山梨", "長野", "岐阜", "静岡", "愛知", "三重",
	"滋賀", "京都", "大阪", "兵庫", "奈良", "和歌山", "鳥取", "島根", "岡山", "広島", "山口", "徳島",
	"香川", "愛媛", "高知", "福岡", "佐賀", "長崎", "熊本", "大分", "宮崎", "鹿児島", "沖縄",
}

func getNearChoices(average float64, choices map[string]float64) []string {
	var left, just, right string
	for k, v := range choices {
		if (left == "" || choices[left] < v) && v < average {
			left = k
		} else if average < v && (right == "" || v < choices[right]) {
			right = k
		} else if average == v {
			just = k
		}
	}
	near := make([]string, 0)
	if left != "" {
		near = append(near, left)
	}
	if right != "" {
		near = append(near, right)
	}
	if just != "" {
		near = append(near, just)
	}
	return near
}

type Range struct {
	From int64
	To   int64
}

func int64ToString(n int64) string {
	return fmt.Sprintf("%d", n)
}

func intmax(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func intmin(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func getRoundedRange(average float64, lower, upper, unit, width int64) Range {
	ind := (average - float64(lower)) / float64(unit)
	floor := int64(math.Floor(ind))
	ceil := int64(math.Ceil(ind))
	bound := (upper - lower) / unit
	if floor == ceil {
		if floor < width {
			return Range{lower, lower + unit*(floor+width)}
		}
		if floor+width >= bound {
			return Range{lower + unit*(bound-width), lower + unit*bound}
		}
		return Range{lower + unit*(floor-width), lower + unit*(floor+width)}
	}
	return Range{lower + unit*intmax(floor-width+1, 0), lower + unit*intmin(ceil+width-1, bound)}
}

func getNearHeight(average float64) models.IdealTypeHeight {
	r := getRoundedRange(average, 135, 200, 5, 2)
	return models.IdealTypeHeight{From: int64ToString(r.From) + "cm", To: int64ToString(r.To) + "cm"}
}

func (p *Coordinate) Distance2(q *Coordinate) float64 {
	return math.Pow(p.Latitude-q.Latitude, 2) + math.Pow(p.Longitude-q.Longitude, 2)
}

func getNearState(x, y float64) []string {
	p := Coordinate{Longitude: x, Latitude: y}
	dist := make(map[string]float64)
	for k, v := range CoordinateMap {
		dist[k] = p.Distance2(&v)
	}
	sort.SliceStable(States, func(i, j int) bool {
		return dist[States[i]] < dist[States[j]]
	})
	ans := make([]string, 10)
	copy(ans, States[:8])
	ans[8] = "東京"
	ans[9] = "大阪"
	return ans
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

var IncomeChoices = []int64{200, 400, 600, 800, 1000, 1500, 2000, 3000}

func getNearAnnualIncome(average float64) models.IdealTypeAnnualIncome {
	if average <= 200 {
		return models.IdealTypeAnnualIncome{
			From: "",
			To:   "400万円"}
	}
	if average >= 3000 {
		return models.IdealTypeAnnualIncome{
			From: "3000万円",
			To:   ""}
	}
	var left int64 = 0
	var right int64 = 7
	for k, v := range IncomeChoices {
		fv := float64(v)
		if (left < 0 || IncomeChoices[left] < v) && fv <= average {
			left = int64(k)
		}
		if average <= fv && (right < 0 || v < IncomeChoices[right]) {
			right = int64(k)
		}
	}
	return models.IdealTypeAnnualIncome{
		From: int64ToString(IncomeChoices[intmax(left-1, 0)]) + "万円",
		To:   int64ToString(IncomeChoices[intmin(right+1, 7)]) + "万円"}
}

func getNearAge(average float64) models.IdealTypeAge {
	now := time.Now()
	birth := time.Unix(int64(round(average)), 0)
	span := now.Sub(birth)
	years := span.Hours() / (24 * 365)
	r := getRoundedRange(years, 18, 65, 1, 6)
	return models.IdealTypeAge{From: int64ToString(r.From) + "歳", To: int64ToString(r.To) + "歳"}
}

var JobID = map[string]int64{"会社員": 0, "学生": 1, "クリエイター": 2}
var Jobs = []string{"会社員", "学生", "クリエイター"}
var HolidayID = map[string]int64{"平日": 0, "土日": 1, "不定期": 2, "その他": 3}
var Holiday = []string{"平日", "土日", "不定期", "その他"}

func getModeJob(freq [3]float64) []string {
	sort.SliceStable(Jobs, func(i, j int) bool {
		return freq[JobID[Jobs[i]]] > freq[JobID[Jobs[j]]]
	})
	ans := make([]string, 2)
	copy(ans, Jobs[:2])
	return ans
}

func getModeHoliday(freq [4]float64) []string {
	sort.SliceStable(Holiday, func(i, j int) bool {
		return freq[HolidayID[Holiday[i]]] > freq[HolidayID[Holiday[j]]]
	})
	ans := make([]string, 2)
	copy(ans, Holiday[:2])
	return ans
}

func (s UserStats) Build() models.IdealType {
	age := getNearAge(s.Birthday)
	height := getNearHeight(s.Height)
	income := getNearAnnualIncome(s.AnnualIncome)
	return models.IdealType{
		Drinking:       getNearChoices(s.Drinking, DrinkingChoices),
		Education:      getNearChoices(s.Education, EducationChoices),
		BodyBuild:      getNearChoices(s.BodyBuild, BodyBuildChoices),
		Smoking:        getNearChoices(s.Smoking, SmokingChoices),
		Age:            &age,
		Height:         &height,
		AnnualIncome:   &income,
		HomeState:      getNearState(s.HomeStateX, s.HomeStateY),
		ResidenceState: getNearState(s.ResidenceStateX, s.ResidenceStateY),
		Job:            getModeJob([3]float64{s.JobEmployee, s.JobStudent, s.JobCreator}),
		Holiday:        getModeHoliday([4]float64{s.HolidayWeekday, s.HolidayWeekend, s.HolidayRandom, s.HolidayOthers})}
}
