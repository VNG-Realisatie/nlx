import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import { withStyles } from '@material-ui/core'
//import { WithSimpleDivaAuthorization } from 'diva-react';

import { Typography } from '@material-ui/core'

import Table from './Table'
import { prepTableData, logGroup } from '../utils/appUtils'

//fake data
import logsIn from '../mockdata/txlog.dev.denhaag-out.json'
//import logsOut from '../store/logs-out';

import SimpleModal from './SimpleModal'
import ClockIcon from '@material-ui/icons/AccessTimeOutlined'
import CalendarIcon from '@material-ui/icons/CalendarToday'

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
        this.getOrganizationInfo()
    }

    shouldComponentUpdate(nextProps, nextState){
        let { cid } = nextProps.match.params,
            { modal } = nextState;

        if (cid === this.state.cid
            && modal.open === this.state.modal.open ){
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

    render() {
        const { classes } = this.props
        const { cid } = this.props.match.params
        const { modal } = this.state
        const data = modal.data

        let modalContent
        if (data) {
            const d = new Date(data['created'])
            const localDate = d.toLocaleDateString()
            const localTime = d.toLocaleTimeString()

            modalContent = (
                <React.Fragment>
                    <Typography variant="title" color="primary" style={{marginLeft: -1, marginBottom: 5}}>
                        {data['attributes'] ? data['attributes'] : "Geen attribuut opgevraagd."}
                    </Typography>
                    <div style={{display: 'flex', justifyContent: 'space-between', flexWrap: 'wrap'}}>
                        <Typography variant="caption">
                            #{data['id']}
                        </Typography>
                        <Typography variant="caption">
                            <CalendarIcon className={classes.calendarIcon} />{localDate}
                            &nbsp;&nbsp;&nbsp;<ClockIcon className={classes.clockIcon} />{localTime}
                        </Typography>
                    </div>
                    <br/>
                    <div style={{display: 'flex', justifyContent: 'space-between'}}>
                        <div>
                            <Typography variant="caption">
                                Opgevraagd door
                            </Typography>
                            <Link to="">{data['destination_organization']}</Link>
                        </div>
                        <div>
                            <Typography variant="caption" align="right">
                            Opgevraagd bij
                            </Typography>
                            <Typography align="right">
                                <Link to="">{data['source_organization']}</Link>
                            </Typography>
                        </div>
                    </div>
                    <br/>
                    <Typography variant="caption">
                        Reden
                    </Typography>
                        {data['reason'] ? data['reason'] : "Geen reden opgegeven."}
                </React.Fragment>
            )
        }

        return (
            <React.Fragment>
                <Typography variant="title" color="primary" noWrap gutterBottom>
                    Selected Organization {cid}
                </Typography>

                {this.createTable()}

                <SimpleModal
                    open={modal.open}
                    closeModal={this.onCloseModal}
                >
                    {modalContent}
                </SimpleModal>
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