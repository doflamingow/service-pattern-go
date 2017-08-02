// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import models "service-pattern-go/models"

// IPlayerService is an autogenerated mock type for the IPlayerService type
type IPlayerService struct {
	mock.Mock
}

// FindById provides a mock function with given fields: playerId
func (_m *IPlayerService) FindById(playerId int) models.Player {
	ret := _m.Called(playerId)

	var r0 models.Player
	if rf, ok := ret.Get(0).(func(int) models.Player); ok {
		r0 = rf(playerId)
	} else {
		r0 = ret.Get(0).(models.Player)
	}

	return r0
}
