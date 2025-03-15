package main

import "reflect"

// modifyTenantIDRecursively は、構造体のフィールドに対して再帰的にテナントIDを設定する
// 注意: この関数は渡された構造体がポインタ型の場合のみ変更が反映される
func modifyTenantIDRecursively(v any, tenantID string) {
	val := reflect.ValueOf(v)

	// ポインタ型の場合はElem()を使ってポインタが指す値を取得
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// 値の種類に応じて処理
	switch val.Kind() {
	case reflect.Struct:
		processStruct(val, tenantID)
	case reflect.Slice:
		processSlice(val, tenantID)
	}
}

// processStruct は構造体の各フィールドを処理する
func processStruct(val reflect.Value, tenantID string) {
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		// TenantIDフィールドの処理
		if fieldName == "TenantID" && field.Type().String() == "string" && field.CanSet() {
			field.SetString(tenantID)
			continue
		}

		// フィールドの種類に応じた処理
		processField(field, tenantID)
	}
}

// processField はフィールドの種類に応じて適切な処理を行う
func processField(field reflect.Value, tenantID string) {
	if !field.IsValid() || !field.CanInterface() {
		return
	}

	switch field.Kind() {
	case reflect.Struct:
		// 構造体フィールドの場合
		if field.CanAddr() {
			modifyTenantIDRecursively(field.Addr().Interface(), tenantID)
		}
	case reflect.Ptr:
		// ポインタフィールドの場合
		if !field.IsNil() {
			modifyTenantIDRecursively(field.Interface(), tenantID)
		}
	case reflect.Slice:
		// スライスフィールドの場合
		processSlice(field, tenantID)
	}
}

// processSlice はスライスの各要素を処理する
func processSlice(sliceVal reflect.Value, tenantID string) {
	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)
		processField(elem, tenantID)
	}
}
