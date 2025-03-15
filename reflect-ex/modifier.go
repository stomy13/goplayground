package main

import "reflect"

// modifyTenantIDRecursively は、構造体のフィールドに対して再帰的にテナントIDを設定する
// 注意: この関数は渡された構造体がポインタ型の場合のみ変更が反映される
func modifyTenantIDRecursively(v any, tenantID string) {
	// reflectパッケージを使ってvの値の型を取得
	val := reflect.ValueOf(v)
	// reflectパッケージを使ってvの値の型がポインタ型かどうかを取得
	if val.Kind() == reflect.Ptr {
		// ポインタ型の場合はElem()を使ってポインタが指す値を取得
		val = val.Elem()
	}

	// 値の種類に応じて処理
	switch val.Kind() {
	case reflect.Struct:
		// 構造体のフィールド数だけループ
		modifyTenantIDFields(val, tenantID)
	case reflect.Slice:
		// スライスの場合、各要素に対して処理
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)
			// スライスの要素が構造体またはポインタの場合
			modifyTenantIDStruct(elem, tenantID)
		}
	}
}

func modifyTenantIDFields(val reflect.Value, tenantID string) {
	for i := 0; i < val.NumField(); i++ {
		// 構造体のフィールドを取得
		field := val.Field(i)
		// フィールド名を取得
		fieldName := val.Type().Field(i).Name

		// フィールドの種類に応じて処理
		switch field.Kind() {
		case reflect.Struct:
			// 構造体のフィールドが構造体の場合は再帰的にテナントIDを設定
			modifyTenantIDRecursively(field.Addr().Interface(), tenantID)
		case reflect.Ptr:
			// ポインタ型で、その先が構造体の場合
			if !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				modifyTenantIDRecursively(field.Interface(), tenantID)
			}
		case reflect.Slice:
			// スライスの場合、各要素に対して処理
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j)
				// スライスの要素が構造体またはポインタの場合
				modifyTenantIDStruct(elem, tenantID)
			}
		default:
			// フィールド名がTenantIDで、型がstringの場合のみ値を設定
			if fieldName == "TenantID" && field.Type().String() == "string" {
				field.SetString(tenantID)
			}
		}
	}
}

func modifyTenantIDStruct(elem reflect.Value, tenantID string) {
	if elem.Kind() == reflect.Struct {
		modifyTenantIDRecursively(elem.Addr().Interface(), tenantID)
	} else if elem.Kind() == reflect.Ptr && !elem.IsNil() {
		modifyTenantIDRecursively(elem.Interface(), tenantID)
	}
}
