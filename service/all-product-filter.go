package service

import "project-app-ecommerce-ahmad-syarifuddin/model"

func (s *ProductService) GetAllProductsFilter(nameTerm, categoryTerm string, page, limit int) ([]*model.Product, int, error) {
	return s.Repo.ProductRepo.GetAllProductsFilter(nameTerm, categoryTerm, page, limit)
}
