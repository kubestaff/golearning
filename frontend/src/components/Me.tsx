import React, { useState, useEffect } from "react";

import User from "./User";
import { response } from "express";

const backendUrl = "http://localhost:34567/users";


export default function Me() {
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

    useEffect(() => fetchUser("5"), []);
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