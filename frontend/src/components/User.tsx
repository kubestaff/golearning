
type User = {
    id: string,
    name: string,
    age: number,
    jobTitle: string
}

export default function User({id, name, age, jobTitle}: User) {
    return <li key={id}>{name} [{age}] - {jobTitle} <a href="/users-change/{id}">Edit</a></li>
}