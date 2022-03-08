export interface User {
    id: number;
    login: string;

    firstName: string;
    lastName: string;
    city: string;
    gender: string;
    age: number; // TODO
}

export class UserM {
    id: number;
    friendshipId: number;
    login: string;

    firstName: string;
    lastName: string;
    city: string;
    gender: string;
    age: number;

    constructor(value: any) {
        this.id = value['id'];
        this.friendshipId = value['friendshipId'];
        this.login = value['login'];
        this.firstName = value['firstName'];
        this.lastName = value['lastName'];
        this.city = value['city'];
        this.gender = value['gender'];
        this.age = value['age'];
    }
}

export interface IFriendList {
    friends: UserM[];
    requested: UserM[];
    waitingForResponse: UserM[];
}
