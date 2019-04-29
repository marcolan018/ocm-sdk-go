/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// ValueBuilder contains the data and logic needed to build 'value' objects.
//
// Numeric value and the unit used to measure it.
//
// Units are not mandatory, and they're not specified for some resources. For
// resources that use bytes, the accepted units are:
//
// - 1 B = 1 byte
// - 1 KB = 10^3 bytes
// - 1 MB = 10^6 bytes
// - 1 GB = 10^9 bytes
// - 1 TB = 10^12 bytes
// - 1 PB = 10^15 bytes
//
// - 1 B = 1 byte
// - 1 KiB = 2^10 bytes
// - 1 MiB = 2^20 bytes
// - 1 GiB = 2^30 bytes
// - 1 TiB = 2^40 bytes
// - 1 PiB = 2^50 bytes
type ValueBuilder struct {
	value *float64
	unit  *string
}

// NewValue creates a new builder of 'value' objects.
func NewValue() *ValueBuilder {
	return new(ValueBuilder)
}

// Value sets the value of the 'value' attribute
// to the given value.
//
//
func (b *ValueBuilder) Value(value float64) *ValueBuilder {
	b.value = &value
	return b
}

// Unit sets the value of the 'unit' attribute
// to the given value.
//
//
func (b *ValueBuilder) Unit(value string) *ValueBuilder {
	b.unit = &value
	return b
}

// Build creates a 'value' object using the configuration stored in the builder.
func (b *ValueBuilder) Build() (object *Value, err error) {
	object = new(Value)
	if b.value != nil {
		object.value = b.value
	}
	if b.unit != nil {
		object.unit = b.unit
	}
	return
}