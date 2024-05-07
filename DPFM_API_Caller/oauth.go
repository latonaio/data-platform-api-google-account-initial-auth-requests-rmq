package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-google-account-initial-auth-requests-rmq/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-google-account-initial-auth-requests-rmq/DPFM_API_Output_Formatter"
	"data-platform-api-google-account-initial-auth-requests-rmq/config"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) GoogleAccountInitialAuth(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) *[]dpfm_api_output_formatter.GoogleAccountInitialAuth {
	var googleAccountInitialAuth []dpfm_api_output_formatter.GoogleAccountInitialAuth

	googleAccountInitialAuth = append(googleAccountInitialAuth, dpfm_api_output_formatter.GoogleAccountInitialAuth{
		URL: conf.OAuth.AuthURL,
	})

	return &googleAccountInitialAuth
}
