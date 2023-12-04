import UsersList from "../components/Users";
import Button from "react-bootstrap/Button";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import { useParams } from 'react-router-dom';
import Form from 'react-bootstrap/Form';
import Alert from 'react-bootstrap/Alert';

import React, { useState, useEffect } from 'react';
const backendUrl = "http://localhost:34567/users?id=";
const changeUserUrl = "http://127.0.0.1:34567/user-change?id=";

export default function Users() {
  return (
    <div>
      <Row>
        <h2>Users List</h2>
      </Row>
      <UsersList />
      <Row>
        <Col>
          <Button variant="primary" href="/users-create">
            Create
          </Button>
        </Col>
      </Row>
    </div>
  );
}

export function Change() {
  let { userId } = useParams();
  const [firstName, setFirstName] = useState('');
  const [age, setAge] = useState(0);
  const [jobTitle, setJobTitle] = useState('');

  const fetchUserData = function() {
    fetch(backendUrl + userId)
      .then(response => {
        return response.json()
      })
      .then(data => {
        setFirstName(data.Name)
        setAge(data.Age)
        setJobTitle(data.JobTitle)
      })
      .catch(error => console.error(error));
  }

  useEffect(() => {
    fetchUserData()
  }, [])

  const handleSubmit = (e: React.SyntheticEvent) => {
    e.preventDefault()
    const data = {
      Age: age,
      Name: firstName,
      JobTitle: jobTitle,
    }
  
    fetch(changeUserUrl + userId, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    }).then(response => {
      console.log(response.json())
    })
  };

  return (
    <>
     <Alert variant={"danger"} dismissible>
         Some error happened
     </Alert>
     <Alert variant={"success"} dismissible>
         User successfully saved
     </Alert>
    <h2>Change User </h2>
    <Form className="bg-gray p-3" onSubmit={handleSubmit}>
      <Form.Group className="mb-3" controlId="name">
        <Form.Control type="text" placeholder="Name" value={firstName} required onChange={e => setFirstName(e.target.value)} />
      </Form.Group>
      <Form.Group className="mb-3" controlId="age">
        <Form.Control type="number" placeholder="Age" value={age} min={0} max={100} required onChange={e => setAge(Number(e.target.value))}/>
      </Form.Group>
      <Form.Group className="mb-3" controlId="job">
        <Form.Control type="text" placeholder="Job Title" required value={jobTitle} onChange={e => setJobTitle(e.target.value)} />
      </Form.Group>
      <Button variant="primary" type="submit">
        Save
      </Button>
    </Form>
    </>
  );
}

export function Create() {
  const [firstName, setFirstName] = useState('');
  const [age, setAge] = useState(0);
  const [jobTitle, setJobTitle] = useState('');

  const handleSubmit = (e: React.SyntheticEvent) => {
    e.preventDefault()
    const data = {
      Age: age,
      Name: firstName,
      JobTitle: jobTitle,
    }
  
    fetch(changeUserUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    }).then(response => {
      console.log(response.json())
    })
  };

  return (
    <>
    
    <h2>Create user</h2>
    <Form className="bg-gray p-3" onSubmit={handleSubmit}>
      <Form.Group className="mb-3" controlId="name">
        <Form.Control type="text" placeholder="Name" required onChange={e => setFirstName(e.target.value)} />
      </Form.Group>
      <Form.Group className="mb-3" controlId="age">
        <Form.Control type="number" placeholder="Age" min={0} max={100} required onChange={e => setAge(Number(e.target.value))}/>
      </Form.Group>
      <Form.Group className="mb-3" controlId="job">
        <Form.Control type="text" placeholder="Job Title" required onChange={e => setJobTitle(e.target.value)} />
      </Form.Group>
      <Button variant="primary" type="submit">
        Save
      </Button>
    </Form>
    </>
  );
}