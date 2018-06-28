package robodb

// 创建方案 struct  SolutionParams
type WeldingParams struct {
	SlnNo         string          `json:"sln_no" binding:"required"`
	UID           int             `json:"-"`
	SlnBasicInfo  *SlnBasicInfo   `json:"sln_basic_info"`
	SlnUserInfo   *SlnUserInfo    `json:"sln_user_info"`
	WeldingInfo   *WeldingInfo    `json:"welding_info"`  //todo
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

