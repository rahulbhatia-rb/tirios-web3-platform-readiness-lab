package lifecycle

type State string

const (
	Draft State = "draft"
	Listed State = "listed"
	Tokenized State = "tokenized"
	Fundraising State = "fundraising"
	Funded State = "funded"
	RevenueActive State = "revenue_active"
	SecondaryMarket State = "secondary_market"
	Exited State = "exited"
)

var transitions = map[State][]State{
	Draft: {Listed},
	Listed: {Tokenized},
	Tokenized: {Fundraising},
	Fundraising: {Funded},
	Funded: {RevenueActive},
	RevenueActive: {SecondaryMarket, Exited},
	SecondaryMarket: {Exited},
}

func CanTransition(from, to State) bool {
	for _, allowed := range transitions[from] {
		if allowed == to {
			return true
		}
	}
	return false
}
