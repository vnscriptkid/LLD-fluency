
export class WeeklySchedule {
    constructor(public day: string, public time: string) {}
}

export class Flight {
    constructor(public flightNumber: string, public schedule: WeeklySchedule) {}

    depart() {
        console.log(`Flight ${this.flightNumber} departing on ${this.schedule.day} at ${this.schedule.time}`);
    }
}
