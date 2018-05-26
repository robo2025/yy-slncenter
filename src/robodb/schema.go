package robodb

import "time"

type SlnBasicInfo struct {
	ID            int       `json:"id"`
	SlnNo         string    `json:"sln_no"`
	CustomerID    int       `json:"-"`
	SupplierID    int       `json:"-"`
	SlnName       string    `json:"sln_name"`
	SlnNum        int       `json:"sln_num"`
	SlnDate       time.Time `json:"sln_date"`
	CustomerPrice float64   `json:"customer_price"`
	SupplierPrice float64   `json:"supplier_price"`
	SlnStatus     string    `json:"sln_status"`
}

type SlnUserInfo struct {
	ID          int    `json:"id"`
	SlnNo       string `json:"sln_no"`
	PayRatio    int    `json:"pay_ratio"`
	WeldingName string `json:"welding_name"`
	WeldingNote string `json:"welding_note"`
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

type WeldingFile struct {
	ID       int    `json:"id"`
	SlnNo    string `json:"sln_no"`
	SlnRole  string `json:"sln_role"`
	FileURL  string `json:"file_url"`
	FileType string `json:"file_type"`
}

type WeldingDevice struct {
	ID          int     `json:"id"`
	SlnNo       string  `json:"sln_no"`
	SlnRole     string  `json:"sln_role"`
	DeviceID    int     `json:"device_id"`
	DeviceType  string  `json:"device_type"`
	DeviceModel string  `json:"device_model"`
	DevicePrice float64 `json:"device_price"`
	DeviceNum   int     `json:"device_num"`
	BrandName   string  `json:"brand_name"`
	DeviceNote  string  `json:"device_note"`
}

type SolutionParams struct {
	SlnNo         string          `json:"sln_no" binding:"required"`
	UID           int             `json:"-"`
	SlnBasicInfo  *SlnBasicInfo   `json:"sln_basic_info"`
	SlnUserInfo   *SlnUserInfo    `json:"sln_user_info"`
	WeldingInfo   *WeldingInfo    `json:"welding_info"`
	WeldingDevice []WeldingDevice `json:"welding_device"`
	WeldingFile   []WeldingFile   `json:"welding_file"`
}
