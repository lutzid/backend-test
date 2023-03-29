const express = require('express');
const { fetchToken, fetchRoleAdmin, fetchRoleNonAdmin, fetch } = require('../controllers');
const router = express.Router();
const { JWTAuthentication } = require('../middleware');


router.get('/api/fetch/without-jwt', fetch)
router.use('/api', JWTAuthentication)

router.get('/api/fetch/role-admin', fetchRoleAdmin)
router.get('/api/fetch/role-non-admin', fetchRoleNonAdmin)
router.get('/api/fetch/token', fetchToken)

module.exports = router;