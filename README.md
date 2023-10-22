# go-microservices-training
this project for learning purpose only.
Try to implement RESTFul API with GIN.
New features are coming soon, wait for it.

/token/get: Get FreeToken
/v1 Group: Add freetoken to the "Authorization" request header

Project flow:
 - Send Get request to /v1/panicCheck/check to check is panic or not ( trigger panic check middleware to switch panik true/false)
 - Send POST request to /token/get to get token. Request body type: form-data with form ID: admin
 - Send POST Request to /v1/login with parameter: user=thinhnx&pass=123. AND Authorization header is Token in Setp 2.
 - Everything in /v1 group need Authorization in Header. Panik Check Path just simple make handler panic and return from middleware.


 HTTPS cert:

openssl genrsa -des3 -out rootCA.key 2048
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024 -out rootCA.pem

create server.csr.cnf and v3.ext

openssl req -new -sha256 -nodes -out server.csr -newkey rsa:2048 -keyout server.key -config ./server.csr.cnf
openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 800 -sha256 -extfile v3.ext