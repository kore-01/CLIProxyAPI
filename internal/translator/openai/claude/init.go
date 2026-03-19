package claude

import (
	. "github.com/kore-01/CLIProxyAPI/v6/internal/constant"
	"github.com/kore-01/CLIProxyAPI/v6/internal/interfaces"
	"github.com/kore-01/CLIProxyAPI/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		Claude,
		OpenAI,
		ConvertClaudeRequestToOpenAI,
		interfaces.TranslateResponse{
			Stream:     ConvertOpenAIResponseToClaude,
			NonStream:  ConvertOpenAIResponseToClaudeNonStream,
			TokenCount: ClaudeTokenCount,
		},
	)
}
