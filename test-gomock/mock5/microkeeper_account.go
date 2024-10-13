package mock5

import (
	"context"
	"encoding/json"
)

type MicrokeeperAccount struct {
	provider            *MicrokeeperProvider
	svcIntegrationsAuth *IntegrationsAuth
}

type UpsertAbnData struct {
	Abn           string `json:"abn"`
	AbnName       string `json:"abn_name"`
	AbnTrading    string `json:"abn_trading"`
	AbnBranch     string `json:"abn_branch"`
	AbnEmail      string `json:"abn_email"`
	AbnAccounting string `json:"abn_accounting"`
	AbnBankId     string `json:"abn_BDID"`
}

func (mke *MicrokeeperAccount) UpsertABN(ctx context.Context) error {
	// Get Microkeeper business key
	businessKey, err := mke.svcIntegrationsAuth.GetAccessToken(ctx)
	if err != nil {
		return err
	}

	// Build the data payload
	mkData := &UpsertAbnData{
		Abn:           "business.BusinessNumber",
		AbnName:       "business.LegalName",
		AbnTrading:    "business.TradingName",
		AbnBranch:     "business.BranchCode",
		AbnEmail:      "business.OwnerEmail",
		AbnAccounting: "business.AccountingSuite",
		AbnBankId:     "business.BankDetails[0].Id", // Currently only one bank detail is supported by MK. TODO confirm this
	}
	payload, err := json.Marshal(mkData)
	if err != nil {
		return err
	}

	// Call to Microkeeper via svc-integrations
	res, err := mke.provider.Request(ctx, "upsertAbnAction", payload, businessKey)
	if err != nil || len(res) < 1 {
		return err
	}

	return nil
}
