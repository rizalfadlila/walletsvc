package wallet

// Mutation Statement
const (
	sqlStore = `insert into wallet (id, owned_by, status, enabled_at, balance) values (:id, :owned_by, :status, :enabled_at, :balance)`
)

var (
	sqlUpdateStatusByUserID = `update wallet set status = $1, %s = $2, %s = '' where owned_by = $3 and status = $4`
)

// Query Statement
const (
	sqlFindByUserID = `select id, owned_by, balance, enabled_at, disabled_at, status from wallet where owned_by = $1`
)
