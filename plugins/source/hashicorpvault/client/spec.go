package client

import "github.com/pkg/errors"

type VaultParameters struct {
	// connection parameters
	Address         string `json:"address"`
	ApproleRoleID   string `json:"approle_role_id"`
	ApproleSecretID string `json:"approle_secret_id"`
}

type Paths struct {
	Pki []string `json:"pki"`
}

type Spec struct {
	Name       string           `json:"name"`
	Parameters *VaultParameters `json:"parameters"`
	Paths      Paths            `json:"paths"`
}

func (s *Spec) Validate() error {
	if s.Name == "" {
		return errors.New("missing vault name in configuration file")
	}
	if s.Parameters == nil {
		return errors.New("missing vault parameters in configuration file")
	}
	if s.Parameters.ApproleRoleID == "" {
		return errors.New("missing vault approle role id in configuration file")
	}
	if s.Parameters.ApproleSecretID == "" {
		return errors.New("missing vault approle secret id in configuration file")
	}
	if s.Paths.Pki == nil {
		return errors.New("missing vault pki paths in configuration file")
	}
	return nil
}
