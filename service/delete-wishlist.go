package service

import "go.uber.org/zap"

func (s *ProductService) DeleteFromWishlist(wishlistID int) error {
	err := s.Repo.ProductRepo.DeleteFromWishlist(wishlistID)
	if err != nil {
		s.Repo.ProductRepo.Logger.Error("Error deleting from wishlist", zap.Error(err))
		return err
	}
	return nil
}
