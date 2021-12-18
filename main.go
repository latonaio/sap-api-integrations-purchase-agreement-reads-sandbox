package main

import (
	sap_api_caller "sap-api-integrations-purchase-scheduling-agreement-reads/SAP_API_Caller"
	"sap-api-integrations-purchase-scheduling-agreement-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Purchase_Scheduling_Item_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Item",
		}
	}

	caller.AsyncGetPurchaseSchedulingAgreement(
		inoutSDC.PurchaseSchedulingAgreement.SchedulingAgreement,
		inoutSDC.PurchaseSchedulingAgreement.PurchaseSchedulingAgreementItem.SchedulingAgreementItem,
		accepter,
	)
}
