package httph

import (
	"net/http"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CanonicalHeaderKeyKey = "CanonicalHeaderKey"

	DetectContentTypeKey = "DetectContentType"

	ErrLineTooLongKey = "ErrLineTooLong"

	ErrorKey = "Error"

	FileServerKey = "FileServer"

	GetKey = "Get"

	HandleKey = "Handle"

	HandleFuncKey = "HandleFunc"

	HeadKey = "Head"

	ListenAndServeKey = "ListenAndServe"

	ListenAndServeTLSKey = "ListenAndServeTLS"

	MaxBytesReaderKey = "MaxBytesReader"

	NewFileTransportKey = "NewFileTransport"

	NewRequestKey = "NewRequest"

	NewServeMuxKey = "NewServeMux"

	NotFoundKey = "NotFound"

	NotFoundHandlerKey = "NotFoundHandler"

	ParseHTTPVersionKey = "ParseHTTPVersion"

	ParseTimeKey = "ParseTime"

	PostKey = "Post"

	PostFormKey = "PostForm"

	ProxyFromEnvironmentKey = "ProxyFromEnvironment"

	ProxyURLKey = "ProxyURL"

	ReadRequestKey = "ReadRequest"

	ReadResponseKey = "ReadResponse"

	RedirectKey = "Redirect"

	RedirectHandlerKey = "RedirectHandler"

	ServeKey = "Serve"

	ServeContentKey = "ServeContent"

	ServeFileKey = "ServeFile"

	ServeTLSKey = "ServeTLS"

	SetCookieKey = "SetCookie"

	StatusTextKey = "StatusText"

	StripPrefixKey = "StripPrefix"

	TimeoutHandlerKey = "TimeoutHandler"
)

func New() hctx.Map {
	return hctx.Map{

		CanonicalHeaderKeyKey: CanonicalHeaderKey,

		DetectContentTypeKey: DetectContentType,

		ErrLineTooLongKey: ErrLineTooLong,

		ErrorKey: Error,

		FileServerKey: FileServer,

		GetKey: Get,

		HandleKey: Handle,

		HandleFuncKey: HandleFunc,

		HeadKey: Head,

		ListenAndServeKey: ListenAndServe,

		ListenAndServeTLSKey: ListenAndServeTLS,

		MaxBytesReaderKey: MaxBytesReader,

		NewFileTransportKey: NewFileTransport,

		NewRequestKey: NewRequest,

		NewServeMuxKey: NewServeMux,

		NotFoundKey: NotFound,

		NotFoundHandlerKey: NotFoundHandler,

		ParseHTTPVersionKey: ParseHTTPVersion,

		ParseTimeKey: ParseTime,

		PostKey: Post,

		PostFormKey: PostForm,

		ProxyFromEnvironmentKey: ProxyFromEnvironment,

		ProxyURLKey: ProxyURL,

		ReadRequestKey: ReadRequest,

		ReadResponseKey: ReadResponse,

		RedirectKey: Redirect,

		RedirectHandlerKey: RedirectHandler,

		ServeKey: Serve,

		ServeContentKey: ServeContent,

		ServeFileKey: ServeFile,

		ServeTLSKey: ServeTLS,

		SetCookieKey: SetCookie,

		StatusTextKey: StatusText,

		StripPrefixKey: StripPrefix,

		TimeoutHandlerKey: TimeoutHandler,
	}
}

// CanonicalHeaderKey returns the canonical format of the
// header key s. The canonicalization converts the first
// letter and any letter following a hyphen to upper case;
// the rest are converted to lowercase. For example, the
// canonical key for &#34;accept-encoding&#34; is &#34;Accept-Encoding&#34;.
// If s contains a space or invalid header field bytes, it is
// returned without modifications.
var CanonicalHeaderKey = http.CanonicalHeaderKey

// DetectContentType implements the algorithm described
// at http://mimesniff.spec.whatwg.org/ to determine the
// Content-Type of the given data. It considers at most the
// first 512 bytes of data. DetectContentType always returns
// a valid MIME type: if it cannot determine a more specific one, it
// returns &#34;application/octet-stream&#34;.
var DetectContentType = http.DetectContentType

var ErrLineTooLong = http.ErrLineTooLong

// Error replies to the request with the specified error message and HTTP code.
// It does not otherwise end the request; the caller should ensure no further
// writes are done to w.
// The error message should be plain text.
var Error = http.Error

// FileServer returns a handler that serves HTTP requests
// with the contents of the file system rooted at root.
//
// To use the operating system&#39;s file system implementation,
// use http.Dir:
//
//     http.Handle(&#34;/&#34;, http.FileServer(http.Dir(&#34;/tmp&#34;)))
//
// As a special case, the returned file server redirects any request
// ending in &#34;/index.html&#34; to the same path, without the final
// &#34;index.html&#34;.
var FileServer = http.FileServer

// Get issues a GET to the specified URL. If the response is one of
// the following redirect codes, Get follows the redirect, up to a
// maximum of 10 redirects:
//
//    301 (Moved Permanently)
//    302 (Found)
//    303 (See Other)
//    307 (Temporary Redirect)
//    308 (Permanent Redirect)
//
// An error is returned if there were too many redirects or if there
// was an HTTP protocol error. A non-2xx response doesn&#39;t cause an
// error.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
//
// Get is a wrapper around DefaultClient.Get.
//
// To make a request with custom headers, use NewRequest and
// DefaultClient.Do.
var Get = http.Get

// Handle registers the handler for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
var Handle = http.Handle

// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
var HandleFunc = http.HandleFunc

// Head issues a HEAD to the specified URL. If the response is one of
// the following redirect codes, Head follows the redirect, up to a
// maximum of 10 redirects:
//
//    301 (Moved Permanently)
//    302 (Found)
//    303 (See Other)
//    307 (Temporary Redirect)
//    308 (Permanent Redirect)
//
// Head is a wrapper around DefaultClient.Head
var Head = http.Head

// ListenAndServe listens on the TCP network address addr
// and then calls Serve with handler to handle requests
// on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
// Handler is typically nil, in which case the DefaultServeMux is
// used.
//
// A trivial example server is:
//
// 	package main
//
// 	import (
// 		&#34;io&#34;
// 		&#34;net/http&#34;
// 		&#34;log&#34;
// 	)
//
// 	// hello world, the web server
// 	func HelloServer(w http.ResponseWriter, req *http.Request) {
// 		io.WriteString(w, &#34;hello, world!\n&#34;)
// 	}
//
// 	func main() {
// 		http.HandleFunc(&#34;/hello&#34;, HelloServer)
// 		log.Fatal(http.ListenAndServe(&#34;:12345&#34;, nil))
// 	}
//
// ListenAndServe always returns a non-nil error.
var ListenAndServe = http.ListenAndServe

// ListenAndServeTLS acts identically to ListenAndServe, except that it
// expects HTTPS connections. Additionally, files containing a certificate and
// matching private key for the server must be provided. If the certificate
// is signed by a certificate authority, the certFile should be the concatenation
// of the server&#39;s certificate, any intermediates, and the CA&#39;s certificate.
//
// A trivial example server is:
//
// 	import (
// 		&#34;log&#34;
// 		&#34;net/http&#34;
// 	)
//
// 	func handler(w http.ResponseWriter, req *http.Request) {
// 		w.Header().Set(&#34;Content-Type&#34;, &#34;text/plain&#34;)
// 		w.Write([]byte(&#34;This is an example server.\n&#34;))
// 	}
//
// 	func main() {
// 		http.HandleFunc(&#34;/&#34;, handler)
// 		log.Printf(&#34;About to listen on 10443. Go to https://127.0.0.1:10443/&#34;)
// 		err := http.ListenAndServeTLS(&#34;:10443&#34;, &#34;cert.pem&#34;, &#34;key.pem&#34;, nil)
// 		log.Fatal(err)
// 	}
//
// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
//
// ListenAndServeTLS always returns a non-nil error.
var ListenAndServeTLS = http.ListenAndServeTLS

// MaxBytesReader is similar to io.LimitReader but is intended for
// limiting the size of incoming request bodies. In contrast to
// io.LimitReader, MaxBytesReader&#39;s result is a ReadCloser, returns a
// non-EOF error for a Read beyond the limit, and closes the
// underlying reader when its Close method is called.
//
// MaxBytesReader prevents clients from accidentally or maliciously
// sending a large request and wasting server resources.
var MaxBytesReader = http.MaxBytesReader

// NewFileTransport returns a new RoundTripper, serving the provided
// FileSystem. The returned RoundTripper ignores the URL host in its
// incoming requests, as well as most other properties of the
// request.
//
// The typical use case for NewFileTransport is to register the &#34;file&#34;
// protocol with a Transport, as in:
//
//   t := &amp;http.Transport{}
//   t.RegisterProtocol(&#34;file&#34;, http.NewFileTransport(http.Dir(&#34;/&#34;)))
//   c := &amp;http.Client{Transport: t}
//   res, err := c.Get(&#34;file:///etc/passwd&#34;)
//   ...
var NewFileTransport = http.NewFileTransport

// NewRequest returns a new Request given a method, URL, and optional body.
//
// If the provided body is also an io.Closer, the returned
// Request.Body is set to body and will be closed by the Client
// methods Do, Post, and PostForm, and Transport.RoundTrip.
//
// NewRequest returns a Request suitable for use with Client.Do or
// Transport.RoundTrip. To create a request for use with testing a
// Server Handler, either use the NewRequest function in the
// net/http/httptest package, use ReadRequest, or manually update the
// Request fields. See the Request type&#39;s documentation for the
// difference between inbound and outbound request fields.
//
// If body is of type *bytes.Buffer, *bytes.Reader, or
// *strings.Reader, the returned request&#39;s ContentLength is set to its
// exact value (instead of -1), GetBody is populated (so 307 and 308
// redirects can replay the body), and Body is set to NoBody if the
// ContentLength is 0.
var NewRequest = http.NewRequest

// NewServeMux allocates and returns a new ServeMux.
var NewServeMux = http.NewServeMux

// NotFound replies to the request with an HTTP 404 not found error.
var NotFound = http.NotFound

// NotFoundHandler returns a simple request handler
// that replies to each request with a ``404 page not found&#39;&#39; reply.
var NotFoundHandler = http.NotFoundHandler

// ParseHTTPVersion parses a HTTP version string.
// &#34;HTTP/1.0&#34; returns (1, 0, true).
var ParseHTTPVersion = http.ParseHTTPVersion

// ParseTime parses a time header (such as the Date: header),
// trying each of the three formats allowed by HTTP/1.1:
// TimeFormat, time.RFC850, and time.ANSIC.
var ParseTime = http.ParseTime

// Post issues a POST to the specified URL.
//
// Caller should close resp.Body when done reading from it.
//
// If the provided body is an io.Closer, it is closed after the
// request.
//
// Post is a wrapper around DefaultClient.Post.
//
// To set custom headers, use NewRequest and DefaultClient.Do.
//
// See the Client.Do method documentation for details on how redirects
// are handled.
var Post = http.Post

// PostForm issues a POST to the specified URL, with data&#39;s keys and
// values URL-encoded as the request body.
//
// The Content-Type header is set to application/x-www-form-urlencoded.
// To set other headers, use NewRequest and DefaultClient.Do.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
//
// PostForm is a wrapper around DefaultClient.PostForm.
//
// See the Client.Do method documentation for details on how redirects
// are handled.
var PostForm = http.PostForm

// ProxyFromEnvironment returns the URL of the proxy to use for a
// given request, as indicated by the environment variables
// HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions
// thereof). HTTPS_PROXY takes precedence over HTTP_PROXY for https
// requests.
//
// The environment values may be either a complete URL or a
// &#34;host[:port]&#34;, in which case the &#34;http&#34; scheme is assumed.
// An error is returned if the value is a different form.
//
// A nil URL and nil error are returned if no proxy is defined in the
// environment, or a proxy should not be used for the given request,
// as defined by NO_PROXY.
//
// As a special case, if req.URL.Host is &#34;localhost&#34; (with or without
// a port number), then a nil URL and nil error will be returned.
var ProxyFromEnvironment = http.ProxyFromEnvironment

// ProxyURL returns a proxy function (for use in a Transport)
// that always returns the same URL.
var ProxyURL = http.ProxyURL

// ReadRequest reads and parses an incoming request from b.
var ReadRequest = http.ReadRequest

// ReadResponse reads and returns an HTTP response from r.
// The req parameter optionally specifies the Request that corresponds
// to this Response. If nil, a GET request is assumed.
// Clients must call resp.Body.Close when finished reading resp.Body.
// After that call, clients can inspect resp.Trailer to find key/value
// pairs included in the response trailer.
var ReadResponse = http.ReadResponse

// Redirect replies to the request with a redirect to url,
// which may be a path relative to the request path.
//
// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
var Redirect = http.Redirect

// RedirectHandler returns a request handler that redirects
// each request it receives to the given url using the given
// status code.
//
// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
var RedirectHandler = http.RedirectHandler

// Serve accepts incoming HTTP connections on the listener l,
// creating a new service goroutine for each. The service goroutines
// read requests and then call handler to reply to them.
// Handler is typically nil, in which case the DefaultServeMux is used.
var Serve = http.Serve

// ServeContent replies to the request using the content in the
// provided ReadSeeker. The main benefit of ServeContent over io.Copy
// is that it handles Range requests properly, sets the MIME type, and
// handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since,
// and If-Range requests.
//
// If the response&#39;s Content-Type header is not set, ServeContent
// first tries to deduce the type from name&#39;s file extension and,
// if that fails, falls back to reading the first block of the content
// and passing it to DetectContentType.
// The name is otherwise unused; in particular it can be empty and is
// never sent in the response.
//
// If modtime is not the zero time or Unix epoch, ServeContent
// includes it in a Last-Modified header in the response. If the
// request includes an If-Modified-Since header, ServeContent uses
// modtime to decide whether the content needs to be sent at all.
//
// The content&#39;s Seek method must work: ServeContent uses
// a seek to the end of the content to determine its size.
//
// If the caller has set w&#39;s ETag header formatted per RFC 7232, section 2.3,
// ServeContent uses it to handle requests using If-Match, If-None-Match, or If-Range.
//
// Note that *os.File implements the io.ReadSeeker interface.
var ServeContent = http.ServeContent

// ServeFile replies to the request with the contents of the named
// file or directory.
//
// If the provided file or directory name is a relative path, it is
// interpreted relative to the current directory and may ascend to parent
// directories. If the provided name is constructed from user input, it
// should be sanitized before calling ServeFile. As a precaution, ServeFile
// will reject requests where r.URL.Path contains a &#34;..&#34; path element.
//
// As a special case, ServeFile redirects any request where r.URL.Path
// ends in &#34;/index.html&#34; to the same path, without the final
// &#34;index.html&#34;. To avoid such redirects either modify the path or
// use ServeContent.
var ServeFile = http.ServeFile

// ServeTLS accepts incoming HTTPS connections on the listener l,
// creating a new service goroutine for each. The service goroutines
// read requests and then call handler to reply to them.
//
// Handler is typically nil, in which case the DefaultServeMux is used.
//
// Additionally, files containing a certificate and matching private key
// for the server must be provided. If the certificate is signed by a
// certificate authority, the certFile should be the concatenation
// of the server&#39;s certificate, any intermediates, and the CA&#39;s certificate.
var ServeTLS = http.ServeTLS

// SetCookie adds a Set-Cookie header to the provided ResponseWriter&#39;s headers.
// The provided cookie must have a valid Name. Invalid cookies may be
// silently dropped.
var SetCookie = http.SetCookie

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
var StatusText = http.StatusText

// StripPrefix returns a handler that serves HTTP requests
// by removing the given prefix from the request URL&#39;s Path
// and invoking the handler h. StripPrefix handles a
// request for a path that doesn&#39;t begin with prefix by
// replying with an HTTP 404 not found error.
var StripPrefix = http.StripPrefix

// TimeoutHandler returns a Handler that runs h with the given time limit.
//
// The new Handler calls h.ServeHTTP to handle each request, but if a
// call runs for longer than its time limit, the handler responds with
// a 503 Service Unavailable error and the given message in its body.
// (If msg is empty, a suitable default message will be sent.)
// After such a timeout, writes by h to its ResponseWriter will return
// ErrHandlerTimeout.
//
// TimeoutHandler buffers all Handler writes to memory and does not
// support the Hijacker or Flusher interfaces.
var TimeoutHandler = http.TimeoutHandler
