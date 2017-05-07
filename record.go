package whois

// Record represents a parsed whois response.

type Contact struct {
	ID            string
	Type          string
	Name          string
	Organization  string
	Street        []string
	City          string
	StateProvince string
	PostalCode    string
	Country       string
	Phone         string
	PhoneExt      string
	Fax           string
	FaxExt        string
	Email         string
}

type Record struct {
	Registrar             string
	Whois                 string
	SponsoringRegistrarID int
	ReferralURL           string
	NameServers           []string
	Status                []string
	LastUpdate            string
	Expire                string
	Creation              string
	DNSSEC                string
	Contacts              []*Contact
}
