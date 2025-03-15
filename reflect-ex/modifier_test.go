package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Simple struct for testing
type SimpleStruct struct {
	ID       string
	Name     string
	TenantID string
	Count    int
	Active   bool
}

// Nested struct for testing
type NestedStruct struct {
	ID       string
	Info     SimpleStruct
	Details  *SimpleStruct
	Numbers  []int
	TenantID string
}

// ComplexStruct for testing
type ComplexStruct struct {
	ID               string
	Name             string
	Count            int
	Ratio            float64
	IsActive         bool
	Tags             []string
	Coordinates      [2]int
	TenantID         string
	Child            ComplexChildStruct
	SimpleStructs    []SimpleStruct
	PtrSimpleStructs []*SimpleStruct
}

type ComplexChildStruct struct {
	SubID      string
	SubName    string
	SubCount   int
	TenantID   string
	GrandChild *SimpleStruct
}

func TestModifyTenantIDRecursively(t *testing.T) {
	// テストケースを定義
	tests := []struct {
		name     string
		input    any
		tenantID string
		want     any
	}{
		// これはポインタ型ではないため、変更が反映されない
		{
			name: "SimpleStruct",
			input: SimpleStruct{
				ID:       "id-1",
				Name:     "test-name",
				TenantID: "old-tenant",
				Count:    10,
				Active:   true,
			},
			tenantID: "new-tenant",
			want: SimpleStruct{
				ID:       "id-1",
				Name:     "test-name",
				TenantID: "old-tenant",
				Count:    10,
				Active:   true,
			},
		},
		{
			name: "SimpleStructWithPointer",
			input: &SimpleStruct{
				ID:       "id-1",
				Name:     "test-name",
				TenantID: "old-tenant",
				Count:    10,
				Active:   true,
			},
			tenantID: "new-tenant",
			want: &SimpleStruct{
				ID:       "id-1",
				Name:     "test-name",
				TenantID: "new-tenant",
				Count:    10,
				Active:   true,
			},
		},
		{
			name: "NestedStruct",
			input: &NestedStruct{
				ID: "parent-id",
				Info: SimpleStruct{
					ID:       "child-id-1",
					Name:     "child-name-1",
					TenantID: "old-tenant-1",
					Count:    100,
					Active:   true,
				},
				Details: &SimpleStruct{
					ID:       "child-id-2",
					Name:     "child-name-2",
					TenantID: "old-tenant-2",
					Count:    200,
					Active:   false,
				},
				Numbers:  []int{1, 2, 3},
				TenantID: "old-parent-tenant",
			},
			tenantID: "new-tenant",
			want: &NestedStruct{
				ID: "parent-id",
				Info: SimpleStruct{
					ID:       "child-id-1",
					Name:     "child-name-1",
					TenantID: "new-tenant",
					Count:    100,
					Active:   true,
				},
				Details: &SimpleStruct{
					ID:       "child-id-2",
					Name:     "child-name-2",
					TenantID: "new-tenant",
					Count:    200,
					Active:   false,
				},
				Numbers:  []int{1, 2, 3},
				TenantID: "new-tenant",
			},
		},
		{
			name: "ComplexStruct",
			input: &ComplexStruct{
				ID:          "complex-id",
				Name:        "complex-name",
				Count:       42,
				Ratio:       3.14,
				IsActive:    true,
				Tags:        []string{"tag1", "tag2", "tag3"},
				Coordinates: [2]int{10, 20},
				TenantID:    "old-tenant",
				Child: ComplexChildStruct{
					SubID:    "sub-id",
					SubName:  "sub-name",
					SubCount: 99,
					TenantID: "old-sub-tenant",
					GrandChild: &SimpleStruct{
						ID:       "grand-id",
						Name:     "grand-name",
						TenantID: "old-grand-tenant",
						Count:    123,
						Active:   true,
					},
				},
				SimpleStructs: []SimpleStruct{
					{
						ID:       "simple-struct-id-1",
						Name:     "simple-struct-name-1",
						TenantID: "old-tenant",
						Count:    123,
						Active:   true,
					},
					{
						ID:       "simple-struct-id-2",
						Name:     "simple-struct-name-2",
						TenantID: "old-tenant",
						Count:    123,
						Active:   true,
					},
				},
				PtrSimpleStructs: []*SimpleStruct{
					{
						ID:       "ptr-simple-struct-id-1",
						Name:     "ptr-simple-struct-name-1",
						TenantID: "old-tenant",
						Count:    123,
						Active:   true,
					},
					{
						ID:       "ptr-simple-struct-id-2",
						Name:     "ptr-simple-struct-name-2",
						TenantID: "old-tenant",
						Count:    123,
						Active:   true,
					},
				},
			},
			tenantID: "new-tenant",
			want: &ComplexStruct{
				ID:          "complex-id",
				Name:        "complex-name",
				Count:       42,
				Ratio:       3.14,
				IsActive:    true,
				Tags:        []string{"tag1", "tag2", "tag3"},
				Coordinates: [2]int{10, 20},
				TenantID:    "new-tenant",
				Child: ComplexChildStruct{
					SubID:    "sub-id",
					SubName:  "sub-name",
					SubCount: 99,
					TenantID: "new-tenant",
					GrandChild: &SimpleStruct{
						ID:       "grand-id",
						Name:     "grand-name",
						TenantID: "new-tenant",
						Count:    123,
						Active:   true,
					},
				},
				SimpleStructs: []SimpleStruct{
					{
						ID:       "simple-struct-id-1",
						Name:     "simple-struct-name-1",
						TenantID: "new-tenant",
						Count:    123,
						Active:   true,
					},
					{
						ID:       "simple-struct-id-2",
						Name:     "simple-struct-name-2",
						TenantID: "new-tenant",
						Count:    123,
						Active:   true,
					},
				},
				PtrSimpleStructs: []*SimpleStruct{
					{
						ID:       "ptr-simple-struct-id-1",
						Name:     "ptr-simple-struct-name-1",
						TenantID: "new-tenant",
						Count:    123,
						Active:   true,
					},
					{
						ID:       "ptr-simple-struct-id-2",
						Name:     "ptr-simple-struct-name-2",
						TenantID: "new-tenant",
						Count:    123,
						Active:   true,
					},
				},
			},
		},
	}

	// テストケースを実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 関数を実行
			modifyTenantIDRecursively(tt.input, tt.tenantID)

			// 結果を検証
			if diff := cmp.Diff(tt.want, tt.input); diff != "" {
				t.Errorf("modifyTenantIDRecursively() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
