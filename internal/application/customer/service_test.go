package customer

import (
	"context"
	"testing"

	"github.com/andredecarli/go-example/internal/domain/customer"
	"github.com/stretchr/testify/suite"
	gomock "go.uber.org/mock/gomock"
)

type CustomerServiceTestSuite struct {
	suite.Suite

	ctrl       *gomock.Controller
	repository *MockCustomerRepository
	service    *service
}

func (s *CustomerServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repository = NewMockCustomerRepository(s.ctrl)
	s.service = NewService(s.repository)
}

func (s *CustomerServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *CustomerServiceTestSuite) TestCreate() {
	s.Run("should create a customer", func() {
		input := &customer.Customer{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}
		ctx := context.Background()

		s.repository.EXPECT().Create(gomock.Any(), input).Return(&customer.Customer{}, nil)

		created, err := s.service.Create(ctx, input)
		s.NoError(err)
		s.NotNil(created)
	})

	s.Run("should return an error if the name is required", func() {
		input := &customer.Customer{
			Email: "john.doe@example.com",
		}
		ctx := context.Background()

		s.repository.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)

		created, err := s.service.Create(ctx, input)
		s.Error(err)
		s.Nil(created)
		s.Equal(ErrNameIsRequired, err)
	})

	s.Run("should return an error if the email is required", func() {
		input := &customer.Customer{
			Name: "John Doe",
		}
		ctx := context.Background()

		s.repository.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)

		created, err := s.service.Create(ctx, input)
		s.Error(err)
		s.Nil(created)
		s.Equal(ErrEmailIsRequired, err)
	})
}

func TestCustomerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerServiceTestSuite))
}
