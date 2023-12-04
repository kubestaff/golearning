import Button from 'react-bootstrap/Button';

type User = {
    id: string,
    name: string,
    age: number,
    jobTitle: string
}

export default function User({id, name, age, jobTitle}: User) {
    return (<div key={id}>{name} [{age}] - {jobTitle} 
        <Button href={`/users-change/${id}`} size="sm" className="ms-1 mb-1" variant="outline-info">Edit</Button>
    </div>
    )
}