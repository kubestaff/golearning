import UsersList from "../components/Users";
import Button from "react-bootstrap/Button";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import { useParams } from 'react-router-dom';
import Form from 'react-bootstrap/Form';
import Alert from 'react-bootstrap/Alert';

import React, { useState, useEffect } from 'react';
const backendUrl = "http://localhost:34567/users";
const uploadUrl = "http://localhost:34567/upload";

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

function updateUser(userId: string | undefined, data: any, setErrorText: any, setSuccessText: any) {
  let updateUserUrl = backendUrl
  if (userId) {
    updateUserUrl = `${updateUserUrl}?id=${userId}`
  }

  fetch(updateUserUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  }).then(response => {
    return response.json()
  }).then(data => {
    if (data.Error) {
      setErrorText(data.Error)
      return
    }
    if (data.Message) {
      setSuccessText(data.Message)
    }
  })
    .catch(error => {
      setErrorText(error.message)
    })
}

export function Change() {
  const [selectedFile, setSelectedFile] = useState<File>();
	const [isFilePicked, setIsFilePicked] = useState(false);

  let { userId } = useParams();
  const [successText, setSuccessText] = useState('')
  const [firstName, setFirstName] = useState('');
  const [age, setAge] = useState(0);
  const [jobTitle, setJobTitle] = useState('');
  const [characteristics, setCharacteristics] = useState<string[]>([]);
  const [likes, setLikes] = useState<string[]>([]);
  const [dislikes, setDislikes] = useState<string[]>([]);
  const [image, setImage] = useState('');
  const [backgroundCol, setBackground] = useState('');
  const [nameCol, setNamecol] = useState('');
  const [jobCol, setJobcol] = useState('');
  const [ageCol, setAgecol] = useState('');
  const [about, setAbout] = useState('');
  const [errorText, setErrorText] = useState('')
  const [imageUrl, setImageUrl] = useState('')

  const fetchUserData = function () {
    const fetchUserUrl = `${backendUrl}?id=${userId}`
    fetch(fetchUserUrl)
      .then(response => {
        return response.json()
      })
      .then(data => {
        if (data.Error) {
          setErrorText(data.Error)
          return
        }

        setFirstName(data.Name)
        setAge(data.Age)
        setJobTitle(data.JobTitle)
        setCharacteristics(data.Characteristics)
        setLikes(data.likes)
        setDislikes(data.Dislikes)
        setBackground(data.backgroundCol)
        setNamecol(data.nameCol)
        setJobcol(data.jobCol)
        setAgecol(data.ageCol)
        setAbout(data.about)
        setImage(data.Image)
        setImageUrl(data.ImageUrl)
      })
      .catch(error => {
        setErrorText(error.message)
      });
  }

  useEffect(() => {
    if (userId) {
      fetchUserData()
    }
  }, [])

  const handleSubmit = (e: React.SyntheticEvent) => {
    e.preventDefault()
    const data = {
      Age: age,
      Name: firstName,
      JobTitle: jobTitle,
      Characteristics: characteristics,
      Likes: likes,
      Dislikes: dislikes,
      Image: image,
      Backgroundcol: backgroundCol,
      Namecol: nameCol,
      Jobcol: jobCol,
      Agecol: ageCol,
      About: about,
    }

    if (isFilePicked) {
      const formData = new FormData();
      if (selectedFile) {
        formData.append("file", selectedFile, selectedFile.name)
        if (image) {
          formData.append("uuid", image)
        }
      }

      fetch(uploadUrl, {
        method: 'POST',
        body: formData
      }).then(response => {
        return response.json()
      }).then(responseData => {
        if (responseData.Error) {
          setErrorText(responseData.Error)
          return
        }
        data.Image = responseData.Uuid
        updateUser(userId, data, setErrorText, setSuccessText)
      })
        .catch(error => {
          setErrorText(error.message)
        })
    } else {
      updateUser(userId, data, setErrorText, setSuccessText)
    }
  };

  const changeHandler = (event: any) => {
		setSelectedFile(event.target.files[0]);
		setIsFilePicked(true);
	};

  const setCharacteristicsFromFlatValue = (characteristics: string) => {
    const characteristicsArray = characteristics.split(",")
    setCharacteristics(characteristicsArray)
  }
  const setLikesFromFlatValue = (likes: string) => {
    const likesArray = likes.split(",")
    setLikes(likesArray)
  }
  const setDislikesFromFlatValue = (dislikes: string) => {
    const dislikesArray = dislikes.split(",")
    setDislikes(dislikesArray)
  }

  return (
    <>
      {errorText !== '' && (<Alert variant={"danger"} dismissible>
        {errorText}
      </Alert>)}
      {successText !== '' && (<Alert variant={"success"} dismissible>
        {successText}
      </Alert>)}
      <h2>Change User </h2>
      <Form className="bg-gray p-3" onSubmit={handleSubmit}>
        <Form.Group className="mb-3" controlId="name">
          <Form.Label>Name</Form.Label>
          <Form.Control type="text" value={firstName} required onChange={e => setFirstName(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="age">
          <Form.Label>Age</Form.Label>
          <Form.Control type="number" value={age} min={0} max={100} required onChange={e => setAge(Number(e.target.value))} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="job">
          <Form.Label>Job Title</Form.Label>
          <Form.Control type="text" required value={jobTitle} onChange={e => setJobTitle(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="characteristics">
          <Form.Label >Characteristics</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" value={characteristics} onChange={e => setCharacteristicsFromFlatValue(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="likes">
          <Form.Label >Likes</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" value={likes} onChange={e => setLikesFromFlatValue(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="dislikes">
          <Form.Label >Dislikes</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" value={dislikes} onChange={e => setDislikesFromFlatValue(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="img">
          <Form.Label>Image</Form.Label>
          <Form.Control type="file" onChange={changeHandler} />
          {isFilePicked && selectedFile && (
             <img src={URL.createObjectURL(selectedFile)} width={100} className={"pt-2"}/>
         )}
         {!isFilePicked && imageUrl !== '' && (
             <img src={imageUrl} width={100} className={"pt-2"}/>
         )}
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
          <Form.Control size="lg" type="text" as="textarea" onChange={e => setAbout(e.target.value)} />
        </Form.Group>
        <Button variant="primary" type="submit">
          Save
        </Button>
      </Form>
    </>
  );
};


export function Create() {
  const [selectedFile, setSelectedFile] = useState<File>();
	const [isFilePicked, setIsFilePicked] = useState(false);

  const [errorText, setErrorText] = useState('')
  const [successText, setSuccessText] = useState('')
  const [firstName, setFirstName] = useState('');
  const [age, setAge] = useState(0);
  const [jobTitle, setJobTitle] = useState('');
  const [characteristics, setCharacteristics] = useState<string[]>([]);
  const [likes, setLikes] = useState<string[]>([]);;
  const [dislikes, setDislikes] = useState<string[]>([]);;
  const [image, setImage] = useState('');
  const [backgroundCol, setBackground] = useState('');
  const [nameCol, setNamecol] = useState('');
  const [jobCol, setJobcol] = useState('');
  const [ageCol, setAgecol] = useState('');
  const [about, setAbout] = useState('');


  const handleSubmit = (e: React.SyntheticEvent) => {
    e.preventDefault()
    const data = {
      Age: age,
      Name: firstName,
      JobTitle: jobTitle,
      Characteristics: characteristics,
      Likes: likes,
      Dislikes: dislikes,
      Image: image,
      Backgroundcol: backgroundCol,
      Namecol: nameCol,
      Jobcol: jobCol,
      Agecol: ageCol,
      About: about,
    }

    fetch(backendUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    }).then(response => {
      return response.json()
    }).then(data => {
      if (data.Error) {
        setErrorText(data.Error)
        return
      }
      if (data.Message && data.Message) {
        setSuccessText(data.Message)
      }
    })
      .catch(error => {
        setErrorText(error.message)
      });
  };

  const setCharacteristicsFromFlatValue = (characteristics: string) => {
    const characteristicsArray = characteristics.split(",")
    setCharacteristics(characteristicsArray)
  }
  const setLikesFromFlatValue = (likes: string) => {
    const likesArray = likes.split(",")
    setLikes(likesArray)
  }
  const setDislikesFromFlatValue = (dislikes: string) => {
    const dislikesArray = dislikes.split(",")
    setDislikes(dislikesArray)
  }

  const changeHandler = (event: any) => {
		setSelectedFile(event.target.files[0]);
		setIsFilePicked(true);
	};

  return (
    <>
      {errorText !== '' && (<Alert variant={"danger"} dismissible>
        {errorText}
      </Alert>)}

      {successText !== '' && (<Alert variant={"success"} dismissible>
        {successText}
      </Alert>)}
      <h2>Create User</h2>
      <Form className="bg-gray p-3" onSubmit={handleSubmit}>
        <Form.Group className="mb-3" controlId="name">
          <Form.Label >Name</Form.Label>
          <Form.Control type="text" required onChange={e => setFirstName(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="age">
          <Form.Label >Age</Form.Label>
          <Form.Control type="number" min={0} max={100} required onChange={e => setAge(Number(e.target.value))} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="job">
          <Form.Label >Job Title</Form.Label>
          <Form.Control type="text" required onChange={e => setJobTitle(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="characteristics">
          <Form.Label >Characteristics</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" onChange={e => setCharacteristicsFromFlatValue(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="likes">
          <Form.Label >Likes</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" onChange={e => setLikesFromFlatValue(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="dislikes">
          <Form.Label >Dislikes</Form.Label>
          <Form.Control size="sm" type="text" as="textarea" onChange={e => setDislikesFromFlatValue(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="img">
          <Form.Label>Image</Form.Label>
          <Form.Control type="file" onChange={changeHandler} />
          {isFilePicked && selectedFile && (
              <div>
                <p>Filename: {selectedFile.name}</p>
                <p>Filetype: {selectedFile.type}</p>
                <p>Size in bytes: {selectedFile.size}</p>
              </div>
         )}
        </Form.Group>
        <Form.Group className="mb-3" controlId="background colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Background color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#FFFFFF" placeholder="Choose your color" onChange={e => setBackground(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="name font colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Name font color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#000000" placeholder="Choose your color" onChange={e => setNamecol(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="job font colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Job font color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#000000" placeholder="Choose your color" onChange={e => setJobcol(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="age font colour">
          <Form.Label htmlFor="exampleColorInput">Choose your Age font color</Form.Label>
          <Form.Control type="color" id="exampleColorInput" defaultValue="#000000" placeholder="Choose your color" onChange={e => setAgecol(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" controlId="about">
          <Form.Label >Tell us about yourself</Form.Label>
          <Form.Control size="lg" type="text" as="textarea" required onChange={e => setAbout(e.target.value)} />
        </Form.Group>
        <Button variant="primary" type="submit">
          Save
        </Button>
      </Form>
    </>
  );
}