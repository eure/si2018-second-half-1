package idealtype

import (
	"fmt"

	"github.com/eure/si2018-second-half-1/libs/token"
	"github.com/eure/si2018-second-half-1/repositories"
	si "github.com/eure/si2018-second-half-1/restapi/summerintern"
	"github.com/go-openapi/runtime/middleware"
)

func GetIdealType(p si.GetIdealTypeParams) middleware.Responder {
	fmt.Println("**************** GetIdealType STRAT ****************")
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
	sr := repositories.NewUserStatsRepository(s)
	stat, err := sr.GetByUserID(me.ID)
	if err != nil {
		return si.NewGetIdealTypeInternalServerError().WithPayload(
			&si.GetIdealTypeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Statsの取得に失敗しました",
			})
	}
	if stat == nil {
		return si.NewGetIdealTypeNotFound().WithPayload(
			&si.GetIdealTypeNotFoundBody{
				Code:    "404",
				Message: "Ideal Type Not Found",
			})
	}

	sEnt := stat.Build()
	fmt.Println("**************** GetIdealType END ****************")
	return si.NewGetIdealTypeOK().WithPayload(&sEnt)
}
