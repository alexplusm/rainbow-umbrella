export class User {
    id: number;
    avatarUrl: string;
    friendshipId: number;
    friendshipStatus: FriendshipStatus;
    login: string;
    firstName: string;
    lastName: string;
    city: string;
    gender: string;
    age: number;
    interests: string[];

    constructor(value: any) {
        this.id = value['id'];
        this.avatarUrl = "https://cdn.quasar.dev/img/boy-avatar.png";
        this.friendshipId = value['friendshipId'];
        this.friendshipStatus = FriendshipStatus.Myself;

        this.login = value['login'];
        this.firstName = value['firstName'];
        this.lastName = value['lastName'];
        this.city = value['city'];
        this.gender = value['gender'];
        this.age = value['age'];
        this.interests = ["one", "two", "azazazazazaz", "lololololololol", "wtf"];
    }
}

export interface IFriendList {
    friends: User[];
    requested: User[];
    waitingForResponse: User[];
}

export enum FriendshipStatus {
    Myself = "myself",
    NotFriend = "notFriend",
    Friend = "friend",
    Wait = "wait"
}
