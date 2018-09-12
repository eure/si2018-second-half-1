package stats

import (
	"strconv"

	"github.com/eure/si2018-second-half-1/entities"
)

func ApplyNewLike(stats *entities.UserStats, user *entities.User) entities.UserStats {
	return entities.UserStats{UserID: stats.UserID, Height: (stats.Height + getHeight(user.Height)) / 2}
}

func getHeight(height string) float64 {
	num, _ := strconv.Atoi(height[0 : len(height)-2])
	return float64(num)
}
