package stats

import (
	"strconv"

	"github.com/eure/si2018-second-half-1/entities"
)

// 新しくいいねした user を, すでにある統計 stats に反映して, 新しい統計を返す
func ApplyNewLike(stats *entities.UserStats, user *entities.User) entities.UserStats {
	return entities.UserStats{UserID: stats.UserID, Height: (stats.Height + getHeight(user.Height)) / 2}
}

// いいねしたユーザー履歴 users から統計をとる
// len(users) >= 10 でなければならない
func GetAverage(users []entities.User) entities.UserStats {
	return entities.UserStats{}
}

func getHeight(height string) float64 {
	num, _ := strconv.Atoi(height[0 : len(height)-2])
	return float64(num)
}
