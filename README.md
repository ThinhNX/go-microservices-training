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