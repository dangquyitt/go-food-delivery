package restaurantstorage

import (
	"context"
	restaurantmodel "food_delivery/module/restaurant/model"
)

func (s *sqlStore) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
