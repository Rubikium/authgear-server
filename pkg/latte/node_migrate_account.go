package latte

import (
	"context"

	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity"
	"github.com/authgear/authgear-server/pkg/lib/workflow"
)

func init() {
	workflow.RegisterNode(&NodeMigrateAccount{})
}

type NodeMigrateAccount struct {
	IdentityMigrateSpecs      []*identity.MigrateSpec      `json:"identity_migrate_specs"`
	AuthenticatorMigrateSpecs []*authenticator.MigrateSpec `json:"authenticator_migrate_specs"`
}

func (n *NodeMigrateAccount) Kind() string {
	return "latte.NodeMigrateAccount"
}

func (n *NodeMigrateAccount) GetEffects(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) (effs []workflow.Effect, err error) {
	return nil, nil
}

func (*NodeMigrateAccount) CanReactTo(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) ([]workflow.Input, error) {
	return nil, workflow.ErrEOF
}

func (*NodeMigrateAccount) ReactTo(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow, input workflow.Input) (*workflow.Node, error) {
	return nil, workflow.ErrIncompatibleInput
}

func (n *NodeMigrateAccount) OutputData(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) (interface{}, error) {
	return nil, nil
}
