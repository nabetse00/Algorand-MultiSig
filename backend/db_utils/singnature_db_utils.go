package db_utils

import (
	"errors"
	"fmt"
	"multisigdb-svc/db"
	"multisigdb-svc/dto"
	"multisigdb-svc/model"

	"github.com/algorand/go-algorand-sdk/types"
)

func UpdateNumberOfSignsRequired(txnId string) bool {

	currentNumberOfSigns, err := GetCurrentNumberOfSignature(txnId)
	if err != nil {
		return false
	}

	if currentNumberOfSigns == 0 {
		return false
	}
	currentNumberOfSigns -= 1

	tx := db.DbConnection.Model(&model.RawTxn{}).Where("txn_id = ?", txnId).Update("number_of_signs_required", currentNumberOfSigns)
	if tx.Error != nil {
		return false
	}
	if tx.RowsAffected == 0 {
		return false
	}
	return true
}

func AddRawTxn(txn model.RawTxn) error {
	resp := db.DbConnection.Create(&txn)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func AddSignersAddrs(txn []model.SignerAddress) error {
	resp := db.DbConnection.Create(&txn)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func AddSignedTxn(txn model.SignedTxn) error {

	resp := db.DbConnection.Create(&txn)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func GetAllSignedTxnOnTxnId(txnId string) (dto.GetAllSingedTxnResponse, error) {
	var signedTxns []model.SignedTxn
	tx := db.DbConnection.Where("txn_id = ?", txnId).Find(&signedTxns)

	if tx.Error != nil {
		return dto.GetAllSingedTxnResponse{
			Message: tx.Error.Error(),
			Success: false,
		}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return dto.GetAllSingedTxnResponse{
			Message: "No Record Found",
			Success: false,
		}, tx.Error
	}

	return dto.GetAllSingedTxnResponse{
		Message: "Txn Found",
		Success: true,
		Txn:     signedTxns,
	}, nil
}

func GetTxnOnTxnId(txnId string) (dto.GetSingedTxnResponse, error) {
	var signedTxn model.SignedTxn
	tx := db.DbConnection.Where("txn_id = ?", txnId).Last(&signedTxn)

	if tx.Error != nil {
		return dto.GetSingedTxnResponse{
			Message: tx.Error.Error(),
			Success: false,
		}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return dto.GetSingedTxnResponse{
			Message: "No Record Found",
			Success: false,
		}, tx.Error
	}

	return dto.GetSingedTxnResponse{
		Message: "Txn Found",
		Success: true,
		Txn:     signedTxn,
	}, nil
}

func GetRawTxnOnTxnId(txnId string) (dto.GetRawTxnResponse, error) {
	var rawTxn model.RawTxn
	tx := db.DbConnection.Where("txn_id = ?", txnId).Last(&rawTxn)

	if tx.Error != nil {
		return dto.GetRawTxnResponse{
			Message: tx.Error.Error(),
			Success: false,
		}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return dto.GetRawTxnResponse{
			Message: "No Record Found",
			Success: false,
		}, tx.Error
	}

	addrs, err := FindAllSignerAddrWithTxnId(txnId)

	if err != nil {
		return dto.GetRawTxnResponse{
			Message: tx.Error.Error(),
			Success: false,
		}, tx.Error
	}

	if len(addrs) != int(rawTxn.NumberOfSignsTotal) {
		return dto.GetRawTxnResponse{
			Message: "Number of signers mismatch ",
			Success: false,
		}, tx.Error
	}

	return dto.GetRawTxnResponse{
		Message:      "Txn Found",
		Success:      true,
		Txn:          rawTxn,
		SignersAddrs: addrs,
	}, nil
}

func GetRawTxnSignersOnTxnId(txnId string) (dto.GetRawTxnSignersAddrsResponse, error) {
	var signersAddrs []model.SignerAddress
	tx := db.DbConnection.Where("sign_txn_id = ?", txnId).Find(&signersAddrs)

	if tx.Error != nil {
		return dto.GetRawTxnSignersAddrsResponse{
			Success: false,
			Message: tx.Error.Error(),
		}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return dto.GetRawTxnSignersAddrsResponse{
			Message: "No Record Found",
			Success: false,
		}, tx.Error
	}

	return dto.GetRawTxnSignersAddrsResponse{
		Success: true,
		Message: "Txn Found",
		Addrs:   signersAddrs,
	}, nil
}

func GetCurrentNumberOfSignature(txnId string) (int64, error) {
	var rawTxn model.RawTxn
	tx := db.DbConnection.Where("txn_id = ?", txnId).Find(&rawTxn)

	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, nil
	}

	return rawTxn.NumberOfSignsRequired, nil
}

func UpdateStatusOfTransaction(txn, status string) error {
	err := db.DbConnection.Model(&model.RawTxn{}).Where("txn_id = ?", txn).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckIfATxnExistsAndSignsNeeded(txn string) bool {
	var rawTxn model.RawTxn
	num := db.DbConnection.Where("txn_id = ?", txn).Where("number_of_signs_required != ?", 0).Find(&rawTxn).RowsAffected
	return num > 0
}

func CheckIfAPubAddressFound(txn, pubAddress string) bool {
	var signedTxn model.SignedTxn
	num := db.DbConnection.Where("txn_id = ?", txn).Where("signer_public_address = ?", pubAddress).Find(&signedTxn).RowsAffected
	return num == 0

}

func FindAllTxnWithTxnStatusReady() []model.RawTxn {
	var rawTxnsWithStatusReady []model.RawTxn
	db.DbConnection.Where("status = ?", "READY").Find(&rawTxnsWithStatusReady)
	return rawTxnsWithStatusReady
}

func UpdateStatusOfTransactionToDone(txn string) error {
	err := db.DbConnection.Model(&model.RawTxn{}).Where("txn_id = ?", txn).Update("status", "BROADCASTED").Error
	if err != nil {
		return err
	}
	return nil
}

func GetTxnIdOnAddr(addr string) (dto.GetTxnIdsResponse, error) {

	_, err := isValidAddress(addr)

	if err != nil {
		return dto.GetTxnIdsResponse{
			Success: false,
			Message: "Error invalid address",
		}, err
	}

	var signers []model.SignerAddress
	err = db.DbConnection.Where("signer_address = ?", addr).Find(&signers).Error
	if err != nil {
		return dto.GetTxnIdsResponse{
			Success: false,
			Message: "Error getting txnIds wrong signer address?",
		}, err
	}

	if len(signers) == 0 {
		return dto.GetTxnIdsResponse{
			Success: true,
			Message: "No txnid found",
			TxnIds:  []string{},
		}, nil
	}

	var addrTxnIds []string

	for _, e := range signers {
		addrTxnIds = append(addrTxnIds, e.SignTxnId)
	}

	return dto.GetTxnIdsResponse{
		Success: true,
		Message: fmt.Sprintf("Valid txnIds for address %s found", addr),
		TxnIds:  addrTxnIds,
	}, nil
}

func FindAllSignerAddrWithTxnId(txnId string) ([]model.SignerAddress, error) {
	var signers []model.SignerAddress
	err := db.DbConnection.Where("sign_txn_id = ?", txnId).Find(&signers).Error
	return signers, err
}

func GetDoneTxn(txnId string) (model.DoneTransaction, error) {
	var doneTxn model.DoneTransaction
	result := db.DbConnection.Where("txn_id = ?", txnId).Last(&doneTxn)

	if result.Error != nil {
		return model.DoneTransaction{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.DoneTransaction{}, errors.New(fmt.Sprintf("No Transaction found for %s", txnId))
	}

	if result.RowsAffected > 1 {
		return model.DoneTransaction{}, errors.New(fmt.Sprintf("More than one Transaction found for %s", txnId))
	}

	return doneTxn, nil
}

func AddDoneTxn(txnId string, doneTxnId string) error {
	var doneTxn = model.DoneTransaction{
		TxnId:         txnId,
		TransactionID: doneTxnId,
	}
	err := db.DbConnection.Create(&doneTxn).Error
	if err != nil {
		return err
	}
	return nil
}

func isValidAddress(addr string) (bool, error) {
	_, err := types.DecodeAddress(addr)
	if err != nil {
		return false, err
	}
	return true, nil
}
