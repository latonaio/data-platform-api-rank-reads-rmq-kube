package dpfm_api_output_formatter

import (
	"data-platform-api-rank-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToRank(rows *sql.Rows) (*[]Rank, error) {
	defer rows.Close()
	rank := make([]Rank, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Rank{}

		err := rows.Scan(
			&pm.RankType,
			&pm.Rank,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &rank, nil
		}

		data := pm
		rank = append(rank, Rank{
			RankType:				data.RankType,
			Rank:					data.Rank,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &rank, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.RankType,
			&pm.Rank,
			&pm.Language,
			&pm.RankName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			RankType:				data.RankType,
			Rank:     				data.Rank,
			Language:          		data.Language,
			RankName:				data.RankName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
