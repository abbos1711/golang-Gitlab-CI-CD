package models

type WorkersByMonthResp struct {
	WorkersResp []WorkersByMonth `json:"workers_resp"`
}

type WorkersResp struct {
	Id          string `json:"id"`
	Img_url     string `json:"img"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	AverageTime string `json:"average_time"`
	// Department  string `json:"department"`
	// Gender      string `json:"gender"`
	// ComeTime    string `json:"come_time"`
	// LeaveTime   string `json:"leave_time"`
}

type WorkersByMonth struct {
	Id                   string `json:"id"`
	Img_url              string `json:"img"`
	Name                 string `json:"name"`
	Surname              string `json:"surname"`
	WorkDayMonth         string `json:"work_day_month"`
	MiddleComeTime       string `json:"middle_come_time"`
	MiddleLeaveTime      string `json:"middle_leave_time"`
	MiddleWorkHoursMonth string `json:"middle_work_hours_month"`
}
