package models

type DailyReq struct {
	Id   string `json:"id"`
	Time string `json:"time"`
}

type DailyRes struct {
	WorkerId     string `json:"workerId"`
	Date         string `json:"date"`
	ComeTime     string `json:"comeTime"`
	LeaveTime    string `json:"leaveTime"`
	WorkDuration string `json:"workDuration"`
	Status       bool   `json:"status"`
}

type AttendancePortion struct {
	Portion float32 `json:"portion"`
}
