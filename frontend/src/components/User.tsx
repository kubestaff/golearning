import { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';


type User = {
    id: string,
    name: string,
    age: number,
    jobTitle: string
}

const backendUrl = "http://localhost:34567/users?id=";

export default function User({ id, name, age, jobTitle }: User) {
    const [showDetails, setShowDetails] = useState(false);

    const fetchUserData = () => {
        fetch(backendUrl + id)
            .then((response) => {
                return response.json();
            })
            .then((data) => {

            })
            .catch((error) => console.error(error));
    };

    useEffect(() => {
        fetchUserData();
    }, [id]);

    const viewUser = () => {
        setShowDetails(!showDetails);
    };

    return (
    <div key={id}>
      {name}
      <Button
        onClick={() => viewUser()}
        size="sm"
        className="ms-1 mb-1"
        variant="outline-info"
      >
        {showDetails ? "Hide" : "View"}
      </Button>
      <Button
        href={`/users-change/${id}`}
        size="sm"
        className="ms-1 mb-1"
        variant="outline-info"
      >
        Edit
      </Button>

      {showDetails && (
        <Card className="position-absolute top-0 end-0 mt-2 me-5">
          <Card.Body>
            <Card.Title>{name}</Card.Title>
            <Card.Text>
              <strong>Age:</strong> {age} <br />
              <strong>Job Title:</strong> {jobTitle}
            </Card.Text>
          </Card.Body>
        </Card>
      )}
    </div>
  );
}