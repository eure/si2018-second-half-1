package idealtype

import (
	"github.com/eure/si2018-second-half-1/libs/token"
	"github.com/eure/si2018-second-half-1/models"
	"github.com/eure/si2018-second-half-1/repositories"
	si "github.com/eure/si2018-second-half-1/restapi/summerintern"
	"github.com/go-openapi/runtime/middleware"
)

func GetIdealType(p si.GetIdealTypeParams) middleware.Responder {
	// バリデーション
	t := p.Token
	if res := ValidateGetIdealType(t); res != nil {
		return res
	}

	s := repositories.NewSession()

	//Meの取得
	me, err := token.GetUserByToken(s, t)
	if err != nil {
		return si.NewGetIdealTypeInternalServerError().WithPayload(
			&si.GetIdealTypeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewGetIdealTypeUnauthorized().WithPayload(
			&si.GetIdealTypeUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: 無効なトークン",
			})
	}

	// ダミーデータ
	var ideal models.IdealType
	ideal.Age = new(models.IdealTypeAge)
	ideal.Age.From = 20
	ideal.Age.To = 25
	ideal.AnnualIncome = new(models.IdealTypeAnnualIncome)
	ideal.AnnualIncome.From = 200
	ideal.AnnualIncome.To = 600
	ideal.BodyBuild = []string{"スリム", "やや細め", "普通"}
	ideal.Drinking = []string{"飲む", "ときどき飲む"}
	ideal.Education = []string{"高校卒", "大学卒"}
	ideal.Height = new(models.IdealTypeHeight)
	ideal.Height.From = 165
	ideal.Height.To = 185
	ideal.Holiday = []string{"土日", "平日"}
	ideal.HomeState = []string{"東京", "千葉", "神奈川", "埼玉"}
	ideal.Job = []string{"会社員", "医師", "弁護士"}
	ideal.ResidenceState = []string{"東京", "神奈川"}
	ideal.Smoking = []string{"吸う", "ときどき吸う", "非喫煙者の前では吸わない", "相手が嫌ならやめる", "吸う(電子タバコ)"}

	/*
		Likeの統計データの取得
		平均値が格納されたDBから理想のお相手を取得します。
	*/

	//sEnt := user.Build()
	return si.NewGetIdealTypeOK().WithPayload(&ideal)
}
