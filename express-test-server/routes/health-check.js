var express = require('express');
const { redirect } = require('express/lib/response');
var router = express.Router();

/* GET home page. */
router.get('/success', function(req, res) {
  res.json({ msg: "Success!"});
});

router.get('/error', function(req, res) {
  res.status(500).json({ msg: "Error!" });
});

router.get('/sketchy', function(req, res) {
  const randomNumber = Math.random();
  // Determine the response based on the random number
  if (randomNumber < 0.9) {
    // Successful response (status 200)
    res.status(200).json({ message: 'Success!' });
  } else {
    // Server error response (status 500)
    res.status(500).json({ message: 'Internal Server Error' });
  }
});

module.exports = router;
