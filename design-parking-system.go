package leetcode

type ParkingSystem struct {
	parkingSpots map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	ps := ParkingSystem{parkingSpots: make(map[int]int)}
	ps.parkingSpots[1] = big
	ps.parkingSpots[2] = medium
	ps.parkingSpots[3] = small
	return ps
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.parkingSpots[carType] == 0 {
		return false
	}
	this.parkingSpots[carType]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
