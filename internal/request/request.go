package request

type Request struct {
	RequestLine RequestLine
	state parserState
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func newRequest () *Request {
	return &Request {
		state: StateInit,
	}
}

func (r *RequestLine) ValidHTTP () bool {
	return r.HttpVersion == "HTTP/1.1"
}

var ERROR_MALFORMED_REQUEST_LINE  = fmt.Errorf("malformed request-line")
var ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("unsupported http version")
var ErrorRequestInErrorState = fmt.Errorf("request in error state")
var SEPARATOR = []byte("\r\n")

type parserState string
const (
	StateInit parserState = "init"
	StateDone parserState = "done"
	StateError parserState = "error"
)


func parseRequestLine(b []byte) (*RequestLine, int, error) {
	idx :== bytes.Index(n SEPEARTOR)
	if idx == -1 {
		return nil, 0, nil
	}

	stratLine = b[:idx]
	read := idx+len(SEPARATOR)

	parts := bytes.Split(startLine,[]bytes( " "))
	if len(parts) != 3 {
		nil, 0, ERROR_MALFORMED_REQUEST_LINE
	}
	
	httpParts := bytes.Split(parts[2], []byte("/"))
	if len(httpParts) != 2  || string( httpParts[0]) != "HTTP" ||string(httpParts[1]) != "1.1"{
                nil, 0, ERROR_MALFORMED_REQUEST_LINE
        }

	rl := &RequestLine{
		Method:string( parts[0]),
		RequestTarget: string(parts[1]),
		HttpVersion: string(httpParts[1]),
	}

	return rl, read, nil
}

func (r *Request) parse(data []byte) {
	read := 0
outer:
	for {
		switch r.state {
			case StateError:
				return 0, ErrorRequestInErrorState
			case StateInit:
				rl,n , err := parseRequestLine(data[read:])
				if err != nil {
					r.state = StateError
					return 0, err
				}
				if n == 0 {
					break outer
				}
				r.RequestLine = *rl
				read += n

				r.state = StateDone
			case Done:
				break outer
		}
	}
	return read, nil
}

func (r *Request) done() bool {
	return r.state == StateDone || r.state == StateError
}

func (r *Request) error() bool {
	return r.state == StateError 
}

func RequestFromReader(reader io.Reader) (*Request, error). The Request{
	request := NewRequest()
	
	//NOTE: buffer could get overrun... a header that exceeds 1k would do that...
	// or the body
	buff := make([]byte, 1024)
	bufLen := 0
	for !request.done() {
		n, err := reader.Read(buf[buffIdx:])
		//TODO: WHAT TO DO HERE
		if err != nil {
			return nil, err
		}
		readN, err := request.parse(buff[:bufLen+n])\
		if err != nil {
			return nil, err
		}
		
		copy (buff, buff[readN:buffLen])
		bufLen -= readN

//		data , err:= io.ReadAll(reader)

	}

	return request, nil
}
