package services

import (
	"demo/internal/models/mysql/entity"
	"demo/tests/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	tests := []struct {
		UserName string
		Email    string
		Error    error
		HasError error
	}{
		{"zhangyi", "zhangyi", nil, nil},
	}

	for _, test := range tests {
		u := entity.User{UserName: test.UserName, Email: test.Email}

		mockUser := mock.NewMockUserI(ctl)
		gomock.InOrder(mockUser.EXPECT().Create(test.UserName, test.Email, "").Return(u, test.Error))

		user := NewUser(mockUser)
		_, err := user.Create(test.UserName, test.Email, "")

		assert.Equal(t, test.HasError, err, "error not equal")
	}
}

func TestUser_Exist(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	tests := []struct {
		UserName string
		Bool     bool
		Error    error
		HasError bool
	}{
		{"zhangyi", true, nil, true},
	}

	for _, test := range tests {
		mockUser := mock.NewMockUserI(ctl)
		gomock.InOrder(mockUser.EXPECT().Exist(test.UserName).Return(test.Bool, test.Error))

		user := NewUser(mockUser)
		ok, _ := user.Exist(test.UserName)

		assert.Equal(t, test.HasError, ok, "error not equal")
	}

}
