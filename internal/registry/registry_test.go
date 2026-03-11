package registry

import (
	"context"
	"testing"

	"github.com/georgirtodorov/tenant-control-service/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestService(t *testing.T) {
	t.Run("Create and Get Tenant", func(t *testing.T) {
		ctx := context.Background()
		srv := New()

		createReq := &api.CreateTenantRequest{
			Tenant: &api.Tenant{Name: "Test Tenant One"},
		}

		createResp, err := srv.CreateTenant(ctx, createReq)
		require.NoError(t, err)
		createdTenantID := createResp.Tenant.Id

		getReq := &api.GetTenantRequest{Id: createdTenantID}
		getResp, err := srv.GetTenant(ctx, getReq)
		require.NoError(t, err)
		assert.Equal(t, createdTenantID, getResp.Tenant.Id)
	})

	t.Run("Update Tenant", func(t *testing.T) {
		ctx := context.Background()
		srv := New()

		createReq := &api.CreateTenantRequest{Tenant: &api.Tenant{Name: "Original Name"}}
		createResp, err := srv.CreateTenant(ctx, createReq)
		require.NoError(t, err)
		createdTenantID := createResp.Tenant.Id

		updatedTenant := &api.Tenant{Id: createdTenantID, Name: "Updated Name"}
		updateReq := &api.UpdateTenantRequest{Tenant: updatedTenant}
		_, err = srv.UpdateTenant(ctx, updateReq)
		require.NoError(t, err)

		getReq := &api.GetTenantRequest{Id: createdTenantID}
		getResp, err := srv.GetTenant(ctx, getReq)
		require.NoError(t, err)
		assert.Equal(t, "Updated Name", getResp.Tenant.Name)

		_, err = srv.UpdateTenant(ctx, &api.UpdateTenantRequest{Tenant: &api.Tenant{Id: "non-existent-id"}})
		require.Error(t, err)
		st, _ := status.FromError(err)
		assert.Equal(t, codes.NotFound, st.Code())
	})

	t.Run("Delete Tenant", func(t *testing.T) {
		ctx := context.Background()
		srv := New()

		// 1. Create a tenant to delete
		createReq := &api.CreateTenantRequest{Tenant: &api.Tenant{Name: "To Be Deleted"}}
		createResp, err := srv.CreateTenant(ctx, createReq)
		require.NoError(t, err)
		createdTenantID := createResp.Tenant.Id

		// 2. Delete the tenant
		deleteReq := &api.DeleteTenantRequest{Id: createdTenantID}
		deleteResp, err := srv.DeleteTenant(ctx, deleteReq)
		require.NoError(t, err, "DeleteTenant should not return an error")
		assert.True(t, deleteResp.Success, "Delete response should indicate success")

		// 3. Try to get the deleted tenant, expecting a NotFound error
		_, err = srv.GetTenant(ctx, &api.GetTenantRequest{Id: createdTenantID})
		require.Error(t, err, "Getting a deleted tenant should return an error")
		st, ok := status.FromError(err)
		require.True(t, ok, "Error should be a gRPC status error")
		assert.Equal(t, codes.NotFound, st.Code(), "Error code should be NotFound after deletion")

		// 4. Test deleting a non-existent tenant
		_, err = srv.DeleteTenant(ctx, &api.DeleteTenantRequest{Id: "non-existent-id"})
		require.Error(t, err, "Deleting a non-existent tenant should return an error")
		st, ok = status.FromError(err)
		require.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code(), "Error code for deleting non-existent tenant should be NotFound")
	})
}
