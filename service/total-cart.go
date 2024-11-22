package service

func (service *ProductService) GetTotalCartQuantity() (int, error) {
	totalQuantity, err := service.Repo.ProductRepo.GetTotalCartQuantity()
	if err != nil {
		return 0, err
	}
	return totalQuantity, nil
}
