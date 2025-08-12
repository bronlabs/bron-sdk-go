package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/orderedmap"
)

type OpenApiSchema struct {
	Type                 string                `json:"type,omitempty"`
	Enum                 []interface{}         `json:"enum,omitempty"`
	Ref                  string                `json:"$ref,omitempty"`
	Properties           orderedmap.OrderedMap `json:"properties,omitempty"`
	AdditionalProperties interface{}           `json:"additionalProperties,omitempty"`
	Required             []string              `json:"required,omitempty"`
	Items                *OpenApiSchema        `json:"items,omitempty"`
	AllOf                []OpenApiSchema       `json:"allOf,omitempty"`
	OneOf                []OpenApiSchema       `json:"oneOf,omitempty"`
	AnyOf                []OpenApiSchema       `json:"anyOf,omitempty"`
}

type OpenApiOperation struct {
	Summary     string   `json:"summary,omitempty"`
	OperationID string   `json:"operationId,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Parameters  []struct {
		Name     string         `json:"name"`
		In       string         `json:"in"`
		Required *bool          `json:"required"`
		Schema   *OpenApiSchema `json:"schema"`
	} `json:"parameters,omitempty"`
	RequestBody *struct {
		Content map[string]struct {
			Schema *OpenApiSchema `json:"schema"`
		} `json:"content"`
	} `json:"requestBody,omitempty"`
	Responses map[string]struct {
		Content map[string]struct {
			Schema *OpenApiSchema `json:"schema"`
		} `json:"content"`
	} `json:"responses,omitempty"`
}

type SchemaEntry struct {
	Name   string
	Schema OpenApiSchema
}

type PathEntry struct {
	Route   string
	Methods orderedmap.OrderedMap
}

type OpenApiSpec struct {
	Components struct {
		Schemas []SchemaEntry `json:"schemas"`
	} `json:"components"`
	Paths []PathEntry `json:"paths"`
}

type Generator struct {
	spec        OpenApiSpec
	schemaOrder []string
	pathOrder   []string
	typesDir    string
	apiDir      string
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: generator <spec-file> <types-dir> <api-dir>")
		os.Exit(1)
	}

	specPath := os.Args[1]
	typesDir := os.Args[2]
	apiDir := os.Args[3]

	generator, err := NewGenerator(specPath, typesDir, apiDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating generator: %v\n", err)
		os.Exit(1)
	}

	if err := generator.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Code generation completed successfully!")
}

func NewGenerator(specPath, typesDir, apiDir string) (*Generator, error) {
	data, err := os.ReadFile(specPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec file: %w", err)
	}

	// Parse JSON with orderedmap to preserve order
	var rawSpec orderedmap.OrderedMap
	if err := json.Unmarshal(data, &rawSpec); err != nil {
		return nil, fmt.Errorf("failed to parse spec file: %w", err)
	}

	// Parse schemas and preserve order
	var schemas []SchemaEntry
	if components, ok := rawSpec.Get("components"); ok {
		if componentsMap, ok := components.(orderedmap.OrderedMap); ok {
			if schemasRaw, ok := componentsMap.Get("schemas"); ok {
				if schemasMap, ok := schemasRaw.(orderedmap.OrderedMap); ok {
					for _, name := range schemasMap.Keys() {
						schemaRaw, _ := schemasMap.Get(name)
						schemaBytes, _ := json.Marshal(schemaRaw)
						var schemaMap orderedmap.OrderedMap
						json.Unmarshal(schemaBytes, &schemaMap)

						// Convert to OpenApiSchema while preserving order
						var schema OpenApiSchema
						if typeVal, ok := schemaMap.Get("type"); ok {
							if typeStr, ok := typeVal.(string); ok {
								schema.Type = typeStr
							}
						}
						if refVal, ok := schemaMap.Get("$ref"); ok {
							if refStr, ok := refVal.(string); ok {
								schema.Ref = refStr
							}
						}
						if enumVal, ok := schemaMap.Get("enum"); ok {
							if enumArr, ok := enumVal.([]interface{}); ok {
								schema.Enum = enumArr
							}
						}
						if requiredVal, ok := schemaMap.Get("required"); ok {
							if requiredArr, ok := requiredVal.([]interface{}); ok {
								for _, req := range requiredArr {
									if reqStr, ok := req.(string); ok {
										schema.Required = append(schema.Required, reqStr)
									}
								}
							}
						}
						if propsVal, ok := schemaMap.Get("properties"); ok {
							if propsMap, ok := propsVal.(orderedmap.OrderedMap); ok {
								schema.Properties = propsMap
							}
						}
						schemas = append(schemas, SchemaEntry{Name: name, Schema: schema})
					}
				}
			}
		}
	}

	// Parse paths and preserve order
	var paths []PathEntry
	if pathsRaw, ok := rawSpec.Get("paths"); ok {
		if pathsMap, ok := pathsRaw.(orderedmap.OrderedMap); ok {
			for _, route := range pathsMap.Keys() {
				methodsRaw, _ := pathsMap.Get(route)
				methodsBytes, _ := json.Marshal(methodsRaw)
				var methods orderedmap.OrderedMap
				json.Unmarshal(methodsBytes, &methods)
				paths = append(paths, PathEntry{Route: route, Methods: methods})
			}
		}
	}

	spec := OpenApiSpec{
		Components: struct {
			Schemas []SchemaEntry `json:"schemas"`
		}{
			Schemas: schemas,
		},
		Paths: paths,
	}

	return &Generator{
		spec:     spec,
		typesDir: typesDir,
		apiDir:   apiDir,
	}, nil
}

func (g *Generator) Generate() error {
	if err := g.generateTypes(); err != nil {
		return fmt.Errorf("failed to generate types: %w", err)
	}

	if err := g.generateAPI(); err != nil {
		return fmt.Errorf("failed to generate API: %w", err)
	}

	return nil
}

func (g *Generator) generateTypes() error {
	if err := os.MkdirAll(g.typesDir, 0755); err != nil {
		return fmt.Errorf("failed to create types directory: %w", err)
	}

	// Clear existing .go files
	if err := g.clearDirectory(g.typesDir, ".go"); err != nil {
		return fmt.Errorf("failed to clear types directory: %w", err)
	}

	for _, entry := range g.spec.Components.Schemas {
		content := g.generateTypeDefinition(entry.Name, entry.Schema)
		filename := filepath.Join(g.typesDir, entry.Name+".go")
		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write type file %s: %w", filename, err)
		}
	}

	return nil
}

func (g *Generator) generateTypeDefinition(name string, schema OpenApiSchema) string {
	if schema.Enum != nil {
		return g.generateEnum(name, schema)
	}
	if len(schema.AllOf) == 1 && schema.AllOf[0].Ref != "" {
		return g.generateTypeAlias(name, g.extractRefName(schema.AllOf[0].Ref))
	}
	if schema.Ref != "" {
		return g.generateTypeAlias(name, g.extractRefName(schema.Ref))
	}
	if g.isEmptyObject(schema) {
		return g.generateMapType(name)
	}
	if schema.Type == "object" || len(schema.Properties.Keys()) > 0 {
		return g.generateStruct(name, schema)
	}
	return g.generateInterfaceType(name)
}

func (g *Generator) generateEnum(name string, schema OpenApiSchema) string {
	var sb strings.Builder
	sb.WriteString("package types\n\n")
	sb.WriteString(fmt.Sprintf("type %s string\n\n", name))
	sb.WriteString("const (\n")

	for _, enumValue := range schema.Enum {
		if str, ok := enumValue.(string); ok {
			constName := strings.ToUpper(strings.ReplaceAll(str, "-", "_"))
			// Fix invalid constant names that start with numbers
			if len(constName) > 0 && constName[0] >= '0' && constName[0] <= '9' {
				constName = "TYPE_" + constName
			}
			// Prefix with type name to avoid conflicts
			constName = name + "_" + constName
			sb.WriteString(fmt.Sprintf("\t%s %s = \"%s\"\n", constName, name, str))
		}
	}

	sb.WriteString(")\n")
	return sb.String()
}

func (g *Generator) generateTypeAlias(name, refName string) string {
	return fmt.Sprintf("package types\n\nimport \"github.com/bronlabs/bron-sdk-go/sdk/types\"\n\ntype %s = %s\n", name, refName)
}

func (g *Generator) generateMapType(name string) string {
	return fmt.Sprintf("package types\n\ntype %s map[string]interface{}\n", name)
}

func (g *Generator) generateInterfaceType(name string) string {
	return fmt.Sprintf("package types\n\ntype %s interface{}\n", name)
}

func (g *Generator) generateStruct(name string, schema OpenApiSchema) string {
	var sb strings.Builder
	sb.WriteString("package types\n\n")
	sb.WriteString(fmt.Sprintf("type %s struct {\n", name))

	if len(schema.Properties.Keys()) > 0 {
		for _, propName := range schema.Properties.Keys() {
			propNameStr := propName
			propSchemaRaw, _ := schema.Properties.Get(propNameStr)
			// Convert to orderedmap to preserve order
			propSchemaBytes, _ := json.Marshal(propSchemaRaw)
			var propSchemaMap orderedmap.OrderedMap
			json.Unmarshal(propSchemaBytes, &propSchemaMap)

			// Extract type, ref, and items directly from the ordered map
			var propType, propRef string
			if typeVal, ok := propSchemaMap.Get("type"); ok {
				if typeStr, ok := typeVal.(string); ok {
					propType = typeStr
				}
			}
			if refVal, ok := propSchemaMap.Get("$ref"); ok {
				if refStr, ok := refVal.(string); ok {
					propRef = refStr
				}
			}

			// Parse items for array types
			var itemsSchema *OpenApiSchema
			if propType == "array" {
				if itemsVal, ok := propSchemaMap.Get("items"); ok {
					itemsBytes, _ := json.Marshal(itemsVal)
					var itemsMap orderedmap.OrderedMap
					json.Unmarshal(itemsBytes, &itemsMap)

					var itemsType, itemsRef string
					if itemsTypeVal, ok := itemsMap.Get("type"); ok {
						if itemsTypeStr, ok := itemsTypeVal.(string); ok {
							itemsType = itemsTypeStr
						}
					}
					if itemsRefVal, ok := itemsMap.Get("$ref"); ok {
						if itemsRefStr, ok := itemsRefVal.(string); ok {
							itemsRef = itemsRefStr
						}
					}

					itemsSchema = &OpenApiSchema{
						Type: itemsType,
						Ref:  itemsRef,
					}
				}
			}

			// Create a schema for type resolution
			propSchema := OpenApiSchema{
				Type:  propType,
				Ref:   propRef,
				Items: itemsSchema,
			}
			goType := g.resolveType(propSchema)
			jsonTag := fmt.Sprintf("`json:\"%s", propNameStr)

			// Check if field is required
			isRequired := false
			for _, required := range schema.Required {
				if required == propName {
					isRequired = true
					break
				}
			}

			// Make field optional if not required
			if !isRequired {
				if goType != "interface{}" {
					goType = "*" + goType
				}
				jsonTag += ",omitempty"
			}
			jsonTag += "\"`"

			sb.WriteString(fmt.Sprintf("\t%s %s %s\n", g.toProperPascalCase(propName), goType, jsonTag))
		}
	}

	sb.WriteString("}\n")
	return sb.String()
}

func (g *Generator) resolveType(schema OpenApiSchema) string {
	if schema.Ref != "" {
		return g.extractRefName(schema.Ref)
	}
	if schema.Enum != nil {
		return "string"
	}

	switch schema.Type {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		if schema.Items != nil {
			return "[]" + g.resolveType(*schema.Items)
		}
		return "[]interface{}"
	case "object":
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}

func (g *Generator) generateAPI() error {
	if err := os.MkdirAll(g.apiDir, 0755); err != nil {
		return fmt.Errorf("failed to create API directory: %w", err)
	}

	// Clear existing .go files
	if err := g.clearDirectory(g.apiDir, ".go"); err != nil {
		return fmt.Errorf("failed to clear API directory: %w", err)
	}

	fileData := make(map[string]*apiFileData)
	// Map to collect query parameters per tag
	queries := make(map[string][]struct {
		name     string
		schema   OpenApiSchema
		required bool
	}) // key = query struct name

	// Process all paths and methods in original order
	for _, pathEntry := range g.spec.Paths {
		for _, method := range pathEntry.Methods.Keys() {
			opRaw, _ := pathEntry.Methods.Get(method)
			opBytes, _ := json.Marshal(opRaw)
			var op OpenApiOperation
			json.Unmarshal(opBytes, &op)
			if op.Summary == "" {
				continue
			}

			fileName := g.getFileName(op)

			funcNameCamel := g.toCamelCase(op.Summary)
			baseName := g.toProperPascalCase(funcNameCamel)
			if strings.HasPrefix(baseName, "Get") {
				baseName = baseName[3:]
			}
			structName := baseName + "Query"

			// Collect query parameters for this handle
			for _, p := range op.Parameters {
				if p.In == "query" && p.Schema != nil {
					queries[structName] = append(queries[structName], struct {
						name     string
						schema   OpenApiSchema
						required bool
					}{name: p.Name, schema: *p.Schema, required: p.Required != nil && *p.Required})
				}
			}
			if fileData[fileName] == nil {
				fileData[fileName] = &apiFileData{
					types:   make(map[string]bool),
					methods: []string{},
				}
			}

			// Generate method
			methodCode := g.generateMethod(op, method, pathEntry.Route)
			fileData[fileName].methods = append(fileData[fileName].methods, methodCode)

			// Add required types
			if returnType := g.getReturnType(op); returnType != "" {
				fileData[fileName].types[returnType] = true
			}
		}
	}

	// Write API files in original order
	for fileName, data := range fileData {
		content := g.generateAPIFile(fileName, data)
		filename := filepath.Join(g.apiDir, fileName+".go")
		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write API file %s: %w", filename, err)
		}
	}

	// Generate query types
	for structName, params := range queries {
		var sb strings.Builder
		sb.WriteString("package types\n\n")
		sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))
		fieldNames := make(map[string]bool)
		for _, p := range params {
			if fieldNames[p.name] {
				continue
			}
			fieldNames[p.name] = true
			fieldName := g.toProperPascalCase(p.name)
			goType := g.resolveType(p.schema)
			jsonTag := fmt.Sprintf("`json:\"%s", p.name)
			if !p.required {
				// optional: pointer & omitempty, except interface{}
				if goType != "interface{}" {
					goType = "*" + goType
				}
				jsonTag += ",omitempty"
			}
			jsonTag += "\"`"
			sb.WriteString(fmt.Sprintf("\t%s %s %s\n", fieldName, goType, jsonTag))
		}
		sb.WriteString("}\n")
		filename := filepath.Join(g.typesDir, structName+".go")
		if err := os.WriteFile(filename, []byte(sb.String()), 0644); err != nil {
			return fmt.Errorf("failed to write query type file %s: %w", filename, err)
		}
	}

	return nil
}

type apiFileData struct {
	types   map[string]bool
	methods []string
}

func (g *Generator) generateMethod(op OpenApiOperation, method, route string) string {
	funcName := g.toCamelCase(op.Summary)
	funcName = strings.ReplaceAll(funcName, "ID", "Id")

	// Determine parameters
	paramType, paramOptional, pathParams := g.processParameters(op)
	// Collect query parameters
	var queryParams []struct{ name string }
	for _, p := range op.Parameters {
		if p.In == "query" {
			queryParams = append(queryParams, struct{ name string }{name: p.Name})
		}
	}

	returnType := g.getReturnType(op)

	// Generate method signature and body
	var sb strings.Builder

	// Method signature
	if returnType != "" {
		sb.WriteString(fmt.Sprintf("func (api *%sAPI) %s(", g.getAPIClassName(op), g.toProperPascalCase(funcName)))
	} else {
		sb.WriteString(fmt.Sprintf("func (api *%sAPI) %s(", g.getAPIClassName(op), g.toProperPascalCase(funcName)))
	}

	// Add path parameters
	if len(pathParams) > 0 {
		for i, param := range pathParams {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%s string", param.name))
		}
	}

	// Add query parameter
	if len(queryParams) > 0 {
		if len(pathParams) > 0 {
			sb.WriteString(", ")
		}
		baseName := g.toProperPascalCase(funcName)
		if strings.HasPrefix(baseName, "Get") {
			baseName = baseName[3:]
		}
		queryTypeName := baseName + "Query"
		sb.WriteString(fmt.Sprintf("query ...*types.%s", queryTypeName))
	}

	// Add body parameter
	if paramType != "interface{}" {
		if len(pathParams) > 0 || len(queryParams) > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("body%s types.%s", paramOptional, paramType))
	}

	sb.WriteString(")")

	// Return type
	responseType := g.getReturnType(op)
	if responseType != "" {
		sb.WriteString(fmt.Sprintf(" (*types.%s, error)", responseType))
	} else {
		sb.WriteString(" error")
	}

	// Method body
	sb.WriteString(" {\n")

	// Build path
	pathExpr := route

	hasWorkspaceParam := strings.Contains(pathExpr, "{workspaceId}")
	pathExpr = strings.ReplaceAll(pathExpr, "{workspaceId}", "%s")
	for _, param := range pathParams {
		pathExpr = strings.ReplaceAll(pathExpr, "{"+param.name+"}", "%s")
	}

	// HTTP request
	sb.WriteString(fmt.Sprintf("\tpath := fmt.Sprintf(\"%s\"", pathExpr))

	if hasWorkspaceParam {
		sb.WriteString(", api.workspaceID")
	}

	for _, param := range pathParams {
		sb.WriteString(fmt.Sprintf(", %s", param.name))
	}
	sb.WriteString(")\n")

	// Return proper types
	if responseType != "" {
		sb.WriteString(fmt.Sprintf("\tvar result types.%s\n", responseType))
	} else {
		sb.WriteString("\tvar result interface{}\n")
	}
	if len(queryParams) > 0 {
		sb.WriteString("\tvar queryParam *types." + g.getQueryTypeName(funcName) + "\n")
		sb.WriteString("\tif len(query) > 0 && query[0] != nil {\n")
		sb.WriteString("\t\tqueryParam = query[0]\n")
		sb.WriteString("\t}\n")
	}
	sb.WriteString("\toptions := http.RequestOptions{\n")
	sb.WriteString(fmt.Sprintf("\t\tMethod: \"%s\",\n", strings.ToUpper(method)))
	sb.WriteString("\t\tPath:   path,\n")
	if len(queryParams) > 0 {
		sb.WriteString("\t\tQuery:  queryParam,\n")
	}
	if paramType != "interface{}" {
		sb.WriteString("\t\tBody:   body,\n")
	}
	sb.WriteString("\t}\n")
	if responseType != "" {
		sb.WriteString("\terr := api.http.Request(&result, options)\n")
		sb.WriteString("\treturn &result, err\n")
	} else {
		sb.WriteString("\terr := api.http.Request(&result, options)\n")
		sb.WriteString("\treturn err\n")
	}

	sb.WriteString("}")
	return sb.String()
}

func (g *Generator) generateAPIFile(fileName string, data *apiFileData) string {
	var sb strings.Builder

	// Package declaration
	sb.WriteString("package api\n\n")

	// Imports
	sb.WriteString("import (\n")
	sb.WriteString("\t\"fmt\"\n\n")
	sb.WriteString("\t\"github.com/bronlabs/bron-sdk-go/sdk/http\"\n")
	sb.WriteString("\t\"github.com/bronlabs/bron-sdk-go/sdk/types\"\n")
	sb.WriteString(")\n\n")

	// API class
	className := g.getAPIClassNameFromFileName(fileName)
	sb.WriteString(fmt.Sprintf("type %sAPI struct {\n", className))
	sb.WriteString("\thttp        *http.Client\n")
	sb.WriteString("\tworkspaceID string\n")
	sb.WriteString("}\n\n")

	// Constructor
	sb.WriteString(fmt.Sprintf("func New%sAPI(http *http.Client, workspaceID string) *%sAPI {\n", className, className))
	sb.WriteString(fmt.Sprintf("\treturn &%sAPI{\n", className))
	sb.WriteString("\t\thttp:        http,\n")
	sb.WriteString("\t\tworkspaceID: workspaceID,\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n\n")

	// Methods
	for _, method := range data.methods {
		sb.WriteString(method)
		sb.WriteString("\n\n")
	}

	return sb.String()
}

func (g *Generator) processParameters(op OpenApiOperation) (paramType, paramOptional string, pathParams []struct{ name string }) {
	// Extract path parameters
	for _, param := range op.Parameters {
		if param.In == "path" && param.Name != "workspaceId" {
			pathParams = append(pathParams, struct{ name string }{name: param.Name})
		}
	}

	// Determine body parameter type
	if op.RequestBody != nil {
		if content, ok := op.RequestBody.Content["application/json"]; ok && content.Schema != nil {
			if content.Schema.Ref != "" {
				paramType = g.extractRefName(content.Schema.Ref)
			} else {
				paramType = "interface{}"
			}
		}
	}

	if paramType == "" {
		paramType = "interface{}"
	}

	if paramType != "interface{}" {
		paramOptional = ""
	} else {
		paramOptional = "?"
	}

	return paramType, paramOptional, pathParams
}

func (g *Generator) getReturnType(op OpenApiOperation) string {
	if op.Responses != nil {
		// Check for both 200 and 201 status codes
		for _, statusCode := range []string{"200", "201"} {
			if response, ok := op.Responses[statusCode]; ok {
				if content, ok := response.Content["application/json"]; ok && content.Schema != nil {
					if content.Schema.Ref != "" {
						return g.extractRefName(content.Schema.Ref)
					}
				}
			}
		}
	}
	return ""
}

func (g *Generator) getFileName(op OpenApiOperation) string {
	if len(op.Tags) > 0 {
		return g.toCamelCase(op.Tags[0])
	}
	return "misc"
}

func (g *Generator) getAPIClassName(op OpenApiOperation) string {
	fileName := g.getFileName(op)
	return g.getAPIClassNameFromFileName(fileName)
}

func (g *Generator) getAPIClassNameFromFileName(fileName string) string {
	if len(fileName) == 0 {
		return "Misc"
	}
	return strings.ToUpper(fileName[:1]) + fileName[1:]
}

func (g *Generator) clearDirectory(dir, extension string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), extension) {
			if err := os.Remove(filepath.Join(dir, entry.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Generator) isEmptyObject(schema OpenApiSchema) bool {
	return (schema.Type == "object" || len(schema.Properties.Keys()) > 0 || schema.AdditionalProperties != nil) &&
		len(schema.Properties.Keys()) == 0 &&
		(schema.AdditionalProperties == nil || schema.AdditionalProperties == true)
}

func (g *Generator) extractRefName(ref string) string {
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

func (g *Generator) toCamelCase(str string) string {
	// Convert to camelCase
	parts := strings.FieldsFunc(str, func(r rune) bool {
		return r == '-' || r == '_' || r == ' '
	})

	if len(parts) == 0 {
		return ""
	}

	result := strings.ToLower(parts[0])
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			result += strings.ToUpper(parts[i][:1]) + strings.ToLower(parts[i][1:])
		}
	}

	return result
}

func (g *Generator) toProperPascalCase(str string) string {
	if str == "" {
		return str
	}

	// Handle camelCase to PascalCase conversion
	if len(str) > 0 {
		return strings.ToUpper(str[:1]) + str[1:]
	}

	return str
}

func (g *Generator) getQueryTypeName(funcName string) string {
	baseName := g.toProperPascalCase(funcName)
	if strings.HasPrefix(baseName, "Get") {
		baseName = baseName[3:]
	}
	return baseName + "Query"
}
