//SPDX-License-Identifier: Apache-2.0
var cors = require('cors');

var cert = require('./controller.js');

module.exports = function(app){

  app.get('/certificates/:id', cors(), function(req, res){
    cert.get_cert(req, res);
  });
  app.post('/certificates', cors(), function(req, res){
    cert.addNewCertificate(req, res);
  });
  app.delete('/certificates', cors(), function(req, res){
    // cert.deleteCertificate(req, res);
  });
  app.get('/certificates', cors(), function(req, res){
    cert.get_all_cert(req, res);
  });
  app.post('/certificates/:certificate_id/transferName', cors(), function(req, res){
    cert.transfer_cert(req, res);
  });
}
