package robodb

// 创建方案 struct  SolutionParams
type WeldingParams struct {
	SlnNo         string          `json:"sln_no" binding:"required"`
	UID           int             `json:"-"`
	SlnBasicInfo  *SlnBasicInfo   `json:"sln_basic_info"`
	SlnUserInfo   *SlnUserInfo    `json:"sln_user_info"`
	WeldingInfo   *WeldingInfo    `json:"sln_info"`  //todo
	SlnDevice 	  []SlnDevice     `json:"sln_device"`
	SlnFile   	  []SlnFile   	  `json:"sln_file"`

}

// 供应商报价 struct
type OfferParams struct {
	SlnNo            string             `json:"sln_no" binding:"required"`
	UID              int                `json:"-"`
	SlnSupplierInfo  *SlnSupplierInfo   `json:"sln_supplier_info"`
	SlnDevice	     []SlnDevice    	`json:"sln_device"`
	WeldingTechParam []WeldingTechParam `json:"welding_tech_param"`
	SlnSupport   	 []SlnSupport		`json:"sln_support"`
	SlnFile      	 []SlnFile     		`json:"sln_file"`
}

// 方案细节页面 Welding  SolutionDetailParams
type WeldingDetailParams struct {
	Customer *WeldingParams `json:"customer"`
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
