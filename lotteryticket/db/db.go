package db

type DoubleColorBallDB struct {
	Id          int32  `grom:"column:id"`
	Week        string `grom:"column:week"`
	Date        string `grom:"column:date"`
	Red1        int    `grom:"column:red1"`
	Red2        int    `grom:"column:red2"`
	Red3        int    `grom:"column:red3"`
	Red4        int    `grom:"column:red4"`
	Red5        int    `grom:"column:red5"`
	Red6        int    `grom:"column:red6"`
	Blue1       int    `grom:"column:blue1"`
	Blue2       int    `grom:"column:blue2"`
	DetailsLink string `grom:"column:detailsLink"`
}

func (db *DoubleColorBallDB) TableName() string {
	return "double_color_ball"
}
