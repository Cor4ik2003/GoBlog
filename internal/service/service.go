package service

import "github.com/zhashkevych/todo-app/pkg/repository"

type Service struct {
	repo *repository.Repository
}