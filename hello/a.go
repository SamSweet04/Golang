package main

import ("fmt" 
        "hello/hello2"
		"github.com/spf13/cobra")

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")

			hello2.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
}