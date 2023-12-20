import UsersList from "../components/Users";
import Button from "react-bootstrap/Button";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import { useParams } from 'react-router-dom';
import Form from 'react-bootstrap/Form';
import Alert from 'react-bootstrap/Alert';
import { BackendUrl } from "../components/Url";
import TagsInput from "../components/MultiValue";

import React, { useState, useEffect } from 'react';
const backendUrl = BackendUrl + "/users";
const uploadUrl = BackendUrl + "/upload";

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
  const [backgroundColor, setBackgroundColor] = useState('#FFFFFF');
  const [nameColor, setNameColor] = useState('#000000');
  const [jobColor, setJobColor] = useState('#000000');
  const [ageColor, setAgeColor] = useState('#000000');
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
        setLikes(data.Likes)
        setDislikes(data.Dislikes)
        if (data.BackgroundColor !== "") {
          setBackgroundColor(data.BackgroundColor)
        }
        
        if (data.NameFontColor !== "") {
          setNameColor(data.NameFontColor)
        }

        if (data.JobFontColor !== "") {
          setJobColor(data.JobFontColor)
        }
        if (data.AgeFontColor !== "") {
          setAgeColor(data.AgeFontColor)
        }
        
        setAbout(data.About)
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
      backgroundColor: backgroundColor,
      NameFontColor: nameColor,
      JobFontColor: jobColor,
      AgeFontColor: ageColor,
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
        <Form.Group className="mb-3" id="name">
          <Form.Label>Name</Form.Label>
          <Form.Control type="text" value={firstName} required onChange={e => setFirstName(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="age">
          <Form.Label>Age</Form.Label>
          <Form.Control type="number" value={age} min={0} max={100} required onChange={e => setAge(Number(e.target.value))} />
        </Form.Group>
        <Form.Group className="mb-3" id="job">
          <Form.Label>Job Title</Form.Label>
          <Form.Control type="text" required value={jobTitle} onChange={e => setJobTitle(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="characteristics">
          <Form.Label >Characteristics</Form.Label>
          <Form.Control type="text" as="textarea" value={characteristics} onChange={e => setCharacteristicsFromFlatValue(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="likes">
          <Form.Label >Likes</Form.Label>
          <Form.Control type="text" as="textarea" value={likes} onChange={e => setLikesFromFlatValue(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="dislikes">
          <Form.Label >Dislikes</Form.Label>
          <Form.Control type="text" as="textarea" value={dislikes} onChange={e => setDislikesFromFlatValue(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="img">
          <Form.Label>Image</Form.Label>
          <Form.Control type="file" onChange={changeHandler} />
          {isFilePicked && selectedFile && (
             <img src={URL.createObjectURL(selectedFile)} width={100} className={"pt-2"}/>
         )}
         {!isFilePicked && imageUrl !== '' && (
             <img src={imageUrl} width={100} className={"pt-2"}/>
         )}
        </Form.Group>
        <Form.Group className="mb-3" id="background colour">
          <Form.Label htmlFor="backgroundColorInput">Choose your Background color</Form.Label>
          <Form.Control type="color" id="backgroundColorInput" placeholder="Choose your color" value={backgroundColor} onChange={e => setBackgroundColor(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="name font colour">
          <Form.Label htmlFor="nameColorInput">Choose your Name font color</Form.Label>
          <Form.Control type="color" id="nameColorInput" placeholder="Choose your color" value={nameColor} onChange={e => setNameColor(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="job font colour">
          <Form.Label htmlFor="jobColorInput">Choose your Job font color</Form.Label>
          <Form.Control type="color" id="jobColorInput" placeholder="Choose your color" value={jobColor} onChange={e => setJobColor(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="age font colour">
          <Form.Label htmlFor="ageColorInput">Choose your Age font color</Form.Label>
          <Form.Control type="color" id="ageColorInput" placeholder="Choose your color" value={ageColor} onChange={e => setAgeColor(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="about">
          <Form.Label >Tell us about yourself</Form.Label>
          <Form.Control size="lg" type="text" as="textarea" value={about} onChange={e => setAbout(e.target.value)} />
        </Form.Group>
        <Form.Group className="mb-3" id="name">
        <TagsInput/>
        </Form.Group>
        <Button variant="primary" type="submit">
          Save
        </Button>
      </Form>
    </>
  );
};