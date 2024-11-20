package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (ps *ProductService) AddToWishlist(productDetailsID int) (*model.Wishlist, error) {
	// Memanggil fungsi repositori untuk menambahkan produk ke wishlist
	wishlist, err := ps.Repo.ProductRepo.AddToWishlist(productDetailsID)
	if err != nil {
		return nil, fmt.Errorf("failed to add to wishlist: %w", err)
	}
	return wishlist, nil
}
