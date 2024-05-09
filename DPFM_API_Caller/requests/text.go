package requests

type Text struct {
	RankType			string	`json:"RankType"`
	Rank				int		`json:"Rank"`
	Language          	string  `json:"Language"`
	RankName			string 	`json:"RankName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
