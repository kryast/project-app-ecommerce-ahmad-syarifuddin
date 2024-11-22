package service

func (service *ProductService) UpdateCartQuantity(productDetailsID, newQuantity int) error {
	err := service.Repo.ProductRepo.UpdateCartQuantity(productDetailsID, newQuantity)
	if err != nil {
		return err
	}
	return nil
}
