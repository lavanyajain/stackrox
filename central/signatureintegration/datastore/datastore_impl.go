package datastore

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/role/resources"
	"github.com/stackrox/rox/central/signatureintegration/store"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/errox"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	signatureSAC = sac.ForResource(resources.SignatureIntegration)
)

type datastoreImpl struct {
	storage store.SignatureIntegrationStore

	lock sync.RWMutex
}

func (d *datastoreImpl) GetSignatureIntegration(ctx context.Context, id string) (*storage.SignatureIntegration, bool, error) {
	if ok, err := signatureSAC.ReadAllowed(ctx); !ok || err != nil {
		return nil, false, err
	}

	return d.storage.Get(id)
}

func (d *datastoreImpl) GetAllSignatureIntegrations(ctx context.Context) ([]*storage.SignatureIntegration, error) {
	if ok, err := signatureSAC.ReadAllowed(ctx); !ok || err != nil {
		return nil, err
	}

	var integrations []*storage.SignatureIntegration
	err := d.storage.Walk(func(integration *storage.SignatureIntegration) error {
		integrations = append(integrations, integration)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return integrations, nil
}

func (d *datastoreImpl) AddSignatureIntegration(ctx context.Context, integration *storage.SignatureIntegration) (*storage.SignatureIntegration, error) {
	if err := sac.VerifyAuthzOK(signatureSAC.WriteAllowed(ctx)); err != nil {
		return nil, err
	}
	if integration.GetId() != "" {
		return nil, errox.Newf(errox.InvalidArgs, "id should be empty but %q provided", integration.GetId())
	}
	integration.Id = GenerateSignatureIntegrationID()
	if err := ValidateSignatureIntegration(integration); err != nil {
		return nil, errox.NewErrInvalidArgs(err.Error())
	}

	// Protect against TOCTOU race condition.
	d.lock.Lock()
	defer d.lock.Unlock()

	if err := d.verifyIntegrationIDDoesNotExist(integration.GetId()); err != nil {
		if errors.Is(err, errox.AlreadyExists) {
			return nil, errors.Wrap(err, "collision in generated signature integration id, try again")
		}
		return nil, err
	}

	err := d.storage.Upsert(integration)
	if err != nil {
		return nil, err
	}
	return integration, nil
}

func (d *datastoreImpl) UpdateSignatureIntegration(ctx context.Context, integration *storage.SignatureIntegration) error {
	if err := sac.VerifyAuthzOK(signatureSAC.WriteAllowed(ctx)); err != nil {
		return err
	}
	if err := ValidateSignatureIntegration(integration); err != nil {
		return errox.NewErrInvalidArgs(err.Error())
	}

	// Protect against TOCTOU race condition.
	d.lock.Lock()
	defer d.lock.Unlock()

	if err := d.verifyIntegrationIDExists(integration.GetId()); err != nil {
		return err
	}

	return d.storage.Upsert(integration)
}

func (d *datastoreImpl) RemoveSignatureIntegration(ctx context.Context, id string) error {
	if err := sac.VerifyAuthzOK(signatureSAC.WriteAllowed(ctx)); err != nil {
		return err
	}

	d.lock.Lock()
	defer d.lock.Unlock()

	if err := d.verifyIntegrationIDExists(id); err != nil {
		return err
	}

	return d.storage.Delete(id)
}

func (d *datastoreImpl) verifyIntegrationIDExists(id string) error {
	_, found, err := d.storage.Get(id)
	if err != nil {
		return err
	} else if !found {
		return errox.Newf(errox.NotFound, "signature integration id=%s doesn't exist", id)
	}
	return nil
}

func (d *datastoreImpl) verifyIntegrationIDDoesNotExist(id string) error {
	_, found, err := d.storage.Get(id)
	if err != nil {
		return err
	} else if found {
		return errox.Newf(errox.AlreadyExists, "signature integration id=%s already exists", id)
	}
	return nil
}