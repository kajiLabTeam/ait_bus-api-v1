package model

type BusState struct {
	IsFirst bool `json:"IsFirst"`
	IsExist bool `json:"IsExist"`
}

type BusStateV2 struct {
	IsFirst bool `json:"isFirst"`
	IsExist bool `json:"isExist"`
}

type NextBusTimeResponse struct {
	Schedule           string   `json:"schedule"`
	State              BusState `json:"busState"`
	NextHourToAIT      int      `json:"nextHourToAIT"`
	NextMinuteToAIT    int      `json:"nextMinuteToAIT"`
	NextHourToYakusa   int      `json:"nextHourToYakusa"`
	NextMinuteToYakusa int      `json:"nextMinuteToYakusa"`
}

type NextBusTimeResponseV2 struct {
	Schedule           string     `json:"schedule"`
	State              BusStateV2 `json:"busState"`
	NextHourToAIT      int        `json:"nextHourToAIT"`
	NextMinuteToAIT    int        `json:"nextMinuteToAIT"`
	NextHourToYakusa   int        `json:"nextHourToYakusa"`
	NextMinuteToYakusa int        `json:"nextMinuteToYakusa"`
}

type FirstBusTimeResponse struct {
	Schedule           string   `json:"schedule"`
	State              BusState `json:"busState"`
	NextHourToAIT      int      `json:"nextHourToAIT"`
	NextMinuteToAIT    int      `json:"nextMinuteToAIT"`
	NextHourToYakusa   int      `json:"nextHourToYakusa"`
	NextMinuteToYakusa int      `json:"nextMinuteToYakusa"`
}

type NotExistBusResponse struct {
	Schedule string   `json:"schedule"`
	State    BusState `json:"busState"`
}

type OperationModeResponse struct {
	OperationMode string `json:"operationMode"`
}

type OperationMode struct {
	Date string `json:"date"`
	Mode string `json:"mode"`
}

type OperationModeAllResponse struct {
	OperationModes map[string]OperationMode `json:"operationModes"`
}

type QueryErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type TimeTable struct {
	Hour    int   `json:"hour"`
	Minutes []int `json:"minutes"`
}

type TimeTableResponse struct {
	TimeTable map[string]TimeTable `json:"timeTable"`
}

type TimeTableAllResponse struct {
	TimeTables map[string]map[string][]int `json:"timeTables"`
}
