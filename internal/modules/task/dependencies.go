package task

import "gorm.io/gorm"

func InitModule(db *gorm.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return handler
}
