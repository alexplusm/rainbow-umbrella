import {UserM} from "@/models/user";
import router from "@/router";


export async function loginApi(login: string, password: string): Promise<string> {
    const data = {login, password};

    return await fetch("/api/v1/users/login", {method: "POST", body: JSON.stringify(data)})
        .then(response => response.json())
        .then(data => data["sessionID"])
        .catch(data => data);
}

export async function retrieveUserAPI(login: string, headers: Headers): Promise<UserM> {
    return await fetch(`/api/v1/users/${login}`, {headers})

        .then(resp => {
            console.log("RESP:" , resp);
            if (resp.status === 401) {
                router.push({name: 'welcome'});
            }
            return resp;
        })
        .then(resp => resp.json())

        .then(data => new UserM(data))
        .then(user => {
            user.age = 666;
            console.log("user", user);
            return user;
        })
        ;
}
