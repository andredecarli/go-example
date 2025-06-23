package customer

import (
	"context"
	"testing"

	"github.com/andredecarli/go-example/internal/domain/customer"
	"github.com/andredecarli/go-example/pkg/util"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "go-example"

type DatabaseConfig struct {
	ctx       context.Context
	container testcontainers.Container
	uri       string
}

type MongoDBRepositoryTestSuite struct {
	suite.Suite

	dbConfig   DatabaseConfig
	db         *mongo.Database
	repository *Repository
}

func (s *MongoDBRepositoryTestSuite) SetupTest() {
	ctx, container, uri, err := util.SetupMongoDBContainer()
	if err != nil {
		s.FailNow("Failed to setup MongoDB container", err)
	}
	s.dbConfig = DatabaseConfig{
		ctx:       ctx,
		container: container,
		uri:       uri,
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		s.FailNow("Failed to connect to MongoDB", err)
	}
	s.db = client.Database(databaseName)
	s.repository = NewRepository(s.db)
}

func (s *MongoDBRepositoryTestSuite) TearDownTest() {
	s.dbConfig.container.Terminate(s.dbConfig.ctx)
}

func (s *MongoDBRepositoryTestSuite) TearDownSubTest() {
	s.db.Drop(s.dbConfig.ctx)
}

func (s *MongoDBRepositoryTestSuite) TestRepository_Create() {
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

func TestMongoDBRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MongoDBRepositoryTestSuite))
}
