const path = require('path');
const axios = require('axios');
const { PactV3, MatchersV3 } = require('@pact-foundation/pact');
const { Publisher } = require('@pact-foundation/pact-node');

const mock_port = 1234;
const mock_server_url = 'http://127.0.0.1:' + mock_port;

const PACT_BROKER_URL = 'http://127.0.0.1:80';
const PACT_DIR = path.resolve(process.cwd(), 'pacts');

const provider = new PactV3({
    consumer: 'NodeService',
    provider: 'GoService',
    port: mock_port,
    dir: PACT_DIR,
    pactBrokerUrl: PACT_BROKER_URL
});

describe('Pact with GoService', () => {
    it('handles POST /execute request', () => {
        provider
            .uponReceiving("a POST request for user")
            .withRequest({
                method: 'POST',
                path: '/user',
                body: {
                    email: 'example@email.com',
                    cpf: '12345678901'
                }
            })
            .willRespondWith({
                status: 200,
                body: MatchersV3.like({
                    message: 'User created successfully!'
                })
            });

        return provider.executeTest(() => {
            return axios.post(mock_server_url + '/user', {
                email: 'example@email.com',
                cpf: '12345678901'
            }).then(response => {
                expect(response.data.message).toBe('User created successfully!');
            });
        });
    });

});