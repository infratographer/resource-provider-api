package api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/resource-provider-api/internal/ent/generated"
	"go.infratographer.com/resource-provider-api/internal/testclient"
)

func TestQuery_resourceProvider(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()
	organizationalUnitID := gidx.MustNewID("testtnt")
	rp1 := (&ResourceProviderBuilder{
		OrganizationalUnitID: organizationalUnitID,
	}).MustNew(ctx)
	rp2 := (&ResourceProviderBuilder{
		Description:          gofakeit.HackerPhrase(),
		OrganizationalUnitID: organizationalUnitID,
	}).MustNew(ctx)

	testCases := []struct {
		name                     string
		queryID                  gidx.PrefixedID
		hasDescription           bool
		hasOrganizationalUnitID  bool
		expextedResourceProvider *generated.ResourceProvider
		errorMsg                 string
	}{
		{
			name:                     "happy path td1",
			queryID:                  rp1.ID,
			expextedResourceProvider: rp1,
		},
		{
			name:                     "happy path td2",
			queryID:                  rp2.ID,
			hasDescription:           true,
			expextedResourceProvider: rp2,
		},
		{
			name:     "invalid-id",
			queryID:  gidx.MustNewID("testing"),
			errorMsg: "resource_provider not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			resp, err := client.GetResourceProvider(ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ResourceProvider)
			assert.EqualValues(t, tc.expextedResourceProvider.ID, resp.ResourceProvider.ID)
			if tc.hasDescription {
				assert.Equal(t, tc.expextedResourceProvider.Description, *resp.ResourceProvider.Description)
			}
			if tc.hasOrganizationalUnitID {
				assert.Equal(t, tc.expextedResourceProvider.OrganizationalUnitID, *&resp.ResourceProvider.OrganizationalUnit.ID)
			}
		})
	}
}

func Test_HappyPath(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()
	organizationalUnitID := gidx.MustNewID("testtnt")

	t.Run("Create + List + Update + Delete", func(t *testing.T) {
		td, err := client.ResourceProviderCreate(ctx, testclient.CreateResourceProviderInput{
			Name:                 gofakeit.JobTitle(),
			Description:          nil,
			OrganizationalUnitID: organizationalUnitID,
		})
		require.NoError(t, err)
		require.NotNil(t, td)

		list, err := client.ListResourceProviders(ctx, organizationalUnitID, nil)
		require.NoError(t, err)
		require.NotNil(t, list)
		assert.Len(t, list.Entities[0].ResourceProvider.Edges, 1)

		tdUpdate, err := client.ResourceProviderUpdate(ctx, td.ResourceProviderCreate.ResourceProvider.ID, testclient.UpdateResourceProviderInput{
			Name:        newString(gofakeit.Name()),
			Description: newString(gofakeit.HackerPhrase()),
		})

		require.NoError(t, err)
		require.NotNil(t, tdUpdate)

		assert.NotEqual(t, td.ResourceProviderCreate.ResourceProvider.Name, tdUpdate.ResourceProviderUpdate.ResourceProvider.Name)
		assert.NotEqual(t, td.ResourceProviderCreate.ResourceProvider.Description, tdUpdate.ResourceProviderUpdate.ResourceProvider.Description)

		deleteID, err := client.ResourceProviderDelete(ctx, td.ResourceProviderCreate.ResourceProvider.ID)
		require.NoError(t, err)
		require.NotNil(t, deleteID)
		assert.EqualValues(t, td.ResourceProviderCreate.ResourceProvider.ID, deleteID.ResourceProviderDelete.DeletedID)
	})
}
