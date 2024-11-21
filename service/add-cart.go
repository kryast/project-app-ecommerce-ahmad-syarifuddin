package service

func (s *ProductService) AddToCart(id int, quantity int) error {
	return s.Repo.ProductRepo.AddToCart(id, quantity)
}
