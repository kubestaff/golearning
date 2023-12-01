import { useState, useEffect } from "react";

const backendUrl = "http://localhost:34567/user";

type User = {
    id: string,
    name: string,
    age: number,
    jobTitle: string
}

export default function User({ id, name, age, jobTitle }: User) {
    return <li key={id}>{name} [{age}] - {jobTitle} <a href="/users-change/{id}">Edit</a></li>
}

export function Me(userid: string) {
    const [user, setUser] = useState<any>();

    const fetchUser = function (userid: string) {
        const apiUrl = `${backendUrl}?id=${userid}`;

        fetch(apiUrl)
            .then((response) => {
                return response.json();
            })
            .then((data) => {
                setUser(data);

            })
            .catch((error) => console.error(error));
    };

    useEffect(() => fetchUser(userid), []);
    return (
        <div>
            <User
                id={user.ID}
                name={user.Name}
                age={user.Age}
                jobTitle={user.jobTitle}
            />
        </div>
    );

}
