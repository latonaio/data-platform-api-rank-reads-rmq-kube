package requests

type Rank struct {
	RankType			string	`json:"RankType"`
	Rank				int		`json:"Rank"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
