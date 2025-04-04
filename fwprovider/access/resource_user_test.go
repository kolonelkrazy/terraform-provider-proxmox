//go:build acceptance || all

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package access_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/require"

	"github.com/bpg/terraform-provider-proxmox/fwprovider/test"
	"github.com/bpg/terraform-provider-proxmox/proxmox/access"
)

func TestAccResourceUser(t *testing.T) {
	t.Parallel()

	te := test.InitEnvironment(t)

	userID := fmt.Sprintf("%s@pve", gofakeit.Username())
	te.AddTemplateVars(map[string]any{
		"UserID": userID,
	})

	tests := []struct {
		name  string
		steps []resource.TestStep
	}{
		{"create and update user", []resource.TestStep{
			{
				Config: te.RenderConfig(`resource "proxmox_virtual_environment_user" "user" {
					  comment  			= "Managed by Terraform"
					  email 			= "{{.UserID}}"
					  enabled 			= true
					  expiration_date 	= "2034-01-01T22:00:00Z"
					  first_name 		= "First"
					  last_name 		= "Last"
					  user_id  			= "{{.UserID}}"
				}`),
				Check: test.ResourceAttributes("proxmox_virtual_environment_user.user", map[string]string{
					"comment":         "Managed by Terraform",
					"email":           userID,
					"enabled":         "true",
					"expiration_date": "2034-01-01T22:00:00Z",
					"first_name":      "First",
					"last_name":       "Last",
					"user_id":         userID,
				}),
			},
			{
				Config: te.RenderConfig(`resource "proxmox_virtual_environment_user" "user" {
					  enabled 			= false
					  expiration_date 	= "2035-01-01T22:00:00Z"
					  user_id  			= "{{.UserID}}"
					  first_name 		= "First One"
				}`),
				Check: test.ResourceAttributes("proxmox_virtual_environment_user.user", map[string]string{
					"enabled":         "false",
					"expiration_date": "2035-01-01T22:00:00Z",
					"first_name":      "First One",
					"user_id":         userID,
				}),
			},
			{
				ResourceName:      "proxmox_virtual_environment_user.user",
				ImportState:       true,
				ImportStateVerify: true,
			},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: te.AccProviders,
				Steps:                    tt.steps,
			})
		})
	}
}

func TestAccResourceUserToken(t *testing.T) {
	t.Parallel()

	te := test.InitEnvironment(t)
	userID := fmt.Sprintf("%s@pve", gofakeit.Username())
	tokenName := gofakeit.Word()

	te.AddTemplateVars(map[string]any{
		"UserID":    userID,
		"TokenName": tokenName,
	})

	tests := []struct {
		name     string
		preCheck func()
		steps    []resource.TestStep
	}{
		{
			"create and update user token",
			func() {
				err := te.AccessClient().CreateUser(context.Background(), &access.UserCreateRequestBody{
					ID:       userID,
					Password: gofakeit.Password(true, true, true, true, false, 8),
				})
				require.NoError(t, err)

				t.Cleanup(func() {
					err = te.AccessClient().DeleteUser(context.Background(), userID)
					require.NoError(t, err)
				})
			},
			[]resource.TestStep{
				{
					Config: te.RenderConfig(`resource "proxmox_virtual_environment_user_token" "user_token" {
						comment  			= "Managed by Terraform"
						token_name 			= "{{.TokenName}}"
						user_id  			= "{{.UserID}}"
					}`),
					Check: test.ResourceAttributes("proxmox_virtual_environment_user_token.user_token", map[string]string{
						"comment": "Managed by Terraform",
						"id":      fmt.Sprintf("%s!%s", userID, tokenName),
						"user_id": userID,
						"value":   fmt.Sprintf("%s!%s=.*", userID, tokenName),
					}),
				},
				{
					Config: te.RenderConfig(`resource "proxmox_virtual_environment_user_token" "user_token" {
						comment  			  = "Managed by Terraform 2"
						expiration_date 	  = "2033-01-01T01:01:01Z"
						privileges_separation = false
						token_name 			  = "{{.TokenName}}"
						user_id  			  = "{{.UserID}}"
					}`),
					Check: resource.ComposeTestCheckFunc(
						test.ResourceAttributes("proxmox_virtual_environment_user_token.user_token", map[string]string{
							"comment":               "Managed by Terraform 2",
							"expiration_date":       "2033-01-01T01:01:01Z",
							"privileges_separation": "false",
							"token_name":            tokenName,
							"user_id":               userID,
							"value":                 fmt.Sprintf("%s!%s=.*", userID, tokenName),
						}),
					),
				},
				{
					Config: te.RenderConfig(`resource "proxmox_virtual_environment_user_token" "user_token" {
						comment  			  = "Managed by Terraform 2"
						privileges_separation = false
						token_name 			  = "{{.TokenName}}"
						user_id  			  = "{{.UserID}}"
					}`),
					Check: resource.ComposeTestCheckFunc(
						test.ResourceAttributes("proxmox_virtual_environment_user_token.user_token", map[string]string{
							"comment":               "Managed by Terraform 2",
							"privileges_separation": "false",
							"token_name":            tokenName,
							"user_id":               userID,
							"value":                 fmt.Sprintf("%s!%s=.*", userID, tokenName),
						}),
						test.NoResourceAttributesSet("proxmox_virtual_environment_user_token.user_token", []string{
							"expiration_date",
						}),
					),
				},
				{
					ResourceName:      "proxmox_virtual_environment_user_token.user_token",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateVerifyIgnore: []string{
						"value",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: te.AccProviders,
				PreCheck:                 tt.preCheck,
				Steps:                    tt.steps,
			})
		})
	}
}
