package registry

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/georgirtodorov/tenant-control-service/api"
)

// Service implements the Registry server
type Service struct {
	api.UnimplementedRegistryServer
	mu      sync.Mutex
	tenants map[string]*api.Tenant
}

// New creates a new Service
func New() *Service {
	return &Service{
		tenants: make(map[string]*api.Tenant),
	}
}

// CreateTenant creates a new tenant
func (s *Service) CreateTenant(ctx context.Context, req *api.CreateTenantRequest) (*api.CreateTenantResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// simple validation
	if req.Tenant == nil {
		return nil, status.Errorf(codes.InvalidArgument, "tenant is required")
	}

	// generate a new UUID for the tenant
	tenantID := uuid.New().String()
	req.Tenant.Id = tenantID

	s.tenants[tenantID] = req.Tenant

	fmt.Printf("Tenant created: %+v\n", req.Tenant)

	return &api.CreateTenantResponse{Tenant: req.Tenant}, nil
}

// GetTenant retrieves a tenant by ID
func (s *Service) GetTenant(ctx context.Context, req *api.GetTenantRequest) (*api.GetTenantResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tenant, ok := s.tenants[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "tenant with id %s not found", req.Id)
	}

	return &api.GetTenantResponse{Tenant: tenant}, nil
}

// UpdateTenant updates an existing tenant
func (s *Service) UpdateTenant(ctx context.Context, req *api.UpdateTenantRequest) (*api.UpdateTenantResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Tenant == nil {
		return nil, status.Errorf(codes.InvalidArgument, "tenant is required")
	}

	_, ok := s.tenants[req.Tenant.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "tenant with id %s not found", req.Tenant.Id)
	}

	s.tenants[req.Tenant.Id] = req.Tenant

	fmt.Printf("Tenant updated: %+v\n", req.Tenant)

	return &api.UpdateTenantResponse{Tenant: req.Tenant}, nil
}

// DeleteTenant deletes a tenant by ID
func (s *Service) DeleteTenant(ctx context.Context, req *api.DeleteTenantRequest) (*api.DeleteTenantResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.tenants[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "tenant with id %s not found", req.Id)
	}

	delete(s.tenants, req.Id)

	fmt.Printf("Tenant deleted: %s\n", req.Id)

	return &api.DeleteTenantResponse{Success: true}, nil
}
