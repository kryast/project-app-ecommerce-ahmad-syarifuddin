package router

import (
	"project-app-ecommerce-ahmad-syarifuddin/database"
	"project-app-ecommerce-ahmad-syarifuddin/handler"
	"project-app-ecommerce-ahmad-syarifuddin/repository"
	"project-app-ecommerce-ahmad-syarifuddin/service"
	"project-app-ecommerce-ahmad-syarifuddin/util"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func InitRouter() (*chi.Mux, *zap.Logger, error) {
	r := chi.NewRouter()

	logger := util.InitLog()

	config, err := util.ReadConfiguration()
	if err != nil {
		return nil, logger, err
	}

	db, err := database.InitDB(config)
	if err != nil {
		return nil, logger, err
	}

	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	Handle := handler.NewAllHandler(service, logger, config)

	r.Post("/register", Handle.AuthHandler.Register)
	r.Post("/login", Handle.AuthHandler.Login)

	r.Get("/products", Handle.ProductHandler.GetAllProducts)
	r.Get("/categories", Handle.ProductHandler.GetAllCategory)
	r.Get("/best-seller", Handle.ProductHandler.GetBestSellers)
	r.Get("/banner", Handle.ProductHandler.GetAllBanner)

	r.Get("/products/{id}", Handle.ProductHandler.GetProductByIdHandler)
	r.Post("/products/{id}/wishlist", Handle.ProductHandler.AddToWishlistHandler)

	r.Delete("/products/{id}/wishlist", Handle.ProductHandler.DeleteFromWishlist)

	r.Get("/products/cart", Handle.ProductHandler.GetAllCart)

	r.Get("/products/recomment", Handle.ProductHandler.GetAllRecomment)

	return r, logger, nil
}
