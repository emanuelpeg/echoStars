var express = require('express');
const { redirect } = require('express/lib/response');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res) {
  res.json({ hello: "Hello Success!"});
});

module.exports = router;
