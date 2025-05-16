/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package access

import (
	"github.com/bpg/terraform-provider-proxmox/proxmox/access"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type realmResourceModel struct {
	Realm types.String `tfsdk:"realm"`
	Type  types.String `tfsdk:"type"`

	AcrValues           types.String `tfsdk:"acr-values"`
	AutoCreate          types.Bool   `tfsdk:"autocreate"`
	BaseDn              types.String `tfsdk:"base_dn"`
	BindDn              types.String `tfsdk:"bind_dn"`
	CaPath              types.String `tfsdk:"capath"`
	CaseSensitive       types.Bool   `tfsdk:"case-sensitive"`
	Cert                types.String `tfsdk:"cert"`
	CertKey             types.String `tfsdk:"certkey"`
	CheckConnection     types.Bool   `tfsdk:"check-connection"`
	ClientId            types.String `tfsdk:"client-id"`
	ClientKey           types.String `tfsdk:"client-key"`
	Comment             types.String `tfsdk:"comment"`
	Default             types.Bool   `tfsdk:"default"`
	Domain              types.String `tfsdk:"domain"`
	Filter              types.String `tfsdk:"filter"`
	GroupClasses        types.String `tfsdk:"group_classes"`
	GroupDn             types.String `tfsdk:"group_dn"`
	GroupFilter         types.String `tfsdk:"group_filter"`
	GroupNameAttr       types.String `tfsdk:"group_name_attr"`
	GroupsAutoCreate    types.Bool   `tfsdk:"groups-autocreate"`
	GroupsClaim         types.Bool   `tfsdk:"groups-claim"`
	GroupsOverwrite     types.Bool   `tfsdk:"groups-overwrite"`
	IssuerUrl           types.String `tfsdk:"issuer-url"`
	Mode                types.String `tfsdk:"mode"`
	Password            types.String `tfsdk:"password"`
	Port                types.Int32  `tfsdk:"port"`
	Prompt              types.String `tfsdk:"prompt"`
	QueryUserInfo       types.Bool   `tfsdk:"query-userinfo"`
	Scopes              types.String `tfsdk:"scopes"`
	Secure              types.Bool   `tfsdk:"secure"`
	Server1             types.String `tfsdk:"server1"`
	Server2             types.String `tfsdk:"server2"`
	SslVersion          types.String `tfsdk:"sslversion"`
	SyncDefaultsOptions types.String `tfsdk:"sync-defaults-options"`
	SyncAttributes      types.String `tfsdk:"sync_attributes"`
	Tfa                 types.String `tfsdk:"tfa"`
	UserAttr            types.String `tfsdk:"user_attr"`
	UserClasses         types.String `tfsdk:"user_classes"`
	UserNameClaim       types.String `tfsdk:"username-claim"`
	Verify              types.Bool   `tfsdk:"verify"`
}

func (r *realmResourceModel) intoUpdateBody() *access.RealmUpdateRequestBody {
	body := &access.RealmUpdateRequestBody{}
	return body
}

func (r *realmResourceModel) intoCreateBody() *access.RealmCreateRequestBody {
	body := &access.RealmCreateRequestBody{}
	return body
}
