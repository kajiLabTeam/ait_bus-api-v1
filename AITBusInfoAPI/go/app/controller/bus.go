package controller

import (
	"AITBusInfo/db"
	"AITBusInfo/model"
	"AITBusInfo/service"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// PingExample godoc
// @Summary Get Next Bus Time
// @Schemes
// @Description 次に来るバスの時刻を返すエンドポイント
// @Tags example
// @Accept json
// @Produce json
// @Param offset query integer false "0の場合は次のバスの時刻を、1の場合は次の次のバスの時刻を返します" default(0)
// @Success 200 {object} model.NextBusTimeResponse "<strong>※その日のバスの運行が終了した or バスの運行が無い日の場合は以下の要素はレスポンスに含まれません。</strong><br> <li>`nextHourToAIT`</li><li>`nextHourToYakusa`</li><li>`nextMinuteToAIT`</li><li>`nextMinuteToYakusa`</li>"
// @Router /nextbus [get]
func GetNextTime(c *gin.Context) {
	// offsetのgetクエリパラメータを取得してintに変換
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	now := time.Now().UTC().Add(9 * time.Hour) // 現在時刻をJSTで取得
	date := now.Format("2006-1-2")             // 今日の日付をYYYY-MM-DDの形式で取得
	hour := now.Hour()                         // 時間を取得
	minute := now.Minute()                     // 分を取得
	mode := db.GetOperationMode()[date]        // 時刻表を取得

	// fmt.Printf("%s %d:%d\n", date, hour, minute)
	// fmt.Println(mode)
	// hour = 18
	// minute = 30
	// mode = "A"
	// operationModeの値が空かどうかをチェック
	if mode == "" {
		// fmt.Println("運行していません")
		notExistBusResponse := model.NotExistBusResponse{
			State: model.BusState{IsFirst: false, IsExist: false},
		}
		c.JSON(200, notExistBusResponse)
		return
	}

	var modeToAIT, modeToYakusa string
	switch mode {
	case "A":
		modeToAIT, modeToYakusa = "A1", "A2"
	case "B":
		modeToAIT, modeToYakusa = "B1", "B2"
	case "C":
		modeToAIT, modeToYakusa = "C1", "C2"
	}

	// 時刻表を取得
	tt := service.TimeTable{Schedule: db.GetTimeTable()}

	// 現在の時刻が0~7時なら始発バスの時刻を返す
	if hour <= 7 && offset == 0 {
		// fmt.Println("始発バスの時刻: ")
		// fmt.Printf("始発->AIT: %d:%d, 始発->Yakusa: %d:%d\n", tt.GetNextTime(modeToAIT, 8, 0, 0).Hour, tt.GetNextTime(modeToAIT, 8, 0, 0).Minute, tt.GetNextTime(modeToYakusa, 8, 0, 0).Hour, tt.GetNextTime(modeToYakusa, 8, 0, 0).Minute)
		// fmt.Printf("始発2->AIT: %d:%d, 始発2->Yakusa: %d:%d\n", tt.GetNextTime(modeToAIT, 8, 0, 1).Hour, tt.GetNextTime(modeToAIT, 8, 0, 1).Minute, tt.GetNextTime(modeToYakusa, 8, 0, 1).Hour, tt.GetNextTime(modeToYakusa, 8, 0, 1).Minute)
		firstBusTimeResponse := model.FirstBusTimeResponse{
			Schedule:           mode,
			State:              model.BusState{IsFirst: true, IsExist: false},
			NextHourToAIT:      tt.GetNextTime(modeToAIT, 8, 0, 0).Hour,
			NextMinuteToAIT:    tt.GetNextTime(modeToAIT, 8, 0, 0).Minute,
			NextHourToYakusa:   tt.GetNextTime(modeToYakusa, 8, 0, 0).Hour,
			NextMinuteToYakusa: tt.GetNextTime(modeToYakusa, 8, 0, 0).Minute,
		}
		c.JSON(200, firstBusTimeResponse)
		return
	} else if hour <= 7 && offset == 1 {
		firstBusTimeResponse := model.FirstBusTimeResponse{
			Schedule:           mode,
			State:              model.BusState{IsFirst: true, IsExist: false},
			NextHourToAIT:      tt.GetNextTime(modeToAIT, 8, 0, 1).Hour,
			NextMinuteToAIT:    tt.GetNextTime(modeToAIT, 8, 0, 1).Minute,
			NextHourToYakusa:   tt.GetNextTime(modeToYakusa, 8, 0, 1).Hour,
			NextMinuteToYakusa: tt.GetNextTime(modeToYakusa, 8, 0, 1).Minute,
		}
		c.JSON(200, firstBusTimeResponse)
		return
	}

	// 次のバスの時刻を取得(offsetが指定されている場合はそれに対応した時刻)
	nextTimeToAIT := tt.GetNextTime(modeToAIT, hour, minute, offset)
	nextTimeToYakusa := tt.GetNextTime(modeToYakusa, hour, minute, offset)

	// もしnextTimeとnextNextTimeの両方が-1ならば、本日の運行は終了している
	if nextTimeToAIT.Minute == -1 && nextTimeToYakusa.Minute == -1 {
		// fmt.Println("本日の運行は終了しました")
		notExistBusResponse := model.NotExistBusResponse{
			Schedule: mode,
			State:    model.BusState{IsFirst: false, IsExist: false},
		}
		c.JSON(200, notExistBusResponse)
		return
	}

	// 形式を整えて返す
	nextBusTimeResponse := model.NextBusTimeResponse{
		Schedule:           mode,
		State:              model.BusState{IsFirst: false, IsExist: true},
		NextHourToAIT:      nextTimeToAIT.Hour,
		NextMinuteToAIT:    nextTimeToAIT.Minute,
		NextHourToYakusa:   nextTimeToYakusa.Hour,
		NextMinuteToYakusa: nextTimeToYakusa.Minute,
	}
	// fmt.Printf("次のバスの時刻:\nAIT: %d:%d, Yakusa: %d:%d\n", nextTimeToAIT.Hour, nextTimeToAIT.Minute, nextTimeToYakusa.Hour, nextTimeToYakusa.Minute)

	c.JSON(200, nextBusTimeResponse)
}

func GetNextTimeV2(c *gin.Context) {
	// offsetのgetクエリパラメータを取得してintに変換
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	now := time.Now().UTC().Add(9 * time.Hour) // 現在時刻をJSTで取得
	date := now.Format("2006-1-2")             // 今日の日付をYYYY-MM-DDの形式で取得
	hour := now.Hour()                         // 時間を取得
	minute := now.Minute()                     // 分を取得
	mode := db.GetOperationMode()[date]        // 時刻表を取得

	// operationModeの値が空かどうかをチェック
	if mode == "" {
		notExistBusResponse := model.NextBusTimeResponseV2{
			Schedule:           mode,
			State:              model.BusStateV2{IsFirst: false, IsExist: false},
			NextHourToAIT:      -1,
			NextMinuteToAIT:    -1,
			NextHourToYakusa:   -1,
			NextMinuteToYakusa: -1,
		}
		c.JSON(200, notExistBusResponse)
		return
	}

	var modeToAIT, modeToYakusa string
	switch mode {
	case "A":
		modeToAIT, modeToYakusa = "A1", "A2"
	case "B":
		modeToAIT, modeToYakusa = "B1", "B2"
	case "C":
		modeToAIT, modeToYakusa = "C1", "C2"
	}

	// 時刻表を取得
	tt := service.TimeTable{Schedule: db.GetTimeTable()}

	// 現在の時刻が0~7時なら始発バスの時刻を返す
	if hour <= 7 && offset == 0 {
		firstBusTimeResponse := model.NextBusTimeResponseV2{
			Schedule:           mode,
			State:              model.BusStateV2{IsFirst: true, IsExist: false},
			NextHourToAIT:      tt.GetNextTime(modeToAIT, 8, 0, 0).Hour,
			NextMinuteToAIT:    tt.GetNextTime(modeToAIT, 8, 0, 0).Minute,
			NextHourToYakusa:   tt.GetNextTime(modeToYakusa, 8, 0, 0).Hour,
			NextMinuteToYakusa: tt.GetNextTime(modeToYakusa, 8, 0, 0).Minute,
		}
		c.JSON(200, firstBusTimeResponse)
		return
	} else if hour <= 7 && offset == 1 {
		firstBusTimeResponse := model.NextBusTimeResponseV2{
			Schedule:           mode,
			State:              model.BusStateV2{IsFirst: true, IsExist: false},
			NextHourToAIT:      tt.GetNextTime(modeToAIT, 8, 0, 1).Hour,
			NextMinuteToAIT:    tt.GetNextTime(modeToAIT, 8, 0, 1).Minute,
			NextHourToYakusa:   tt.GetNextTime(modeToYakusa, 8, 0, 1).Hour,
			NextMinuteToYakusa: tt.GetNextTime(modeToYakusa, 8, 0, 1).Minute,
		}
		c.JSON(200, firstBusTimeResponse)
		return
	}

	// 次のバスの時刻を取得(offsetが指定されている場合はそれに対応した時刻)
	nextTimeToAIT := tt.GetNextTime(modeToAIT, hour, minute, offset)
	nextTimeToYakusa := tt.GetNextTime(modeToYakusa, hour, minute, offset)

	// もしnextTimeとnextNextTimeの両方が-1ならば、本日の運行は終了している
	if nextTimeToAIT.Minute == -1 && nextTimeToYakusa.Minute == -1 {
		notExistBusResponse := model.NextBusTimeResponseV2{
			Schedule:           mode,
			State:              model.BusStateV2{IsFirst: false, IsExist: false},
			NextHourToAIT:      -1,
			NextMinuteToAIT:    -1,
			NextHourToYakusa:   -1,
			NextMinuteToYakusa: -1,
		}
		c.JSON(200, notExistBusResponse)
		return
	}

	// 形式を整えて返す
	nextBusTimeResponse := model.NextBusTimeResponseV2{
		Schedule:           mode,
		State:              model.BusStateV2{IsFirst: false, IsExist: true},
		NextHourToAIT:      nextTimeToAIT.Hour,
		NextMinuteToAIT:    nextTimeToAIT.Minute,
		NextHourToYakusa:   nextTimeToYakusa.Hour,
		NextMinuteToYakusa: nextTimeToYakusa.Minute,
	}

	c.JSON(200, nextBusTimeResponse)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary Get Operation Mode
// @Schemes
// @Description 指定した日の運行ダイヤを返すエンドポイント
// @Tags example
// @Accept json
// @Produce json
// @Param date query string false "例：2024-06-05"
// @Success 200 {object} model.OperationModeResponse
// @Failure 400 {object} model.QueryErrorResponse
// @Router /mode [get]
func GetOperationMode(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().Format("2006-1-2")) // dateのデフォルト値を今日にする

	// 日付の形式をチェック
	convertedDate, err := time.Parse("2006-1-2", date)
	if err != nil {
		// fmt.Println(err)
		c.JSON(400, model.QueryErrorResponse{ErrorMessage: "日付の形式が正しくありません"})
		return
	}

	// 与えられた日付をもとに運行ダイヤを取得
	mode := db.GetOperationMode()[convertedDate.Format("2006-1-2")]
	c.JSON(200, model.OperationModeResponse{OperationMode: mode})
}

// @BasePath /api/v1
// PingExample godoc
// @Summary Get All Operation Mode
// @Schemes
// @Description 1年分全ての運行ダイヤを返すエンドポイント
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} model.OperationModeAllResponse
// @Router /modeall [get]
func GetOperationModeAll(c *gin.Context) {
	// 与えられた日付をもとに運行ダイヤを取得
	mode := db.GetOperationMode()

	// modeのkeyを日付の昇順にソート
	var keys []string
	for k := range mode {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 日付のフォーマットをYYYY-MM-DDに変更
	formatedMode := map[string]string{}
	for _, k := range keys {
		t, _ := time.Parse("2006-1-2", k)
		formatedMode[t.Format("2006-01-02")] = mode[k]
	}

	// 形式を整えて返す
	modeMap := map[string]model.OperationMode{}
	for k, v := range formatedMode {
		modeMap[k] = model.OperationMode{Date: k, Mode: v}
	}

	c.JSON(200, model.OperationModeAllResponse{OperationModes: modeMap})
}

// @BasePath /api/v1
// PingExample godoc
// @Summary Get TimeTable
// @Schemes
// @Description 指定した運行ダイヤの時刻表を返すエンドポイント
// @Tags example
// @Accept json
// @Produce json
// @Param mode query string false "例：A_toAIT | A_toYakusa | B_toAIT | B_toYakusa | C_toAIT | C_toYakusa"
// @Success 200 {object} model.TimeTableResponse
// @Failure 400 {object} model.QueryErrorResponse
// @Router /timetable [get]
func GetTimeTable(c *gin.Context) {
	selectedMode := c.DefaultQuery("mode", "A_toAIT")

	var mode string
	switch selectedMode {
	case "A_toAIT":
		mode = "A1"
	case "A_toYakusa":
		mode = "A2"
	case "B_toAIT":
		mode = "B1"
	case "B_toYakusa":
		mode = "B2"
	case "C_toAIT":
		mode = "C1"
	case "C_toYakusa":
		mode = "C2"
	default:
		mode = ""
	}

	if mode == "" {
		c.JSON(400, model.QueryErrorResponse{ErrorMessage: "modeが指定されていないか、正しくありません"})
		return
	}

	// 時刻表を取得
	tt := service.TimeTable{Schedule: db.GetTimeTable()}
	timeTable := tt.GetTimeTable(mode)

	// timeTableの形式を整えて返す
	timeTableMap := map[string]model.TimeTable{}
	for i := 0; i <= 23; i++ {
		timeTableMap[strconv.Itoa(i)] = model.TimeTable{Hour: i, Minutes: timeTable[strconv.Itoa(i)]}
	}
	c.JSON(200, model.TimeTableResponse{TimeTable: timeTableMap})
}

// @BasePath /api/v1
// PingExample godoc
// @Summary Get All TimeTable
// @Schemes
// @Description 1年分の全時刻表を返すエンドポイント
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} model.TimeTableResponse
// @Router /timetableall [get]
func GetTimeTableAll(c *gin.Context) {
	// 時刻表を取得
	timeTable := db.GetTimeTable()

	timeTableMap := map[string]map[string][]int{}
	timeTableMap["A_toAIT"] = timeTable["A1"]
	timeTableMap["A_toYakusa"] = timeTable["A2"]
	timeTableMap["B_toAIT"] = timeTable["B1"]
	timeTableMap["B_toYakusa"] = timeTable["B2"]
	timeTableMap["C_toAIT"] = timeTable["C1"]
	timeTableMap["C_toYakusa"] = timeTable["C2"]

	c.JSON(200, model.TimeTableAllResponse{TimeTables: timeTableMap})
}
