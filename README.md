go-better-sessions 
==============

Go wrapper for gorilla sessions
http://www.gorillatoolkit.org/pkg/sessions

# Installation

```
$ go get github.com/gorilla/sessions
$ go get github.com/Ulbora/go-better-sessions

```

# Usage

## Initialize and use
```
    import usess github.com/Ulbora/go-better-sessions
    var s usess.Session
    func main() {
	    s.MaxAge = 5 * 60
	    s.Name = "user-session-test"
	    s.SessionKey = "554dfgdffdd11dfgf1ff1f"
    }

    func handleSomething(res http.ResponseWriter, req *http.Request) {
	    s.InitSessionStore(res, req)
        session, err := s.GetSession(req)
	    if err != nil {
		    http.Error(res, err.Error(), http.StatusInternalServerError)
	    }
	    user := session.Values["username"]

        session.Values["someVal"] = "someValue"
        session.Values["someOtherVal"] = 55
        //do something
    }
```

# Important Note:
If you aren't using gorilla/mux, you need to wrap your handlers with context.ClearHandler as or else you will leak memory! An easy way to do this is to wrap the top-level mux when calling http.ListenAndServe: