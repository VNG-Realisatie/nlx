import React, { Component } from 'react'

import { withStyles } from '@material-ui/core'
//import { WithSimpleDivaAuthorization } from 'diva-react';
//import { Route } from 'react-router-dom';

import { Typography } from '@material-ui/core'

import Table from './Table'
import { prepTableData, logGroup } from '../utils/appUtils'

//fake data
import logsIn from '../mockdata/txlog.dev.denhaag-out.json'
//import logsOut from '../store/logs-out';

//import SimpleModal from './SimpleModal'
import LogModal from './LogModal';
import IrmaPage from './IrmaVerify';
//import IrmaPage from './IrmaTest';
//import QRPage from './QRPage';
//import ClockIcon from '@material-ui/icons/AccessTimeOutlined'
//import CalendarIcon from '@material-ui/icons/CalendarToday'

const styles = theme => ({
	calendarIcon: {
		fontSize: 14,
		marginBottom: -2,
		marginRight: 3
	},
	clockIcon: {
		fontSize: 15,
		marginBottom: -3,
		marginRight: 2
	}
})

class OrganizationPage extends Component {
    state={
        cid: null,
        //jwt to use at api point
        jwt: null,
        loggedIn: false,
        //column definitions to pass to table component
        colDef: [
            { id: 'date', label: 'Datum', width: 100, src:'created', type:"date", disablePadding: true},
            { id: 'time', label: 'Tijd', src:'created', type:"time", disablePadding: false},
            { id: 'source', label: 'Opgevraagd door', src:'source_organization', type:"string", disablePadding: false},
            { id: 'destination', label: 'Opgevraagd bij', src:'destination_organization', type:"string", disablePadding: false},
            { id: 'service', label: 'Reden', src:'service_name', type:"string", disablePadding: false }
        ],
        //data to pass to table component
        data: [],
        //modal on the page
        modal: {
            open: false,
            data: null
        }
    }

    componentDidMount(){
        //let { cid } = this.props.match;
        //let { jwt } = this.state;
        logGroup({
            title: "Organization",
            method: "componentDidMount",
            props: this.props,
            state: this.state
        });
        /*debugger 
        if (jwt){
            this.getOrganizationInfo();
        } else {
            console.log("no jwt...", jwt)
        }*/
    }
    /*
    shouldComponentUpdate(nextProps, nextState){
        let { cid, jwt } = nextProps.match.params,
            { modal } = nextState;
        
        debugger 
        
        if (cid === this.state.cid
            && modal.open === this.state.modal.open
            && jwt === this.state.jwt ){
            return false
        } else {
            return true
        }
    }*/
    componentDidUpdate(){
        logGroup({
            title:"Organization",
            method:"componentDidUpdate",
            props: this.props,
            state: this.state
        });
        //this.getOrganizationLogs();
    }
    /**
     * Get Organization log info.
     * NOTE: The initial idea is to place api call here.
     * WARNING: Currently mock data is DIRECTLY LOADED
     */
    getOrganizationLogs = jwt =>{
        //debugger

        this.setState({
            cid: this.props.match.params.cid,
            jwt: jwt,
            data: this.prepData(logsIn.records)
        })
    }

    /**
     * Convert raw log data to MUI table format
     * @param {object} rawData - array of objects
     */
    prepData = rawData => {
        let prepData = prepTableData({
            colDef: this.state.colDef,
            rawData: rawData
        })        
        return prepData;
    }

    getDetails = id => {
        let row = this.state.data[id];
        this.setState({
            modal: {
                open:true,
                data:row
            }
        })
    }

    createTable(){
        let { colDef, data } = this.state;
        if ( data.length > 0 ){
            return (
                <Table
                    cols={colDef}
                    data={data}
                    onDetails={this.getDetails}
                />
            )
        } else {
            return (
                <h1>No information avaliable</h1>
            )
        }
    }

    onCloseModal = () =>{
        this.setState({
            modal: {
                open: false,
                data: null
            }
        })
    }

    onJWT = jwt => {
        //debugger
        logGroup({
            title: "Organization",
            method: "onJWT",
            jwt: jwt,
            props: this.props,
            state: this.state
        });
        //get logs
        this.getOrganizationLogs(jwt)
    }

    onCancelVerification = () =>{
        //debugger
        let { history } = this.props; 
        logGroup({
            title: "Organization",
            method: "onCancelVerification",
            action: "push to /",
            props: this.props,
            state: this.state
        });
        history.push('/');
    }

    getContent = () => {
        //debugger 
        let { jwt } = this.state;
        const { modal } = this.state;
        if (jwt){
            return (
                <section>
                    { this.createTable() }
                    { modal.data && <LogModal
                        open={modal.open}
                        closeModal={this.onCloseModal}
                        data={modal.data} />                                    }
                </section>                
            )
        }else{
            return(<IrmaPage 
                onJWT={this.onJWT}
                onCancel={this.onCancelVerification}/>
            )
        }
    }

    render() {        
        
        const { cid } = this.props.match.params  

        return (
            <React.Fragment>
                <Typography variant="title" color="primary" noWrap gutterBottom>
                    Selected Organization {cid}
                </Typography>
                
                { this.getContent() }

            </React.Fragment>
        )
    }
}

export default withStyles(styles)(OrganizationPage)