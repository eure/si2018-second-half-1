package idealtype

import (
	"github.com/go-openapi/runtime/middleware"

	si "github.com/eure/si2018-second-half-1/restapi/summerintern"
)

func ValidateGetIdealType(t string) middleware.Responder {
	if len(t) == 0 {
		return si.NewGetProfileByUserIDUnauthorized().WithPayload(
			&si.GetProfileByUserIDUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: Tokenを指定してください",
			})
	}

	return nil
}
