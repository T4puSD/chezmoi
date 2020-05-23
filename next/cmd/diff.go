package cmd

import (
	"github.com/spf13/cobra"

	"github.com/twpayne/chezmoi/next/internal/chezmoi"
)

var diffCmd = &cobra.Command{
	Use:     "diff [targets...]",
	Short:   "Print the diff between the target state and the destination state",
	Long:    mustGetLongHelp("diff"),
	Example: getExample("diff"),
	RunE:    config.runDiffCmd,
}

type diffCmdConfig struct {
	include   *chezmoi.IncludeSet
	recursive bool
	NoPager   bool
	Pager     string
}

func init() {
	rootCmd.AddCommand(diffCmd)

	persistentFlags := diffCmd.PersistentFlags()
	persistentFlags.VarP(config.Diff.include, "include", "i", "include entry types")
	persistentFlags.BoolVar(&config.Diff.NoPager, "no-pager", config.Diff.NoPager, "disable pager")
	persistentFlags.BoolVarP(&config.Diff.recursive, "recursive", "r", config.Diff.recursive, "recursive")

	markRemainingZshCompPositionalArgumentsAsFiles(diffCmd, 1)
}

func (c *Config) runDiffCmd(cmd *cobra.Command, args []string) error {
	gitDiffSystem := chezmoi.NewGitDiffSystem(chezmoi.NewDiscardWritesSystem(c.system), c.DestDir+chezmoi.PathSeparatorStr, c.colored)
	if err := c.applyArgs(gitDiffSystem, c.DestDir, args, c.Diff.include, c.Diff.recursive); err != nil {
		return err
	}
	return c.writeOutputString(gitDiffSystem.String())
}
