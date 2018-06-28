package robodb


//sewage_info 表
type SewageInfo struct {
	ID                 int     `json:"-"`
	SlnNo              string  `json:"sln_no"`
	SewageBusiness     string  `json:"sewage_business"`
	SewageScenario     string  `json:"sewage_scenario"`
	TechMethod         string  `json:"tech_method"`
	GeneralNorm        string  `json:"general_norm"`
	OtherNorm          string  `json:"other_norm"`
	DailyCapacity      float32 `json:"daily_capacity"`
	Disinfector        int     `json:"disinfector"`
	Valve              int     `json:"valve"`
	Blower             int     `json:"blower"`
	Stirrer            int     `json:"stirrer"`
	AuxEquipmentNubs   string  `json:"aux_equipment_nubs"`
	TotalEquipmentNubs string  `json:"total_equipment_nubs"`
	Pump			   int     `json:"pump"`
	Doser			   int     `json:"doser"`
	OperatingSize	   string  `json:"operating_size"`
}
