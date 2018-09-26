import React, { Component } from 'react';

import './IrmaTest.css';
import { logGroup } from '../utils/appUtils';


class IrmaPage extends Component {
  state={
    jwt:null,
    cancel:false,
    error:null
  }
  render() {
    let {jwt, cancel, error} = this.state;
    logGroup({
      title:"IrmaPage",
      method:"render",
      state: this.state 
    })
    return ( 
      <div>
          <h1>Irma test page</h1>
          {jwt && <h2>Token received</h2>}
          {cancel && <h2>Verification session CANCELED</h2>}
          {error && <h2>Token error: {error}</h2>}
      </div>
    );
  }
  componentDidMount(){
    this.initIRMA();
  }
  initIRMA = () =>{
    let irma = window.IRMA.init("https://demo.irmacard.org/tomcat/irma_api_server/api/v2/");
    //let irma = window.IRMA.init("http://insight-api.test.haarlem.commonground.nl/getDataSubjects")

    this.createUnsignedVerificationJWT([
      {
        "label": "over12",
        "attributes": [
          "irma-demo.MijnOverheid.ageLower.over12"
        ]
      },{
        "label": "over1",
        "attributes": [
          "irma-demo.MijnOverheid.ageLower.over12"
        ]
      }
    ])
  }
  /*
  initIRMA = () =>{
    //let irma = window.IRMA.init("https://demo.irmacard.org/tomcat/irma_api_server/api/v2/");
    let irma = window.IRMA.init("http://insight-api.test.haarlem.commonground.nl/getDataSubjects")

    this.createUnsignedVerificationJWT([
      {{
        "label": "over18",
        "attributes": [
          "irma-demo.MijnOverheid.ageLower.over12"
        ]
      }
    ])
  }*/
  createUnsignedVerificationJWT(attributes){
    let request = {
      "validity": 60,
      "request": {
          "content": [
              ...attributes
          ]
      }
    }
    debugger
    let jwt = window.IRMA.createUnsignedVerificationJWT(request);
    //return jwt;
    this.irmaVerify(jwt).then((jwt)=>{
      //console.log("success...", jwt);
      this.setState({
        jwt: jwt,
        error:null,
        cancel: false
      })
    }).catch(e=>{
      if (e==="CANCELED"){
        this.setState({
          jwt: null,
          error: null,
          cancel: true
        })
        //console.log("CANCELED...go home")
      }else{
        //console.log("ERROR, SHOW ERROR...", e);
        this.setState({
          jwt: null,
          error: e,
          cancel: false
        })
      }
    })
  }
  /**
   * Performs verification using irma.js
   * returns promise which resolves with JWT
   * reject has CANCEL and ERROR states
   * CANCEL state occures when user cancels 
   * QR code scanning session
   */
  irmaVerify = (jwt) =>{
    return new Promise( (res, rej)=>{
      //debugger
      window.IRMA.verify(jwt, 
        (jwt)=>{
          //console.log("success..",jwt);
          res(jwt);
        }, 
        (w)=>{
          //console.log("warning", w);
          rej(w);
        }, 
        (e)=>{
          //console.log("error", e);
          rej(e)
        }
      );
    })
  }
}

export default IrmaPage;