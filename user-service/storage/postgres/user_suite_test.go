package postgres

// import (
// 	"github.com/stretchr/testify/suite"
// 	"template-service/config"
// 	pb "template-service/genproto"
// 	"template-service/pkg/db"
// 	"template-service/storage/repo"
// 	"testing"
// )

// type UserRepositoryTestSuite struct {
// 	suite.Suite
// 	CleanUpFunc func()
// 	Repository  repo.UserStorageI
// }

// func (s *UserRepositoryTestSuite) SetupSuite() {
// 	pgPool, cleanUp := db.ConnectDBForSuite(config.Load())
// 	s.Repository = NewUserRepo(pgPool)
// 	s.CleanUpFunc = cleanUp
// }

// func (s *UserRepositoryTestSuite) TestUserCRUD() {
// 	user := pb.User{
// 		Name:     "Test User",
// 		LastName: "Suite Test User",
// 	}

// 	createdUser, err := s.Repository.Create(&user)
// 	s.Suite.NotNil(createdUser)
// 	s.Suite.NoError(err)
// 	s.Suite.Equal(user.Name, createdUser.Name)
// 	s.Suite.Equal(user.LastName, createdUser.LastName)

// 	getUser, err := s.Repository.GetUser(&user.Id)
// 	s.Suite.NotNil(getUser)
// 	s.Suite.NoError(err)
// 	s.Suite.Equal(user.Name, getUser.Name)
// 	s.Suite.Equal(user.LastName, getUser.LastName)

// 	createdUser.Name = "Updated User Name"
// 	createdUser.LastName = "Updated Last Name"

// 	updatedUser, err := s.Repository.GetUser(&user.Id)
// 	s.Suite.NotNil(updatedUser)
// 	s.Suite.NoError(err)

// 	getUpdatedUser, err := s.Repository.GetUser(&user.Id)
// 	s.Suite.NotNil(getUpdatedUser)
// 	s.Suite.NoError(err)
// 	s.Suite.NotEqual(createdUser.Name, getUpdatedUser.Name)
// 	s.Suite.NotEqual(createdUser.LastName, getUpdatedUser.LastName)

// 	allUsers, err := s.Repository.GetAll(&pb.GetAllRequest{Page: 1, Limit: 10})
// 	s.Suite.NotNil(allUsers)
// 	s.Suite.NoError(err)

// 	deleteUser, err := s.Repository.Delete(getUpdatedUser)
// 	s.Suite.NotNil(deleteUser)
// 	s.Suite.NoError(err)
// }

// func (s *UserRepositoryTestSuite) TearDownSuite() {
// 	s.CleanUpFunc()
// }

// func TestUserRepositoryTestSuite(t *testing.T) {
// 	suite.Run(t, new(UserRepositoryTestSuite))
// }
