const fetch = require('node-fetch');

module.exports = {
    fetchResource: async function () {
      const response = await fetch('https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list');
      const result = await response.json();
      const resource = result.filter(res => res.uuid !== null);

      return resource;
    },
    fetchUSDToIDR: async function () {
      const response = await fetch('https://api.currencyfreaks.com/v2.0/rates/latest?apikey=3f9150d231414c139267df15abc9d74d&symbols=IDR');
      const result = await response.json();
      console.log(result.rates.IDR)
      return result.rates.IDR;
    }
  };