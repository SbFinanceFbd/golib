package constant

const (
	YES = "YES"
	NO  = "NO"

	PAY_TYPES_DAILY       = "DAILY"
	PAY_TYPES_WEEKLY      = "WEEKLY"
	PAY_TYPES_MONTHLY     = "MONTHLY"
	PAY_TYPES_HALF_YEARLY = "HALF_YEARLY"
	PAY_TYPES_QUARTERLY   = "QUARTERLY"
	PAY_TYPES_YEARLY      = "YEARLY"

	// Loan Pack status
	LOAN_PACK_ACTIVE   = "ACTIVE"
	LOAN_PACK_INACTIVE = "INACTIVE"

	// RD Pack status
	RD_PACK_ACTIVE   = "ACTIVE"
	RD_PACK_INACTIVE = "INACTIVE"

	// FD Pack status
	FD_PACK_ACTIVE   = "ACTIVE"
	FD_PACK_INACTIVE = "INACTIVE"

	// Pigmy Pack Status
	PIGMY_PACK_ACTIVE   = "ACTIVE"
	PIGMY_PACK_INACTIVE = "INACTIVE"

	//Customer status
	CUST_ACTIVE   = "ACTIVE"
	CUST_INACTIVE = "INACTIVE"

	// Loan status
	LOAN_REGULAR        = "REGULAR"        // REGULAR Payment
	LOAN_IRREGULAR      = "IRREGULAR"      // not REGULAR
	LOAN_RENEWAL_PERIOD = "RENEWAL_PERIOD" //
	LOAN_OVER_DUE       = "OVER_DUE"       // number of emis or number days completed.
	LOAN_CLOSED         = "CLOSED"         // Loan Account Closed

	// RD status
	NO_FIRST_PAYMENT = "NO_FIRST_PAYMENT"
	RD_ACTIVE        = "ACTIVE"   // once payment made change status to ACTIVE
	RD_INACTIVE      = "INACTIVE" // Customer can not get Matruity amount, interest amount
	RD_HOLD          = "HOLD"     // Hold for Loan purpose or time took to transafer amount to customer
	RD_CLOSED        = "CLOSED"   // Rd Account Closed with payment of maturity amount

	// FD status
	NOT_DEPOSITED = "NOT_DEPOSITED"
	FD_ACTIVE     = "ACTIVE"   // once DEPOSITED change status to ACTIVE
	FD_INACTIVE   = "INACTIVE" // Customer can not get Matruity amount, interest amount
	FD_HOLD       = "HOLD"     // Hold for Loan purpose or time took to transafer amount to customer
	FD_CLOSED     = "CLOSED"   // Fd Account Closed with payment of maturity amount

	// Pigmy status
	PIGMY_ACTIVE      = "ACTIVE"
	PIGMY_INACTIVE    = "INACTIVE"
	PIGMY_HAS_LOAN    = "HAS_LOAN"
	PIGMY_LOAN_CLOSED = "LOAN_CLOSED"
	PIGMY_CLOSED      = "CLOSED"
)
