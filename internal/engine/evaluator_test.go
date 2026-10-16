package engine
import (
"testing"
"github.com/tracolerd/schema-drift-observatory/internal/model"
)
func TestCompareSpecs_SDO001_PathRemoved(t *testing.T) {
baseline := &model.NormalizedSpec{
Endpoints: map[string]*model.Endpoint{
"GET /users": {Method: "GET", Path: "/users"},
},
}
candidate := &model.NormalizedSpec{
Endpoints: map[string]*model.Endpoint{}, // Endpoint removed
}
result := CompareSpecs(baseline, candidate)
if result.Compatible {
t.Errorf("Expected result to be incompatible")
}
if result.Counts[CategoryBreaking] != 1 {
t.Errorf("Expected 1 breaking finding, got %d", result.Counts[CategoryBreaking])
}
if len(result.Findings) == 0 || result.Findings[0].RuleID != "SDO-001" {
t.Errorf("Expected SDO-001 finding, got %v", result.Findings)
}
}
func TestCompareSpecs_SDO002_ParamAddedRequired(t *testing.T) {
baseline := &model.NormalizedSpec{
Endpoints: map[string]*model.Endpoint{
"GET /users": {
Method: "GET", Path: "/users",
Parameters: map[string]*model.Parameter{},
},
},
}
candidate := &model.NormalizedSpec{
Endpoints: map[string]*model.Endpoint{
"GET /users": {
Method: "GET", Path: "/users",
Parameters: map[string]*model.Parameter{
"query:id": {Name: "id", In: "query", Required: true},
},
},
},
}
result := CompareSpecs(baseline, candidate)
if result.Compatible {
t.Errorf("Expected result to be incompatible")
}
if result.Counts[CategoryBreaking] != 1 {
t.Errorf("Expected 1 breaking finding for new required param")
}
if len(result.Findings) == 0 || result.Findings[0].RuleID != "SDO-002" {
t.Errorf("Expected SDO-002 finding")
}
}
func TestCompareSpecs_SDO006_OptionalResponsePropertyAdded(t *testing.T) {
baseline := &model.NormalizedSpec{
Endpoints: map[string]*model.Endpoint{
"GET /users": {
Method: "GET", Path: "/users",
Responses: map[string]*model.Response{
"200": {Content: &model.Schema{Properties: map[string]*model.Schema{"id": {Type: "integer"}}}},
},
},
},
}
candidate := &model.NormalizedSpec{
Endpoints: map[string]*model.Endpoint{
"GET /users": {
Method: "GET", Path: "/users",
Responses: map[string]*model.Response{
"200": {Content: &model.Schema{Properties: map[string]*model.Schema{
"id":   {Type: "integer"},
"name": {Type: "string"}, // New optional property
}}},
},
},
},
}
result := CompareSpecs(baseline, candidate)
if !result.Compatible {
t.Errorf("Expected result to be compatible for optional response additions")
}
if result.Counts[CategoryNonBreaking] != 1 {
t.Errorf("Expected 1 non_breaking finding")
}
if len(result.Findings) == 0 || result.Findings[0].RuleID != "SDO-006" {
t.Errorf("Expected SDO-006 finding")
}
}
