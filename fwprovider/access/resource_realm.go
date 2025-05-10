/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package access

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/bpg/terraform-provider-proxmox/fwprovider/config"
	"github.com/bpg/terraform-provider-proxmox/proxmox"
)

var (
	_ resource.Resource                     = (*realmResource)(nil)
	_ resource.ResourceWithConfigure        = (*realmResource)(nil)
	_ resource.ResourceWithImportState      = (*realmResource)(nil)
	_ resource.ResourceWithConfigValidators = (*realmResource)(nil)
)

type realmResource struct {
	client proxmox.Client
}

// NewRealmResource creates a new Realm resource.
func NewRealmResource() resource.Resource {
	return &realmResource{}
}

func (r *realmResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	cfg, ok := req.ProviderData.(config.Resource)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected config.Resource, got: %T", req.ProviderData),
		)

		return
	}

	r.client = cfg.Client
}

func (r *realmResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
			path.MatchRoot("realm"),
		),
	}
}

func (r *realmResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages Realms on the Proxmox cluster",
		MarkdownDescription: "Manages Realms on the Proxmox cluster.\n\n" +
			"Realms are used for user authentication to a Proxmox node or cluster.\n",
		Attributes: map[string]schema.Attribute{
			"realm": schema.StringAttribute{
				Description: "The ID of the realm which is of type string",
				Optional:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				Description: "Realm type.",
				Optional:    false,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^ad$|^ldap$|^openid$|^pam$|^pve$`),
						"must be one of [`ad`, `ldap`, `openid`, `pam`, `pve`]",
					),
				},
			},
			"acr-values": schema.StringAttribute{
				Description: "Specifies the Authentication Context Class Reference values" +
					"that theAuthorization Server is being requested to use for the Auth Request.",
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[^\x00-\x1F\x7F <>#"]*$`),
						"Not a valid ACR value",
					),
				},
			},
			"autocreate": schema.BoolAttribute{
				Description: "Automatically create users if they do not exist.\n" +
					"default is 0 or false",
				Optional: true,
			},
			"base_dn": schema.StringAttribute{
				Description: "LDAP base domain name",
				Optional:    true,
			},
			"bind_dn": schema.StringAttribute{
				Description: "LDAP bind domain name",
				Optional:    true,
			},
			"capath": schema.StringAttribute{
				Description: "Path to the CA certificate store",
				Optional:    true,
			},
			"case-sensitive": schema.BoolAttribute{
				Description: "username is case-sensitive\n" +
					"default is 1 or true",
				Optional: true,
			},
			"cert": schema.StringAttribute{
				Description: "Path to the client certificate",
				Optional:    true,
			},
			"certkey": schema.StringAttribute{
				Description: "Path to the client certificate key",
				Optional:    true,
			},
			"check-connection": schema.BoolAttribute{
				Description: "Check bind connection to the server. \n" +
					"default is 0 or false",
				Optional: true,
			},
			"client-id": schema.StringAttribute{
				Description: "OpenID Client ID",
				Optional:    true,
			},
			"client-key": schema.StringAttribute{
				Description: "OpenID Client Key",
				Optional:    true,
			},
			"comment": schema.StringAttribute{
				Description: "Description in Proxmox",
				Optional:    true,
			},
			"domain": schema.BoolAttribute{
				Description: "Use this as default realm",
				Optional:    true,
			},
			"filter": schema.StringAttribute{
				Description: "LDAP filter for user sync.",
				Optional:    true,
			},
			"group_classes": schema.StringAttribute{
				Description: "The objectclasses for groups." +
					"default is `groupOfNames, group, univentionGroup, ipausergroup`",
				Optional: true,
			},
			"group_dn": schema.StringAttribute{
				Description: "LDAP base domain name for group sync. If not set, the base_dn will be used.",
				Optional:    true,
			},
			"group_filter": schema.StringAttribute{
				Description: "LDAP filter for group sync.",
				Optional:    true,
			},
			"group_name_attr": schema.StringAttribute{
				Description: "LDAP attribute representing a groups name. If not set or found," +
					"the first value of the DN will be used as name.",
				Optional: true,
			},
			"groups-autocreate": schema.BoolAttribute{
				Description: "Automatically create groups if they do not exist. \n" +
					"default is 0 or false",
				Optional: true,
			},
			"groups-claim": schema.StringAttribute{
				Description: "OpenID claim used to retrieve groups with.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`[A-Za-z0-9._-]+`), // from API docs (?^:[A-Za-z0-9.-_]+) in perl syntax
						"alphanumeric, period, hyphen, and underscore are the only valid characters",
					),
				},
			},
			"groups-overwrite": schema.BoolAttribute{
				Description: "All groups will be overwritten for the user on login. \n" +
					"default is 0 or false",
				Optional: true,
			},
			"issuer-url": schema.StringAttribute{
				Description: "OpenID Issuer Url",
				Optional:    true,
			},
			"mode": schema.StringAttribute{
				Description: "LDAP protocol mode.\n" +
					"one of [`ldap`,`ldaps`,`ldap+starttls`]\n" +
					"default is `ldap`",
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^ldap$|^ldaps$|^ldap\+starttls$`),
						"must be one of [`ldap`,`ldaps`,`ldap+starttls`]",
					),
				},
			},
			"password": schema.StringAttribute{
				Description: "LDAP bind password. Will be stored in '/etc/pve/priv/realm/<REALM>.pw'.",
				Optional:    true,
			},
			"port": schema.Int32Attribute{
				Description: "Server port.",
				Optional:    true,
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
					int32validator.AtLeast(65535),
				},
			},
			"promt": schema.StringAttribute{
				Description: "Specifies whether the Authorization Server prompts the End-User" +
					"for reauthentication and consent.",
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`(?:none|login|consent|select_account|\S+)`),
						"must be a string",
					),
				},
			},
			"query-userinfo": schema.BoolAttribute{
				Description: "Enables querying the userinfo endpoint for claims values. \n" +
					"default is 1 or true",
				Optional: true,
			},
			"scopes": schema.StringAttribute{
				Description: "Specifies the scopes (user details) that should be authorized and returned," +
					"for example 'email' or 'profile'.\n" +
					"default is `email profile`",
				Optional: true,
			},
			"secure": schema.BoolAttribute{
				Description: "Use secure LDAPS protocol. DEPRECATED: use 'mode' instead.",
				Optional:    true,
			},
			"server1": schema.StringAttribute{
				Description: "Server IP address (or DNS name)",
				Optional:    true,
			},
			"server2": schema.StringAttribute{
				Description: "Fallback Server IP address (or DNS name)",
				Optional:    true,
			},
			"sslversion": schema.StringAttribute{
				Description: "LDAPS TLS/SSL version. It's not recommended to use version older than 1.2!",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^tlsv1$|^tlsv1_1$|^tlsv1_2$|^tlsv1_3$`),
						"must be one of [`tlsv1`, `tlsv1_1`, `tlsv1_2`, `tlsv1_3`]",
					),
				},
			},
			"sync-defaults-options": schema.StringAttribute{
				Description: "The default options for behavior of synchronizations." +
					"[enable-new=<1|0>]\n" +
					"[,full=<1|0>]\n" +
					"[,purge=<1|0>]\n" +
					"[,remove-vanished=([acl];[properties];[entry])|none]\n" +
					"[,scope=<users|groups|both>]",
				Optional: true,
			},
			"sync_attributes": schema.StringAttribute{
				Description: "Comma separated list of key=value pairs for specifying which\n" +
					"LDAP attributes map to which PVE user field. For example, to map the LDAP attribute\n" +
					"'mail' to PVEs 'email', write  'email=mail'. By default, each PVE user field is\n" +
					"represented  by an LDAP attribute of the same name.",
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`\w+=[^,]+(,\s*\w+=[^,]+)*`),
						"must be a comma separated list of key=vault pairs",
					),
				},
			},
			"tfa": schema.StringAttribute{
				Description: "Use Two-factor authentication. \n" +
					"type=<TFATYPE>\n" +
					"[,digits=<COUNT>]\n" +
					"[,id=<ID>]\n" +
					"[,key=<KEY>]\n" +
					"[,step=<SECONDS>]\n" +
					"[,url=<URL>]",
				Optional: true,
			},
			"user_attr": schema.StringAttribute{
				Description: "LDAP user attribute name",
				Optional:    true,
			},
			"user_classes": schema.StringAttribute{
				Description: "The objectclasses for users.\n" +
					"default is `inetorgperson, posixaccount, person, user`",
				Optional: true,
			},
			"username-claim": schema.StringAttribute{
				Description: "OpenID claim used to generate the unique username.",
				Optional:    true,
			},
			"verify": schema.BoolAttribute{
				Description: "Verify the server's SSL certificate" +
					"default is 1 or true",
				Optional: true,
			},
		},
	}
}

func (r *realmResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_realm"
}

func (r *realmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan realmResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body := plan.intoCreateBody()

	err := r.client.Access().CreateRealm(ctx, body)
	if err != nil {
		resp.Diagnostics.AddError("Unable to create Realm", apiCallFailed+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *realmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state realmResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	realms, err := r.client.Access().GetRealm(ctx, state.Realm.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable read Realm", apiCallFailed+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &realms)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *realmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var (
		state realmResourceModel
		plan  realmResourceModel
	)

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	stateBody := state.intoUpdateBody()

	err := r.client.Access().UpdateRealm(ctx, state.Realm.ValueString(), stateBody)
	if err != nil {
		resp.Diagnostics.AddError("Unable to delete old Realm", apiCallFailed+err.Error())
		return
	}

	planBody := plan.intoUpdateBody()

	err = r.client.Access().UpdateRealm(ctx, state.Realm.ValueString(), planBody)
	if err != nil {
		resp.Diagnostics.AddError("Unable to create Realm", apiCallFailed+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *realmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state realmResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Access().DeleteRealm(ctx, state.Realm.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to delete old Realm", apiCallFailed+err.Error())
		return
	}
}

func (r *realmResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("realm"), req, resp)
}
