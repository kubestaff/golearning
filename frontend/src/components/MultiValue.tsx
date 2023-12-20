import { useState } from 'react'
import Button from 'react-bootstrap/Button'
import Card from 'react-bootstrap/Card';

export default function TagsInput(){
    const [multipleValues, setMultipleValues] = useState<string[]>([])

    const setMultiValues = (e: React.KeyboardEvent<HTMLInputElement>) => {
        // If user did not press enter key, return
        if(e.key !== 'Enter') return
        // Get the value of the input
        const value = e.currentTarget.value
        // If the value is empty, return
        if(!value.trim()) return
        // Add the value to the tags array
        setMultipleValues([...multipleValues, value])
        // Clear the input
        e.currentTarget.value = ''
    }

    const removeElement = (e: React.MouseEvent<HTMLButtonElement>) => {
        const filteredArray: string[] = multipleValues.filter((value, index, self) => {
            return value !== e.currentTarget.value;
        })
        setMultipleValues([...filteredArray])
    }

    const uniqueArray: string[] = multipleValues.filter((value, index, self) => {
        return self.indexOf(value) === index;
    })
      
    return (
        <>
        <label className="form-label">Tell us about yourself</label>
        <Card bg="Light" className='p-2'>
            <div className='p-1 d-flex align-content-start flex-wrap'>
        { uniqueArray.map((tag, index) => (
            <Button key={index} value={tag} className='me-1 mt-1 mb-1' variant='secondary' onClick={removeElement}>
                {tag} &times;
            </Button>
        )) }
        <input type="text" onKeyDown={setMultiValues} className='p-1' />
        </div>
        </Card>
        </>
    )
}