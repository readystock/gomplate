package test

type Account struct {
	UpdatedBy uint64
}

func (account *Account) SetUpdatedBy(id uint64) {
	account.UpdatedBy = id
}

type Accounts []*Account
