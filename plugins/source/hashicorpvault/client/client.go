package client

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	vault "github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/api/auth/approle"
	"github.com/rs/zerolog"
)

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	logger     zerolog.Logger
	name       string
	Paths      Paths
	parameters VaultParameters
	Client     *vault.Client
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.name
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	vaultSpec := &Spec{}
	if err := s.UnmarshalSpec(vaultSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gitlab spec: %w", err)
	}
	if err := vaultSpec.Validate(); err != nil {
		return nil, err
	}

	config := vault.DefaultConfig() // modify for more granular configuration
	config.Address = vaultSpec.Parameters.Address

	cli, err := vault.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize vault client: %w", err)
	}

	c := &Client{
		logger:     logger,
		Client:     cli,
		parameters: *vaultSpec.Parameters,
		Paths:      vaultSpec.Paths,
		name:       vaultSpec.Name,
	}

	_, err = c.login(ctx)
	if err != nil {
		return nil, fmt.Errorf("vault login error: %w", err)
	}

	return c, nil
}

func (v *Client) login(ctx context.Context) (*vault.Secret, error) {
	approleSecretID := &approle.SecretID{
		FromString: v.parameters.ApproleSecretID,
	}

	appRoleAuth, err := approle.NewAppRoleAuth(
		v.parameters.ApproleRoleID,
		approleSecretID,
		approle.WithWrappingToken(), // only required if the SecretID is response-wrapped
	)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize approle authentication method: %w", err)
	}

	v.Client.Logical().List("secret/data")
	authInfo, err := v.Client.Auth().Login(ctx, appRoleAuth)
	if err != nil {
		return nil, fmt.Errorf("unable to login using approle auth method: %w", err)
	}
	if authInfo == nil {
		return nil, fmt.Errorf("no approle info was returned after login")
	}
	return authInfo, nil
}
