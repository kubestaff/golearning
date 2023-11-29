
type User = {
    id: string,
    name: string,
    age: number,
    jobTitle: string
}

export default function User({id, name, age, jobTitle}: User) {
    return <div> <li key={id}>{name} [{age}] - {jobTitle}</li> </div>
}