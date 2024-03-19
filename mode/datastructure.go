package mode

type DeviceCheckInfo struct {
	Ip     string
	Name   string
	Note   string
	Status string
}

type DataDetail struct {
	ID             string `json:"id"`
	Assn           string `json:"assn"`
	AssMemo        string `json:"ass_memo"`
	AType          string `json:"atype"`
	IdcId          string `json:"idc_id"`
	IDCName        string `json:"idc_name"`
	ZoneID         string `json:"zone_id"`
	Region         string `json:"region"`
	Zone           string `json:"zone"`
	Name           string `json:"NAME"`
	UserName       string `json:"user_name"`
	GroupName      string `json:"group_name"`
	ManagerID      string `json:"manager_id"`
	ManagerGroupID string `json:"manager_group_id"`
	IPInner        string `json:"ip_inner"`
	IPOuter        string `json:"ip_outer"`
	Flag           string `json:"flag"`
}
type TextMessage struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}
