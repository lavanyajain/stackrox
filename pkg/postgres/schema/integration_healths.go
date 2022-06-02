// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableIntegrationHealthsStmt holds the create statement for table `integration_healths`.
	CreateTableIntegrationHealthsStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists integration_healths (
                   Id varchar,
                   serialized bytea,
                   PRIMARY KEY(Id)
               )
               `,
		GormModel: (*IntegrationHealths)(nil),
		Indexes:   []string{},
		Children:  []*postgres.CreateStmts{},
	}

	// IntegrationHealthsSchema is the go schema for table `integration_healths`.
	IntegrationHealthsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("integration_healths")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.IntegrationHealth)(nil)), "integration_healths")
		RegisterTable(schema)
		return schema
	}()
)

const (
	IntegrationHealthsTableName = "integration_healths"
)

// IntegrationHealths holds the Gorm model for Postgres table `integration_healths`.
type IntegrationHealths struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}