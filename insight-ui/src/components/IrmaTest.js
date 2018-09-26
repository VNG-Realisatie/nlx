import React, { Component } from 'react';

import './IrmaTest.css';
import { logGroup } from '../utils/appUtils';


class IrmaPage extends Component {
  state={
    jwt:null,
    cancel:false,
    error:null,
    irma:{
      //"https://demo.irmacard.org/tomcat/irma_api_server/api/v2/"
      server:"http://irma-api.test.haarlem.commonground.nl/api/v2/",
      /*attributes:[{
          "label": "over12",
          "attributes": [
            "irma-demo.MijnOverheid.ageLower.over12"
          ]
        },{
          "label": "over18",
          "attributes": [
            "irma-demo.MijnOverheid.ageLower.over18"
          ]
        }
      ]*/      
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
          <h1>Irma test page</h1>
          {jwt && <h2>Token received</h2>}
          {cancel && <h2>Verification session CANCELED</h2>}
          {error && <h2>Token error: { JSON.stringify(error) }</h2>}
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
    //this.performSteps();
    //just send JWT back
    this.props.onJWT("asbnasdasdasd");
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
    debugger 
    //1. init IRMA server
    this.initIRMA(server)
    .then( d => {
      //debugger
      return this.getDataSubjects("/haarlem/getDataSubjects")
    })
    .then( ds => {
      //debugger
      //create Verification JWT
      return this.createUnsignedVerificationJWT(attributes);
    })
    .then((jwt1)=>{
      debugger
      //verify
      return this.irmaVerify(jwt1);
    })
    .then((jwt2)=>{
      debugger
      //console.log("success...", jwt);
      this.setState({
        jwt: jwt2,
        error:null,
        cancel: false
      })
    })
    .catch(e=>{
      debugger
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
    debugger
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
      debugger
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
      debugger
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