import React, { Component } from 'react';
import './IrmaVerify.css';

class IrmaPage extends Component {
  constructor(props) {
    super(props)

    this.state = {
      jwt:null,
      cancel:false,
      error:null,
      irma:{
        server: this.props.server,
        attributes: this.props.attributes
      }
    }
  }

  render() {
    let { jwt, cancel, error } = this.state

    return (
      <div>
          {jwt && <h2>Token received</h2>}
          {cancel && <h2>Verification session CANCELED</h2>}
          {error && <h2>Verification error: { JSON.stringify(error) }</h2>}
      </div>
    )
  }

  componentDidMount() {
    //live
    this.performSteps();
    //test sending JWT back
    //this.props.onJWT("asbnasdasdasd");
    //test cancel verification
    //this.props.onCancel();
  }

  performSteps = () =>{
    let {server, attributes } = this.state.irma;

    //1. init IRMA server
    this.initIRMA(server)
    .then( d => {
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
      //debugger
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