import { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button';


type User = {
    id: string,
    name: string,
    age: number,
    jobTitle: string
}

const backendUrl = "http://localhost:34567/user";

export default function User({ id, name, age, jobTitle }: User) {
    return (<div key={id}>{name} [{age}] - {jobTitle}
        <Button href={`/users-change/${id}`} size="sm" className="ms-1 mb-1" variant="outline-info">Edit</Button>
    </div>
    )
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
