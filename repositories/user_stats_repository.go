package repositories

import (
	"time"

	"github.com/eure/si2018-second-half-1/entities"
	"github.com/go-openapi/strfmt"
)

type UserStatsRepository struct {
	RootRepository
}

func NewUserStatsRepository(s *Session) UserStatsRepository {
	return UserStatsRepository{NewRootRepository(s)}
}

func (r *UserStatsRepository) Create(ent entities.UserStats) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

func (r *UserStatsRepository) Update(ent entities.UserStats) error {
	now := strfmt.DateTime(time.Now())

	s := r.GetSession().Where("user_id = ?", ent.UserID)
	ent.UpdatedAt = now

	if _, err := s.Update(ent); err != nil {
		return err
	}
	return nil
}

func (r *UserStatsRepository) GetByUserID(userID int64) (*entities.UserStats, error) {
	var ent = entities.UserStats{UserID: userID}

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
