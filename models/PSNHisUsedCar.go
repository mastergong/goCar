package models

type PSNHisUsedCar struct {
	UsedId         string `json:"UsedId"`
	UsedPerID      string `json:"UsedPerID"`
	UsedDocDate    string `json:"UsedDocDate"`
	UsedCarID      string `json:"UsedCarID"`
	UsedBeginDate  string `json:"UsedBeginDate"`
	UsedBeginTime  string `json:"UsedBeginTime"`
	UsedEndDate    string `json:"UsedEndDate"`
	UsedEndTime    string `json:"UsedEndTime"`
	UsedDetail     string `json:"UsedDetail"`
	UsedStatus     string `json:"UsedStatus"`
	UsedAllowPerID string `json:"UsedAllowPerID"`
	UsedMemo       string `json:"UsedMemo"`
}