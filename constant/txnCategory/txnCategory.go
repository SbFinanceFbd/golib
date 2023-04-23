package txnCategory

const (
	INF = "INF" // Initial fund (CR)
	// INFM       = "INFM"       // Initial fund Mistake (CR/DB)
	IVF = "IVF" // Invest fund (CR)
	// IVFM       = "IVFM"       // Invest fund (CR)
	IVWF = "IVWF" // Invest Withdraw fund (CR)
	// IVWFM      = "IVWFM"      // Invest Withdraw fund (CR)
	LS        = "LS"        // Loan Sanction (DB)
	LSM       = "LSM"       // Loan Sanction Mistake (CR)
	LI        = "LI"        // Loan Interest (CR)
	LC        = "LC"        // Loan Collection (CR)
	LCM       = "LCM"       // Loan Collection Mistake (DB)
	LPC       = "LPC"       // Loan Penalty Collection (CR)
	LDP       = "LDP"       // Loan DisCount Pay (DB)
	RC        = "RC"        // RD Collection (CR)
	RCM       = "RCM"       // RD Collection Mistake (DB)
	RM        = "RM"        // RD Maturity (DB)
	RMM       = "RMM"       // RD Maturity Mistake (CR)
	RI        = "RI"        // RD Interest (DB)
	FC        = "FC"        // FD Collection (CR)
	FCM       = "FCM"       // FD Collection Mistake (DB)
	FI        = "FI"        // FD Interest (DB)
	FIFR      = "FIFR"      // FD Interest Force close Rollback (CR)
	FP        = "FP"        // FD Principle (DB)
	FPM       = "FPM"       // FD Principle Mistake (CR)
	PC        = "PC"        // Pigmy Collection (CR)
	PCM       = "PCM"       // Pigmy Collection Mistake (DB)
	PI        = "PI"        // Pigmy Interest (DB)
	PP        = "PP"        // Pigmy Maturity (DB)
	PPM       = "PPM"       // Pigmy Maturity Mistake (CR)
	BT        = "BT"        // Branch Transfer (BOTH CR DB)
	EMP_SAL   = "EMP_SAL"   // Employee Salary
	EMP_SAL_M = "EMP_SAL_M" // Employee Salary Mistake/Reversal
	OTHER     = "OTHER"     // OTHER Expenditure (DB)
)
