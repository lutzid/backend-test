const jwt = require('jsonwebtoken');

module.exports = {
  JWTAuthentication: function (req, res, next) {
    const authHeader = req.headers['authorization'];
    const token = authHeader && authHeader.split(' ')[1];

    // Check if token is null
    if (token == null) {
      return res.status(401).json({
        "status": "error",
        "message": "You are unauthorized to access this api.",
      });
    }
    
    // Verify the jwt token
    jwt.verify(token, 'verysecretkey', (err, user) => {
      if (err) {
        return res.status(401).json({
          "status": "error",
          "message": "You are prohibited to access this api.",
        });
      }

      req.user = user;
      console.log("authenticate")
      next();
    });
  }
}