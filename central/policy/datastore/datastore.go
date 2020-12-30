package datastore

import (
	"context"

	clusterDS "github.com/stackrox/rox/central/cluster/datastore"
	notifierDS "github.com/stackrox/rox/central/notifier/datastore"
	"github.com/stackrox/rox/central/policy/index"
	"github.com/stackrox/rox/central/policy/search"
	"github.com/stackrox/rox/central/policy/store"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	searchPkg "github.com/stackrox/rox/pkg/search"
)

// DataStore is an intermediary to PolicyStorage.
//go:generate mockgen-wrapper
type DataStore interface {
	Search(ctx context.Context, q *v1.Query) ([]searchPkg.Result, error)
	SearchPolicies(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error)
	SearchRawPolicies(ctx context.Context, q *v1.Query) ([]*storage.Policy, error)

	GetPolicy(ctx context.Context, id string) (*storage.Policy, bool, error)
	GetAllPolicies(ctx context.Context) ([]*storage.Policy, error)
	GetPolicies(ctx context.Context, ids []string) ([]*storage.Policy, []int, []error, error)
	GetPolicyByName(ctx context.Context, name string) (*storage.Policy, bool, error)

	AddPolicy(context.Context, *storage.Policy) (string, error)
	UpdatePolicy(context.Context, *storage.Policy) error
	RemovePolicy(ctx context.Context, id string) error
	RenamePolicyCategory(ctx context.Context, request *v1.RenamePolicyCategoryRequest) error
	DeletePolicyCategory(ctx context.Context, request *v1.DeletePolicyCategoryRequest) error
	// This method is allowed to return a v1 proto because it is in the baseline in tools/storedprotos/storeinterface/storeinterface.go
	//   (TODO: update "baseline" to inclusive language when updating actual code)
	ImportPolicies(ctx context.Context, policies []*storage.Policy, overwrite bool) (responses []*v1.ImportPolicyResponse, allSucceeded bool, err error)
}

// New returns a new instance of DataStore using the input store, indexer, and searcher.
func New(storage store.Store, indexer index.Indexer, searcher search.Searcher, clusterDatastore clusterDS.DataStore, notifierDatastore notifierDS.DataStore) DataStore {
	return &datastoreImpl{
		storage:           storage,
		indexer:           indexer,
		searcher:          searcher,
		clusterDatastore:  clusterDatastore,
		notifierDatastore: notifierDatastore,
	}
}
