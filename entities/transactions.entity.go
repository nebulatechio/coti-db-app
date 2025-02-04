package entities

import (
	"database/sql"
	"github.com/coti-io/coti-db-app/dto"
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID                             int32               `json:"id" gorm:"column:id;type:int(11) NOT NULL AUTO_INCREMENT"`
	Hash                           string              `json:"hash" gorm:"column:hash;type:varchar(100) COLLATE utf8_unicode_ci NOT NULL;index:hash_INDEX"`
	Index                          *int32              `json:"index" gorm:"column:index;type:int(11) DEFAULT NULL;index:index_INDEX"`
	Amount                         decimal.Decimal     `json:"amount" gorm:"column:amount;type:decimal(25,10) NOT NULL"`
	AttachmentTime                 decimal.Decimal     `json:"attachmentTime" gorm:"column:attachmentTime;type:decimal(20,6) NOT NULL;index:attachmentTime_INDEX"`
	IsValid                        sql.NullBool        `json:"isValid" gorm:"column:isValid;type:tinyint(4) DEFAULT NULL"`
	IsSafeBalanceIncreased         bool                `json:"isSafeBalanceIncreased" gorm:"column:IsSafeBalanceIncreased;type:tinyint(4) DEFAULT NULL;index:isSafeBalanceIncreased_INDEX;index:second_composite_INDEX,priority:1;"`
	IsSafeBalanceDecrease          bool                `json:"isSafeBalanceDecrease" gorm:"column:IsSafeBalanceDecrease;type:tinyint(4) DEFAULT NULL;index:isSafeBalanceDecrease_INDEX"`
	TransactionCreateTime          decimal.Decimal     `json:"transactionCreateTime" gorm:"column:transactionCreateTime;type:decimal(20,6) NOT NULL"`
	LeftParentHash                 *string             `json:"leftParentHash" gorm:"column:leftParentHash;type:varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL"`
	RightParentHash                *string             `json:"rightParentHash" gorm:"column:rightParentHash;type:varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL"`
	NodeHash                       *string             `json:"nodeHash" gorm:"column:nodeHash;type:varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL"`
	SenderHash                     *string             `json:"senderHash" gorm:"column:senderHash;type:varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL"`
	SenderTrustScore               float64             `json:"senderTrustScore" gorm:"column:senderTrustScore;type:decimal(25,10) NOT NULL"`
	TransactionConsensusUpdateTime decimal.NullDecimal `json:"transactionConsensusUpdateTime" gorm:"column:transactionConsensusUpdateTime;type:decimal(20,6);index:composite_INDEX,priority:3;index:transactionConsensusUpdateTime_INDEX;index:second_composite_INDEX,priority:4;"`
	TransactionDescription         *string             `json:"transactionDescription" gorm:"column:transactionDescription;type:varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL"`
	TrustChainConsensus            bool                `json:"trustChainConsensus" gorm:"column:trustChainConsensus;type:tinyint(4) DEFAULT NULL;index:trustChainConsensus_INDEX"`
	TrustChainTrustScore           decimal.Decimal     `json:"trustChainTrustScore" gorm:"column:trustChainTrustScore;type:decimal(25,10) NOT NULL"`
	Type                           *string             `json:"type" gorm:"column:type;type:varchar(100) COLLATE utf8_unicode_ci NOT NULL;index:type_INDEX;index:composite_INDEX,priority:2;index:second_composite_INDEX,priority:3;"`
	IsProcessed                    bool                `json:"isProcessed" gorm:"column:isProcessed;type:tinyint(4) DEFAULT false;index:isProcessed_INDEX;index:composite_INDEX,priority:1;index:second_composite_INDEX,priority:2;"`
	CreateTime                     time.Time           `json:"createTime" gorm:"column:createTime;type:timestamp NOT NULL;default:CURRENT_TIMESTAMP;"`
	UpdateTime                     time.Time           `json:"updateTime" gorm:"column:updateTime;type:timestamp NOT NULL;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index:updateTime_INDEX"`
}

func NewTransaction(tx *dto.TransactionResponse) *Transaction {
	instance := new(Transaction)
	instance.Hash = tx.Hash
	instance.Index = tx.Index
	instance.Amount = tx.Amount
	instance.AttachmentTime = tx.AttachmentTime
	instance.TransactionCreateTime = tx.CreateTime
	instance.LeftParentHash = tx.LeftParentHash
	instance.RightParentHash = tx.RightParentHash
	instance.NodeHash = tx.NodeHash
	instance.SenderHash = tx.SenderHash
	instance.SenderTrustScore = tx.SenderTrustScore
	instance.TransactionConsensusUpdateTime = tx.TransactionConsensusUpdateTime
	instance.TransactionDescription = tx.TransactionDescription
	instance.TrustChainConsensus = tx.TrustChainConsensus
	instance.TrustChainTrustScore = tx.TrustChainTrustScore
	instance.Type = tx.Type
	return instance
}
