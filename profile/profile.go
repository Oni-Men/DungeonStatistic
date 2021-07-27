package profile

type Profile struct {
	MCID *string
	UUID *string
}

var profiles = make(map[string]*Profile)

func AddProfile(p *Profile) {
	profiles[*p.UUID] = p
}

func GetProfile(uuid *string) *Profile {
	p, ok := profiles[*uuid]

	if ok {
		return p
	}

	return nil
}
