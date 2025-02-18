// Code originally generated by pg-bindings generator.

package n5ton6

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_05_to_n_06_postgres_active_components/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_05_to_n_06_postgres_active_components/postgres"
	"github.com/stackrox/rox/migrator/types"
	rawDackbox "github.com/stackrox/rox/pkg/dackbox/raw"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 5,
		VersionAfter:   storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 6},
		Run: func(databases *types.Databases) error {
			legacyStore := legacy.New(rawDackbox.GetGlobalDackBox(), rawDackbox.GetKeyFence())
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving active_components from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize      = 3000
	imageBatchSize = 1000

	schema = pkgSchema.ActiveComponentsSchema
	log    = loghelper.LogWrapper{}
)

type imageIDAndOs struct {
	ID                  string `gorm:"column:id;type:varchar;primaryKey"`
	ScanOperatingSystem string `gorm:"column:scan_operatingsystem;type:varchar"`
}

func move(gormDB *gorm.DB, postgresDB *pgxpool.Pool, legacyStore legacy.Store) error {
	imageTable := gormDB.Table(pkgSchema.ImagesSchema.Table).Model(pkgSchema.CreateTableImagesStmt.GormModel)
	var imageCount int64
	if err := imageTable.Count(&imageCount).Error; err != nil {
		return err
	}
	imageBuf := make([]imageIDAndOs, 0, imageBatchSize)
	imageToOsMap := make(map[string]string, imageCount)
	result := imageTable.FindInBatches(&imageBuf, imageBatchSize, func(_ *gorm.DB, batch int) error {
		for _, sub := range imageBuf {
			imageToOsMap[sub.ID] = sub.ScanOperatingSystem
		}
		return nil
	})
	if result.Error != nil {
		return result.Error
	}
	log.WriteToStderrf("Found %d images", result.RowsAffected)
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(postgresDB)
	pkgSchema.ApplySchemaForTable(context.Background(), gormDB, schema.Table)
	var activeComponents []*storage.ActiveComponent
	err := walk(ctx, legacyStore, func(obj *storage.ActiveComponent) error {
		activeComponents = append(activeComponents, convertActiveVuln(imageToOsMap, obj)...)
		if len(activeComponents) == batchSize {
			if err := store.UpsertMany(ctx, activeComponents); err != nil {
				log.WriteToStderrf("failed to persist active_components to store %v", err)
				return err
			}
			activeComponents = activeComponents[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(activeComponents) > 0 {
		if err = store.UpsertMany(ctx, activeComponents); err != nil {
			log.WriteToStderrf("failed to persist active_components to store %v", err)
			return err
		}
	}
	return nil
}

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.ActiveComponent) error) error {
	return storeWalk(ctx, s, fn)
}

func storeWalk(ctx context.Context, s legacy.Store, fn func(obj *storage.ActiveComponent) error) error {
	ids, err := s.GetIDs(ctx)
	if err != nil {
		return err
	}

	for i := 0; i < len(ids); i += batchSize {
		end := i + batchSize

		if end > len(ids) {
			end = len(ids)
		}
		objs, _, err := s.GetMany(ctx, ids[i:end])
		if err != nil {
			return err
		}
		for _, obj := range objs {
			if err = fn(obj); err != nil {
				return err
			}
		}
	}
	return nil
}

func init() {
	migrations.MustRegisterMigration(migration)
}
