package first

// Using Disable* booleans
// pros
// - seems like a more-go way to do that
// cons:
// - user have to explicitly use Disable* field to disable default value
type GrantOpts struct {
	Read    bool   // default is false
	Write   bool   // default is false
	Ttl     int    // default is 0
	AuthKey string // default is ""

	DisableRead  bool
	DisableWrite bool
}

func Run() {
	// allow only read permissions and do not change write
	grant(&Opts{
		Read:         true,
		DisableWrite: true,
		AuthKey:      "me",
	})

	// will generate request like:
	// read=true&authKey=me&ttl=0

	// omitting DisableWrite will generate write=false
	grant(&Opts{
		Read:    true,
		AuthKey: "me",
	})

	// will generate request like:
	// read=true&write=false&authKey=me&ttl=0
}
