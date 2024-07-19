package main

// PrizeGrade 定义奖级信息结构体
type PrizeGrade struct {
	Type      int    `json:"type"`
	TypeNum   string `json:"typenum"`
	TypeMoney string `json:"typemoney"`
}

// Result 定义结果信息结构体
type Result struct {
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	DetailsLink string       `json:"detailsLink"`
	VideoLink   string       `json:"videoLink"`
	Date        string       `json:"date"`
	Week        string       `json:"week"`
	Red         string       `json:"red"`
	Blue        string       `json:"blue"`
	Blue2       string       `json:"blue2"`
	Sales       string       `json:"sales"`
	PoolMoney   string       `json:"poolmoney"`
	Content     string       `json:"content"`
	AddMoney    string       `json:"addmoney"`
	AddMoney2   string       `json:"addmoney2"`
	Msg         string       `json:"msg"`
	Z2Add       string       `json:"z2add"`
	M2Add       string       `json:"m2add"`
	PrizeGrades []PrizeGrade `json:"prizegrades"`
}

// Data 定义主结构体
type Data struct {
	State    int      `json:"state"`
	Message  string   `json:"message"`
	Total    int      `json:"total"`
	PageNum  int      `json:"pageNum"`
	PageNo   int      `json:"pageNo"`
	PageSize int      `json:"pageSize"`
	Tflag    int      `json:"Tflag"`
	Result   []Result `json:"result"`
}
