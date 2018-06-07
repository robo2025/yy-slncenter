package robodb

// sln_basic_info 表
type SlnBasicInfo struct {
	ID            int     `json:"id"`
	SlnNo         string  `json:"sln_no"`
	SlnName       string  `json:"sln_name"`
	SlnType       string  `json:"sln_type"`
	SlnDate       int     `json:"sln_date"`
	SlnExpired    int     `json:"sln_expired"`
	CustomerID    int     `json:"customer_id"`
	CustomerPrice float64 `json:"customer_price"`
	SupplierID    int     `json:"supplier_id"`
	SupplierPrice float64 `json:"supplier_price"`
	SlnStatus     string  `json:"sln_status"`
}

// sln_user_info 表
type SlnUserInfo struct {
	ID          int    `json:"id"`
	SlnNo       string `json:"sln_no"`
	PayRatio    int    `json:"pay_ratio"`
	WeldingName string `json:"welding_name"`
	WeldingNote string `json:"welding_note"`
}

// welding_info 表
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

// welding_file 表
type WeldingFile struct {
	ID       int    `json:"id"`
	SlnNo    string `json:"sln_no"`
	UserID   int    `json:"user_id"`
	SlnRole  string `json:"sln_role"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileURL  string `json:"file_url"`
}

// welding_device 表
type WeldingDevice struct {
	ID          int     `json:"id"`
	SlnNo       string  `json:"sln_no"`
	UserID      int     `json:"user_id"`
	SlnRole     string  `json:"sln_role"`
	DeviceID    string  `json:"device_id"`
	DeviceType  string  `json:"device_type"`
	DeviceName  string  `json:"device_name"`
	DeviceModel string  `json:"device_model"`
	DevicePrice float64 `json:"device_price"`
	DeviceNum   int     `json:"device_num"`
	BrandName   string  `json:"brand_name"`
	DeviceNote  string  `json:"device_note"`
}

// sln_supplier_info 表
type SlnSupplierInfo struct {
	ID           int     `json:"id"`
	SlnNo        string  `json:"sln_no"`
	UserID       int     `json:"user_id"`
	TotalPrice   float64 `json:"total_price"`
	FreightPrice float64 `json:"freight_price"`
	PayRatio     int     `json:"pay_ratio"`
	ExpiredDate  int     `json:"expired_date"`
	DeliveryDate int     `json:"delivery_date"`
	SlnDesc      string  `json:"sln_desc"`
	SlnNote      string  `json:"sln_note"`
}

// welding_support 表
type WeldingSupport struct {
	ID     int     `json:"id"`
	SlnNo  string  `json:"sln_no"`
	UserID int     `json:"user_id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Note   string  `json:"note"`
}

// welding_tech_param 表
type WeldingTechParam struct {
	ID       int    `json:"id"`
	SlnNo    string `json:"sln_no"`
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	UnitName string `json:"unit_name"`
}
