package elon

import "fmt"

// TODO: define the 'Drive()' method
func(c *Car) Drive() {
    if c.battery > c.batteryDrain {
        c.distance += c.speed
    	c.battery -= c.batteryDrain
    }
}

// TODO: define the 'DisplayDistance() string' method
func(c *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func(c *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func(c *Car) CanFinish(trackDistance int) bool {
    batteryCycles := (c.battery / c.batteryDrain)
    return batteryCycles * c.speed >= trackDistance
}
