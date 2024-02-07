import React, { useState, useEffect } from 'react';
import './App.css';
import ListGroup from 'react-bootstrap/ListGroup';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

const backendUrl = "http://127.0.0.1:345567";

function Users() {
  const [users, setUsers] = useState<any[]>([])

  const fetchUserData = function () {
    fetch(backendUrl)
      .then(response => {
        return response.json()
      })
      .then(usersInJSONFormat => {
        setUsers(usersInJSONFormat)
      })
      .catch(error => console.error(error));
  }

  useEffect(() => {
    fetchUserData();
  }, [])

  return (
    <ListGroup>
      {users.length > 0 && (
        <>
          {users.map(user => (
            <ListGroup.Item action href={"/user?id=" + user.ID}>
              {user.Name}
              </ListGroup.Item>
          ))}
        </>
      )}
    
    </ListGroup>
  );
}

function App() {
  return (
    <div className="App">
      <Container>
        <Row>
          <Col>
            <Users />
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default App;











