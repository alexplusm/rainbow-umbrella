import {UserM} from "@/models/user";

export async function loginApi(login: string, password: string): Promise<string> {
    const data = {login, password};

    return await fetch("/api/v1/users/login", {method: "POST", body: JSON.stringify(data)})
        .then(response => response.json())
        .then(data => data["sessionID"])
}

export async function retrieveUserAPI(login: string, headers: Headers): Promise<UserM> {
    return await fetch(`/api/v1/users/${login}`, {headers})
        .then(resp => resp.json())
        .then(data => new UserM(data))
        .then(user => {
            user.age = 666;
            console.log("user", user);
            return user;
        });
}
