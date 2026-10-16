package parser

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/tracolerd/schema-drift-observatory/internal/model"
)

func ExtractSpec(doc libopenapi.Document) (*model.NormalizedSpec, error) {
	spec := &model.NormalizedSpec{
		Version:   doc.GetVersion(),
		Endpoints: make(map[string]*model.Endpoint),
	}
	docModel, err := doc.BuildV3Model()
	if err != nil {
		return nil, fmt.Errorf("error building v3 model: %v", err)
	}
	if docModel.Model.Paths == nil || docModel.Model.Paths.PathItems == nil {
		return spec, nil
	}
	for pair := docModel.Model.Paths.PathItems.First(); pair != nil; pair = pair.Next() {
		path := pair.Key()
		pathItem := pair.Value()
		extractOperation(spec, "GET", path, pathItem.Get)
		extractOperation(spec, "POST", path, pathItem.Post)
		extractOperation(spec, "PUT", path, pathItem.Put)
		extractOperation(spec, "DELETE", path, pathItem.Delete)
		extractOperation(spec, "PATCH", path, pathItem.Patch)
	}
	return spec, nil
}
func extractOperation(spec *model.NormalizedSpec, method, path string, op *v3.Operation) {
	if op == nil {
		return
	}
	key := fmt.Sprintf("%s %s", method, path)
	endpoint := &model.Endpoint{
		Method:     method,
		Path:       path,
		Parameters: make(map[string]*model.Parameter),
		Responses:  make(map[string]*model.Response),
	}
	// Extract Parameters
	for _, param := range op.Parameters {
		if param == nil {
			continue
		}
		required := false
		if param.Required != nil {
			required = *param.Required
		}
		paramKey := fmt.Sprintf("%s:%s", param.In, param.Name)
		endpoint.Parameters[paramKey] = &model.Parameter{
			Name:     param.Name,
			In:       param.In,
			Required: required,
		}
	}
	// Extract Request Body (JSON only for MVP)
	if op.RequestBody != nil && op.RequestBody.Content != nil {
		if jsonContent, ok := op.RequestBody.Content.Get("application/json"); ok {
			endpoint.RequestBody = extractSchema(jsonContent.Schema.Schema())
		}
	}
	// Extract Responses (200 JSON only for MVP)
	if op.Responses != nil && op.Responses.Codes != nil {
		if resp200, ok := op.Responses.Codes.Get("200"); ok && resp200.Content != nil {
			if jsonContent, ok := resp200.Content.Get("application/json"); ok {
				endpoint.Responses["200"] = &model.Response{
					StatusCode: "200",
					Content:    extractSchema(jsonContent.Schema.Schema()),
				}
			}
		}
	}
	spec.Endpoints[key] = endpoint
}

// extractSchema recursively converts libopenapi Schema to our internal Schema model.
func extractSchema(schema *base.Schema) *model.Schema {
	if schema == nil {
		return nil
	}
	s := &model.Schema{
		Type:       "",
		Properties: make(map[string]*model.Schema),
		Required:   schema.Required,
	}
	// OpenAPI 3.0 type is a slice of strings in libopenapi's high-level model
	if len(schema.Type) > 0 {
		s.Type = schema.Type[0]
	}
	if schema.Properties != nil {
		for pair := schema.Properties.First(); pair != nil; pair = pair.Next() {
			s.Properties[pair.Key()] = extractSchema(pair.Value().Schema())
		}
	}
	return s
}
