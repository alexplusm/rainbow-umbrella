export async function loginApi(login: string, password: string): Promise<string> {
    const data = {login, password};

    return await fetch("/api/v1/users/login", {method: "POST", body: JSON.stringify(data)})
        .then(response => response.json())
        .then(data => data["sessionID"])
}