// +build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/fixtures"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type SecretsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestSecretsStore(t *testing.T) {
	suite.Run(t, new(SecretsStoreSuite))
}

func (s *SecretsStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *SecretsStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *SecretsStoreSuite) TestStore() {

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	if err != nil {
		panic(err)
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	s.NoError(err)
	defer pool.Close()

	Destroy(pool)
	store := New(pool)

	secret := fixtures.GetSecret()
	foundSecret, exists, err := store.Get(secret.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundSecret)

	s.NoError(store.Upsert(secret))
	foundSecret, exists, err = store.Get(secret.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(secret, foundSecret)

	secret.Name = "name_changed"
	s.NoError(store.Upsert(secret))

	foundSecret, exists, err = store.Get(secret.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(secret, foundSecret)

	s.NoError(store.Delete(secret.GetId()))
	foundSecret, exists, err = store.Get(secret.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundSecret)
}