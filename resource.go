package instantolib

type Resource struct {
	FileNameHash int64  `json:"filename_hash"`
	DataHash     int64  `json:"data_hash"`
	FileName     string `json:"filename"`
	Mime         int64  `json:"mime"`
	Size         int64  `json:"size"`
	Private      bool   `json:"private"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}
