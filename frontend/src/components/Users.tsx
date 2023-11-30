import React, { useState, useEffect } from 'react';

import User from './User';

const backendUrl = "http://localhost:34567/users";

export default function Users() {
  const [users, setUsers] = useState<any[]>([])
  const fetchUserData = function() {
    fetch(backendUrl)
      .then(response => {
        return response.json()
      })
      .then(data => {
        setUsers(data)
      })
      .catch(error => console.error(error));
  }

  useEffect(() => {
    fetchUserData()
  }, [])

  return (
    <div>
      {users.length > 0 && (
        <ul>
          {users.map(user => (
            <li> <User id={user.ID} name={user.Name} age={user.Age} jobTitle={user.JobTitle}/> </li>
          ))}
        </ul>
      )}
    </div>
  );
}