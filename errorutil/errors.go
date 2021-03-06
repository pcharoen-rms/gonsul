package errorutil

const ErrorDeleteNotAllowed				= 10
const ErrorBadParams 					= 20
const ErrorFailedConsulConnection 		= 30
const ErrorFailedConsulTxn 				= 31
const ErrorFailedReadingResponse 		= 40
const ErrorFailedJsonEncode 			= 50
const ErrorFailedJsonDecode 			= 51
const ErrorFailedCloning 				= 60
const ErrorFailedMustache 				= 70
const ErrorFailedHTTPServer				= 80

type GonsulError struct {
	Code int
}

func ExitError(err error, errorCode int, logger *Logger) {
	if err.Error() != "" {
		logger.PrintError(err.Error())
	}
	panic(GonsulError{Code: errorCode})
}
