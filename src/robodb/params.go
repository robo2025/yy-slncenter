package robodb

// 创建方案 struct
type SolutionParams struct {
	SlnNo         string          `json:"sln_no" binding:"required"`
	UID           int             `json:"-"`
	SlnBasicInfo  *SlnBasicInfo   `json:"sln_basic_info"`
	SlnUserInfo   *SlnUserInfo    `json:"sln_user_info"`
	WeldingInfo   *WeldingInfo    `json:"welding_info"`
	WeldingDevice []WeldingDevice `json:"welding_device"`
	WeldingFile   []WeldingFile   `json:"welding_file"`
}

// 供应商报价 struct
type OfferParams struct {
	SlnNo            string             `json:"sln_no" binding:"required"`
	UID              int                `json:"-"`
	SlnSupplierInfo  *SlnSupplierInfo   `json:"sln_supplier_info"`
	WeldingDevice    []WeldingDevice    `json:"welding_device"`
	WeldingTechParam []WeldingTechParam `json:"welding_tech_param"`
	WeldingSupport   []WeldingSupport   `json:"welding_support"`
	WeldingFile      []WeldingFile      `json:"welding_file"`
}

// 方案细节页面
type SolutionDetailParams struct {
	Customer *SolutionParams `json:"customer"`
	Supplier *OfferParams    `json:"supplier"`
}

// 方案列表 RPC 请求
type SolutionRPCReqParams struct {
	UID      int      `json:"uid" binding:"required"`
	Solution []string `json:"solution"`
}

// 方案列表 RPC 回复
type SolutionRPCParams struct {
	SlnBasicInfo    *SlnBasicInfo    `json:"sln_basic_info"`
	SlnSupplierInfo *SlnSupplierInfo `json:"sln_supplier_info"`
	Success         bool             `json:"success"`
	ErrorMsg        string           `json:"error_msg"`
}

// 方案整体状态
type SlnStatus string

const (
	SlnStatusSave    SlnStatus = "S"
	SlnStatusPublish SlnStatus = "P"
	SlnStatusOffer   SlnStatus = "M"
	SlnStatusExpired SlnStatus = "E"
)
