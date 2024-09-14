const express = require('express');
const path = require('path');

const app = express();
const port = 3000;

app.use(express.static('.'));

//app.get('/', (req, res) => {
//  res.sendFile(__dirname + '/main.html');
//});
//
//app.get('/main.js', (req,res) => {
//  res.sendFile(__dirname + '/main.js');
//});
//
//app.get('/main.wasm', (req, res) => {
//  res.sendFile(__dirname + '/main.wasm');
//});

app.listen(port, () => {
  console.log(`listening on port http://127.0.0.1:${port}`);
});


