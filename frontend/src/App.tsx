import React, { useState, useEffect } from "react";
import "./App.css";

const backendUrl = "http://localhost:34567";

function Users() {
  const [users, setUsers] = useState<any[]>([]);

  const fetchUserData = () => {
    fetch(backendUrl)
      .then((response) => {
        return response.json();
      })
      .then((data) => {
        setUsers(data);
      })
      .catch((error) => console.error(error));
  };

  useEffect(() => {
    fetchUserData();
  }, []);

  return (
    <div>
      {users.length > 0 && (
        <ul>
          {users.map((user) => (
            <li key={user.ID}>{user.Name}</li>
          ))}
        </ul>
      )}
    </div>
  );
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Users />
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
