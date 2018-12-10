# go-meli-toolkit

Go meli toolkit is a set of libraries aimed at writing Go code within the Meli ecosystem easier

## Package index

We have the following categories:

* Clients
    * gokvsclient: KVS client
    * godsclient: DS client
    * gobigqueue: BQ client
    * golockclient: Distributed lock client
    * restful/rest: Base rest client
    * gosequence: Sequence service client
    * gomemcached: Memcached client
    * datadog-go: Datadog metrics client
    * goosclient: Object storage service client

* Metrics
    * godog: Datadog metrics handler & aggregator
    * gorelic: NewRelic segments support

* Security
    * godin: ODIN MiddleEnd auth & autorization
    * mlauth: Meli request authentication

* Utilities
    * babel: Babel i18n support
    * gingonic: Gingonic handlers & helpers
    * gomelipass: Fury env variable provider
    * goutils: Miscellaneous utilities (logging, errors, helpers)

## Documentation

Each package has its corresponding documentation in its README.md file. Any additions are welcome

## Testing
A test runner script is located at the root level of the toolkit:
```bash  
# usage: sh test.sh [-c] [-v] [-m] [-p <specific_package>]
# -c produces coverage report
# -v runs go tests with verbose option
# -p <specific_package> runs go tests only from the specified package
# -u updates dependencies for each package using dep ensure

# example: sh test.sh -v -c -p "godog"
# parses for options -c, -v, -m and -t
./test.sh -u
```

## Installation
Packages may be installed normally via get: ```go get github.com/mercadolibre/go-meli-toolkit/___package___```, but authentication is required to be able to access our organization

If you're getting a ```Repository not found``` error, it is very likely that it is an auth problem. Check the following:

* For SSH auth, see [this](https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/) and [this](https://help.github.com/articles/testing-your-ssh-connection/)
* For HTTPS, see [this](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) and [this](https://help.github.com/articles/providing-your-2fa-authentication-code/#when-youll-be-asked-for-a-personal-access-token-as-a-password)

## Dependencies
[Go Dep](https://github.com/golang/dep#setup) is used to manage dependencies. Each package has a corresponding Gopkg.toml file that contains the versions used (when applicable). To download the dependencies into a local /vendor directory, use the command ```dep ensure```

## Mock and production modes
All the [clients](https://github.com/mercadolibre/go-meli-toolkit#package-index) use an underlying [rest client](https://github.com/mercadolibre/go-meli-toolkit/tree/master/restful) to communicate with their respective services. This is important when considering the possible mock modes:

* Productive client + Productive rest client
* Productive client + Mock rest client
* Mock client + Productive rest client
* Mock client + Mock rest client

These modes will be set on the ```init()``` function of the clients, and will depend on the value of the ```GO_ENVIRONMENT``` environment variable

For example, this is the init logic for the kvs client:

```go
    var MakeKvsClient func(string, KvsClientConfig) Client

    func init() {
        containers = make(map[kvsClient]*containerConfig)
        if os.Getenv("GO_ENVIRONMENT") == "production" || os.Getenv("GO_ENVIRONMENT") == "gokvsclient_test" {
            MakeKvsClient = makeRealKvsClient
        } else {
            rest.StartMockupServer()
            MakeKvsClient = makeMockKvsClient
        }
    }
```

If we start a mockup server with ```GO_ENVIRONMENT``` set to ```production```, we will get the *Productive client + Mock* rest client mode

## Questions?

[fury@mercadolibre.com](mailto:fury@mercadolibre.com)
