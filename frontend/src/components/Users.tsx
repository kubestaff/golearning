import React, { useState, useEffect } from 'react';
import Alert from 'react-bootstrap/Alert';

import Stack from "react-bootstrap/Stack";

import User from './User';

import { BackendUrl } from './Url';

const backendUrl = BackendUrl + "/users";

export default function Users() {
  const [errorText, setErrorText] = useState('')
  const [users, setUsers] = useState<any[]>([])
  //todo find out how to hide details of current user
  const [userIdWithDetails, setShowDetails] = useState(0)

  useEffect(() => {
    fetch(backendUrl)
      .then(response => {
        return response.json()
      })
      .then(data => {
        if (data.Error) {
          setErrorText(data.Error)
          return;
        }
        setUsers(data)
      })
      .catch(error => {
        setErrorText(error.message)
      });
  }, [])

  return (
    <>
    {errorText !== '' && (<Alert variant={"danger"} dismissible>
        {errorText}
      </Alert>)}
    <Stack>
      {users.length > 0 && (
        <>
          {users.map(user => (
            <div key= {user.ID}>
              <User id={user.ID} name={user.Name} age={user.Age} jobTitle={user.JobTitle} showDetails={userIdWithDetails === user.ID} setShowDetails={setShowDetails}/> 
            </div>
          ))}
         </>
      )}
    </Stack>
    </>
  );
}