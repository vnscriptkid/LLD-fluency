export class User {
    id: number;
    name: string;
    email: string;
    memberships: Membership[] = [];
    createdGroups: Group[] = [];

    constructor(id: number, name: string, email: string) {
        this.id = id;
        this.name = name;
        this.email = email;
    }

    joinGroup(group: Group): void {
        if (!this.isMemberOf(group)) {
            const membership = new Membership(Membership.generateId(), this, group);
            this.memberships.push(membership);
            group.addMembership(membership);
        }
    }

    leaveGroup(group: Group): void {
        this.memberships = this.memberships.filter(membership => membership.group.id !== group.id);
        group.removeMembership(this);
    }

    createGroup(name: string): Group {
        const group = new Group(Group.generateId(), name, this);
        this.createdGroups.push(group);
        this.joinGroup(group);
        return group;
    }

    private isMemberOf(group: Group): boolean {
        return this.memberships.some(membership => membership.group.id === group.id);
    }
}

export class Group {
    id: number;
    name: string;
    createdBy: User;
    memberships: Membership[] = [];

    constructor(id: number, name: string, createdBy: User) {
        this.id = id;
        this.name = name;
        this.createdBy = createdBy;
    }

    static generateId(): number {
        return Math.floor(Math.random() * 1000000);
    }

    addMembership(membership: Membership): void {
        if (!this.memberships.includes(membership)) {
            this.memberships.push(membership);
        }
    }

    removeMembership(user: User): void {
        this.memberships = this.memberships.filter(membership => membership.user.id !== user.id);
    }

    getMembers(): User[] {
        return this.memberships.map(membership => membership.user);
    }
}

export class Membership {
    id: number;
    user: User;
    group: Group;

    constructor(id: number, user: User, group: Group) {
        this.id = id;
        this.user = user;
        this.group = group;
    }

    static generateId(): number {
        return Math.floor(Math.random() * 1000000);
    }
}

// Example usage:
const user1 = new User(1, 'Alice', 'alice@example.com');
const user2 = new User(2, 'Bob', 'bob@example.com');

const group1 = user1.createGroup('Group1');
user2.joinGroup(group1);

console.log(user1);
console.log(user2);
console.log(group1);
console.log(group1.getMembers());
