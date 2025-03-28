// Package testfixture provides a convenient way to build test fixtures for Go tests
package testfixture

import (
	"reflect"
	"sync"
)

// Factory represents a fixture factory that can create objects of type T
type Factory[T any] struct {
	defaultAttrs  map[string]any
	sequenceAttrs map[string]func(seq int) any
	mu            sync.Mutex
	sequence      int
}

// NewFactory creates a new fixture factory for type T
func NewFactory[T any]() *Factory[T] {
	return &Factory[T]{
		defaultAttrs:  make(map[string]any),
		sequenceAttrs: make(map[string]func(seq int) any),
		sequence:      0,
	}
}

// WithDefault sets a default value for a field
func (f *Factory[T]) WithDefault(fieldName string, value any) *Factory[T] {
	f.defaultAttrs[fieldName] = value
	return f
}

// WithSequence sets a sequence generator for a field
func (f *Factory[T]) WithSequence(fieldName string, seqFunc func(seq int) any) *Factory[T] {
	f.sequenceAttrs[fieldName] = seqFunc
	return f
}

// Build creates a new instance of T with the configured attributes
func (f *Factory[T]) Build(overrides ...map[string]any) T {
	f.mu.Lock()
	defer f.mu.Unlock()

	var result T
	resultValue := reflect.ValueOf(&result).Elem()
	resultType := resultValue.Type()

	// Apply default attributes
	for fieldName, value := range f.defaultAttrs {
		f.setField(resultValue, resultType, fieldName, value)
	}

	// Apply sequence attributes
	for fieldName, seqFunc := range f.sequenceAttrs {
		f.setField(resultValue, resultType, fieldName, seqFunc(f.sequence))
	}

	// Apply overrides
	for _, override := range overrides {
		for fieldName, value := range override {
			f.setField(resultValue, resultType, fieldName, value)
		}
	}

	f.sequence++
	return result
}

// BuildMany creates multiple instances of T
func (f *Factory[T]) BuildMany(count int, overrides ...map[string]any) []T {
	results := make([]T, count)
	for i := 0; i < count; i++ {
		results[i] = f.Build(overrides...)
	}
	return results
}

// setField sets the value of a field by name
func (f *Factory[T]) setField(resultValue reflect.Value, resultType reflect.Type, fieldName string, value any) {
	// Find the field by name
	for i := 0; i < resultType.NumField(); i++ {
		field := resultType.Field(i)
		if field.Name == fieldName {
			fieldValue := resultValue.Field(i)
			if fieldValue.CanSet() {
				valueToSet := reflect.ValueOf(value)

				// Handle type conversion if needed
				if valueToSet.Type().ConvertibleTo(fieldValue.Type()) {
					fieldValue.Set(valueToSet.Convert(fieldValue.Type()))
				}
			}
			return
		}
	}
}
