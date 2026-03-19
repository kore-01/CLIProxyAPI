package responses

import (
	. "github.com/kore-01/CLIProxyAPI/v6/internal/constant"
	"github.com/kore-01/CLIProxyAPI/v6/internal/interfaces"
	"github.com/kore-01/CLIProxyAPI/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		OpenaiResponse,
		OpenAI,
		ConvertOpenAIResponsesRequestToOpenAIChatCompletions,
		interfaces.TranslateResponse{
			Stream:    ConvertOpenAIChatCompletionsResponseToOpenAIResponses,
			NonStream: ConvertOpenAIChatCompletionsResponseToOpenAIResponsesNonStream,
		},
	)
}
