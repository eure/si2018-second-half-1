package repositories

import (
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-second-half-1/entities"
	"github.com/eure/si2018-second-half-1/libs/user"
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

// limit / offset / 検索対象の性別 でユーザーを取得
// idsには取得対象に含めないUserIDを入れる (いいね/マッチ/ブロック済みなど)
func (r *UserRepository) FindWithCondition(limit, offset int, gender string, ids []int64, searchCondition models.IdealType) ([]entities.User, error) {
	var users []entities.User

	s := r.GetSession()
	s.Where("gender = ?", gender)

	// if len(ids) > 0 {
	// IDの挿入
	s.NotIn("id", ids)
	//}

	// 年収の挿入
	if searchCondition.AnnualIncome.From != "" || searchCondition.AnnualIncome.To != "" {
		s.In("", user.ReAnnualIncome(searchCondition.Age.From, searchCondition.Age.To))
	}

	// 体型の挿入
	if len(searchCondition.BodyBuild) == 0 {
		s.In("", searchCondition.BodyBuild)
	}

	// お酒を飲むかの挿入
	if len(searchCondition.Drinking) == 0 {
		s.In("", searchCondition.Drinking)
	}

	// 学歴の挿入
	if len(searchCondition.Education) == 0 {
		s.In("", searchCondition.Education)
	}

	// 身長の挿入
	if searchCondition.Height.From != "" || searchCondition.Height.To != "" {
		s.In("", user.ReAnnualIncome(searchCondition.Height.From, searchCondition.Age.To))
	}

	// 休日の挿入
	if len(searchCondition.Holiday) == 0 {
		s.In("", searchCondition.Holiday)
	}

	// 出身地の挿入
	if len(searchCondition.HomeState) == 0 {
		s.In("", searchCondition.HomeState)
	}

	// 仕事の挿入
	if len(searchCondition.Job) == 0 {
		s.In("job", searchCondition.Job)
	}

	// 居住地の挿入
	if len(searchCondition.ResidenceState) == 0 {
		s.In("", searchCondition.ResidenceState)
	}

	// タバコの挿入
	if len(searchCondition.Smoking) == 0 {
		s.In("", searchCondition.Smoking)
	}

	// 年齢の挿入
	if searchCondition.Age.From != "" || searchCondition.Age.To != "" {
		if searchCondition.Age.From != "" {
			s.And("birthday < ?", user.ReAge(searchCondition.Age.From))
		}

		if searchCondition.Age.To != "" {
			s.And("birthday > ?", user.ReAge(searchCondition.Age.To))
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
