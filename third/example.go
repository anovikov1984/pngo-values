package third

// Using builders
// pros
// - can safely omit values you don't want to affect
// cons:
// - builders aren't a usual go pattern, so SDK will look strange among the
// others
// - user still couldn't pass a nil
type GrantOpts struct {
	Read    interface{} // default is nil
	Write   interface{} // default is nil
	Ttl     int         // default is 0
	AuthKey string      // default is ""
}

func Run() {
	// allow only read permissions and do not change write
	grant().Write(true).AuthKey("me")
	// will generate request like:
	// read=true&authKey=me&ttl=0
}
