package robodb

// welding_info 表
type WeldingInfo struct {
	ID                int    `json:"-"`
	SlnNo             string `json:"sln_no"`
	WeldingBusiness   string `json:"welding_business"` // todo
	WeldingScenario   string `json:"welding_scenario"` //todo
	WeldingMetal      string `json:"welding_metal"`
	WeldingEfficiency string `json:"welding_efficiency"`
	WeldingSplash     string `json:"welding_splash"`
	WeldingModel      string `json:"welding_model"`
	WeldingMethod     string `json:"welding_method"`
	WeldingGas        string `json:"welding_gas"`
	GasCost           string `json:"gas_cost"`
	MaxHeight         string `json:"max_height"`
	MaxRadius         string `json:"max_radius"`
}

// welding_tech_param 表
type WeldingTechParam struct {
	ID       int    `json:"-"`
	SlnNo    string `json:"sln_no"`
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	UnitName string `json:"unit_name"`
}
