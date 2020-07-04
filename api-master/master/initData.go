package master

import (
	"database/sql"
	"liveCodeAPI/api-master/master/controllers"
	"liveCodeAPI/api-master/master/repositories/categoryrepository"
	"liveCodeAPI/api-master/master/repositories/productrepository"
	"liveCodeAPI/api-master/master/repositories/transactionrepository"
	"liveCodeAPI/api-master/master/repositories/userrepository"
	"liveCodeAPI/api-master/master/usecases/categoryusecase"
	"liveCodeAPI/api-master/master/usecases/productusecase"
	"liveCodeAPI/api-master/master/usecases/transactionusecase"
	"liveCodeAPI/api-master/master/usecases/userusecase"

	"github.com/gorilla/mux"
)

func InitData(r *mux.Router, db *sql.DB) {
	userRepo := userrepository.InitUserRepoImpl(db)
	userUsecase := userusecase.InitUserUsecase(userRepo)
	controllers.UserController(r, userUsecase)

	productRepo := productrepository.InitProductRepoImpl(db)
	productUsecase := productusecase.InitProductUsecase(productRepo)
	controllers.ProductController(r, productUsecase)

	categoryRepo := categoryrepository.InitCategoryRepoImpl(db)
	categoryUsecase := categoryusecase.InitCategoryUsecase(categoryRepo)
	controllers.CategoryController(r, categoryUsecase)

	transactionRepo := transactionrepository.InitTransactionRepoImpl(db)
	transactionUsecase := transactionusecase.InitTransactionUsecase(transactionRepo)
	controllers.TransactionController(r, transactionUsecase)

}
