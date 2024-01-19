// Copyright 2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/darwishdev/bzns_pro_api/common/auth"
	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/rms/v1"
	"github.com/darwishdev/bzns_pro_api/common/redisclient"
	"github.com/iancoleman/strcase"
	"github.com/rs/zerolog/log"
	"github.com/tangzero/inflector"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (api *Api) authorizeUser(header http.Header) (*auth.Payload, *redisclient.AuthSession, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return nil, nil, fmt.Errorf("missing metadata")
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, nil, fmt.Errorf("invalid authorization header format")
	}

	log.Debug().Interface("fields", fields).Msg("fields")

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := api.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid access token: %s", err)
	}

	authSession, err := api.redisClient.AuthSessionFind(context.Background(), payload.Username)
	if err != nil {
		return nil, nil, fmt.Errorf("canot get the cache: %s", err)
	}

	log.Debug().Interface("authSession", authSession).Msg("hola")
	return payload, authSession, nil
}

func (api *Api) authorizeCustomer(header http.Header) (*auth.Payload, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing metadata")
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	log.Debug().Interface("fields", fields).Msg("fields")

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := api.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}

func (api *Api) GetAccessableActionsForGroup(header http.Header, group string) (*rmsv1.ListDataOptions, error) {
	resp := rmsv1.ListDataOptions{
		Title:       fmt.Sprintf("%s_list", group),
		Description: fmt.Sprintf("%s_description", group),
	}
	var (
		singularizedGroup            string = inflector.Singularize(group)
		redirectRoute                string = fmt.Sprintf("%s_list", group)
		requestProperty              string = fmt.Sprintf("%sId", singularizedGroup)
		deleteRestoreRequestProperty string = inflector.Pluralize(requestProperty)
		create                       string = fmt.Sprintf("%s_create", singularizedGroup)
		update                       string = fmt.Sprintf("%s_update", singularizedGroup)
		deleteRestore                string = fmt.Sprintf("%s_delete_restore", singularizedGroup)
	)
	_, authSession, err := api.authorizeUser(header)
	if err != nil {
		return nil, err
	}
	if authSession.Permissions == nil {
		return &resp, nil
	}

	authorities := authSession.Permissions[group]
	if len(authorities) == 0 {
		return &resp, nil
	}
	log.Debug().Interface("create2", strcase.ToCamel(create)).Str("create", create).Msg("gello")

	if authorities[strcase.ToCamel(create)] {
		resp.CreateHandler = &rmsv1.CreateHandler{
			RedirectRoute: redirectRoute,
			Title:         create,
			Endpoint:      strcase.ToLowerCamel(create),
			RouteName:     create,
		}
		// resp.ImportHandler = &rmsv1.ImportHandler{
		// 	Endpoint:           importEndpoint,
		// 	ImportTemplateLink: importTemplateLink,
		// }
	}
	if authorities[strcase.ToCamel(update)] {
		resp.UpdateHandler = &rmsv1.UpdateHandler{
			RedirectRoute: redirectRoute,
			Title:         update,
			Endpoint:      strcase.ToLowerCamel(update),
			RouteName:     update,
		}
	}
	if authorities[strcase.ToCamel(deleteRestore)] {
		resp.DeleteRestoreHandler = &rmsv1.DeleteRestoreHandler{
			Endpoint:        strcase.ToLowerCamel(deleteRestore),
			RequestProperty: deleteRestoreRequestProperty,
		}
	}
	return &resp, nil
}
