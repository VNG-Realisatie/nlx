import React from 'react';
import axios from 'axios';
import QRCode from 'qrcode.react';

import { logGroup } from '../utils/appUtils';
import ErrorMsg from './ErrorMsg';

import qrpStyles from '../styles/QRPage';
import { withStyles } from '@material-ui/core';

class QRPage extends React.Component {
  state = {
    mode:"DEMO",
    pageTitle:"IRMA QR code",    
    //internal action types to track state during poling
    //values: INIT, SET_QRC_VALUE, WAITING,  
    action:"INIT", 
    api:{
      baseUriLong:'http://insight-api.dev.denhaag.minikube:30080/irma/',  
      baseUriShort:'/irma',
      getQRvalues:'verification-start',
      getVerificationState:'verification',
      verification:'verification'      
    },    
    //token value recevived from irma api 
    sessionId: null,
    //stringified and prepared QRC value
    qrcValue: null,   
    //size of qr image
    qrcSize: 256,
    //token received upon succefull verification
    jwt:null,
    //error - not succefull verification
    qrcError: null,    
    //reference to interval object
    timer: null,
    //keep track of attempts
    attempt:{
      //try every xms
      interval: 1000,
      //keep count
      count: 0,
      //max attempts
      max: 5,
      //message to show when max reached
      maxMsg:"Timeout expired"
    },
    //page message 
    pageDesc:`
      Please scan this QR code with your IRMA smartphone app. 
    `,
  }
  /**
   * NOTE! this is DEMO function. It is used when mode is set to DEMO.
   * Change state.mode to any other value (LIVE eg.) after backend 
   * is implemented. 
   **/
  demoSignature(){ 
    let qrcVal = {
      "irmaqr":"disclosing",
      "u":"0Cogjkq2MLL5s8NEXQele5Am6wsS6ulvJ2gOxW6ymk",
      "v":"2.0",
      "vmax":"2.3"
    }
    //prepare QRval for QRcode
    this.prepQRCode(qrcVal);
  } 
  prepQRCode = (rawQR) =>{
    //debugger    
    let qrcVal = {
      ...rawQR,
      u: `${this.state.api.baseUriLong}/${this.state.api.getVerificationState}/${rawQR.u}` 
    }
    this.setState({
      action: 'SET_QRC_VALUE',
      sessionId: rawQR.u,      
      qrcValue: JSON.stringify(qrcVal)
    });
  }
  /**
   * Sign in to IRMA backend api and receive object
   * which needs to be loaded in the QRcode component.
   * Received IRMA sessionId should be used to check 
   * status of singing process.
   **/
  getQRValues(){
    let uri = `${this.state.api.baseUriShort}/${this.state.api.getQRvalues}`;
    //debugger
    if (uri){
      axios.get(uri)
      .then((res)=>{
        debugger 
        this.prepQRCode(res);
      },(e)=>{
        debugger
        this.setState({
          action:"ERROR",
          qrcError: e,
          qrcValue: null, 
          sessionId: null,
          jwt: null 
        });
        //console.error(e);
      })
    }else{
      this.setState({
        action:"ERROR",
        qrcError:{
          id: 500,                  
          message:'QRPage.getQRValues: Signing url undefined.'
        }
      });
    }
  }
  getStatus = () => {
    //debugger     
    let { attempt } = this.state;   

    if (attempt.count > attempt.max){
      //we reached maximum no. attempts
      if (this.state.mode==="DEMO"){        
        //in demo mode we simply send JWT
        this.setState({
          action:"JWT",          
          jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",          
        });
      }else{        
        this.setState({
          action:"ERROR",
          qrcValue: null,                  
          jwt: null,
          qrcError:{
            id: 500,          
            message:'Timeout expired.'
          }
        });
      }      
      return;      
    }
    console.log("waiting...");
    //we wait specified amount of ms
    setTimeout(()=>{ 
      if (this.state.mode==="DEMO"){
        //demo just increase counter
        this.setState({
          action:"WAITING",
          attempt:{
            ...attempt,
            count: attempt.count + 1 
          }
        })
      } else {
        //make api call
        this.getVerificationState();  
      }
    }, attempt.interval);
  }
  getVerificationState = () =>{
    //debugger
    let uri = `${this.state.api.baseUriShort}/${this.state.api.getQRvalues}`;
    let options = {
      method:"POST",
      body: this.state.sessionId
    };
    
    fetch(uri, options)
      .then(resp => {
        return resp.json();
      })
      .then(data => {
        debugger 
        if (data.state === "JWT"){
          this.setState({
            action:"JWT",
            jwt: data.payload
          })
        } else {
          this.setState({
            action:"WAITING",
            attempt:{
              ...this.state.attempt,
              count: this.attempt.count + 1 
            }
          })
        }
      })
      .catch(e => {
        this.setState({
          action:"ERROR",
          qrcError: e
        })
      });
  }
  authenticated = (jwt) =>{
    //pass JWT to parent
    /*
    if (this.props.onJWT){
      this.props.onJWT(jwt);
    }*/
    let url = `company/2/${jwt}`;
    this.props.history.push(url);
  }
  render(){
    let {qrcValue, qrcError, qrcSize, pageTitle, pageDesc } = this.state;
    let {classes} = this.props;
    return (
      <div className={classes.root}>
        <h1>{ pageTitle }</h1>
        {
          qrcValue && <QRCode value={qrcValue} size={qrcSize}/>
        }{          
          qrcError && <ErrorMsg {...qrcError}/>          
        }        
        <p>
          <br/><br/>
          { qrcValue ? pageDesc : null }
        </p>
      </div>
    );
  }
  componentDidMount(){
    logGroup({
      title:"QRPage",
      method:"componentDidMount",
      state: this.state,
      props: this.props
    })    
    //if in demo mode we skip QRcode api call
    if (this.state.mode==="DEMO"){
      this.demoSignature();
    }else{
      this.getQRValues();
    }    
  }
  componentDidUpdate(){
    logGroup({
      title:"QRPage",
      method:"componentDidUpdate",
      state: this.state,
      props: this.props
    })
    //check the state of component
    this.pageReducer();
  }
  pageReducer = () => {   
    //debugger 
    switch(this.state.action.toUpperCase()){
      case "SET_QRC_VALUE":
      case "WAITING":        
        //start new cycle
        this.getStatus(); 
        break;              
      case "JWT": 
        //we received JWT       
        debugger 
        console.log("JWT received...");        
        this.authenticated(this.state.jwt);
        break;      
      case "ERROR":        
      default:               
        //do nothing error portion is shown  
    }
  }
  removeInterval=()=>{
    //remove interval
    let {timer} = this.state;
    if (timer){
      clearInterval(timer)
    }    
  }
  componentWillUnmount(){
    //debugger 
    this.removeInterval();
  }
};

export default withStyles(qrpStyles)(QRPage);