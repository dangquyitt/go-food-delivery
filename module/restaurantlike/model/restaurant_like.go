package restaurantlikemodel

import (
	"food_delivery/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int                `json:"user_id" gorm:"column:user_id"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    *time.Time         `json:"updated_at" gorm:"column:updated_at"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot like this restaurant",
		"ErrCannotLikeRestaurant")
}

func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot unlike this restaurant",
		"ErrCannotUnlikeRestaurant")
}

func (l Like) GetRestaurantId() int {
	return l.RestaurantId
}
