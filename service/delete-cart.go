package service

import "fmt"

func (service *ProductService) DeleteCartItem(productDetailsID int) error {
	err := service.Repo.ProductRepo.DeleteFromCart(productDetailsID)
	if err != nil {
		return fmt.Errorf("failed to delete product from cart: %w", err)
	}

	return nil
}
