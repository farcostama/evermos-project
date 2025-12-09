package service

import (
	"evermos-project/internal/entity"
	"evermos-project/internal/repository"
)

type CategoryService interface {
	GetAllCategories() ([]entity.Category, error)
	GetCategoryByID(id uint) (*entity.Category, error)
	CreateCategory(req *entity.Category) error
	UpdateCategory(id uint, req *entity.Category) error
	DeleteCategory(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) GetAllCategories() ([]entity.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetCategoryByID(id uint) (*entity.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) CreateCategory(req *entity.Category) error {
	return s.repo.Create(req)
}

func (s *categoryService) UpdateCategory(id uint, req *entity.Category) error {
	cat, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	// Update field
	if req.NamaCategory != "" {
		cat.NamaCategory = req.NamaCategory
	}
	return s.repo.Update(cat)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}