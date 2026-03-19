package geminiCLI

import (
	. "github.com/kore-01/CLIProxyAPI/v6/internal/constant"
	"github.com/kore-01/CLIProxyAPI/v6/internal/interfaces"
	"github.com/kore-01/CLIProxyAPI/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		GeminiCLI,
		Claude,
		ConvertGeminiCLIRequestToClaude,
		interfaces.TranslateResponse{
			Stream:     ConvertClaudeResponseToGeminiCLI,
			NonStream:  ConvertClaudeResponseToGeminiCLINonStream,
			TokenCount: GeminiCLITokenCount,
		},
	)
}
