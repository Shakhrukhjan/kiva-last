package Loan

type Entreps struct {
	Amount    int    `json:"amount" form:"amount"`
	ClientID  string `json:"client_id" form:"client_id"`
	FirstName string `json:"first_name" form:"first_name"`
	Gender    string `json:"gender" form:"gender"`
	LastName  string `json:"last_name" form:"last_name"`
	LoanID    string `json:"loan_id" form:"loan_id"`
}
type Schedule struct {
	Date      string `json:"date" form:"date"`
	Interest  float64    `json:"interest" form:"interest"`
	Principal float64    `json:"principal" form:"principal"`
}
type IndividualLoan struct {
	ActivityID            int        `json:"activity_id" form:"activity_id"`                         // ---- ID for the activity the loan is associated with.
	ClientWaiverSigned    bool       `json:"client_waiver_signed" form:"client_waiver_signed"`       // ---- Flag indicating whether or not the client has signed a waiver.
	Currency              string     `json:"currency" form:"currency"`                               // ---- Currency code for the loan.
	Description           string     `json:"description" form:"description"`                         // ---- Description of the loan.
	DescriptionLanguageID int        `json:"description_language_id" form:"description_language_id"` // ---- ID for the language used in the description.
	DisburseTime          string     `json:"disburse_time" form:"disburse_time"`                     // ---- Date the loan will be disbursed.
	Entreps               []Entreps  `json:"entreps" form:"entreps"`                                 // ---- List of borrowers.
	Loanuse               string     `json:"loanuse" form:"loanuse"`                                 // ---- What the loan will be used for.
	Location              string     `json:"location" form:"location"`                               // ---- Location to associate with the loan.
	Schedule              []Schedule `json:"schedule" form:"schedule"`
	ThemeTypeID           int        `json:"theme_type_id" form:"theme_type_id"` // ---- ID for the activity the loan theme.
}

func (loan *IndividualLoan) SetActivityID(activity_id int) {
	loan.ActivityID = activity_id
}

func (loan *IndividualLoan) SetClient_waiver_signed(client_waiver_signed bool) {
	loan.ClientWaiverSigned = client_waiver_signed
}

func (loan *IndividualLoan) SetCurrency(currency string) {
	loan.Currency = currency
}

func (loan *IndividualLoan) SetDescription(description string) {
	loan.Description = description
}

func (loan *IndividualLoan) SetDescriptionLanguageID(description_language_id int) {
	loan.DescriptionLanguageID = description_language_id
}

func (loan *IndividualLoan) SetDisburseTime(disburse_time string) {
	loan.DisburseTime = disburse_time
}

func (loan *IndividualLoan) SetEntreps(entreps Entreps) {
	loan.Entreps = append(loan.Entreps, entreps)
}

func (loan *IndividualLoan) SetLoanuse(loanuse string) {
	loan.Loanuse = loanuse
}

func (loan *IndividualLoan) SetLocation(location string) {
	loan.Location = location
}

func (loan *IndividualLoan) SetSchedule(schedule []Schedule) {
	loan.Schedule = schedule
}

func (loan *IndividualLoan) SetThemeTypeID(theme_type_id int) {
	loan.ThemeTypeID = theme_type_id
}
