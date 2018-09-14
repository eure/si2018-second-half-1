package stats

import (
	"fmt"

	"github.com/eure/si2018-second-half-1/entities"
	"github.com/eure/si2018-second-half-1/libs/user"
)

func ToString(s entities.UserStats) string {
	return fmt.Sprintf("id = %v, birthday = %v, home_state = (%v, %v), residence_state = (%v, %v), "+
		"education = %v, annual_income = %v, height = %v, body_build = %v, smoking = %v, drinking = %v, "+
		"holiday = (%v, %v, %v, %v), job = (%v, %v, %v)",
		s.UserID, s.Birthday, s.HomeStateX, s.HomeStateY, s.ResidenceStateX, s.ResidenceStateY,
		s.Education, s.AnnualIncome, s.Height, s.BodyBuild, s.Smoking, s.Drinking,
		s.HolidayWeekday, s.HolidayWeekend, s.HolidayRandom, s.HolidayOthers, s.JobEmployee, s.JobCreator, s.JobStudent)
}

// 新しくいいねした user を, すでにある統計 stats に反映して, 新しい統計を返す
func ApplyNewLike(s *entities.UserStats, u *entities.User) entities.UserStats {
	ans := s.Multiply(0.9).Add(user.MakeStat(u), 0.1)
	fmt.Println(ToString(ans))
	return ans
}

// いいねしたユーザー履歴 users から統計をとる
// len(users) >= 10 でなければならない
func GetAverage(users []entities.User) entities.UserStats {
	var s entities.UserStats
	for i, u := range users {
		us := user.MakeStat(&u)
		switch {
		case i == 0:
			s = us.Multiply(7.0 / 85)
		case i < 10:
			s = s.Add(us, (21.0+float64(i))/255)
		default:
			s = s.Multiply(8.0/9).Add(us, 1.0/9)
		}
	}
	fmt.Println(ToString(s))
	return s
}
