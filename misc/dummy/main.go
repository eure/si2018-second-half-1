package main

import (
	"math/rand"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-second-half-1/entities"
	"github.com/eure/si2018-second-half-1/repositories"
)

const (
	firstUserID   = 1
	maleIDStart   = 1
	maleIDEnd     = 5000
	femaleIDStart = 5001
	femaleIDEnd   = 10000
	lastUserID    = 10000

	maleMessageUserID   = 222
	femaleMessageUserID = 5222
	likeUserID          = 5111
	matchUserID         = 5112
)

func main() {
	// dummyToken()             // Token for each Users
	dummyUser() // Male 1-10000, Female 10001-20000
	// dummyImage()             // Images for each Users
	dummyManyMessageCouple() // マッチ & メッセージしてるカップル M222 と F2222
	dummyManyGotLikeUser()   // F11111 に M1〜M100の男性からの被いいね
	dummyManyMatchUser()     // F11112 UID 1〜200の男性が UID 1112の女性とマッチ
}

func dummyManyGotLikeUser() {
	s := repositories.NewSession()
	r := repositories.NewUserLikeRepository(s)

	for i := 1; i <= 100; i++ {
		rand.Seed(time.Now().UnixNano())
		createdDaysAgo := rand.Intn(600)
		minute1 := rand.Intn(1440)
		randTime := strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(minute1) * time.Minute))

		ent := entities.UserLike{
			UserID:    int64(i),
			PartnerID: likeUserID,
			CreatedAt: randTime,
			UpdatedAt: randTime,
		}
		r.Create(ent)
	}
}

func dummyManyMatchUser() {
	s := repositories.NewSession()
	firstLikeDate := strfmt.DateTime(time.Now().AddDate(0, 0, -3))
	responseLikeDate := strfmt.DateTime(time.Now().AddDate(0, 0, -2))

	lr := repositories.NewUserLikeRepository(s)
	mr := repositories.NewUserMatchRepository(s)

	// Male 1-100 & Female 11112
	// =====================================================

	// first like
	for i := 1; i <= 100; i++ {
		ent := entities.UserLike{
			UserID:    int64(i),
			PartnerID: matchUserID,
			CreatedAt: firstLikeDate,
			UpdatedAt: firstLikeDate,
		}
		lr.Create(ent)
	}

	// resp like
	for i := 1; i <= 100; i++ {
		ent := entities.UserLike{
			UserID:    matchUserID,
			PartnerID: int64(i),
			CreatedAt: responseLikeDate,
			UpdatedAt: responseLikeDate,
		}
		lr.Create(ent)
	}

	// match
	for i := 1; i <= 100; i++ {
		ent := entities.UserMatch{
			UserID:    int64(i),
			PartnerID: matchUserID,
			CreatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Duration(i) * time.Minute)), // Paginationのため分刻みでイテレート
			UpdatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Duration(i) * time.Minute)), // Paginationのため分刻みでイテレート
		}
		mr.Create(ent)
	}

	// Male 101-200 & Female 1112
	// =====================================================

	// first like
	for i := 101; i <= 200; i++ {
		ent := entities.UserLike{
			UserID:    matchUserID,
			PartnerID: int64(i),
			CreatedAt: firstLikeDate,
			UpdatedAt: firstLikeDate,
		}
		lr.Create(ent)
	}

	// resp like
	for i := 101; i <= 200; i++ {
		ent := entities.UserLike{
			UserID:    int64(i),
			PartnerID: matchUserID,
			CreatedAt: responseLikeDate,
			UpdatedAt: responseLikeDate,
		}
		lr.Create(ent)
	}

	// match
	for i := 101; i <= 200; i++ {
		ent := entities.UserMatch{
			UserID:    matchUserID,
			PartnerID: int64(i),
			CreatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Duration(i) * time.Minute)), // Paginationのため分刻みでイテレート
			UpdatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Duration(i) * time.Minute)), // Paginationのため分刻みでイテレート
		}
		mr.Create(ent)
	}
}

func dummyManyMessageCouple() {
	s := repositories.NewSession()
	today := strfmt.DateTime(time.Now())
	yesterday := strfmt.DateTime(time.Now().AddDate(0, 0, -1))

	// マッチの前提として相互いいねが必要
	lr := repositories.NewUserLikeRepository(s)
	lr.Create(
		entities.UserLike{
			UserID:    maleMessageUserID,
			PartnerID: femaleMessageUserID,
			CreatedAt: yesterday,
			UpdatedAt: yesterday,
		})
	lr.Create(
		entities.UserLike{
			UserID:    femaleMessageUserID,
			PartnerID: maleMessageUserID,
			CreatedAt: today,
			UpdatedAt: today,
		})

	// メッセージの前提としてマッチが必要
	mr := repositories.NewUserMatchRepository(s)
	mr.Create(
		entities.UserMatch{
			UserID:    maleMessageUserID,
			PartnerID: femaleMessageUserID,
			CreatedAt: today,
			UpdatedAt: today,
		})

	rand.Seed(time.Now().UnixNano())
	createdDaysAgo := rand.Intn(600)
	minute1 := rand.Intn(1440)
	randTime := strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(minute1) * time.Minute))

	// 双方向に130件ずつメッセージ
	msgr := repositories.NewUserMessageRepository(s)
	for i := 0; i <= 130; i++ {
		ent := entities.UserMessage{
			UserID:    maleMessageUserID,
			PartnerID: femaleMessageUserID,
			Message:   "hello",
			CreatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(i*60) * time.Minute)),
			UpdatedAt: randTime,
		}
		msgr.Create(ent)
	}
	for i := 0; i <= 130; i++ {
		ent := entities.UserMessage{
			UserID:    femaleMessageUserID,
			PartnerID: maleMessageUserID,
			Message:   "hi!",
			CreatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(i*60+1) * time.Minute)),
			UpdatedAt: randTime,
		}
		msgr.Create(ent)
	}
}

// covert "1994-12-24" style string to strfmt.Date
func stringToStrFmtDate(str string) strfmt.Date {
	var date strfmt.Date
	date.UnmarshalText([]byte(str))
	return date
}
