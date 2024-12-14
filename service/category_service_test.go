package service

import (
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var CategoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: CategoryRepository}

func TestCategoryService_Get(t *testing.T) {
	CategoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
