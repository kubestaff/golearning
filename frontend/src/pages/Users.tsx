import UsersList from "../components/Users";
import Button from "react-bootstrap/Button";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import { useParams } from 'react-router-dom';
import Form from 'react-bootstrap/Form';

import React, { useState, useEffect } from 'react';
const backendUrl = "http://localhost:34567/users?id=";

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
  const [characteristics, setCharacteristics] = useState('');
  const [likes, setLikes] = useState('');
  const [dislikes, setDislikes] = useState('');
  const [image, setImage] = useState('');
  const [backgroundCol, setBackground] = useState('');
  const [nameCol, setNamecol] = useState('');
  const [jobCol, setJobcol] = useState('');
  const [ageCol, setAgecol] = useState('');
  const [about, setAbout] = useState('');

  const fetchUserData = function () {
    fetch(backendUrl + userId)
      .then(response => {
        return response.json()
      })
      .then(data => {
        setFirstName(data.Name)
        setAge(data.Age)
        setJobTitle(data.JobTitle)
        setCharacteristics(data.Characteristics)
        setLikes(data.likes)
        setDislikes(data.Dislikes)
        setImage(data.Image)
        setBackground(data.backgroundCol)
        setNamecol(data.nameCol)
        setJobcol(data.jobCol)
        setAgecol(data.ageCol)
        setAbout(data.about)
      })
      .catch(error => console.error(error));
  }

  useEffect(() => {
    fetchUserData()
  }, [])

  const handleSubmit = (e: React.SyntheticEvent) => {
    e.preventDefault();
  };

  return (
    <>
      <h2>Change User </h2>
      <Form className="bg-gray p-3" onSubmit={handleSubmit}>
        <Form.Group className="mb-3" controlId="name">
          <Form.Label >Name</Form.Label>
          <Form.Control type="text" value={firstName} required onChange={e => setFirstName(e.target.value)} />
          <Form.Label >Age</Form.Label>
        </Form.Group>
        <Form.Group className="mb-3" controlId="age">
          <Form.Control type="number" value={age} min={0} max={100} required onChange={e => setAge(Number(e.target.value))} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="job">
          <Form.Label >Job Title</Form.Label>
          <Form.Control type="text" required value={jobTitle} onChange={e => setJobTitle(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="characteristics">
          <Form.Label >Characteristics</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" value={characteristics} onChange={e => setCharacteristics(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="likes">
          <Form.Label >Likes</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" value={likes} onChange={e => setLikes(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="dislikes">
          <Form.Label >Dislikes</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" value={dislikes} onChange={e => setDislikes(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="img">
          <Form.Label>Image</Form.Label>
          <Form.Control type="file" value={image} onChange={e => setImage(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="background colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Background color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#FFFFFF" placeholder="Choose your color" value={backgroundCol} onChange={e => setBackground(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="name font colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Name font color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#000000" placeholder="Choose your color" value={nameCol} onChange={e => setNamecol(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="job font colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Job font color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#000000" placeholder="Choose your color" value={jobCol} onChange={e => setJobcol(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="age font colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Age font color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#000000" placeholder="Choose your color" value={ageCol} onChange={e => setAgecol(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="about">
          <Form.Label >Tell us about yourself</Form.Label>
          <Form.Control size="lg" type="text" as="textarea" value={about} onChange={e => setAbout(e.target.value)} />
        </Form.Group>
        <Button variant="primary" type="submit">
          Save
        </Button>
      </Form>
    </>
  );
}

export function Create() {
  const handleClick = () => {
    alert("Button is clicked")
  };

  return (
    <>
      <h2>Create User</h2>
      <Form className="bg-gray p-3" onSubmit={handleClick}>
        <Form.Group className="mb-3" controlId="name">
          <Form.Label >Name</Form.Label>
          <Form.Control type="text" required />
        </Form.Group>
        <Form.Group className="mb-3" controlId="age">
          <Form.Label >Age</Form.Label>
          <Form.Control type="number" min={0} max={100} required />
        </Form.Group>
        <Form.Group className="mb-3" controlId="job">
          <Form.Label >Job Title</Form.Label>
          <Form.Control type="text" required />
        </Form.Group>
        <Form.Group className="mb-3" controlId="characteristics">
          <Form.Label >Characteristics</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" />
        </Form.Group>
        <Form.Group className="mb-3" controlId="likes">
          <Form.Label >Likes</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" />
        </Form.Group>
        <Form.Group className="mb-3" controlId="dislikes">
          <Form.Label >Dislikes</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" />
        </Form.Group>
        <Form.Group className="mb-3" controlId="img">
          <Form.Label>Image</Form.Label>
          <Form.Control type="file" />
        </Form.Group>
        <Form.Group className="mb-3" controlId="background colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Background color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#FFFFFF" placeholder="Choose your color" required />
        </Form.Group>
        <Form.Group className="mb-3" controlId="name font colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Name font color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#000000" placeholder="Choose your color" required />
        </Form.Group>
        <Form.Group className="mb-3" controlId="job font colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Job font color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#000000" placeholder="Choose your color" required />
        </Form.Group>
        <Form.Group className="mb-3" controlId="age font colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Age font color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#000000" placeholder="Choose your color" required />
        </Form.Group>
        <Form.Group className="mb-3" controlId="about">
          <Form.Label >Tell us about yourself</Form.Label>
          <Form.Control size="lg" type="text" as="textarea" required />
        </Form.Group>
        <Button variant="primary" type="submit">
          Save
        </Button>
      </Form>
    </>
  );
}