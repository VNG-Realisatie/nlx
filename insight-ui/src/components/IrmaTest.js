import React, { Component } from 'react';

//import $ from 'jquery'; 
//import * as IRMA from '../utils/irma';

class IrmaPage extends Component {
  render() {
    return (
      <div>
          <h1>Irma test page</h1>
      </div>
    );
  }
  componentDidMount(){
    this.initIRMA();
  }
  initIRMA = () =>{
    
    let irma = window.IRMA.init("https://demo.irmacard.org/tomcat/irma_api_server/api/v2/");
    this.createUnsignedVerificationJWT()
  }
  createUnsignedVerificationJWT(){
    let request = {
      "validity": 60,
      "request": {
          "content": [
              {
                  "label": "over12",
                  "attributes": ["irma-demo.MijnOverheid.ageLower.over12"]
              },
          ]
      }
    }
    let jwt = window.IRMA.createUnsignedVerificationJWT(request);
    //return jwt;
    this.IrmaVerify(jwt);
  }

  IrmaVerify = (jwt) =>{
    return new Promise( (res, rej)=>{
      debugger
      window.IRMA.verify(jwt, 
        (jwt)=>{
          console.log("success..",jwt);
          res(jwt);
        }, 
        (w)=>{
          console.log("warning", w);
          rej(w);
        }, 
        (e)=>{
          console.log("error", e);
          rej(e)
        }
      );
    })
  }


}

export default IrmaPage;