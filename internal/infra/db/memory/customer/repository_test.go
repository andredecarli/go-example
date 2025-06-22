package customer

import (
	"context"
	"testing"

	"github.com/andredecarli/go-example/internal/domain/customer"
	"github.com/stretchr/testify/suite"
)

type MemoryRepositoryTestSuite struct {
	suite.Suite
	repository *repository
}

func (s *MemoryRepositoryTestSuite) SetupTest() {
	s.repository = NewRepository()
}

func (s *MemoryRepositoryTestSuite) TestRepository_Create() {
	customer := &customer.Customer{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	ctx := context.Background()

	created, err := s.repository.Create(ctx, customer)
	s.NoError(err)
	s.NotNil(created)
	s.NotEmpty(created.ID)
	s.Equal(customer.Name, created.Name)
	s.Equal(customer.Email, created.Email)
	s.NotEmpty(created.CreatedAt)
	s.NotEmpty(created.UpdatedAt)
}

func TestMemoryRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryRepositoryTestSuite))
}
