package carapace

import (
	"os"

	"github.com/spf13/cobra"
)

// ActionCallback invokes a go function during completion.
func ActionCallback(callback CompletionCallback) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// ActionExecCommand executes an external command.
//
//	carapace.ActionExecCommand("git", "remote")(func(output []byte) carapace.Action {
//	  lines := strings.Split(string(output), "\n")
//	  return carapace.ActionValues(lines[:len(lines)-1]...)
//	})
func ActionExecCommand(name string, arg ...string) func(f func(output []byte) Action) Action {
	_ = "STUB: not implemented"
	return nil
}

// ActionExecCommandE is like ActionExecCommand but with custom error handling.
//
//	carapace.ActionExecCommandE("supervisorctl", "--configuration", path, "status")(func(output []byte, err error) carapace.Action {
//		if err != nil {
//			const NOT_RUNNING = 3
//			if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != NOT_RUNNING {
//				return carapace.ActionMessage(err.Error())
//			}
//		}
//		return carapace.ActionValues("success")
//	})
func ActionExecCommandE(name string, arg ...string) func(f func(output []byte, err error) Action) Action {
	_ = "STUB: not implemented"
	return nil
}

// seems this needs to be set manually due to stdout being collected?

// ActionImport parses the json output from export as Action
//
//	carapace.Gen(rootCmd).PositionalAnyCompletion(
//		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
//			args := []string{"_carapace", "export", ""}
//			args = append(args, c.Args...)
//			args = append(args, c.Value)
//			return carapace.ActionExecCommand("command", args...)(func(output []byte) carapace.Action {
//				return carapace.ActionImport(output)
//			})
//		}),
//	)
func ActionImport(output []byte) Action { _ = "STUB: not implemented"; return *new(Action) }

// ActionExecute executes completion on an internal command
// TODO example.
func ActionExecute(cmd *cobra.Command) Action { _ = "STUB: not implemented"; return *new(Action) }

// ActionDirectories completes directories.
func ActionDirectories() Action { _ = "STUB: not implemented"; return *new(Action) }

// TODO duplicated from ActionFiles

// ActionFiles completes files with optional suffix filtering.
func ActionFiles(suffix ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

// ActionValues completes arbitrary keywords (values).
func ActionValues(values ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

// ActionStyledValues is like ActionValues but also accepts a style.
func ActionStyledValues(values ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

// ActionValuesDescribed completes arbitrary key (values) with an additional description (value, description pairs).
func ActionValuesDescribed(values ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

// ActionStyledValuesDescribed is like ActionValues but also accepts a style.
func ActionStyledValuesDescribed(values ...string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// ActionMessage displays a help messages in places where no completions can be generated.
func ActionMessage(msg string, args ...any) Action { _ = "STUB: not implemented"; return *new(Action) }

// ActionMultiParts completes parts of an argument separated by sep.
func ActionMultiParts(sep string, callback func(c Context) Action) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// ActionMultiPartsN is like ActionMultiParts but limits the number of parts to `n`.
func ActionMultiPartsN(sep string, n int, callback func(c Context) Action) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// ActionStyleConfig completes style configuration
//
//	carapace.Value=blue
//	carapace.Description=magenta
func ActionStyleConfig() Action { _ = "STUB: not implemented"; return *new(Action) }

// Actionstyles completes styles
//
//	blue
//	bg-magenta
func ActionStyles(styles ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

// ActionExecutables completes executables either from PATH or given directories
//
//	nvim
//	chmod
func ActionExecutables(dirs ...string) Action { _ = "STUB: not implemented"; return *new(Action) }

// TODO allow additional descriptions to be registered somewhere for carapace-bin (key, value,...)

func actionDirectoryExecutables(dir string, prefix string, manDescriptions map[string]string) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// TODO trim slash suffix from dir | backslash path possible? (windows)

func isExecAny(mode os.FileMode) bool { _ = "STUB: not implemented"; return false }

// ActionPositional completes positional arguments for given command ignoring `--` (dash).
// TODO: experimental - likely gives issues with preinvoke (does not have the full args)
//
//	carapace.Gen(cmd).DashAnyCompletion(
//		carapace.ActionPositional(cmd),
//	)
func ActionPositional(cmd *cobra.Command) Action { _ = "STUB: not implemented"; return *new(Action) }

// ActionCommands completes (sub)commands of given command.
// `Context.Args` is used to traverse the command tree further down. Use `Action.Shift` to avoid this.
//
//	carapace.Gen(helpCmd).PositionalAnyCompletion(
//		carapace.ActionCommands(rootCmd),
//	)
func ActionCommands(cmd *cobra.Command) Action { _ = "STUB: not implemented"; return *new(Action) }

// cmd.Find is too lenient

// skip `_carapace` subcommand

// skip all hidden commands

// skip deprecated commands

// alias -> actual name

// ActionCora bridges given cobra completion function.
func ActionCobra(f func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)) Action {
	_ = "STUB: not implemented"
	return *new(Action)
}

// ensure cmd is never nil even if context does not contain one
