module github.com/RevitalS/someone-to-run-with-app/backend/shaynaservice

require (
	github.com/RevitalS/someone-to-run-with-app/backend/foundation v0.0.0
	github.com/julienschmidt/httprouter v1.3.0
)

replace github.com/RevitalS/someone-to-run-with-app/backend/foundation => ../foundation

go 1.15
