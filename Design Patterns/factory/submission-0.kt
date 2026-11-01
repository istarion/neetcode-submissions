interface Vehicle {
    fun getType(): String
}

class Car : Vehicle {
    override fun getType(): String {
        return "Car"
    }
}

class Bike : Vehicle {
    override fun getType(): String {
        return "Bike"
    }
}

class Truck : Vehicle {
    override fun getType(): String {
        return "Truck"
    }
}

abstract class VehicleFactory {
    abstract fun createVehicle(): Vehicle
}

class CarFactory : VehicleFactory() {
    override fun createVehicle(): Vehicle {
        return Car()
    }
}

class BikeFactory : VehicleFactory() {
    override fun createVehicle(): Vehicle {
        return Bike()
    }
}

class TruckFactory : VehicleFactory() {
    override fun createVehicle(): Vehicle {
        return Truck()
    }
}
