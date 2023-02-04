package restaurantlikegin

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantstorage "food_delivery/module/restaurant/storage"
	restaurantlikebusiness "food_delivery/module/restaurantlike/business"
	restaurantlikemodel "food_delivery/module/restaurantlike/model"
	restaurantlikestorage "food_delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /restaurants/:id/like

func UserLikeRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.TokenPayloadInJWTRequest).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStorage(ctx.GetMainDBConnection())
		incStore := restaurantstorage.NewSqlStore(ctx.GetMainDBConnection())
		bsn := restaurantlikebusiness.NewCreateBusiness(store, incStore)
		if err := bsn.CreateLikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
