#include <unordered_map>
using namespace std;

class ParkingSystem {
public:
    unordered_map<int, int> parkingSpots;
    ParkingSystem(int big, int medium, int small) {
        parkingSpots[1] = big;
        parkingSpots[2] = medium;
        parkingSpots[3] = small;
    }
    
    bool addCar(int carType) {
        if (this->parkingSpots[carType] == 0) {
            return false;
        }
        this->parkingSpots[carType]--;
        return true;
    }
};
