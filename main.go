package main

import (
	"fmt"
	"os"

	"github.com/iMartyn/animal/src"
	"github.com/spf13/cobra"
)

func main() {
	var animalname string
	animal.AddAnimals("etc/animals.json")
	var rootCmd = &cobra.Command{
		Use:   "animal",
		Short: "animal is a quick go app to use in k8s turorials",
		Long:  `animal takes an animal name and displays a statement`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(animalname) > 0 {
				animalFound := animal.FindAnimal(animalname)
				if animalFound.AnimalName == "" {
					fmt.Printf("I don't know what sound a %s makes.\n", animalname)
				} else {
					fmt.Printf("The %s %s\n", animalFound.AnimalName, animalFound.AnimalSound)
				}
			} else {
				fmt.Println("Enter valid input. Hint, there isn't one!")
				cmd.Help()
			}
		},
	}
	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve http requests",
		Long:  "Run the webserver to serve http requests",
		Run: func(cmd *cobra.Command, args []string) {
			animal.HandleHTTP()
		},
	}

	rootCmd.Flags().StringVarP(&animalname, "animalname", "p", "", "Animal name")
	serveCmd.Flags().StringVarP(&animal.AnimalName, "animalname", "p", "", "Animal name")
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
