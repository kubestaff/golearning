import React, { useState, useEffect } from "react";

import User from "./User";
import { response } from "express";

const backendUrl = "http://localhost:34567/users";

interface AprResponse {
    ID: string;
    Name: string;
    Age: number;
    JobTitle: string;
}

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

    useEffect(() => fetchUser("User"), []);
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