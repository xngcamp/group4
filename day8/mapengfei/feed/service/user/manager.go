package user

//import "github.com/simplejia/clog"

type Manager struct {
	emails map[string]bool
}

var EmailManager = NewContainer()

func NewContainer() *Manager {
	return &Manager{
		emails: map[string]bool{},
	}
}

func (c *Manager) Exists(email string) (ok bool) {
	_, ok = c.emails[email]
	return
}

func (c *Manager) AddEmail(email string) {
	c.emails[email] = true
	//	clog.Debug("Rigister Done: %v", c.emails)
}
