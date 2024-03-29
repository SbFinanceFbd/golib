package constant

const (
	YES = "YES"
	NO  = "NO"

	BEFORE = "BEFORE"
	AFTER  = "AFTER"

	INITIATED = "INITIATED"
	PENDING   = "PENDING"
	APPROVED  = "APPROVED"
	DECLINED  = "DECLINED"
	CLOSED    = "CLOSED"

	ACTIVE   = "ACTIVE"
	INACTIVE = "INACTIVE"

	FORCE_CLOSE = "FORCE_CLOSE"
	CLOSE       = "CLOSE"

	PAY_TYPES_DAILY       = "DAILY"
	PAY_TYPES_WEEKLY      = "WEEKLY"
	PAY_TYPES_MONTHLY     = "MONTHLY"
	PAY_TYPES_QUARTERLY   = "QUARTERLY"
	PAY_TYPES_HALF_YEARLY = "HALF_YEARLY"
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

	//Member status
	CUST_ACTIVE   = "ACTIVE"
	CUST_INACTIVE = "INACTIVE"

	// Loan status
	LOAN_REGULAR        = "REGULAR"        // REGULAR Payment
	LOAN_IRREGULAR      = "IRREGULAR"      // not REGULAR
	LOAN_RENEWAL_PERIOD = "RENEWAL_PERIOD" //
	LOAN_OVER_DUE       = "OVER_DUE"       // number of emis or number days completed.
	LOAN_CLOSED         = "CLOSED"         // Loan Account Closed

	RD_DEPOSIT_PERIOD  = "DEPOSIT_PERIOD"  // Member desposit amount RD account
	RD_MATURITY_PERIOD = "MATURITY_PERIOD" // Member Deposit period completed . and waiting for maturity amount
	RD_HOLD            = "HOLD"            // Hold for Loan purpose or time took to transafer amount to member
	RD_CLOSE           = "CLOSE"           // Rd Account Closed with payment of maturity amount
	RD_FORCE_CLOSE     = "FORCE_CLOSE"     // Rd Account Force Closed(Break) with payment of maturity amount
	RD_INACTIVE        = "INACTIVE"        // Member can not get Matruity amount, interest amount

	// FD status
	FD_ACTIVE   = "ACTIVE"   // once DEPOSITED change status to ACTIVE
	FD_INACTIVE = "INACTIVE" // Member can not get Matruity amount, interest amount
	FD_HOLD     = "HOLD"     // Hold for Loan purpose or time took to transafer amount to member
	FD_CLOSED   = "CLOSED"   // Fd Account Closed with payment of maturity amount
	//FD_FORCE_CLOSED     = "FORCE_CLOSED"   // Fd Account Closed with payment of maturity amount and without Interest amount

	// Pigmy status
	PIGMY_ACTIVE      = "ACTIVE"
	PIGMY_INACTIVE    = "INACTIVE"
	PIGMY_HAS_LOAN    = "HAS_LOAN"
	PIGMY_LOAN_CLOSED = "LOAN_CLOSED"
	PIGMY_CLOSED      = "CLOSED"

	DATE_FORMATE = "^\\d{4}-(0[1-9]|1[012])-(0[1-9]|[12]\\d|3[01])$" // YYYYMMDD
)
