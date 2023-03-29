const oneHour = 3600
const { fetchResource, fetchUSDToIDR } = require('../utils');
const moment = require('moment');
const ss = require('simple-statistics');
const nodeCache = require('node-cache');
const priceCache = new nodeCache({ stdTTL: oneHour});

module.exports = {
    fetch: async function (req, res, next) {
        const result = await fetchResource();

        res.json({
            "status": "success",
            "message": "Successfully Fetch Resources",
            "data": result
          });
    },

    fetchRoleNonAdmin: async function (req, res, next) {
        const resource = await fetchResource();
        let USDToIDR = priceCache.get('USDToIDR');

        if (!USDToIDR) {
            USDToIDR = await fetchUSDToIDR();
            priceCache.set('USDToIDR', USDToIDR, oneHour)
        }
        const result = [];
    
        for (let i = 0; i < resource.length; i++) {
          const item = resource[i];
    
          const modified = await new Promise(resolve => {
            if (!item.price) item.price = '0';

            priceUSD = (parseInt(item.price) / USDToIDR).toString();
            resolve({
            ...item,
            price_usd: priceUSD
            });
          });
    
          result.push(modified);
        }
    
        res.json({
          "status": "success",
          "message": "Successfully Fetch Resources",
          "data": result
        });
    },

    fetchRoleAdmin: async function (req, res, next) {
        if (req.user.role !== 'admin') {
          return res.status(403).json({
            "status": "Error",
            "message": "Forbidden",
          });
        };
    
        const resource = await fetchResource();
        const result = [];
        const objTmp = {};
    
        for (let i = 0; i < resource.length; i++) {
          const item = resource[i];
    
          if (!item.tgl_parsed) item.tgl_parsed = new Date();
          const tglParsed = moment(new Date(item.tgl_parsed)).format('YYYY-MM-DD HH:mm:ss');
    
          const modifiedItem = {
            ...item,
            tgl_parsed: tglParsed
          };
    
          const dateOnly = moment(new Date(tglParsed)).format('YYYY-MM-DD');
          const keyObj = `${item.area_provinsi}_${dateOnly}`;
    
          if (!objTmp[keyObj]) objTmp[keyObj] = [];
    
          objTmp[keyObj].push(modifiedItem);
        }
    
        for (const [key, value] of Object.entries(objTmp)) {
          const keyArr = key.split('_');
          const provinceName = keyArr[0]
          const date = keyArr[1]

          const year = parseInt(moment(new Date(date)).format('YYYY'));
          const month = parseInt(moment(new Date(date)).format('MM'));
          const week = parseInt(moment(new Date(date)).format('w'));
          console.log(date, week)
        
          const sizeArr = [];
          const priceArr = [];
          value.forEach(item => {
            sizeArr.push(parseInt(item.size));
            priceArr.push(parseInt(item.price));
          })
    
          const minSize = ss.min(sizeArr);
          const maxSize = ss.max(sizeArr);
          const medianSize = ss.median(sizeArr);
          const avgSize = ss.mean(sizeArr);
    
          const minPrice = ss.min(priceArr);
          const maxPrice = ss.max(priceArr);
          const medianPrice = ss.median(priceArr);
          const avgPrice = ss.mean(priceArr);
          
          result.push({
            year,
            month,
            week,
            province_area: provinceName,
            size_stat: {
              min: minSize,
              max: maxSize,
              median: medianSize,
              avg: avgSize
            },
            price_stat: {
              min: minPrice,
              max: maxPrice,
              median: medianPrice,
              avg: avgPrice
            }
          });
        }
    
        res.json({
          "status": "success",
          "message": "Successfully Fetch Resources",
          "data": result
        });
    },

    fetchToken: async function (req, res, next) {
        res.json({
          "status": "success",
          "message": "Successfully Claim Token",
          "data": req.user
        });
    }
}