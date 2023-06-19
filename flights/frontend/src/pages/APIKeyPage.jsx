import Typography from '@mui/material/Typography'
import Button from '@mui/material/Button'
import { Checkbox, TextField } from '@mui/material';
import { Stack } from '@mui/system';
import { useState } from 'react';
import { axiosInstance } from 'config/interceptor';

const APIKeyPage = () => {
    const [apiKey, setApiKey] = useState('')
    const [keyObject, setKeyObject] = useState()
    const [permanent, setPermanent] = useState(false)

    const changePermanent = () => {
        setPermanent((value) => !value)
    }

    const getAPIKey = () => {
        axiosInstance.get(`/api-key?temporary=${!permanent}`)
            .catch((err) => {
                console.error(err)
            })
            .then((response) => {
                setApiKey(response.data.data.key)
                setKeyObject(response.data.data)
                axiosInstance.post(`/api-key`, response.data.data)
                    .catch((e) => {
                        console.error(e)
                    })
                    .then((res) => {
                        console.log(res);
                    })
            })
    }

    return (
        <Stack width={'50%'} margin='8rem auto'>
            <Typography variant="h5" color="initial">
                This is a page where you can generate you own api key, that allows access on you behalf to other users!
            </Typography>
            <TextField value={apiKey} />
            <Stack direction={'row'}>
                <Typography variant="body1" color="initial" marginTop={'0.5rem'}>
                    Should the key be permanent?
                </Typography>
                <Checkbox label='Should the key be permanent?' onChange={changePermanent} />
            </Stack>
            <Button variant="contained" color="primary" onClick={getAPIKey}>
                Get API key
            </Button>
        </Stack>
    );
}

export default APIKeyPage;