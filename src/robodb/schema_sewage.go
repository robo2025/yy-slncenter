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
	AuxEquipmentNubs   int     `json:"aux_equipment_nubs"`
	TotalEquipmentNubs int     `json:"total_equipment_nubs"`
	Pump			   int     `json:"pump"`
	Doser			   int   	`json:"doser"`
	OperatingSize	   float64	`json:"operating_size"`
}
/*
    "SewageBusiness": "工业",
    "SewageScenario": "场景",
    "TechMethod": "为什么没汉字",
    "GeneralNorm": "为什么没汉字",
    "OtherNorm": "为什么没汉字",
    "DailyCapacity": 10.5,
    "Disinfector": 1,
    "Valve": 2,
    "Blower": 10,
    "Stirrer": 12,
    "AuxEquipmentNubs": 12,
    "TotalEquipmentNubs": 12,
    "Pump": 12,
    "Doser": 12,
    "OperatingSize":1.5
*/