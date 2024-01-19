package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/meloneg/mln_rms_core/common/random"
)

type roleCreateTXTest struct {
	name      string
	params    *RoleCreateTXParams
	expectErr bool
}

func TestRoleCreateTX(t *testing.T) {
	// Define a slice of test cases
	testcases := []roleCreateTXTest{
		// Test for a valid role creation.
		{
			name: "ValidRole",
			params: &RoleCreateTXParams{
				RoleParams: RoleCreateParams{
					RoleName:        random.RandomString(10),
					RoleDescription: pgtype.Text{String: random.RandomString(30), Valid: true},
				},
				PermissionsParams: []RolePermissionsBulkCreateParams{
					{PermissionID: random.RandomInt32(1, 60)},
					{PermissionID: random.RandomInt32(1, 60)},
					{PermissionID: random.RandomInt32(1, 60)},
				},
			},
			expectErr: false,
		},
		// Test for an error when role name is not unique.
		{
			name: "UniqueNameError",
			params: &RoleCreateTXParams{
				RoleParams: RoleCreateParams{
					RoleName:        "waiter",
					RoleDescription: pgtype.Text{String: random.RandomString(30), Valid: true},
				},
				PermissionsParams: []RolePermissionsBulkCreateParams{
					{PermissionID: random.RandomInt32(1, 60)},
					{PermissionID: random.RandomInt32(1, 60)},
					{PermissionID: random.RandomInt32(1, 60)},
				},
			},
			expectErr: true,
		},
		// Test for an error when role name is too long.
		{
			name: "TooLongRoleName",
			params: &RoleCreateTXParams{
				RoleParams: RoleCreateParams{
					RoleName:        random.RandomString(220),
					RoleDescription: pgtype.Text{String: random.RandomString(30), Valid: true},
				},
				PermissionsParams: []RolePermissionsBulkCreateParams{
					{PermissionID: random.RandomInt32(1, 60)},
					{PermissionID: random.RandomInt32(1, 60)},
					{PermissionID: random.RandomInt32(1, 60)},
				},
			},
			expectErr: true,
		},
		// Test for an error when creating a role with an existing name.
		{
			name: "ExistingRecordName",
			params: &RoleCreateTXParams{
				RoleParams: RoleCreateParams{
					RoleName:        "waiter",
					RoleDescription: pgtype.Text{String: random.RandomString(30), Valid: true},
				},
				PermissionsParams: []RolePermissionsBulkCreateParams{
					{PermissionID: random.RandomInt32(1, 60)},
					{PermissionID: random.RandomInt32(1, 60)},
					{PermissionID: random.RandomInt32(1, 60)},
				},
			},
			expectErr: true,
		},
		// Test for an error when using an unknown permission ID.
		{
			name: "UnknownPermissionId",
			params: &RoleCreateTXParams{
				RoleParams: RoleCreateParams{
					RoleName:        random.RandomString(10),
					RoleDescription: pgtype.Text{String: random.RandomString(30), Valid: true},
				},
				PermissionsParams: []RolePermissionsBulkCreateParams{
					{PermissionID: random.RandomInt32(70, 80)},
					{PermissionID: random.RandomInt32(1, 60)},
					{PermissionID: random.RandomInt32(1, 60)},
				},
			},
			expectErr: true,
		},
	}

	// Loop through the test cases and test each one
	ctx := context.Background()
	for _, tc := range testcases {
		fmt.Println("role")
		t.Run(tc.name, func(t *testing.T) {

			// Call the RoleCreate function with the role data from the current test case
			role, err := store.RoleCreateTX(ctx, *tc.params)
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none %s", tc.name)
			}

			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			if !tc.expectErr {
				permissionIDs := make([]int32, 0)
				for _, v := range tc.params.PermissionsParams {
					permissionIDs = append(permissionIDs, v.PermissionID)
				}
				permissionQuery := `
					SELECT permission_id
					FROM accounts_schema.role_permissions
					WHERE role_id = $1
				`
				rows, err := connPool.Query(context.Background(), permissionQuery, role.Role.RoleID)
				if err != nil {
					t.Errorf("Error fetching permission IDs: %v", err)
					return
				}
				defer rows.Close()
				for rows.Next() {
					var permissionID int32
					if err := rows.Scan(&permissionID); err != nil {
						t.Errorf("Error scanning permission ID: %v", err)
						return
					}
					if !intInSlice(permissionID, permissionIDs) {
						t.Errorf("not matched permission id wanted : %v get :%d ", permissionIDs, permissionID)
					}

				}

				// Remove the role with CASCADE option to delete associated records
				deletePermissionsRoleQuery := `
						DELETE FROM accounts_schema.role_permissions
						WHERE role_id = $1
					`
				_, err = connPool.Exec(context.Background(), deletePermissionsRoleQuery, role.Role.RoleID)
				if err != nil {
					t.Errorf("Error deleting role with cascade: %v", err)
				}
				deleteRoleQuery := `
						DELETE FROM accounts_schema.roles
						WHERE role_id = $1
					`
				_, err = connPool.Exec(context.Background(), deleteRoleQuery, role.Role.RoleID)
				if err != nil {
					t.Errorf("Error deleting role with cascade: %v", err)
				}

			}

		})
	}
}
