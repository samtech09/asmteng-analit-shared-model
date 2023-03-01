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
	Subject     string
	Subjid      int16
	Topic       string
	Correctans  string
	Response    string
	Timetaken   int16
	Iscorrect   bool
	Timestat    int8
	Review      bool
	Visited     int
	Marks       float32
	TimestatStr string //Super/Perfect/Weak/...
}

type UserPerformance struct {
	Testid  int
	Userid  int64
	Subjid  int16
	Topicid int16
	Score   float32

	//Totalques      int16
	Answered       int16
	Correct        int16
	Review         int16
	Reviewanswered int16
	Notanswered    int16
	Notvisited     int16

	//Totaltime int16
	Giventime int16
	OntimeC   int16 // correct + ontime
	OntimeW   int16 // wrong + ontime

	OvertimeC int16 // correct + overtime
	OvertimeW int16 // wrong + overtime

	TooQuickC int16 // Too quick + Correct
	TooQuickW int16 // Too quick + Wrong

	CMarks float32 // total correct marks of correct-questions without subtracting negative marks of wrong questions
	WMarks float32 // total correct marks of wrong-questions. It is sum of currect marks of those questions, instead of negative.

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
	Testid int
	Qid    int
	//CorrMarks    float32 // uint16 encoded small float value
	Tottime      int
	Totcorrect   int
	Totattempt   int
	Totreview    int
	TotrevAns    int // review and answered
	TotontimeC   int
	TotontimeW   int
	TotovertimeC int
	TotovertimeW int
	TotquickC    int
	TotquickW    int
	// Totontime   int
	// Totovertime int
	// Totguessed  int
	// Totwasted   int
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
