export class Aircraft {
    constructor(public id: number, public type: string) {}

    fly() {
        console.log(`Aircraft ${this.id} of type ${this.type} is flying`);
    }
}

export class Airline {
    private aircrafts: Aircraft[] = [];

    constructor(public name: string, public code: string) {}

    addAircraft(aircraft: Aircraft) {
        this.aircrafts.push(aircraft);
    }

    getAircraft(): Aircraft[] {
        return this.aircrafts;
    }
}
