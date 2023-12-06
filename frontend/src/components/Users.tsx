import React, { useState, useEffect } from 'react';

import Stack from "react-bootstrap/Stack";

import User from './User';

const backendUrl = "http://localhost:34567/users";

export default function Users() {
  const [users, setUsers] = useState<any[]>([])
  const [userIdWithDetails, setShowDetails] = useState(0)
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
    <Stack>
      {users.length > 0 && (
        <>
          {users.map(user => (
            <div> <User id={user.ID} name={user.Name} age={user.Age} jobTitle={user.JobTitle} showDetails={userIdWithDetails === user.ID} setShowDetails={setShowDetails}/> </div>
          ))}
         </>
      )}
    </Stack>
  );
}