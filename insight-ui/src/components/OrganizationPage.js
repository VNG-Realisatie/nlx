import React, { Component } from 'react'

import { withStyles } from '@material-ui/core'
//import { WithSimpleDivaAuthorization } from 'diva-react';

import { Typography } from '@material-ui/core'

import Table from './Table'
import { prepTableData, logGroup } from '../utils/appUtils'

//fake data
import logsIn from '../mockdata/txlog.dev.denhaag-out.json'
//import logsOut from '../store/logs-out';

//import SimpleModal from './SimpleModal'
import LogModal from './LogModal';
import QRPage from './QRPage';
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
        logGroup({
            title: "Organization",
            method: "componentDidMount",
            props:this.props,
            state: this.state
        });
        this.getOrganizationInfo();
    }
    
    shouldComponentUpdate(nextProps, nextState){
        let { cid, jwt } = nextProps.match.params,
            { modal } = nextState;
        
        //debugger 
        
        if (cid === this.state.cid
            && modal.open === this.state.modal.open
            && jwt == this.state.jwt ){
            return false
        } else {
            return true
        }
    }

    componentDidUpdate(){
        logGroup({
            title:"Organization",
            method:"componentDidUpdate",
            props:this.props,
            state: this.state
        });
        this.getOrganizationInfo()
    }
    /**
     * Get Organization log info.
     * NOTE: The initial idea is to place api call here.
     * WARNING: Currently mock data is ONLY LOADED if
     * cid==2
     */
    getOrganizationInfo(){
        //debugger
        if (this.props.match.params.cid==="2"){
            this.setState({
                cid: this.props.match.params.cid,
                data: this.prepData(logsIn.records)
            })
        } else {
            this.setState({
                cid: this.props.match.params.cid,
                data: []
            })
        }
    }

    /**
     * Convert raw log data to MUI table format
     * @param {object} rawData - array of objects
     */

    prepData(rawData){
        let prepData = prepTableData({
            colDef: this.state.colDef,
            rawData: rawData
        })        
        return prepData
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
        if (data.length>0){
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

    getContent = () => {
        //debugger 
        let { jwt } = this.props.match.params;
        const { modal } = this.state;
        //not completed 
        //if (jwt){
            return (
                <section>
                    { this.createTable() }
                    { modal.data && <LogModal
                        open={modal.open}
                        closeModal={this.onCloseModal}
                        data={modal.data} />                                    }
                </section>                
            )
        /*}else{
            return (
                <QRPage {...this.props}/>
            )
        }*/
    }

    render() {        
        
        const { cid } = this.props.match.params  

        return (
            <React.Fragment>
                <Typography variant="title" color="primary" noWrap gutterBottom>
                    Selected Organization {cid}
                </Typography>

                {this.getContent()}

            </React.Fragment>
        )
    }
}
/*
export default WithSimpleDivaAuthorization(
    {},
    'pbdf.pbdf.email.email',
    'Email'
)(Organization);
*/
export default withStyles(styles)(OrganizationPage)