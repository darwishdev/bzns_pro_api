package adapter

import (
	"reflect"
	"sort"
	"testing"

	rmsv1 "github.com/meloneg/mln_rms_core/common/pb/bznspro/v1"
)

func TestPermissionsListGrpcFromSql(t *testing.T) {
	// Create a test input with JSON data.
	data1 := []byte(`{"permission_group":"products","permissions":[{"permission_id": 51, "permission_name": "create product", "permission_description": "Permission to create a product"}]}`)
	data2 := []byte(`{"permission_group":"sales","permissions":[{"permission_id": 52, "permission_name": "update product", "permission_description": "Permission to update a product"}]}`)
	data3 := []byte(`{"permission_group":"marketing","permissions":[{"permission_id": 53, "permission_name": "list products", "permission_description": "Permission to list products"}]}`)
	data4 := []byte(`{"permission_group":"finance","permissions":[{"permission_id": 54, "permission_name": "find product", "permission_description": "Permission to find a product"}]}`)

	resp := [][]byte{data1, data2, data3, data4}

	// Create an instance of UsersAdapter.
	usersAdapter := &AccountsAdapter{}

	// Run the function.
	response, err := usersAdapter.PermissionsListGrpcFromSql(&resp)
	if err != nil {
		t.Fatalf("PermissionsListGrpcFromSql error: %v", err)
	}

	// Verify the response.
	expectedResponse := &rmsv1.PermissionsListResponse{
		Records: []*rmsv1.PermissionGroup{
			{
				PermissionGroup: "products",
				Permissions: []*rmsv1.Permission{
					{
						PermissionId:          51,
						PermissionName:        "create product",
						PermissionDescription: "Permission to create a product",
					},
				},
			},
			{
				PermissionGroup: "sales",
				Permissions: []*rmsv1.Permission{
					{
						PermissionId:          52,
						PermissionName:        "update product",
						PermissionDescription: "Permission to update a product",
					},
				},
			},
			{
				PermissionGroup: "marketing",
				Permissions: []*rmsv1.Permission{
					{
						PermissionId:          53,
						PermissionName:        "list products",
						PermissionDescription: "Permission to list products",
					},
				},
			},
			{
				PermissionGroup: "finance",
				Permissions: []*rmsv1.Permission{
					{
						PermissionId:          54,
						PermissionName:        "find product",
						PermissionDescription: "Permission to find a product",
					},
				},
			},
		},
	}

	if CompareSlices(response, expectedResponse) {
		t.Errorf("Expected response does not match actual response got")
		t.Logf("Actual Response: %+v", response)
		t.Logf("Expected Response: %+v", expectedResponse)
	}
}

// CompareSlices compares two slices while ignoring the order of elements.
func CompareSlices(a, b interface{}) bool {
	// Convert slices to reflect.Values
	va, vb := reflect.ValueOf(a), reflect.ValueOf(b)

	// Check if both are slices
	if va.Kind() != reflect.Slice || vb.Kind() != reflect.Slice {
		return false
	}

	// Get the lengths of both slices
	lenA, lenB := va.Len(), vb.Len()

	// If lengths are different, slices are not equal
	if lenA != lenB {
		return false
	}

	// Create sorted slices to compare
	sortedA := make([]interface{}, lenA)
	sortedB := make([]interface{}, lenB)

	for i := 0; i < lenA; i++ {
		sortedA[i] = va.Index(i).Interface()
		sortedB[i] = vb.Index(i).Interface()
	}

	sort.Slice(sortedA, func(i, j int) bool {
		return sortedA[i].(int) < sortedA[j].(int)
	})

	sort.Slice(sortedB, func(i, j int) bool {
		return sortedB[i].(int) < sortedB[j].(int)
	})

	// Compare sorted slices
	return reflect.DeepEqual(sortedA, sortedB)
}

func TestCompareSlices(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 2, 1}
	c := []int{1, 2, 4}

	if !CompareSlices(a, b) {
		t.Error("Slices a and b should be equal")
	}

	if CompareSlices(a, c) {
		t.Error("Slices a and c should not be equal")
	}
}
