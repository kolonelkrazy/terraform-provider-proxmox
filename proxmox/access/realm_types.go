/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package access

// RealmCreateRequestBody contains the data for a realm create request.
type RealmCreateRequestBody struct {
	Realm               string `json:"realm"                           url:"realm"`
	Type                string `json:"type"                            url:"type"` /*enum*/
	AcrValues           string `json:"acr-values,omitempty"            url:"acr-values,omitempty"`
	AutoCreate          bool   `json:"autocreate,omitempty"            url:"autocreate,omitempty,int"`
	BaseDn              string `json:"base_dn,omitempty"               url:"base_dn,omitempty"`
	BindDn              string `json:"bind_dn,omitempty"               url:"bind_dn,omitempty"`
	CaPath              string `json:"capath,omitempty"                url:"capath,omitempty"`
	CaseSensitive       bool   `json:"case-sensitive,omitempty"        url:"case-sensitive,omitempty,int"`
	Cert                string `json:"cert,omitempty"                  url:"cert,omitempty"`
	CertKey             string `json:"certkey,omitempty"               url:"certkey,omitempty"`
	CheckConnection     bool   `json:"check-connection,omitempty"      url:"check-connection,omitempty,int"`
	ClientId            string `json:"client-id,omitempty"             url:"client-id,omitempty"`
	ClientKey           string `json:"client-key,omitempty"            url:"client-key,omitempty"`
	Comment             string `json:"comment,omitempty"               url:"comment,omitempty"`
	Default             bool   `json:"default,omitempty"               url:"default,omitempty,int"`
	Domain              string `json:"domain,omitempty"                url:"domain,omitempty"`
	Filter              string `json:"filter,omitempty"                url:"filter,omitempty"`
	GroupClasses        string `json:"group_classes,omitempty"         url:"group_classes,omitempty"`
	GroupDn             string `json:"group_dn,omitempty"              url:"group_dn,omitempty"`
	GroupFilter         string `json:"group_filter,omitempty"          url:"group_filter,omitempty"`
	GroupNameAttr       string `json:"group_name_attr,omitempty"       url:"group_name_attr,omitempty"`
	GroupsAutoCreate    bool   `json:"groups-autocreate,omitempty"     url:"groups-autocreate,omitempty"`
	GroupsClaim         bool   `json:"groups-claim,omitempty"          url:"groups-claim,omitempty"`
	GroupsOverwrite     bool   `json:"groups-overwrite,omitempty"      url:"groups-overwrite,omitempty"`
	IssuerUrl           string `json:"issuer-url,omitempty"            url:"issuer-url,omitempty"`
	Mode                string `json:"mode,omitempty"                  url:"mode,omitempty"` /*enum*/
	Password            string `json:"password,omitempty"              url:"password,omitempty"`
	Port                int    `json:"port,omitempty"                  url:"port,omitempty,int"`
	Prompt              string `json:"prompt,omitempty"                url:"prompt,omitempty"`
	Scopes              string `json:"scopes,omitempty"                url:"scopes,omitempty"`
	Secure              bool   `json:"secure,omitempty"                url:"secure,omitempty,int"`
	Server1             string `json:"server1,omitempty"               url:"server1,omitempty"`
	Server2             string `json:"server2,omitempty"               url:"server2,omitempty"`
	SslVersion          string `json:"sslversion,omitempty"            url:"sslversion,omitempty"` /*enum*/
	SyncDefaultsOptions string `json:"sync-defaults-options,omitempty" url:"sync-defaults-options,omitempty"`
	SyncAttributes      string `json:"sync_attributes,omitempty"       url:"sync_attributes,omitempty"`
	Tfa                 string `json:"tfa,omitempty"                   url:"tfa,omitempty"` /*enum*/
	UserAttr            string `json:"user_attr,omitempty"             url:"user_attr,omitempty"`
	UserClasses         string `json:"user_classes,omitempty"          url:"user_classes,omitempty"`
	UsernameClaim       string `json:"username-claim,omitempty"        url:"username-claim,omitempty"`
	Verify              bool   `json:"verify,omitempty"                url:"verify,omitempty,int"`
}

type RealmGetResponseBody struct {
	Data []*RealmGetResponseData `json:"data,omitempty"`
}

// RealmGetResponseBody contains the body from an a realm get response.
type RealmGetResponseData struct {
	AcrValues           string `json:"acr-values,omitempty"`
	AutoCreate          bool   `json:"autocreate,omitempty"`
	BaseDn              string `json:"base_dn,omitempty"`
	BindDn              string `json:"bind_dn,omitempty"`
	CaPath              string `json:"capath,omitempty"`
	CaseSensitive       bool   `json:"case-sensitive,omitempty"`
	Cert                string `json:"cert,omitempty"`
	CertKey             string `json:"certkey,omitempty"`
	CheckConnection     bool   `json:"check-connection,omitempty"`
	ClientId            string `json:"client-id,omitempty"`
	ClientKey           string `json:"client-key,omitempty"`
	Comment             string `json:"comment,omitempty"`
	Default             bool   `json:"default,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Filter              string `json:"filter,omitempty"`
	GroupClasses        string `json:"group_classes,omitempty"`
	GroupDn             string `json:"group_dn,omitempty"`
	GroupFilter         string `json:"group_filter,omitempty"`
	GroupNameAttr       string `json:"group_name_attr,omitempty"`
	GroupsAutoCreate    bool   `json:"groups-autocreate,omitempty"`
	GroupsClaim         bool   `json:"groups-claim,omitempty"`
	GroupsOverwrite     bool   `json:"groups-overwrite,omitempty"`
	IssuerUrl           string `json:"issuer-url,omitempty"`
	Mode                string `json:"mode,omitempty"` /*enum*/
	Password            string `json:"password,omitempty"`
	Port                int    `json:"port,omitempty"`
	Prompt              string `json:"prompt,omitempty"`
	Scopes              string `json:"scopes,omitempty"`
	Secure              bool   `json:"secure,omitempty"`
	Server1             string `json:"server1,omitempty"`
	Server2             string `json:"server2,omitempty"`
	SslVersion          string `json:"sslversion,omitempty"` /*enum*/
	SyncDefaultsOptions string `json:"sync-defaults-options,omitempty"`
	SyncAttributes      string `json:"sync_attributes,omitempty"`
	Tfa                 string `json:"tfa,omitempty"` /*enum*/
	Type                string `json:"type"`          /*enum*/
	UserAttr            string `json:"user_attr,omitempty"`
	UserClasses         string `json:"user_classes,omitempty"`
	UsernameClaim       string `json:"username-claim,omitempty"`
	Verify              bool   `json:"verify,omitempty"`
}

// RealmListResponseBody contains the body from an a realm list response.
type RealmListResponseBody struct {
	Data []*RealmListResponseData `json:"data,omitempty"`
}

// RealmListResponseData contains the data from an a realm list response.
type RealmListResponseData struct {
	Realm   string `json:"realm"             url:"realm"`
	Type    string `json:"type"              url:"type"`
	Comment string `json:"comment,omitempty" url:"comment,omitempty"`
	Tfa     string `json:"tfa,omitempty"     url:"tfa,omitempty"`
}

// RealmUpdateRequestBody contains the data for an a realm update request.
type RealmUpdateRequestBody struct {
	AcrValues           string `json:"acr-values,omitempty"            url:"acr-values,omitempty"`
	AutoCreate          bool   `json:"autocreate,omitempty"            url:"autocreate,omitempty,int"`
	BaseDn              string `json:"base_dn,omitempty"               url:"base_dn,omitempty"`
	BindDn              string `json:"bind_dn,omitempty"               url:"bind_dn,omitempty"`
	CaPath              string `json:"capath,omitempty"                url:"capath,omitempty"`
	CaseSensitive       bool   `json:"case-sensitive,omitempty"        url:"case-sensitive,omitempty,int"`
	Cert                string `json:"cert,omitempty"                  url:"cert,omitempty"`
	CertKey             string `json:"certkey,omitempty"               url:"certkey,omitempty"`
	CheckConnection     bool   `json:"check-connection,omitempty"      url:"check-connection,omitempty,int"`
	ClientId            string `json:"client-id,omitempty"             url:"client-id,omitempty"`
	ClientKey           string `json:"client-key,omitempty"            url:"client-key,omitempty"`
	Comment             string `json:"comment,omitempty"               url:"comment,omitempty"`
	Default             bool   `json:"default,omitempty"               url:"default,omitempty,int"`
	Domain              string `json:"domain,omitempty"                url:"domain,omitempty"`
	Filter              string `json:"filter,omitempty"                url:"filter,omitempty"`
	GroupClasses        string `json:"group_classes,omitempty"         url:"group_classes,omitempty"`
	GroupDn             string `json:"group_dn,omitempty"              url:"group_dn,omitempty"`
	GroupFilter         string `json:"group_filter,omitempty"          url:"group_filter,omitempty"`
	GroupNameAttr       string `json:"group_name_attr,omitempty"       url:"group_name_attr,omitempty"`
	GroupsAutoCreate    bool   `json:"groups-autocreate,omitempty"     url:"groups-autocreate,omitempty"`
	GroupsClaim         bool   `json:"groups-claim,omitempty"          url:"groups-claim,omitempty"`
	GroupsOverwrite     bool   `json:"groups-overwrite,omitempty"      url:"groups-overwrite,omitempty"`
	IssuerUrl           string `json:"issuer-url,omitempty"            url:"issuer-url,omitempty"`
	Mode                string `json:"mode,omitempty"                  url:"mode,omitempty"`
	Password            string `json:"password,omitempty"              url:"password,omitempty"`
	Port                int    `json:"port,omitempty"                  url:"port,omitempty,int"`
	Prompt              string `json:"prompt,omitempty"                url:"prompt,omitempty"`
	Scopes              string `json:"scopes,omitempty"                url:"scopes,omitempty"`
	Secure              bool   `json:"secure,omitempty"                url:"secure,omitempty,int"`
	Server1             string `json:"server1,omitempty"               url:"server1,omitempty"`
	Server2             string `json:"server2,omitempty"               url:"server2,omitempty"`
	SslVersion          string `json:"sslversion,omitempty"            url:"sslversion,omitempty"`
	SyncDefaultsOptions string `json:"sync-defaults-options,omitempty" url:"sync-defaults-options,omitempty"`
	SyncAttributes      string `json:"sync_attributes,omitempty"       url:"sync_attributes,omitempty"`
	Tfa                 string `json:"tfa,omitempty"                   url:"tfa,omitempty"`
	UserAttr            string `json:"user_attr,omitempty"             url:"user_attr,omitempty"`
	UserClasses         string `json:"user_classes,omitempty"          url:"user_classes,omitempty"`
	UsernameClaim       string `json:"username-claim,omitempty"        url:"username-claim,omitempty"`
	Verify              bool   `json:"verify,omitempty"                url:"verify,omitempty,int"`
}
