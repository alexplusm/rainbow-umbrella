import type {IFriendList} from "@/models/user";
import {User} from "@/models/user";
import router from "@/router";

export interface IApiResponse<T=null> {
    hasError: boolean;
    notifyMessage: string | null;
    data: T | null;
}

function buildApiResponse<T=null>(): IApiResponse<T> {
    return {
        hasError: false,
        notifyMessage: null,
        data: null
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

export async function loginApi(login: string, password: string): Promise<string> {
    const data = {login, password};

    return await fetch("/api/v1/users/login", {method: "POST", body: JSON.stringify(data)})
        .then(response => response.json())
        .then(data => data["sessionID"])
        .catch(data => data);
}

export async function retrieveUserAPI(login: string, headers: Headers): Promise<User> {
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

export async function retrieveFriendListApi(login: string, headers: Headers): Promise<IFriendList> {
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

export async function userListApi() :Promise<User[]> {
    return fetch("/api/v1/users")
        .then(data => data.json())
        .then(body => body['users'])
        .then(users => {
            console.log("Users", users);

            // TODO: wtf?
            return (users as []).map((user) => new User(user))
        });
}


export const api = {
    registerUser
}
