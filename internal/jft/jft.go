package jft

import (
	"fmt"

	"github.com/NasSilverBullet/jft/internal/db"
	"github.com/NasSilverBullet/jft/internal/model"
	"github.com/spf13/cobra"
)

func Exec() *cobra.Command {
	cmd := &cobra.Command{
		// TODO: 時間があれば、説明を充実する
		Use:   "jft",
		Short: "calendar cli tool, just for today",
		Long:  ``,
	}
	return cmd
}

func Add() *cobra.Command {
	var description string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add today's new plans",
		// TODO: 時間があれば、説明を充実する,
		Long: ``,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			p, err := model.NewPlan(db, args[0], args[1], args[2], description)
			if err != nil {
				return err
			}
			fmt.Println("added a new plan!!")
			fmt.Println(p)
			return err
		},
	}
	cmd.Flags().StringVarP(&description, "description", "d", "", "detailed description")
	return cmd
}

func Update() *cobra.Command {
	var start, end, title, description string
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update today's each plan, please give me id",
		// TODO: 時間があれば、説明を充実する,
		Long: `hoge`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			p, err := model.GetPlan(db, args[0])
			if err != nil {
				return err
			}
			p, err = p.Update(db, start, end, title, description)
			if err != nil {
				return err
			}
			fmt.Println("updated the plan!!")
			fmt.Println(p)
			return err
		},
	}
	cmd.Flags().StringVarP(&start, "start", "s", "", "start time")
	cmd.Flags().StringVarP(&end, "end", "e", "", "end time")
	cmd.Flags().StringVarP(&title, "title", "t", "", "short description")
	cmd.Flags().StringVarP(&description, "desc", "d", "", "detailed description")
	return cmd
}

func Delete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete today's each plan, please give me id",
		// TODO: 時間があれば、説明を充実する,
		Long: ``,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			p, err := model.GetPlan(db, args[0])
			if err != nil {
				return err
			}
			p, err = p.Delete(db)
			if err != nil {
				return err
			}
			fmt.Println("deleted the plan!!")
			fmt.Println(p)
			return err
		},
	}
	return cmd
}

func List() *cobra.Command {
	var month string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "show plans list",
		// TODO: 時間があれば、説明を充実する,
		Long: ``,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := db.New()
			if err != nil {
				return err
			}
			model.MigratePlan(db)
			defer func() {
				err = db.Close()
			}()
			ps, err := model.FindPlans(db, month)
			if err != nil {
				return err
			}
			var lineFeed string
			for _, p := range ps {
				fmt.Print(lineFeed)
				fmt.Println(p)
				lineFeed = "\n"
			}
			return err
		},
	}
	cmd.Flags().StringVarP(&month, "month", "m", "", "choose month")
	return cmd
}
