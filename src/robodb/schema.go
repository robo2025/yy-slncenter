package robodb

type SlnBasicInfo struct {
	ID        int    `json:"id"`
	SlnNo     string `json:"sln_no"`
	SlnStatus string `json:"sln_status"`
}

type WeldingExtraInfo struct {
	ID           int     `json:"id"`
	SlnNo        string  `json:"sln_no"`
	WeldingPrice float64 `json:"welding_price"`
	WeldingSet   int     `json:"welding_set"`
	PayRatio     int     `json:"pay_ratio"`
	WeldingName  string  `json:"welding_name"`
	WeldingNote  string  `json:"welding_note"`
}

type WeldingFile struct {
	ID       int    `json:"id"`
	SlnNo    string `json:"sln_no"`
	FileURL  string `json:"file_url"`
	FileType string `json:"file_type"`
}

type WeldingInfo struct {
	ID                int     `json:"id"`
	SlnNo             string  `json:"sln_no"`
	WeldingBusiness   string  `json:"welding_business"`
	WeldingScenario   string  `json:"welding_scenario"`
	WeldingMetal      string  `json:"welding_metal"`
	WeldingEfficiency string  `json:"welding_efficiency"`
	WeldingSplash     string  `json:"welding_splash"`
	WeldingModel      string  `json:"welding_model"`
	WeldingMethod     string  `json:"welding_method"`
	WeldingGas        string  `json:"welding_gas"`
	GasCost           string  `json:"gas_cost"`
	MaxHeight         float64 `json:"max_height"`
	MaxRadius         float64 `json:"max_radius"`
}

type WeldingDevices struct {
	ID         int     `json:"id"`
	SlnNo      int     `json:"sln_no"`
	SlnType    string  `json:"sln_type"`
	DeviceID   int     `json:"device_id"`
	DeviceType string  `json:"device_type"`
	BrandName  string  `json:"brand_name"`
	Model      string  `json:"model"`
	Price      float64 `json:"price"`
	DeviceNum  int     `json:"device_num"`
	Note       string  `json:"note"`
}

type SolutionParams struct {
	SlnNo            string            `json:"sln_no" binding:"required"`
	WeldingInfo      WeldingInfo       `json:"welding_info" binding:"required"`
	//WeldingExtraInfo WeldingExtraInfo  `json:"welding_extra_info" binding:"required"`
	WeldingFile      []*WeldingFile    `json:"welding_file" binding:"required"`
	//WeldingDevices   []*WeldingDevices `json:"welding_devices" binding:"required"`
}
