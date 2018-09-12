package idealtype

import (
	"github.com/eure/si2018-second-half-1/libs/token"
	"github.com/eure/si2018-second-half-1/repositories"
	si "github.com/eure/si2018-second-half-1/restapi/summerintern"
	"github.com/go-openapi/runtime/middleware"
)

func GetIdealType(p si.GetProfileByUserIDParams) middleware.Responder {
	// バリデーション
	t := p.Token
	if res := ValidateGetIdealType(t); res != nil {
		return res
	}

	s := repositories.NewSession()

	// Meの取得
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

	/*
		Likeの統計データの取得
		平均値が格納されたDBから理想のお相手を取得します。
	*/

	//sEnt := user.Build()
	return si.NewGetIdealTypeOK() //.WithPayload(&sEnt)
}
