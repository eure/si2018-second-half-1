package stats

import (
	"github.com/eure/si2018-second-half-1/entities"
	"github.com/eure/si2018-second-half-1/libs/user"
)

// 新しくいいねした user を, すでにある統計 stats に反映して, 新しい統計を返す
func ApplyNewLike(s *entities.UserStats, u *entities.User) entities.UserStats {
	return s.Multiply(0.9).Add(user.MakeStat(u).Multiply(0.1))
}

// いいねしたユーザー履歴 users から統計をとる
// len(users) >= 10 でなければならない
func GetAverage(users []entities.User) entities.UserStats {
	var s entities.UserStats
	for i, u := range users {
		us := user.MakeStat(&u)
		switch {
		case i == 0:
			s = us
		case i < 9:
			s = s.Add(us)
		case i == 9:
			s = s.Add(us).Multiply(0.1)
		default:
			s = s.Multiply(0.9).Add(us.Multiply(0.1))
		}
	}
	return s
}
