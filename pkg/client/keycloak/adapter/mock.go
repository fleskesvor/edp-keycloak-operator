package adapter

import (
	"context"

	"github.com/Nerzal/gocloak/v8"
	"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1"
	"github.com/epam/edp-keycloak-operator/pkg/client/keycloak/dto"
	"github.com/epam/edp-keycloak-operator/pkg/model"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
	ExportTokenResult []byte
	ExportTokenErr    error
}

func (m *Mock) PutDefaultIdp(realm *dto.Realm) error {
	return m.Called(realm).Error(0)
}

func (m *Mock) ExistRealm(realm string) (bool, error) {
	args := m.Called(realm)
	if args.Get(0) == nil {
		return false, args.Error(1)
	}

	return args.Bool(0), args.Error(1)
}

func (m *Mock) CreateRealmWithDefaultConfig(realm *dto.Realm) error {
	args := m.Called(realm)
	return args.Error(0)
}

func (m *Mock) ExistCentralIdentityProvider(realm *dto.Realm) (bool, error) {
	args := m.Called(realm)
	if args.Get(0) == nil {
		return false, args.Error(1)
	}
	res := args.Bool(0)
	return res, args.Error(1)
}

func (m *Mock) CreateCentralIdentityProvider(realm *dto.Realm, client *dto.Client) error {
	return m.Called(realm, client).Error(0)
}

func (m *Mock) ExistClient(clientID, realm string) (bool, error) {
	args := m.Called(clientID, realm)
	if args.Get(0) == nil {
		return false, args.Error(1)
	}
	res := args.Bool(0)
	return res, args.Error(1)
}

func (m *Mock) CreateClient(client *dto.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *Mock) ExistClientRole(role *dto.Client, clientRole string) (bool, error) {
	panic("implement me")
}

func (m *Mock) CreateClientRole(role *dto.Client, clientRole string) error {
	panic("implement me")
}

func (m *Mock) ExistRealmRole(realm string, role string) (bool, error) {
	args := m.Called(realm, role)
	return args.Bool(0), args.Error(1)
}

func (m *Mock) CreateIncludedRealmRole(realm string, role *dto.IncludedRealmRole) error {
	args := m.Called(realm, role)
	return args.Error(0)
}

func (m *Mock) CreatePrimaryRealmRole(realm string, role *dto.PrimaryRealmRole) (string, error) {
	args := m.Called(realm, role)
	return args.String(0), args.Error(1)
}

func (m *Mock) ExistRealmUser(realmName string, user *dto.User) (bool, error) {
	called := m.Called(realmName, user)
	return called.Bool(0), called.Error(1)
}

func (m *Mock) CreateRealmUser(realmName string, user *dto.User) error {
	return m.Called(realmName, user).Error(0)
}

func (m *Mock) HasUserClientRole(realmName string, clientId string, user *dto.User, role string) (bool, error) {
	called := m.Called(realmName, clientId, user, role)
	return called.Bool(0), called.Error(1)
}

func (m *Mock) GetOpenIdConfig(realm *dto.Realm) (string, error) {
	args := m.Called(realm)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}

	return args.String(0), args.Error(1)
}

func (m *Mock) AddClientRoleToUser(realmName string, clientId string, user *dto.User, role string) error {
	return m.Called(realmName, clientId, user, role).Error(0)
}

func (m *Mock) GetClientID(clientID, realm string) (string, error) {
	args := m.Called(clientID, realm)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	res := args.String(0)
	return res, args.Error(1)
}

func (m *Mock) MapRoleToUser(realmName string, user dto.User, role string) error {
	panic("implement me")
}

func (m *Mock) ExistMapRoleToUser(realmName string, user dto.User, role string) (*bool, error) {
	panic("implement me")
}

func (m *Mock) AddRealmRoleToUser(ctx context.Context, realmName, username, roleName string) error {
	return m.Called(realmName, username, roleName).Error(0)
}

func (m *Mock) DeleteClient(kkClientID string, realName string) error {
	return m.Called(kkClientID, realName).Error(0)
}

func (m *Mock) DeleteRealmRole(realm, roleName string) error {
	return m.Called(realm, roleName).Error(0)
}

func (m *Mock) DeleteRealm(realmName string) error {
	return m.Called(realmName).Error(0)
}

func (m *Mock) GetClientScope(scopeName, realmName string) (*model.ClientScope, error) {
	called := m.Called(scopeName, realmName)
	if err := called.Error(1); err != nil {
		return nil, err
	}

	return called.Get(0).(*model.ClientScope), nil
}

func (m *Mock) HasUserRealmRole(realmName string, user *dto.User, role string) (bool, error) {
	called := m.Called(realmName, user, role)
	return called.Bool(0), called.Error(1)
}

func (m *Mock) LinkClientScopeToClient(clientName, scopeId, realmName string) error {
	panic("implement me")
}

func (m *Mock) PutClientScopeMapper(clientName, scopeId, realmName string) error {
	panic("implement me")
}

func (m *Mock) SyncClientProtocolMapper(
	client *dto.Client, crMappers []gocloak.ProtocolMapperRepresentation, addOnly bool) error {
	return m.Called(client, crMappers, addOnly).Error(0)
}

func (m *Mock) SyncRealmRole(realmName string, role *dto.PrimaryRealmRole) error {
	return m.Called(realmName, role).Error(0)
}

func (m *Mock) SyncServiceAccountRoles(realm, clientID string, realmRoles []string,
	clientRoles map[string][]string, addOnly bool) error {
	return m.Called(realm, clientID, realmRoles, clientRoles, addOnly).Error(0)
}

func (m *Mock) SyncRealmGroup(realmName string, spec *v1alpha1.KeycloakRealmGroupSpec) (string, error) {
	called := m.Called(realmName, spec)
	return called.String(0), called.Error(1)
}

func (m *Mock) DeleteGroup(realm, groupName string) error {
	return m.Called(realm, groupName).Error(0)
}

func (m *Mock) SyncRealmIdentityProviderMappers(realmName string,
	mappers []dto.IdentityProviderMapper) error {
	return m.Called(realmName, mappers).Error(0)
}

func (m *Mock) DeleteAuthFlow(realmName, alias string) error {
	return m.Called(realmName, alias).Error(0)
}

func (m *Mock) SyncAuthFlow(realmName string, flow *KeycloakAuthFlow) error {
	return m.Called(realmName, flow).Error(0)
}

func (m *Mock) SetRealmBrowserFlow(realmName string, flowAlias string) error {
	return m.Called(realmName, flowAlias).Error(0)
}
func (m *Mock) UpdateRealmSettings(realmName string, realmSettings *RealmSettings) error {
	return m.Called(realmName, realmSettings).Error(0)
}

func (m *Mock) SyncRealmUser(ctx context.Context, realmName string, user *KeycloakUser, addOnly bool) error {
	return m.Called(realmName, user, addOnly).Error(0)
}

func (m *Mock) SetServiceAccountAttributes(realm, clientID string, attributes map[string]string, addOnly bool) error {
	return m.Called(realm, clientID, attributes, addOnly).Error(0)
}

func (m *Mock) CreateClientScope(ctx context.Context, realmName string, scope *ClientScope) (string, error) {
	called := m.Called(realmName, scope)
	if err := called.Error(1); err != nil {
		return "", err
	}

	return called.String(0), nil
}

func (m *Mock) DeleteClientScope(ctx context.Context, realmName, scopeID string) error {
	return m.Called(realmName, scopeID).Error(0)
}

func (m *Mock) UpdateClientScope(ctx context.Context, realmName, scopeID string, scope *ClientScope) error {
	return m.Called(realmName, scopeID, scope).Error(0)
}

func (m *Mock) SetRealmEventConfig(realmName string, eventConfig *RealmEventConfig) error {
	return m.Called(realmName, eventConfig).Error(0)
}

func (m *Mock) ExportToken() ([]byte, error) {
	return m.ExportTokenResult, m.ExportTokenErr
}
