import React, { Component } from 'react';

import './IrmaVerify.css';
import { logGroup } from '../utils/appUtils';

class IrmaPage extends Component {
  state={
    jwt:null,
    cancel:false,
    error:null,
    irma:{
      server:"https://demo.irmacard.org/tomcat/irma_api_server/api/v2/",
      //server:"http://irma-api.test.haarlem.commonground.nl/api/v2/",
      attributes:[{
          "label": "over18",
          "attributes": [
            "irma-demo.MijnOverheid.ageLower.over18"
          ]
        }
      ]      
    }
  }
  render() {
    let { jwt, cancel, error } = this.state;
    logGroup({
      title:"IrmaPage",
      method:"render",
      state: this.state,
      props: this.props 
    })
    return ( 
      <div>
          {jwt && <h2>Token received</h2>}
          {cancel && <h2>Verification session CANCELED</h2>}
          {error && <h2>Verification error: { JSON.stringify(error) }</h2>}
      </div>
    );
  }
  
  componentDidMount(){
    logGroup({
      title:"IrmaPage",
      method:'componentDidMount',
      props: this.props,
      state: this.state
    })
    //live
    this.performSteps();
    //test sending JWT back
    //this.props.onJWT("asbnasdasdasd");
    //test cancel verification
    //this.props.onCancel();
  }

  componentDidUpdate(){
    logGroup({
      title:"IrmaPage",
      method:'componentDidUpdate',
      props: this.props,
      state: this.state
    })
  }

  performSteps = () =>{
    let {server, attributes } = this.state.irma;
    //debugger 
    //1. init IRMA server
    this.initIRMA(server)
    .then( d => {
      //debugger
      //return this.getDataSubjects("/haarlem/getDataSubjects");
      //temporary for test
      return this.state.attributes;
    })
    .then( ds => {
      //debugger
      //create Verification JWT
      return this.createUnsignedVerificationJWT(attributes);
    })
    .then( jwt1 =>{
      //debugger
      //verify
      return this.irmaVerify(jwt1);
    })
    .then( jwt2 => {
      //debugger
      //console.log("success...", jwt);
      this.props.onJWT(jwt2);
    })
    .catch( e => {
      debugger
      if (e==="CANCELED"){
        //debugger
        this.props.onCancel(true);
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
  
  initIRMA = server =>{
    return new Promise((res,rej)=>{
      try{
        //debugger
        //intialize IRMA using irma.js
        window.IRMA.init(server);
        res(true);
      }catch(e){
        rej(e)
      }
    })
  }

  getDataSubjects = url => {
    //debugger
    return fetch(url)
      .then(d=>{
        debugger
        return d.json()
      })
  }

  createUnsignedVerificationJWT = attributes =>{
    let request = {
      "validity": 60,
      "request": {
          "content": [
              ...attributes
          ]
      }
    }    
    return new Promise ((res,rej)=>{
      //debugger
      try{
        let jwt = window.IRMA.createUnsignedVerificationJWT(request);
        res(jwt);
      }catch (e) { 
        rej(e);
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
  irmaVerify = jwt =>{
    return new Promise( (res, rej)=>{
      //debugger
      window.IRMA.verify(jwt, 
        (d)=>{
          //console.log("success..",jwt);
          res(d);
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