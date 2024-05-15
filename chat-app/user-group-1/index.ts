export class User {
    id: number;
    name: string;
    email: string;
    groups: Group[] = [];
    createdGroups: Group[] = [];

    constructor(id: number, name: string, email: string) {
        this.id = id;
        this.name = name;
        this.email = email;
    }

    joinGroup(group: Group): void {
        if (!this.groups.includes(group)) {
            this.groups.push(group);
            group.addMember(this);
        }
    }

    leaveGroup(group: Group): void {
        this.groups = this.groups.filter(g => g.id !== group.id);
        group.removeMember(this);
    }

    createGroup(name: string): Group {
        const group = new Group(Group.generateId(), name, this);
        this.createdGroups.push(group);
        this.joinGroup(group);
        return group;
    }
}

class Group {
    id: number;
    name: string;
    createdBy: User;
    members: User[] = [];

    constructor(id: number, name: string, createdBy: User) {
        this.id = id;
        this.name = name;
        this.createdBy = createdBy;
    }

    static generateId(): number {
        return Math.floor(Math.random() * 1000000);
    }

    addMember(user: User): void {
        if (!this.members.includes(user)) {
            this.members.push(user);
        }
    }

    removeMember(user: User): void {
        this.members = this.members.filter(member => member.id !== user.id);
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
