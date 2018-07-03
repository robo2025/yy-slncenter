package robodb

//sewage_info 表
type SewageInfo struct {
	ID                 int    `json:"-"`
	SlnNo              string `json:"sln_no"`
	SewageBusiness     string `json:"sewage_business"`
	SewageScenario     string `json:"sewage_scenario"`
	TechMethod         string `json:"tech_method"`
	GeneralNorm        string `json:"general_norm"`
	OtherNorm          string `json:"other_norm"`
	DailyCapacity      string `json:"daily_capacity"`
	Disinfector        string `json:"disinfector"`
	Valve              string `json:"valve"`
	Blower             string `json:"blower"`
	Stirrer            string `json:"stirrer"`
	AuxEquipmentNubs   string `json:"aux_equipment_nubs"`
	TotalEquipmentNubs string `json:"total_equipment_nubs"`
	Pump               string `json:"pump"`
	Doser              string `json:"doser"`
	OperatingSize      string `json:"operating_size"`
}
