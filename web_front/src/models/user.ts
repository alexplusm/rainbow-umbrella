export class User {
    id: number;
    friendshipId: number;
    login: string;
    firstName: string;
    lastName: string;
    city: string;
    gender: string;
    age: number;
    interests: string[];

    constructor(value: any) {
        this.id = value['id'];
        this.friendshipId = value['friendshipId'];
        this.login = value['login'];
        this.firstName = value['firstName'];
        this.lastName = value['lastName'];
        this.city = value['city'];
        this.gender = value['gender'];
        this.age = value['age'];
        this.interests = ["one", "two"];
    }
}

export interface IFriendList {
    friends: User[];
    requested: User[];
    waitingForResponse: User[];
}
