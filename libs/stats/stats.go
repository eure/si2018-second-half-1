package stats

import (
	"github.com/eure/si2018-second-half-1/entities"
	"github.com/eure/si2018-second-half-1/libs/user"
)

// 新しくいいねした user を, すでにある統計 stats に反映して, 新しい統計を返す
func ApplyNewLike(s *entities.UserStats, u *entities.User) entities.UserStats {
	return s.Multiply(0.9).Add(user.MakeStat(u), 0.1)
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
	return s
}
