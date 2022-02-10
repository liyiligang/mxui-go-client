/*
 * Copyright 2021 liyiligang.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schema

import (
	"reflect"
	"strings"
)

type TableSchema struct {
	Name 		string
	Width 		string
}

func ParseTagWithTableSchema(field reflect.StructField) TableSchema {
	tableSchema := TableSchema{}
	schema := field.Tag.Get("schema")
	schemaTags := strings.Split(schema, ",")
	for _, v := range schemaTags {
		parseTagValue(v, &tableSchema)
	}
	if tableSchema.Name == "" {
		tableSchema.Name = field.Name
	}
	return tableSchema
}

func parseTagValue(str string, tableSchema *TableSchema) {
	kvStr := strings.Split(str, "=")
	if len(kvStr) != 2 {
		return
	}
	val := kvStr[1]
	switch kvStr[0] {
	case "title":
		tableSchema.Name = val
		break
	case "width":
		tableSchema.Width = val
		break
	}
}