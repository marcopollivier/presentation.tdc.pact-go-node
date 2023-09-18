const express = require('express');
const axios = require('axios');
const bodyParser = require('body-parser');

const app = express();
const PORT = 3000;
const GO_SERVER_URL = 'http://localhost:8080/user';

app.use(bodyParser.json());

app.post('/execute', async (req, res) => {
    try {
        const response = await axios.post(GO_SERVER_URL, req.body);
        res.status(response.status).send(response.data);
    } catch (error) {
        res.status(error.response ? error.response.status : 500).send(error.message);
    }
});

app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
