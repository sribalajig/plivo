1. Install Golang 1.7 => https://golang.org/doc/install
2. Install beego which is a micro framework for Go.  Execute "go get github.com/astaxie/beego" in the command line.
3. Install redis.v3 which is the redis client for Go.  Execute "go get gopkg.in/redis.v3" in the command line.
4. Install libpq which is the PostGre driver for Go.  Execute "go get github.com/lib/pq" in the command line.
5. Navigate to /plivo/plivo.sms.api and execute the following command :  "bee run". A server will be started which listens to connections on port 3000.
6. To run the tests, navigate to the root folder (/plivo) and execute the following command : "go test ./...". This command will run all tests in all subfolders under the root.