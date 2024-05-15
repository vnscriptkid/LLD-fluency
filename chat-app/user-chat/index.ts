enum ThreadType {
    ONE_TO_ONE,
    GROUP
}

export class User {
    id: number;
    name: string;
    email: string;
    groups: Group[] = [];
    createdGroups: Group[] = [];
    threads: Thread[] = [];

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
        group.thread = new Thread(Thread.generateId(), ThreadType.GROUP);
        group.thread.addParticipant(this);
        return group;
    }

    sendMessage(thread: Thread, content: string): Message {
        const message = new Message(Message.generateId(), this, thread, content);
        thread.addMessage(message);
        return message;
    }
}

class Group {
    id: number;
    name: string;
    createdBy: User;
    members: User[] = [];
    thread: Thread;

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
            this.thread.addParticipant(user);
        }
    }

    removeMember(user: User): void {
        this.members = this.members.filter(member => member.id !== user.id);
        this.thread.participants = this.thread.participants.filter(participant => participant.id !== user.id);
    }
}

class Thread {
    id: number;
    type: ThreadType;
    participants: User[] = [];
    messages: Message[] = [];

    constructor(id: number, type: ThreadType) {
        this.id = id;
        this.type = type;
    }

    static generateId(): number {
        return Math.floor(Math.random() * 1000000);
    }

    addMessage(message: Message): void {
        this.messages.push(message);
    }

    addParticipant(user: User): void {
        if (!this.participants.includes(user)) {
            this.participants.push(user);
        }
    }
}

class Message {
    id: number;
    sender: User;
    thread: Thread;
    content: string;
    timestamp: Date;

    constructor(id: number, sender: User, thread: Thread, content: string) {
        this.id = id;
        this.sender = sender;
        this.thread = thread;
        this.content = content;
        this.timestamp = new Date();
    }

    static generateId(): number {
        return Math.floor(Math.random() * 1000000);
    }
}

// Example usage:
const user1 = new User(1, 'Alice', 'alice@example.com');
const user2 = new User(2, 'Bob', 'bob@example.com');
const group1 = user1.createGroup('Group1');

user1.sendMessage(group1.thread, 'Hello Group!');
user2.joinGroup(group1);
user2.sendMessage(group1.thread, 'Hi Alice!');

console.log(user1);
console.log(user2);
console.log(group1);
console.log(group1.thread);
