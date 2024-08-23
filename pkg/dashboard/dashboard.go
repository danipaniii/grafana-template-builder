package dashboard

type CreateDashboard struct {
	Dashboard Dashboard `json:"dashboard"`
	FolderUid string    `json:"folderUid"`
	Message   string    `json:"message"`
	Overwrite bool      `json:"overwrite"`
}

type Dashboard struct {
	Title    string        `json:"title"`
	Editable bool          `json:"editable"`
	Panels   []interface{} `json:"panels"`
}
