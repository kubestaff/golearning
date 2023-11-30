import { useState, useEffect } from "react";

import User from "./User";

const backendUrl = "http://localhost:34567/user";


export default function Me(_userid: any) {
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

    useEffect(() => fetchUser(""), []);
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