package robodb


// 创建污水方案 struct
type SewageParams struct {
	SlnNo         string          `json:"sln_no" binding:"required"`
	UID           int             `json:"-"`
	SewageInfo    *SewageInfo     `json:"sewage_info"`  //todo sln_info
	SlnBasicInfo  *SlnBasicInfo   `json:"sln_basic_info"`
	SlnUserInfo   *SlnUserInfo    `json:"sln_user_info"`
	SlnDevice 	  []SlnDevice 	  `json:"sln_device"`
	SlnFile   	  []SlnFile   	  `json:"sln_file"`
}

// 污水方案细节页面

//type SlnDetailMaster struct {
//	Sewage *SewageDetailParams
//	Welding *SolutionDetailParams
//}
type SewageDetailParams struct {
	Customer *SewageParams `json:"customer"`
	Supplier *OfferParams    `json:"supplier"`
}