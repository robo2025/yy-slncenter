package robodb

// sln_support 表  sln_support
type SlnSupport struct {
	ID     int     `json:"-"`
	SlnNo  string  `json:"sln_no"`
	UserID int     `json:"user_id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Note   string  `json:"note"`
}

// sln_supplier_info 表
type SlnSupplierInfo struct {
	ID           int     `json:"-"`
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

// sln_device 表 sln_device表
type SlnDevice struct {
	ID              int     `json:"-"`
	SlnNo           string  `json:"sln_no"`
	UserID          int     `json:"user_id"`
	SlnRole         string  `json:"sln_role"`
	DeviceID        string  `json:"device_id"`
	DeviceType      string  `json:"device_type"`
	DeviceComponent string  `json:"device_component"`
	DeviceName      string  `json:"device_name"`
	DeviceModel     string  `json:"device_model"`
	DevicePrice     float64 `json:"device_price"`
	DeviceNum       int     `json:"device_num"`
	BrandName       string  `json:"brand_name"`
	DeviceNote      string  `json:"device_note"`
	DeviceOrigin    string  `json:"device_origin"`
}

// sln_file 表 Sln File
type SlnFile struct {
	ID       int    `json:"-"`
	SlnNo    string `json:"sln_no"`
	UserID   int    `json:"user_id"`
	SlnRole  string `json:"sln_role"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileURL  string `json:"file_url"`
}

// sln_user_info 表
type SlnUserInfo struct {
	ID          int    `json:"-"`
	SlnNo       string `json:"sln_no"`
	PayRatio    int    `json:"pay_ratio"`
	WeldingName string `json:"welding_name"`
	SlnNote     string `json:"sln_note"`
}

// sln_basic_info 表
type SlnBasicInfo struct {
	ID            int     `json:"-"`
	SlnNo         string  `json:"sln_no"`
	SlnName       string  `json:"sln_name"`
	SlnType       string  `json:"sln_type"`
	SlnDate       int     `json:"sln_date"`
	SlnExpired    int     `json:"sln_expired"`
	CustomerID    int     `json:"customer_id"`
	CustomerName  string  `json:"customer_name"` //
	CustomerPrice float64 `json:"customer_price"`
	SupplierID    int     `json:"supplier_id"`
	SupplierName  string  `json:"supplier_name"` //
	SupplierPrice float64 `json:"supplier_price"`
	SlnStatus     string  `json:"sln_status"`
	SpDate        int     `json:"sp_date"`       //报价时间
	AssignStatus  string  `json:"assign_status"` // 指派状态
	SlnMsg        string  `json:"sln_msg"`
}

// sln_assign 表
type SlnAssign struct {
	ID         int    `json:"-"`
	SlnNo      string `json:"sln_no"`
	SupplierId int    `json:"supplier_id"`
	AddTime    int    `json:"add_time"`
}

// operation_log
type OperationLog struct {
	ID            int    `json:"-"`
	SlnNo         string `json:"sln_no"`
	OperationType string `json:"operation_type"`
	Operator      string `json:"operator"`
	Content       string `json:"content"`
	Remark        string `json:"remark"`
	AddTime       int    `json:"add_time"`
	OperatorId    int    `json:"operator_id"`
}
type OfferOperation struct {
	ID            int    `json:"-"`
	SlnNo         string `json:"sln_no"`
	SbmNo         int    `json:"sbm_no"`
	Role          int    `json:"role"`
	OperatingPart string `json:"operating_part"`
	OperatingType string `json:"operating_type"`
	Content       string `json:"content"`
	AddTime       int    `json:"add_time"`
}
