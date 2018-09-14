package repositories

import (
	"time"

	"github.com/go-openapi/strfmt"

	"strconv"
	"strings"

	"fmt"

	"github.com/eure/si2018-second-half-1/entities"
	"github.com/eure/si2018-second-half-1/models"
)

type UserRepository struct {
	RootRepository
}

func NewUserRepository(s *Session) *UserRepository {
	return &UserRepository{NewRootRepository(s)}
}

func (r *UserRepository) Create(ent entities.User) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(ent *entities.User) error {
	now := strfmt.DateTime(time.Now())
	s := r.GetSession().Where("id = ?", ent.ID)
	ent.UpdatedAt = now
	if _, err := s.Update(ent); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByUserID(userID int64) (*entities.User, error) {
	var ent = entities.User{ID: userID}

	s := r.GetSession()
	has, err := s.Get(&ent)
	if err != nil {
		return nil, err
	}

	if has {
		return &ent, nil
	}

	return nil, nil
}

func toString(t models.IdealType) string {
	return fmt.Sprintf("age = %v ~ %v, annual_income = %v ~ %v, body_build = %v, drinking = %v, education = %v, height = %v ~ %v, holiday = %v, home_state = %v, job = %v, residence_state = %v, smoking = %v",
		t.Age.From, t.Age.To, t.AnnualIncome.From, t.AnnualIncome.To, t.BodyBuild, t.Drinking, t.Education, t.Height.From, t.Height.To, t.Holiday, t.HomeState, t.Job, t.ResidenceState, t.Smoking)
}

// limit / offset / 検索対象の性別 でユーザーを取得
// idsには取得対象に含めないUserIDを入れる (いいね/マッチ/ブロック済みなど)
func (r *UserRepository) FindWithCondition(limit, offset int, gender string, ids []int64, searchCondition models.IdealType) ([]entities.User, error) {
	fmt.Println("***************** FindWithConditionのpamars *******************")
	fmt.Println("limit", limit, "offset", offset, "gender", gender, "exclude", ids)
	fmt.Println(toString(searchCondition))
	fmt.Println("***************************************************************")

	var users []entities.User

	s := r.GetSession()
	s.Where("gender = ?", gender)

	if len(ids) > 0 {
		s.NotIn("id", ids)
	}

	// 年収の挿入
	if searchCondition.AnnualIncome.From != "" || searchCondition.AnnualIncome.To != "" {
		s.In("annual_income", ReAnnualIncome(searchCondition.Age.From, searchCondition.Age.To))
	}

	// 体型の挿入
	if len(searchCondition.BodyBuild) != 0 {
		s.In("body_build", searchCondition.BodyBuild)
	}

	// お酒を飲むかの挿入
	if len(searchCondition.Drinking) != 0 {
		s.In("drinking", searchCondition.Drinking)
	}

	// 学歴の挿入
	if len(searchCondition.Education) != 0 {
		s.In("education", searchCondition.Education)
	}

	// 身長の挿入
	if searchCondition.Height.From != "" || searchCondition.Height.To != "" {
		s.In("height", ReHeight(searchCondition.Height.From, searchCondition.Height.To))
	}

	// 休日の挿入
	if len(searchCondition.Holiday) != 0 {
		s.In("holiday", searchCondition.Holiday)
	}

	// 出身地の挿入
	if len(searchCondition.HomeState) == 0 {
		s.In("home_state", searchCondition.HomeState)
	}

	// 仕事の挿入
	if len(searchCondition.Job) != 0 {
		s.In("job", searchCondition.Job)
	}

	// 居住地の挿入
	if len(searchCondition.ResidenceState) != 0 {
		s.In("residence_state", searchCondition.ResidenceState)
	}

	// タバコの挿入
	if len(searchCondition.Smoking) != 0 {
		s.In("smoking", searchCondition.Smoking)
	}

	// 年齢の挿入
	if searchCondition.Age.From != "" || searchCondition.Age.To != "" {
		if searchCondition.Age.From != "" {
			s.And("birthday < ?", ReAge(searchCondition.Age.From))
		}

		if searchCondition.Age.To != "" {
			s.And("birthday > ?", ReAge(searchCondition.Age.To))
		}
	}

	s.Limit(limit, offset)
	s.Desc("id")

	err := s.Find(&users)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *UserRepository) FindByIDs(ids []int64) ([]entities.User, error) {
	var users []entities.User

	err := engine.In("id", ids).Find(&users)
	if err != nil {
		return users, err
	}

	return users, nil
}

func ReHeight(from, to string) []string {
	var heights []string
	f, _ := strconv.Atoi(strings.Replace(from, "cm", "", -1))
	t, _ := strconv.Atoi(strings.Replace(to, "cm", "", -1))

	rs := t - f

	heights = append(heights, from)
	for rs > 0 {
		f++
		heights = append(heights, strconv.Itoa(f)+"cm")
		rs--
	}

	return heights
}

func ReAge(age string) time.Time {
	now := time.Now()
	a, _ := strconv.Atoi(strings.Replace(age, "歳", "", -1))

	ageTime := now.AddDate(-a, 0, 0)
	return ageTime
}

func ReAnnualIncome(from, to string) []string {
	var annualIncomes []string
	var f, t int

	f, _ = strconv.Atoi(strings.Replace(from, "万円", "", -1))
	t, _ = strconv.Atoi(strings.Replace(to, "万円", "", -1))

	if from == "" {
		f = 200
		annualIncomes = append(annualIncomes, strconv.Itoa(f)+"万円未満")
	}

	if to == "" {
		t = 3000
	}

	i := f

	for i < t {
		switch {
		case i >= 2000:
			i += 1000
		case i >= 1000:
			i += 500
		case i < 1000:
			i += 200
		}

		annualIncomes = append(annualIncomes, strconv.Itoa(f)+"万円以上"+"〜"+strconv.Itoa(i)+"万円未満")

		switch {
		case f >= 2000:
			f += 1000
		case f >= 1000:
			f += 500
		case f < 1000:
			f += 200
		}
	}

	if to == "" {
		annualIncomes = append(annualIncomes, strconv.Itoa(t)+"万円以上")
	}

	return annualIncomes
}
