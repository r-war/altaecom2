package response

import "AltaEcom/business/transcation"

type TranscationsResponse struct {
	Transcation []TranscationResponse `json:"transcations"`
}

func GetTranscationsResponse(trans []transcation.Transcation) TranscationsResponse {
	getTranscations := TranscationsResponse{}

	var transcations TranscationResponse
	for _, val := range trans {
		transcations.ID = val.ID
		transcations.OrderID = val.OrderID
		transcations.Status = string(val.Status)
		transcations.Total = val.Total
		transcations.TypePayment = string(val.TypePayment)
		transcations.DatePayment = val.DatePayment
		transcations.Comment = val.Comment

		getTranscations.Transcation = append(getTranscations.Transcation, transcations)
	}

	if getTranscations.Transcation != nil {
		getTranscations.Transcation = []TranscationResponse{}
	}

	return getTranscations
}
