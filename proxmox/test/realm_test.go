/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

// file deepcode ignore NoHardcodedCredentials/test: test file

package test

import (
	"log"
	"testing"
)

func TestListRealms(t *testing.T) {
	t.Parallel()

	te := InitEnvironment(t)

	acc, err := te.AccessClient().ListRealms(te.t.Context())
	if err != nil {
		log.Printf("%s", err)
		t.Fail()
	}

	for _, obj := range acc {
		log.Printf("%s", obj)
	}

	_ = acc
}

func TestGetRealm(t *testing.T) {
	t.Parallel()

	te := InitEnvironment(t)

	realmList := [...]string{"pve", "pam"}

	for _, realm := range realmList {

		acc, err := te.AccessClient().GetRealm(te.t.Context(), realm)
		if err != nil {
			log.Printf("%s", err)
			t.Fail()
		}

		log.Printf("%+v\n", acc)

	}
}

// func TestCreateRealm(t *testing.T) {
// 	t.Parallel()

// 	parseTime := func(s string) time.Time {
// 		stime, err := time.Parse(time.RFC3339, s)
// 		require.NoError(t, err)

// 		return stime.UTC()
// 	}

// 	realmTests := []struct {
// 		name    string
// 		want    RealmCreateRequestBody
// 		wantErr bool
// 	}{
// 		{
// 			name: "ldap creation",
// 			want: RealmCreateRequestBody{
// 				Realm:    "test",
// 				Type:     "ldap",
// 				BaseDn:   "DN=,CN=",
// 				UserAttr: "freeipausers",
// 				Server1:  "192.168.1.60",
// 				Port:     666,
// 			},
// 		},
// 	}
// 	for _, tt := range realmTeststests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			got, err := CreateRealm(tt.taskID)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ParseTaskID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ParseTaskID() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
