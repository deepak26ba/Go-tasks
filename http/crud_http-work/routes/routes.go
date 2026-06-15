package routes

import (
	"http/handlers"
	"http/repository"
	"http/service"
	"net/http"

	"gorm.io/gorm"
)

func Routes(db *gorm.DB) {

	repo := repository.New(db)
	services := service.New(repo)
	handler := handlers.New(services)

	http.HandleFunc("/user-role/create", handler.CreateUserRole)
	http.HandleFunc("/user-role/get", handler.GetUserRole)
	http.HandleFunc("/user-role/get-by-id", handler.GetByIdUserRole)
	http.HandleFunc("/user-role/updates-by-id", handler.UpdateByIdUserRole)
	http.HandleFunc("/user-role/update-by-id", handler.UpdateUserRole)
	http.HandleFunc("/user-role/delete-by-id", handler.DeleteByIdUserRole)

}
