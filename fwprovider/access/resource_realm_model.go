/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package access

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/bpg/terraform-provider-proxmox/proxmox/access"
)

type realmResourceModel struct {
	Realm types.String `tfsdk:"realm"`
	Type  types.String `tfsdk:"type"`

	AcrValues           types.String `tfsdk:"acr-values"`
	AutoCreate          bool         `tfsdk:"autocreate"`
	BaseDn              types.String `tfsdk:"base_dn"`
	BindDn              types.String `tfsdk:"bind_dn"`
	CaPath              types.String `tfsdk:"capath"`
	CaseSensitive       bool         `tfsdk:"case-sensitive"`
	Cert                types.String `tfsdk:"cert"`
	CertKey             types.String `tfsdk:"certkey"`
	CheckConnection     bool         `tfsdk:"check-connection"`
	ClientId            types.String `tfsdk:"client-id"`
	ClientKey           types.String `tfsdk:"client-key"`
	Comment             types.String `tfsdk:"comment"`
	Default             bool         `tfsdk:"default"`
	Domain              types.String `tfsdk:"domain"`
	Filter              types.String `tfsdk:"filter"`
	GroupClasses        types.String `tfsdk:"group_classes"`
	GroupDn             types.String `tfsdk:"group_dn"`
	GroupFilter         types.String `tfsdk:"group_filter"`
	GroupNameAttr       types.String `tfsdk:"group_name_attr"`
	IssuerUrl           types.String `tfsdk:"issuer-url"`
	Mode                types.String `tfsdk:"mode"`
	Password            types.String `tfsdk:"password"`
	Port                int          `tfsdk:"port"`
	Prompt              types.String `tfsdk:"promptly"`
	Scopes              types.String `tfsdk:"scopes"`
	Secure              bool         `tfsdk:"secure"`
	Server1             types.String `tfsdk:"server1"`
	Server2             types.String `tfsdk:"server2"`
	SslVersion          types.String `tfsdk:"sslversion"`
	SyncDefaultsOptions types.String `tfsdk:"sync-defaults-options"`
	SyncAttributes      types.String `tfsdk:"sync_attributes"`
	Tfa                 types.String `tfsdk:"tfa"`
	UserAttr            types.String `tfsdk:"user_attr"`
	UserClasses         types.String `tfsdk:"user_classes"`
	UserNameClaim       types.String `tfsdk:"username-claim"`
	Verify              bool         `tfsdk:"verify"
}

const RealmPromptFormat = `(?:none|login|consent|select_account|\S+)`
const RealmSyncAttributesFormat = `\w+=[^,]+(,\s*\w+=[^,]+)*`
const RealmAcrValuesFormat = `^[^\x00-\x1F\x7F <>#"]*$`
const RealmDomainFormat = `\S+`

func parseValidateMode(id string) (*realmResourceModel, error) {
	/*
	 *	ldap | ldaps | ldap+starttls
	 *
	 */

	model := &realmResourceModel{}
	return model, nil
}

func parseValidateSslVersion(id string) (*realmResourceModel, error) {
	/*
	 *	tlsv1 | tlsv1_1 | tlsv1_2 | tlsv1_3
	 *
	 */

	model := &realmResourceModel{}
	return model, nil
}

func parseValidateSyncDefaultOptions(id string) (*realmResourceModel, error) {
	/*
	 *	[enable-new=<1|0>] [,full=<1|0>] [,purge=<1|0>] [,remove-vanished=([acl];[properties];[entry])|none] [,scope=<users|groups|both>]
	 *
	 */

	model := &realmResourceModel{}
	return model, nil
}

func parseValidatePort(id string) (*realmResourceModel, error) {
	/*
	 * (1 - 65535)
	 *
	 */

	model := &realmResourceModel{}
	return model, nil
}

func parseValidateTFA(id string) (*realmResourceModel, error) {
	/*
	 *	yubico | oath
	 * 	type=<TFATYPE> [,digits=<COUNT>] [,id=<ID>] [,key=<KEY>] [,step=<SECONDS>] [,url=<URL>]
	 *
	 */

	model := &realmResourceModel{}
	return model, nil
}

func parseValidateType(id string) (*realmResourceModel, error) {
	/*
	 *	ad | ldap | openid | pam | pve
	 *
	 */

	model := &realmResourceModel{}
	return model, nil
}

func parseValidateDomain(id string) (*realmResourceModel, error) {

	model := &realmResourceModel{}
	return model, nil
}

func parseValidatePrompt(id string) (*realmResourceModel, error) {

	model := &realmResourceModel{}
	return model, nil
}

func parseValidateACRValues(id string) (*realmResourceModel, error) {

	model := &realmResourceModel{}
	return model, nil
}

func (r *realmResourceModel) intoUpdateBody() *access.RealmUpdateRequestBody {
	body := &access.RealmUpdateRequestBody{
		// TODO
	}
	return body
}
