package main

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/Dansukoffc/go-devices/electronic"
	"github.com/Dansukoffc/go-devices/computing"
	"github.com/Dansukoffc/go-devices/phone"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clearScreen()

	// Electronic Device
	device1 := electronic.NewElectronicDevice("Smart Watch", 499, "Built-in Battery", true)
	device1.PrintInfo()
	device1.PrintResale()

	// Computing Device
	computer1 := computing.NewComputingDevice("Gaming Laptop", 1499, "AC Adapter", false, 8500, 64, "35.6 x 24.5 x 2.3 cm")
	computer1.PrintInfo()
	computer1.PrintResale()

	// Phone
	phone1 := phone.NewPhone("iPhone 15 Pro", 1099, "Li-Ion Battery", true, 7200, 64, "14.6 x 7.1 x 0.8 cm",
		"iPhone 15 Pro", 2023, true)
	phone1.PrintInfo()
	phone1.PrintResale()

	phone2 := phone.NewPhone("Galaxy S24 Ultra", 1299, "Li-Ion Battery", true, 8500, 64, "16.3 x 7.9 x 0.9 cm",
		"Samsung Galaxy S24 Ultra", 2024, true)
	phone2.PrintInfo()
	phone2.PrintResale()

	fmt.Println("=== Getters ===")
	fmt.Println("Electronic Device:", device1.Name(), "| Price:", device1.Price())
	fmt.Println("Computing Device:", computer1.Name(), "| Power:", computer1.ComputingPower(), "GFLOPS")
	fmt.Println("Phone:", phone1.ModelName(), "| Year:", phone1.ReleaseYear())
}
