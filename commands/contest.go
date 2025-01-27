package commands

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var contestCommand = &cobra.Command{
	Use:   "contest",
	Short: "Get contests",
	RunE:  fetchContests,
}

func fetchContests(cmd *cobra.Command, args []string) error {
	api, err := contestsApi()
	if err != nil {
		return fmt.Errorf("could not connect to the API; %w", err)
	}

	if contestId != "" {
		c, err := api.ContestById(contestId)

		if err != nil {
			return fmt.Errorf("could not retrieve contest; %w", err)
		}

		fmt.Printf(" %10s: %s\n", c.Id, c.Name)
		fmt.Printf("             %v starting at %v\n", c.Duration, c.StartTime)
	} else {
		c, err := api.Contests()

		if err != nil {
			return fmt.Errorf("could not retrieve contests; %w", err)
		}

		// sort by start time
		sort.Slice(c, func(i, j int) bool {
			return c[i].StartTime.Time().Before(c[j].StartTime.Time())
		})

		// output
		fmt.Printf("Contests (%d):\n", len(c))
		for _, o := range c {
			fmt.Printf(" %10s: %s\n", o.Id, o.Name)
			fmt.Printf("             %v starting at %v\n", o.Duration, o.StartTime)
		}
	}

	return nil
}
