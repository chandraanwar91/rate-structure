// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import gin "github.com/gin-gonic/gin"
import mock "github.com/stretchr/testify/mock"
import models "misteraladin.com/jasmine/rate-structure/models"

// IRateBackofficeUseCase is an autogenerated mock type for the IRateBackofficeUseCase type
type IRateBackofficeUseCase struct {
	mock.Mock
}

// CheckAvailable provides a mock function with given fields: c
func (_m *IRateBackofficeUseCase) CheckAvailable(c *gin.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gin.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: c
func (_m *IRateBackofficeUseCase) Fetch(c *gin.Context) ([]*models.RateBackoffice, *models.Pagination, error) {
	ret := _m.Called(c)

	var r0 []*models.RateBackoffice
	if rf, ok := ret.Get(0).(func(*gin.Context) []*models.RateBackoffice); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.RateBackoffice)
		}
	}

	var r1 *models.Pagination
	if rf, ok := ret.Get(1).(func(*gin.Context) *models.Pagination); ok {
		r1 = rf(c)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Pagination)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*gin.Context) error); ok {
		r2 = rf(c)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: c
func (_m *IRateBackofficeUseCase) GetByID(c *gin.Context) (*models.RateBackoffice, error) {
	ret := _m.Called(c)

	var r0 *models.RateBackoffice
	if rf, ok := ret.Get(0).(func(*gin.Context) *models.RateBackoffice); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RateBackoffice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: c
func (_m *IRateBackofficeUseCase) Store(c *gin.Context) (*models.RateBackoffice, error) {
	ret := _m.Called(c)

	var r0 *models.RateBackoffice
	if rf, ok := ret.Get(0).(func(*gin.Context) *models.RateBackoffice); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RateBackoffice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Test provides a mock function with given fields: name
func (_m *IRateBackofficeUseCase) Test(name string) string {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Update provides a mock function with given fields: c
func (_m *IRateBackofficeUseCase) Update(c *gin.Context) (*models.RateBackoffice, error) {
	ret := _m.Called(c)

	var r0 *models.RateBackoffice
	if rf, ok := ret.Get(0).(func(*gin.Context) *models.RateBackoffice); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RateBackoffice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
