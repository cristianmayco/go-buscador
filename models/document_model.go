package models

type DocumentModel struct {
	MasterContextName   string  `json:"masterContextName"`
	MasterPk            string  `json:"masterPk"`
	MasterId            int64   `json:"masterId"`
	DocumentContextName string  `json:"documentContextName"`
	DocumentPk          string  `json:"documentPk"`
	DocumentId          int64   `json:"documentId"`
	MasterName          string  `json:"masterName"`
	OriginalName        string  `json:"originalName"`
	IndexId             string  `json:"indexId"`
	File                []byte  `json:"file"`
	AbsolutePath        string  `json:"absolutePath"`
	ExtractedText       string  `json:"extractedText"`
	ModulePath          string  `json:"modulePath"`
	RelevanceScore      float64 `json:"relevanceScore"`
}
