package models

type WorkerCreate struct {
	Img_url    string `json:"img"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Position   string `json:"position"`
	Department string `json:"department"`
	Gender     string `json:"gender"`
	Contact    string `json:"contact"`
	Birthday   string `json:"birthday"`
	ComeTime   string `json:"come_time"`
}

type WorkerResp struct {
	Id         string `json:"id"`
	Img_url    string `json:"img"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Position   string `json:"position"`
	Department string `json:"department"`
	Gender     string `json:"gender"`
	Contact    string `json:"contact"`
	Birthday   string `json:"birthday"`
	ComeTime   string `json:"come_time"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type WorkerUpdate struct {
	Id         string `json:"id"`
	Img_url    string `json:"img"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Position   string `json:"position"`
	Department string `json:"department"`
	Gender     string `json:"gender"`
	Contact    string `json:"contact"`
	Birthday   string `json:"birthday"`
	ComeTime   string `json:"come_time"`
}

type AllWorkers struct {
	Total_amount int `json:"total_amount"`
	Male         int `json:"male"`
	Female       int `json:"female"`
	Workers      []WorkerResp
}


type AllWorkersFilter struct {
	Workers []WorkerResp
}

type SortingWorker struct {
	AvarHour      string `json:"avarage_hour_at_work"`
	AvarComeTime  string `json:"avarage_come_time"`
	AvarLeaveTime string `json:"avarage_leave_time"`
	Id            string `json:"id"`
	Img_url       string `json:"img"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Position      string `json:"position"`
	Department    string `json:"department"`
	Gender        string `json:"gender"`
	Contact       string `json:"contact"`
	Birthday      string `json:"birthday"`
	ComeTime      string `json:"come_time"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type TopWorkers struct {
	TopBest []TopTen `json:"topBest"`
	TopBad  []TopTen `json:"topBad"`
}

type TopTen struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Img_url    string `json:"image_url"`
	Gender     string `json:"gender"`
	Position   string `json:"position"`
	Department string `json:"wdepartment"`
	ComeTime   string `json:"comeTime"`
}
