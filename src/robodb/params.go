package robodb

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
	SlnStatusOffer   SlnStatus = "M"   //string(SlnStatusOffer)
	SlnStatusExpired SlnStatus = "E"
)
type AssignStatus string

const (
	AssignStatusW  AssignStatus = "W"	//未指派
	AssignStatusY AssignStatus = "Y"    //已指派
)

type AssignParams struct {
	SlnAssign *SlnAssign `json:"sln_assign"`
}