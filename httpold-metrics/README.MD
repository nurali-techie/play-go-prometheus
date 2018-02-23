**Info:** Example shows how to use metrics for http_requests and http_response provided by prometheus.

**Run:** go run main.go

**Output:**

HELP http_requests_total Total number of HTTP requests made.
TYPE http_requests_total counter
http_requests_total{code="200",handler="myapp",method="get"} 5
http_requests_total{code="200",handler="prometheus",method="get"} 4

**Explain:**

For handler="myapp", we have instrumented "myapp" handler using prometheus metrics.
All request/response to "myapp" handler will be track by prometheus.InstrumentHandler to record http_request and http_response metrics.

**Code:**

http.Handle("/myapp", prometheus.InstrumentHandler("myapp", http.HandlerFunc(myAppHander)))

Here, myAppHander is actual http handler which belong to myapp and which we are instrumenting using prometheus.
