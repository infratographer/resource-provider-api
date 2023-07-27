package api_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/testing/auth"
)

func TestJWTEnabledLoadbalancerGETWithAuthClient(t *testing.T) {
	oauthCLI, issuer, oAuthClose := auth.OAuthTestClient("urn:test:resourceprovider", "")
	defer oAuthClose()

	srv, err := newTestServer(
		withAuthConfig(
			&echojwtx.AuthConfig{
				Issuer: issuer,
			},
		),
		withPermissions(
			permissions.WithDefaultChecker(permissions.DefaultAllowChecker),
		),
	)

	require.NoError(t, err)
	require.NotNil(t, srv)

	defer srv.Close()

	ctx := context.Background()
	lb1 := (&ResourceProviderBuilder{}).MustNew(ctx)

	resp, err := graphTestClient(
		withGraphClientHTTPClient(oauthCLI),
		withGraphClientServerURL(srv.URL+"/query"),
	).GetResourceProvider(ctx, lb1.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, lb1.ID, resp.ResourceProvider.ID)
}

func TestJWTENabledLoadbalancerGETWithDefaultClient(t *testing.T) {
	_, issuer, oAuthClose := auth.OAuthTestClient("urn:test:loadbalancer", "")
	defer oAuthClose()

	srv, err := newTestServer(
		withAuthConfig(
			&echojwtx.AuthConfig{
				Issuer: issuer,
			},
		),
		withPermissions(
			permissions.WithDefaultChecker(permissions.DefaultAllowChecker),
		),
	)

	require.NoError(t, err)
	require.NotNil(t, srv)

	defer srv.Close()

	ctx := context.Background()
	rp1 := (&ResourceProviderBuilder{}).MustNew(ctx)

	resp, err := graphTestClient(
		withGraphClientHTTPClient(http.DefaultClient),
		withGraphClientServerURL(srv.URL+"/query"),
	).GetResourceProvider(ctx, rp1.ID)

	require.Error(t, err, "Expected an authorization error")
	require.Nil(t, resp)
	assert.ErrorContains(t, err, `{"networkErrors":{"code":401`)
}
