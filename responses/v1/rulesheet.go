package v1

import "github.com/bancodobrasil/featws-api/models"

// Rulesheet ...
type Rulesheet struct {
	ID         string             `json:"id,omitempty"`
	Name       string             `json:"name,omitempty"`
	Version    string             `json:"version,omitempty"`
	Features   *[]interface{}     `json:"features,omitempty"`
	Parameters *[]interface{}     `json:"parameters,omitempty"`
	Rules      *map[string]string `json:"rules,omitempty"`
}

// NewRulesheet ...
func NewRulesheet(entity *models.Rulesheet) Rulesheet {
	return Rulesheet{
		ID:         entity.ID.Hex(),
		Name:       entity.Name,
		Version:    entity.Version,
		Features:   entity.Features,
		Parameters: entity.Parameters,
		Rules:      entity.Rules,
	}
}
