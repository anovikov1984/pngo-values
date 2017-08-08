package second

// Using interfaces
// pros
// - can safely omit values you don't want to affect
// cons:
// - type of expected value is dynamic and incorrect type assignation
// cannot be detected at compile time
// - hard to define which type should user use for specific field
type GrantOpts struct {
	Read    interface{} // default is nil
	Write   interface{} // default is nil
	Ttl     int         // default is 0
	AuthKey string      // default is ""
}

func Run() {
	// allow only read permissions and do not change write
	// omitting Write will set it to nil, so we can easily check for nil and
	// ignore it in a request
	grant(&GrantOpts{
		Read:    true,
		AuthKey: "me",
	})

	// will generate request like:
	// read=true&authKey=me&ttl=0
}
