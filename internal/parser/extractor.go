package parser
import (
"fmt"
"github.com/pb33f/libopenapi"
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
spec.Endpoints[key] = endpoint
}
