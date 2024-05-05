import { Aircraft, Airline } from "./aggregation";
import { Flight, WeeklySchedule } from "./composition";

function main() {
    /////////// Aggregation ///////////
    console.log('== Aggregation ==')
    const airline = new Airline("Acme Airlines", "ACM");

    const aircraft1 = new Aircraft(101, "Boeing 737");
    const aircraft2 = new Aircraft(102, "Airbus A320");

    airline.addAircraft(aircraft1);
    airline.addAircraft(aircraft2);

    airline.getAircraft().forEach(aircraft => aircraft.fly());

    /////////// Composition ///////////
    console.log('== Composition ==')
    const schedule = new WeeklySchedule("Monday", "10:00 AM");
    const flight = new Flight("ACM123", schedule);

    flight.depart();
}

main();