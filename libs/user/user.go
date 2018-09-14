package user

import (
	"strconv"

	"time"

	"github.com/eure/si2018-second-half-1/entities"
	"github.com/eure/si2018-second-half-1/repositories"
	si "github.com/eure/si2018-second-half-1/restapi/summerintern"
)

func BuildUserEntityByModel(meID int64, p si.PutProfileBody) entities.User {
	return entities.User{
		ID: meID,

		Nickname:       p.Nickname,
		ImageURI:       p.ImageURI,
		Tweet:          p.Tweet,
		Introduction:   p.Introduction,
		ResidenceState: p.ResidenceState,
		HomeState:      p.HomeState,
		Education:      p.Education,
		Job:            p.Job,
		AnnualIncome:   p.AnnualIncome,
		Height:         p.Height,
		BodyBuild:      p.BodyBuild,
		MaritalStatus:  p.MaritalStatus,
		Child:          p.Child,
		WhenMarry:      p.WhenMarry,
		WantChild:      p.WantChild,
		Smoking:        p.Smoking,
		Drinking:       p.Drinking,
		Holiday:        p.Holiday,
		HowToMeet:      p.HowToMeet,
		CostOfDate:     p.CostOfDate,
		NthChild:       p.NthChild,
		Housework:      p.Housework,
	}
}

func SetUsersImage(s *repositories.Session, users []entities.User) ([]entities.User, error) {
	imageMap := make(map[int64]entities.User, len(users))
	var userIDs []int64
	for _, m := range users {
		userIDs = append(userIDs, m.ID)
		imageMap[m.ID] = m
	}

	rp := repositories.NewUserImageRepository(s)
	userImage, err := rp.GetByUserIDs(userIDs)
	if err != nil {
		return nil, err
	}

	userList := make([]entities.User, len(users))
	for _, m := range userImage {
		user := imageMap[m.UserID]
		user.ImageURI = m.Path
		imageMap[m.UserID] = user
	}

	for i, m := range users {
		userList[i] = imageMap[m.ID]
	}
	return userList, err
}

func SetUserImage(s *repositories.Session, user entities.User) (entities.User, error) {
	rp := repositories.NewUserImageRepository(s)
	userImage, err := rp.GetByUserID(user.ID)
	if err != nil {
		return entities.User{}, err
	}

	user.ImageURI = userImage.Path
	return user, err
}

func GetHeight(u *entities.User) float64 {
	num, _ := strconv.Atoi(u.Height[0 : len(u.Height)-2])
	return float64(num)
}

func GetAnnualIncome(u *entities.User) float64 {
	num, _ := strconv.Atoi(u.AnnualIncome[0 : len(u.AnnualIncome)-2])
	return float64(num)
}

func OneHot(value string, items []string) []float64 {
	vec := make([]float64, 0)
	for _, v := range items {
		hot := 0.0
		if value == v {
			hot = 1.0
		}
		vec = append(vec, hot)
	}
	return vec
}

func MakeStat(u *entities.User) entities.UserStats {
	home := entities.CoordinateMap[u.HomeState]
	resid := entities.CoordinateMap[u.ResidenceState]
	hol := OneHot(u.Holiday, entities.Holiday)
	job := OneHot(u.Job, entities.Jobs)
	return entities.UserStats{
		Birthday:        float64(time.Time(u.Birthday).Unix()),
		HomeStateX:      home.Longitude,
		HomeStateY:      home.Latitude,
		ResidenceStateX: resid.Longitude,
		ResidenceStateY: resid.Latitude,
		Education:       entities.EducationChoices[u.Education],
		AnnualIncome:    GetAnnualIncome(u),
		Height:          GetHeight(u),
		BodyBuild:       entities.BodyBuildChoices[u.BodyBuild],
		Smoking:         entities.SmokingChoices[u.Smoking],
		Drinking:        entities.DrinkingChoices[u.Drinking],
		HolidayWeekday:  hol[0],
		HolidayWeekend:  hol[1],
		HolidayRandom:   hol[2],
		HolidayOthers:   hol[3],
		JobEmployee:     job[0],
		JobStudent:      job[1],
		JobCreator:      job[2]}
}
