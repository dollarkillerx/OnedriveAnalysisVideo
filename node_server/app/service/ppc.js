const request = require('request');

request('https://www.baidu.com', function (error, response, body) {
    console.error('error:', error);
    console.log('statusCode:', response && response.statusCode);
    console.log('body:', body);
});