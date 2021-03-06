//+build db_test

package services

import (
	"fmt"
	"os"
	"testing"

	"github.com/alephshahor/Mirlo/server/utils"

	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"

	"github.com/alephshahor/Mirlo/server/repositories"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceDBTestSuite struct {
	suite.Suite
}

func (suite *UserServiceDBTestSuite) SetupTest() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
}

func (suite *UserServiceDBTestSuite) TestNewUserService() {
	var userService = NewUserService(repositories.Repositories())

	assert.NotNil(suite.T(), userService)
}

func (suite *UserServiceDBTestSuite) TestCreateUser() {
	var userService = Services().User()

	var user models.User
	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	user, err = userService.Create(newUserReq)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), user)

	assert.Equal(suite.T(), newUserReq.UserName, user.UserName)
	assert.Equal(suite.T(), newUserReq.Email, user.Email)
	assert.NotEqual(suite.T(), newUserReq.Password, user.Password)

	assert.True(suite.T(), utils.PasswordMatch(newUserReq.Password, user.Password))
}

func (suite *UserServiceDBTestSuite) TestFindByUserName() {
	var userService = Services().User()

	var user models.User
	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	user, err = userService.Create(newUserReq)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), user)

	var foundUser models.User
	foundUser, err = userService.FindByUserName(newUserReq.UserName)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), foundUser)

	assert.Equal(suite.T(), newUserReq.UserName, foundUser.UserName)
	assert.Equal(suite.T(), newUserReq.Email, foundUser.Email)
}

func (suite *UserServiceDBTestSuite) TestFindByEmail() {
	var userService = Services().User()

	var user models.User
	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	user, err = userService.Create(newUserReq)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), user)

	var foundUser models.User
	foundUser, err = userService.FindByEmail(newUserReq.Email)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), foundUser)

	assert.Equal(suite.T(), newUserReq.UserName, foundUser.UserName)
	assert.Equal(suite.T(), newUserReq.Email, foundUser.Email)
}

func TestUserServiceDBTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceDBTestSuite))
}
