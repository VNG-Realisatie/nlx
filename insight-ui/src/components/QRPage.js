import React from 'react';
import axios from 'axios';
import QRCode from 'qrcode.react';

import { logGroup } from '../utils/appUtils';
import ErrorMsg from './ErrorMsg';

import qrpStyles from '../styles/QRPage';
import {withStyles} from '@material-ui/core';

class QRPage extends React.Component {
  state={
    pageTitle:"IRMA QR code",    
    uri:"/start-irma-session",
    data:{"type":"ISSUE","credentialType":"FIELDLAB"},
    sessionId: null,
    qrcValue: null,
    qrcSize: 256,
    qrcError: null,
    pageDesc:`
      Please scan this QR code with your IRMA smartphone app. 
    `
  }
  /**
   * NOTE! this is test function it should be 
   * replaced with actual api call to IRMA
   */
  setSignature(){ 
    this.setState({
      qrcValue: JSON.stringify({
        "qrContent": {
          "u": "https://int-diva-irma-api-server.appx.cloud/api/v2/issue/yE0S826ETydH3bKP0I09BZ169pccRJTS350zcrqviR9W",
          "v": "2.0",
          "vmax": "2.3",
          "irmaqr": "issuing"
        }
      })
    })
  }
  /**
   * Sign in to IRMA backend api and receive content
   * which needs to be loaded in QRcode component.
   * In addition IRMA sessionId is received which 
   * should be used to check status of singing process
   */
  getSignature(){
    let {uri,data} = this.state;
    debugger
    if (uri){
      axios.post(uri,data)
      .then((res)=>{
        debugger 
        this.setState({
          sessionId: data.sessionId,
          qrcValue: JSON.stringify(data.qrContent)
        })
      },(e)=>{
        this.setState({
          qrcError: e,
          qrcValue: null 
        });
        //console.error(e);
      })
    }else{
      console.warn("QRPage.getSignature: Singing url missing. Please check uri definition.");
    }
  }
  render(){
    let {qrcValue, qrError, qrcSize, pageTitle, pageDesc } = this.state;
    let {classes} = this.props;
    return (
      <div className={classes.root}>
        <h1>{ pageTitle }</h1>
        {
          qrcValue && <QRCode value={ qrcValue } size={qrcSize}/>
        }{          
          qrError && <ErrorMsg {...qrError}/>          
        }        
        <p>
          <br/><br/>
          { pageDesc }
        </p>
      </div>
    );
  }
  componentDidMount(){
    logGroup({
      title:"QRPage",
      method:"componentDidMount",
      state: this.state 
    })
    // use api call
    //this.getSignature();    
    //temp test
    this.setSignature(); 
  }
};

export default withStyles(qrpStyles)(QRPage);