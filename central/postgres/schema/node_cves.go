// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/central/globaldb"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableNodeCvesStmt holds the create statement for table `node_cves`.
	CreateTableNodeCvesStmt = &postgres.CreateStmts{
		Table: `
               create table if not exists node_cves (
                   Id varchar,
                   Cve varchar,
                   Cvss numeric,
                   ImpactScore numeric,
                   PublishedOn timestamp,
                   CreatedAt timestamp,
                   Suppressed bool,
                   SuppressExpiry timestamp,
                   Severity integer,
                   serialized bytea,
                   PRIMARY KEY(Id)
               )
               `,
		Indexes:  []string{},
		Children: []*postgres.CreateStmts{},
	}

	// NodeCvesSchema is the go schema for table `node_cves`.
	NodeCvesSchema = func() *walker.Schema {
		schema := globaldb.GetSchemaForTable("node_cves")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.CVE)(nil)), "node_cves")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_NODE_VULNERABILITIES, "cve", (*storage.CVE)(nil)))
		globaldb.RegisterTable(schema)
		return schema
	}()
)
