package adapter

import (
	"encoding/json"
	"sync"
	"sync/atomic"

	"github.com/meloneg/mln_rms_core/common/convertor"
	db "github.com/meloneg/mln_rms_core/common/db/gen"
	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

// func parsePermission(res []byte, resultch chan *rmsv1.PermissionGroup, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	var record *rmsv1.PermissionGroup
// 	err := json.Unmarshal(res, &record)
// 	if err != nil {
// 		// Handle the error, e.g., log it or send it to an error channel.
// 		return
// 	}

// 	resultch <- record
// }

func (a *AccountsAdapter) PermissionsListGrpcFromSql(resp *[][]byte) (*rmsv1.PermissionsListResponse, error) {
	// resultch := make(chan *rmsv1.PermissionGroup, 500)
	response := make([]*rmsv1.PermissionGroup, len(*resp))
	// mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	resultch := make(chan *rmsv1.PermissionGroup, 1024)
	defer close(resultch)

	// Create an atomic counter to keep track of the next index to append to.
	var nextIndex int32
	go func() {
		for record := range resultch {
			response = append(response, record)
		}
	}()

	for _, v := range *resp {
		wg.Add(1)
		go func(data []byte) {
			defer wg.Done()
			var record *rmsv1.PermissionGroup
			err := json.Unmarshal(data, &record)
			if err != nil {
				// Handle the error, e.g., log it or send it to an error channel.
				return
			}

			index := atomic.AddInt32(&nextIndex, 1) - 1
			response[index] = record
		}(v)
	}
	wg.Wait()
	response = response[:nextIndex]
	return &rmsv1.PermissionsListResponse{
		Records: response,
	}, nil
}

//list

func (a *AccountsAdapter) rolesListRowGrpcFromSql(resp *db.RolesListRow) *rmsv1.RolesListRow {
	return &rmsv1.RolesListRow{
		RoleId:           resp.RoleID,
		RoleName:         resp.RoleName,
		RoleDescription:  resp.RoleDescription.String,
		PermissionsCount: int32(resp.PermissionsCount),
		UsersCount:       int32(resp.UsersCount),
		CreatedAt:        resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:        resp.DeletedAt.Time.Format(a.dateFormat),
	}
}

func (a *AccountsAdapter) RolesListGrpcFromSql(resp *[]db.RolesListRow) *rmsv1.RolesListResponse {
	records := make([]*rmsv1.RolesListRow, 0)
	deletedRecords := make([]*rmsv1.RolesListRow, 0)
	for _, v := range *resp {
		record := a.rolesListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &rmsv1.RolesListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *AccountsAdapter) RolesInputListGrpcFromSql(resp *[]db.RolesInputListRow) *rmsv1.RolesInputListResponse {
	// RolesInputListGrpcFromSql
	records := make([]*rmsv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.RoleID, v.RoleName)
		records = append(records, record)
	}
	return &rmsv1.RolesInputListResponse{
		Options: records,
	}
}

func (a *AccountsAdapter) RoleCreateSqlFromGrpc(req *rmsv1.RoleCreateRequest) *db.RoleCreateTXParams {
	permissionsParams := make([]db.RolePermissionsBulkCreateParams, 0)
	for _, v := range req.Permissions {
		rolePermission := db.RolePermissionsBulkCreateParams{
			PermissionID: v,
		}
		permissionsParams = append(permissionsParams, rolePermission)
	}
	return &db.RoleCreateTXParams{
		RoleParams: db.RoleCreateParams{
			RoleName:        req.RoleName,
			RoleDescription: convertor.ToPgType(req.RoleDescription),
		},
		PermissionsParams: permissionsParams,
	}
}
func (a *AccountsAdapter) RoleEntityGrpcFromSql(resp *db.AccountsSchemaRole) *rmsv1.Role {
	return &rmsv1.Role{
		RoleId:          int32(resp.RoleID),
		RoleName:        resp.RoleName,
		RoleDescription: resp.RoleDescription.String,
		CreatedAt:       resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:       resp.DeletedAt.Time.Format(a.dateFormat),
	}

}
func (a *AccountsAdapter) RoleCreateGrpcFromSql(resp *db.AccountsSchemaRole) *rmsv1.RoleCreateResponse {
	return &rmsv1.RoleCreateResponse{
		Role: a.RoleEntityGrpcFromSql(resp),
	}
}

func (a *AccountsAdapter) RoleFindForUpdateGrpcFromSql(resp *[]byte) (*rmsv1.RoleUpdateRequest, error) {
	var response rmsv1.RoleUpdateRequest
	err := json.Unmarshal([]byte(*resp), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *AccountsAdapter) RoleUpdateSqlFromGrpc(req *rmsv1.RoleUpdateRequest) *db.RoleUpdateTXParams {
	permissionsParams := make([]db.RolePermissionsBulkCreateParams, 0)
	for _, v := range req.Permissions {
		rolePermission := db.RolePermissionsBulkCreateParams{
			PermissionID: v,
		}
		permissionsParams = append(permissionsParams, rolePermission)
	}
	return &db.RoleUpdateTXParams{
		RoleParams: db.RoleUpdateParams{
			RoleID:          req.RoleId,
			RoleName:        req.RoleName,
			RoleDescription: convertor.ToPgType(req.RoleDescription),
		},
		PermissionsParams: permissionsParams,
	}
}
func (a *AccountsAdapter) RoleUpdateGrpcFromSql(resp *db.AccountsSchemaRole) *rmsv1.RoleUpdateResponse {
	return &rmsv1.RoleUpdateResponse{
		Role: a.RoleEntityGrpcFromSql(resp),
	}
}
