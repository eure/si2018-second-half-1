package userlike

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"fmt"

	"github.com/eure/si2018-second-half-1/entities"
	"github.com/eure/si2018-second-half-1/libs/stats"
	"github.com/eure/si2018-second-half-1/libs/token"
	userlib "github.com/eure/si2018-second-half-1/libs/user"
	"github.com/eure/si2018-second-half-1/repositories"
	si "github.com/eure/si2018-second-half-1/restapi/summerintern"
)

func GetLikes(p si.GetLikesParams) middleware.Responder {
	// バリデーション
	t := p.Token
	limit := p.Limit
	offset := p.Offset
	v := NewGetValidator(t, limit, offset)
	if res := v.Validate(); res != nil {
		return res
	}

	s := repositories.NewSession()

	// meチェック
	me, err := token.GetUserByToken(s, t)
	if err != nil {
		return si.NewGetLikesInternalServerError().WithPayload(
			&si.GetLikesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewGetLikesUnauthorized().WithPayload(
			&si.GetLikesUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}

	// LikeのレスポンスにはMatch済みのお相手を返さないので not in するためにmatchIDsを取得する
	mr := repositories.NewUserMatchRepository(s)
	matchIDs, err := mr.FindAllByUserID(me.ID)

	// GotLikeをlimit offsetで取得
	r := repositories.NewUserLikeRepository(s)
	likes, err := r.FindGotLikeWithLimitOffset(me.ID, int(limit), int(offset), matchIDs)
	if err != nil {
		return si.NewGetLikesInternalServerError().WithPayload(
			&si.GetLikesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: likeの取得に失敗",
			})
	}

	// 取得したlikesをrangeで回す
	// keyにお相手のUID, valueがLikeUserResponseのlikeUserMapを作る
	// ( LikeUserResponse.liked_at はここで入れる )
	likeUserIDList := make([]int64, len(likes))
	likeUserMap := make(map[int64]entities.LikeUserResponse, len(likes))
	for i, l := range likes {
		if me.ID == l.UserID {
			likeUserIDList[i] = l.PartnerID
			likeUserMap[l.PartnerID] = entities.LikeUserResponse{
				LikedAt: l.CreatedAt,
			}
			continue
		}
		likeUserIDList[i] = l.UserID
		likeUserMap[l.UserID] = entities.LikeUserResponse{
			LikedAt: l.CreatedAt,
		}
	}

	// ユーザーを取得
	ur := repositories.NewUserRepository(s)
	users, err := ur.FindByIDs(likeUserIDList)
	if err != nil {
		return si.NewGetLikesInternalServerError().WithPayload(
			&si.GetLikesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Userの取得に失敗",
			})
	}

	// UserImage.URIをUser.ImageURIに付与
	users, err = userlib.SetUsersImage(s, users)
	if err != nil {
		return si.NewGetLikesInternalServerError().WithPayload(
			&si.GetLikesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: UserImageの取得に失敗",
			})
	}

	// Like順に, UserデータをlikeUserMapに入れていく
	for _, u := range users {
		respUser := likeUserMap[u.ID]
		respUser.ApplyUser(u)
		likeUserMap[u.ID] = respUser
	}

	// MapをSliceに変換
	sortedRespList := make([]entities.LikeUserResponse, len(likeUserIDList))
	for i, m := range likeUserIDList {
		sortedRespList[i] = likeUserMap[m]
	}

	// SwaggerのModelに変換してreturnする
	responses := entities.LikeUserResponses(sortedRespList)
	swaggerResponses := responses.Build()
	return si.NewGetLikesOK().WithPayload(swaggerResponses)
}

func toString(t si.PostLikeParams) string {
	return fmt.Sprintf("id = %v, params = {token = %v}",
		t.UserID, t.Params.Token)
}

func PostLike(p si.PostLikeParams) middleware.Responder {
	fmt.Println("**************** PostLike STRAT ****************")
	fmt.Println("**************** ", toString(p), " ****************")
	// リクエストパラメータのバリデーション
	t := p.Params.Token
	v := NewPostValidator(t, p.UserID)
	if resp := v.Validate(); resp != nil {
		return resp
	}

	s := repositories.NewSession()

	// meチェック
	me, err := token.GetUserByToken(s, t)
	if err != nil {
		fmt.Println("**************** !!!PostLike ERROR!!! 500 Meの取得に失敗しました ****************")
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		fmt.Println("**************** !!!PostLike ERROR!!! 401 トークンが無効です ****************")
		return si.NewPostLikeUnauthorized().WithPayload(
			&si.PostLikeUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}

	partnerID := p.UserID
	ur := repositories.NewUserRepository(s)
	partner, err := ur.GetByUserID(partnerID)
	if err != nil {
		fmt.Println("**************** !!!PostLike ERROR!!! 500 Partnerの取得に失敗しました ****************")
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Partnerの取得に失敗しました",
			})
	}
	if partner == nil {
		fmt.Println("**************** !!!PostLike ERROR!!! 400 お相手が存在しません ****************")
		return si.NewPostLikeBadRequest().WithPayload(
			&si.PostLikeBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: お相手が存在しません",
			})
	}
	if me.Gender == partner.Gender {
		fmt.Println("**************** !!!PostLike ERROR!!! 400 同性へのいいねはできません ****************")
		return si.NewPostLikeBadRequest().WithPayload(
			&si.PostLikeBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: 同性へのいいねはできません",
			})
	}

	// Like (Me -> Partner) の存在チェック
	// 既に送ってたら400で返す.
	lr := repositories.NewUserLikeRepository(s)
	sendLike, err := lr.GetLikeBySenderIDReceiverID(me.ID, partnerID)
	if err != nil {
		fmt.Println("**************** !!!PostLike ERROR!!! 500 Likeの取得に失敗しました ****************")
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Likeの取得に失敗しました",
			})
	}
	if sendLike != nil {
		fmt.Println("**************** !!!PostLike ERROR!!! 400 既にLike送信済みです ****************")
		return si.NewPostLikeBadRequest().WithPayload(
			&si.PostLikeBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: 既にLike送信済みです",
			})
	}

	// ここからスピードワゴン機能
	sr := repositories.NewUserStatsRepository(s)
	stat, err := sr.GetByUserID(me.ID)
	if err != nil {
		fmt.Println("**************** !!!PostLike ERROR!!! 500 Statsの取得に失敗しました ****************")
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Statsの取得に失敗しました",
			})
	}

	repositories.TransactionBegin(s)
	// Like (Me -> Partner) レコードのInsert
	now := strfmt.DateTime(time.Now())
	err = lr.Create(
		entities.UserLike{
			UserID:    me.ID,
			PartnerID: partnerID,
			CreatedAt: now,
			UpdatedAt: now,
		})
	if err != nil {
		repositories.TransactionRollBack(s)
		fmt.Println("**************** !!!PostLike ERROR!!! 500 LikeのInsertに失敗しました ****************")
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: LikeのInsertに失敗しました",
			})
	}

	if stat == nil {

		ids, err := lr.FindMeLikeAll(me.ID)
		if err != nil {
			repositories.TransactionRollBack(s)
			fmt.Println("**************** !!!PostLike ERROR!!! 500 何人Likeしているかのカウントに失敗しました ****************")
			return si.NewPostLikeInternalServerError().WithPayload(
				&si.PostLikeInternalServerErrorBody{
					Code:    "500",
					Message: "Internal Server Error :: 何人Likeしているかのカウントに失敗しました",
				})
		}

		if len(ids) >= 10 {
			fmt.Println("ユーザー", me.ID, "は", len(ids), "人にいいねしたので, 統計データを作成します")
			us, err := ur.FindByIDs(ids)
			if err != nil {
				repositories.TransactionRollBack(s)
				fmt.Println("**************** !!!PostLike ERROR!!! 500 MeがLikeしているユーザー情報の取得に失敗しました ****************")
				return si.NewPostLikeInternalServerError().WithPayload(
					&si.PostLikeInternalServerErrorBody{
						Code:    "500",
						Message: "Internal Server Error :: MeがLikeしているユーザー情報の取得に失敗しました",
					})
			}

			stats := stats.GetAverage(us)
			stats.UserID = me.ID
			stats.CreatedAt = now
			stats.UpdatedAt = now
			sr.Create(stats)
		}
	} else {
		fmt.Println("ユーザー", me.ID, "の統計データを更新します")
		sr.Update(stats.ApplyNewLike(stat, partner))
	}

	// Like (Partner -> Me) の存在チェック
	// 向こうからまだLikeが来てなければ, LikeをInsertするだけでreturnしてよい
	getLike, err := lr.GetLikeBySenderIDReceiverID(partnerID, me.ID)
	if getLike == nil {
		repositories.TransactionCommit(s)
		fmt.Println("**************** PostLike OK 200 Likeが送信されました ****************")
		return si.NewPostLikeOK().WithPayload(
			&si.PostLikeOKBody{
				Code:    "200",
				Message: "OK :: Likeが送信されました",
			})
	}

	// 向こうからLikeがきていた場合,Like送信して,さらにMatchもさせる
	mr := repositories.NewUserMatchRepository(s)
	err = mr.Create(
		entities.UserMatch{
			UserID:    me.ID,
			PartnerID: partnerID,
			CreatedAt: now,
			UpdatedAt: now,
		})
	if err != nil {
		repositories.TransactionRollBack(s)
		fmt.Println("**************** !!!PostLike ERROR!!! 500 MatchのInsertに失敗しました ****************")
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: MatchのInsertに失敗しました",
			})
	}
	repositories.TransactionCommit(s)

	fmt.Println("**************** PostLike END ****************")
	return si.NewPostLikeOK().WithPayload(
		&si.PostLikeOKBody{
			Code:    "200",
			Message: "OK :: Likeを送信し,お相手とMatchしました",
		})
}
