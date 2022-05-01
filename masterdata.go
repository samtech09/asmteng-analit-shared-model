package asmteng_analit_shared_model

type Qmeta struct {
	Qid       int
	Subjid    int16
	Topicid   int16
	Corrans   int16
	Stdtime   int16
	Mintime   int16
	Lengthy   int16
	Difflevel int16
	Tricky    byte
	CorrMarks float32 // extra field not exist in database
}

type Language struct {
	ID       byte
	Langcode string
	Title    string
}

//type Subject struct {
//	ID    int16
//	Title string
//}

//type SubjAlias struct {
//	Subjid  int16
//	Alias string
//}

type Topic struct {
	ID       int16
	Subjid   int16
	Title    string
	Videoimg string
}

type Concept struct {
	ID      int16
	Subjid  int16
	Topicid int16
	Concept string
}

type QuesConcept struct {
	Qid       int
	Conceptid int16
}

type Test struct {
	ID         int
	Title      string
	TestType   string
	TotalMarks float32
}

// type TestSubject struct {
//         testid int
//         subjid int16
//         aliasid int16
//         duration int16
//         seqno byte
// }
