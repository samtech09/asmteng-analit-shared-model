package asmteng_analit_shared_model

//type UserResp struct {
//	Qid       int
//	Corr      byte
//	Marks     float32
//	Review    byte
//	Giventime int16
//	Timestat  byte
//
//	// it is not user response
//	// it tell whether user answered the ques
//	// or not
//	Answered byte
//}

type UserResponse struct {
	Qid         int
	Seqno       int16
	Sectid      int
	Subjid      int16
	Subject     string
	Topic       string
	Correctans  string
	Response    string
	Timetaken   int16
	Iscorrect   bool
	Timestat    int8
	Review      bool
	Visited     int
	Marks       float32
	Timestr     string
	TimestatStr string //Super/Perfect/Weak/...
}

type UserPerformance struct {
	Testid  int     `json:"testid"`
	Userid  int64   `json:"userid"`
	Sectid  int     `json:"sectid"`
	Subjid  int16   `json:"subjid"`
	Topicid int16   `json:"topicid"`
	Score   float32 `json:"score"`

	Totques     int16 `json:"totques"`
	Answered    int16 `json:"answered"`
	Correct     int16 `json:"correct"`
	Review      int16 `json:"review"`
	Reviewans   int16 `json:"reviewans"`
	Notanswered int16 `json:"notanswered"`
	Notvisited  int16 `json:"notvisited"`

	TotStdTime int16   `json:"totstdtime"`
	Giventime  int16   `json:"giventime"`
	OntimeC    int16   `json:"ontimec"`   // correct + ontime
	OntimeW    int16   `json:"ontimew"`   // wrong + ontime
	OvertimeC  int16   `json:"overtimec"` // correct + overtime
	OvertimeW  int16   `json:"overtimew"` // wrong + overtime
	QuickC     int16   `json:"quickc"`    // Too quick + Correct
	QuickW     int16   `json:"quickw"`    // Too quick + Wrong
	CMarks     float32 `json:"cmarks"`    // total correct marks of correct-questions without subtracting negative marks of wrong questions
	WMarks     float32 `json:"wmarks"`    // total correct marks of wrong-questions. It is sum of currect marks of those questions, instead of negative.

	// Super Attempts	= OntimeC 	(correct + ontime)
	// Perfect Attempts	= OvertimeC	(correct + overtime)
	// Weak Attempts	= OntimeW	(wrong + ontime)
	// Wasted Attempts 	= OvertimeW	(wrong + overtime)
	// Guessed Attempts 	= TooQuickC	(correct + too-quick)
	// Bogus Attempts	= TooQuickW	(wrong + too-quick)
}

type TestScore struct {
	Testid int
	Userid int64
	Rank   int
	Marks  float32
	Langid byte
}

type TestPerformance struct {
	Testid       int     `json:"testid"`
	Qid          int     `json:"qid"`
	Sectid       int     `json:"sectid"`
	CorrMarks    float32 `json:"corrmarks"` // uint16 encoded small float value
	Tottime      int     `json:"tottime"`
	Totcorrect   int     `json:"totcorrect"`
	Totattempt   int     `json:"totattempt"`
	Totreview    int     `json:"totreview"`
	Totreviewans int     `json:"totreviewans"` // review and answered
	TotontimeC   int     `json:"totontimec"`
	TotontimeW   int     `json:"totontimew"`
	TotovertimeC int     `json:"totovertimec"`
	TotovertimeW int     `json:"totovertimew"`
	TotquickC    int     `json:"totquickc"`
	TotquickW    int     `json:"totquickw"`
}

type UserRespData struct {
	TestID int
	UserID int64
	Data   []UserResponse
}

// type QMarks struct {
// 	Qid       int
// 	CorrMarks float32
// }

// //GetKey - parentid1 should be Testid, parentid2 should be Userid
// func (s UserResponse) GetKey(parentid1, parentid2 string) string {
// 	var buffer bytes.Buffer
// 	buffer.WriteString("RPT:RESP:TID")
// 	buffer.WriteString(parentid1) // parentid1 should be Testid
// 	buffer.WriteString(":UID")
// 	buffer.WriteString(parentid2) // parentid2 should be Userid
// 	return buffer.String()
// }
// func (s UserResponse) GetMasterKey() string {
// 	return "RPT:RESP:TID<testid>:UID<userid>"
// }
// func (s UserResponse) GetExpiration() time.Duration {
// 	// cache for 10 days 240 hours
// 	return time.Duration(240) * time.Hour
// }
