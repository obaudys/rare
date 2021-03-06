package expressions

import "path/filepath"

func kfPathManip(manipulator func(string) string) func([]KeyBuilderStage) KeyBuilderStage {
	return func(args []KeyBuilderStage) KeyBuilderStage {
		if len(args) != 1 {
			return stageLiteral(ErrorArgCount)
		}
		return KeyBuilderStage(func(context KeyBuilderContext) string {
			return manipulator(args[0](context))
		})
	}
}

var (
	kfPathBase = kfPathManip(filepath.Base)
	kfPathDir  = kfPathManip(filepath.Dir)
	kfPathExt  = kfPathManip(filepath.Ext)
)
