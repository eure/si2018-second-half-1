package user

import (
	"strconv"

	"strings"

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
