package solana

import (
	"github.com/0xjeffro/tx-parser/solana/programs/computeBudget"
	ComputeBudgetParsers "github.com/0xjeffro/tx-parser/solana/programs/computeBudget/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	PumpfunParsers "github.com/0xjeffro/tx-parser/solana/programs/pumpfun/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/systemProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	TokenProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/tokenProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func router(result *types.ParsedResult, i int) (types.Action, error) {
	programID := result.AccountList[result.RawTx.Transaction.Message.Instructions[i].ProgramIDIndex]
	instruction := result.RawTx.Transaction.Message.Instructions[i]
	switch programID {
	case systemProgram.Program:
		return SystemProgramParsers.InstructionRouter(result, instruction)
	case tokenProgram.Program:
		return TokenProgramParsers.InstructionRouter(result, instruction)
	case computeBudget.Program:
		return ComputeBudgetParsers.InstructionRouter(result, instruction)
	case pumpfun.Program:
		return PumpfunParsers.InstructionRouter(result, instruction)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       programID,
				ProgramName:     "Unknown",
				InstructionName: "Unknown",
			},
		}, nil
	}
}
