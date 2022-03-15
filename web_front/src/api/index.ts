import router from "@/router";
import type { IFriendList } from "@/models/user";
import { User } from "@/models/user";

export interface IApiResponse<T=null> {
    hasError: boolean;
    notifyMessage: string;
    data: T;
}

function buildApiResponse<T=null>(): IApiResponse<T> {
    return {
        hasError: false,
        notifyMessage: "",
        data: {}
    } as IApiResponse<T>;
}

async function registerUser(formData: FormData): Promise<IApiResponse> {
    return fetch("/api/v1/users/register", {
        method: "POST",
        body: formData
    }).then(async resp => {
        const response = buildApiResponse();

        if (resp.status >= 299) {
            response.notifyMessage = await resp.text();
            response.hasError = true;
        }

        return response;
    })
}

async function login(login: string, password: string): Promise<IApiResponse<string>> {
    const data = {login, password};

    return await fetch("/api/v1/users/login", {method: "POST", body: JSON.stringify(data)})
        .then(async resp => {
            const response = buildApiResponse<string>();
            response.hasError = resp.status > 399;

            if (resp.status === 404) {
                response.notifyMessage = "User not found"
                return response;
            }

            const data = await resp.json()

            response.data = data["sessionID"];

            return response;
        });
}

async function retrieveUser(login: string, headers: Headers): Promise<User> {
    return await fetch(`/api/v1/users/${login}`, {headers})
        .then(resp => {
            console.log("RESP:" , resp);
            if (resp.status === 401) {
                router.push({name: 'welcome'});
            }
            return resp;
        })
        .then(resp => resp.json())

        .then(data => new User(data))
        .then(user => {
            user.age = 666;
            console.log("user", user);
            return user;
        });
}

async function retrieveFriendList(login: string, headers: Headers): Promise<IFriendList> {
    return fetch(`/api/v1/friendship/${login}`, {headers})
        .then(data => data.json())
        .then(value => {
            return {
                friends: value["friends"] || [],
                requested: value["requested"] || [],
                waitingForResponse: value["waitingForResponse"] || [],
            } as IFriendList;
        });
}

async function userList(): Promise<User[]> {
    return fetch("/api/v1/users")
        .then(data => data.json())
        .then(body => body['users'])
        .then(users => {
            console.log("Users", users);

            // TODO: wtf?
            return (users as []).map((user) => new User(user))
        });
}

async function createFriendship(requestingUserId: number, targetingUserId: number): Promise<IApiResponse> {
    const body = JSON.stringify({
        requestingUserId,
        targetingUserId
    });

    return fetch("/api/v1/friendships", {method: "POST",  body})
        .then(resp => {
            const response = buildApiResponse();

            return response;
        });
}

export const api = {
    registerUser,
    login,
    retrieveUser,
    retrieveFriendList,
    userList,

    createFriendship,
}
