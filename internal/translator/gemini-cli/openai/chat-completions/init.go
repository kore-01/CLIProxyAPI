package chat_completions

import (
	. "github.com/kore-01/CLIProxyAPI/v6/internal/constant"
	"github.com/kore-01/CLIProxyAPI/v6/internal/interfaces"
	"github.com/kore-01/CLIProxyAPI/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		OpenAI,
		GeminiCLI,
		ConvertOpenAIRequestToGeminiCLI,
		interfaces.TranslateResponse{
			Stream:    ConvertCliResponseToOpenAI,
			NonStream: ConvertCliResponseToOpenAINonStream,
		},
	)
}
