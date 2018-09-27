import React, { Component } from 'react'
import axios from 'axios'
import { withRouter } from 'react-router-dom'
import { withStyles } from '@material-ui/core'
import { Typography } from '@material-ui/core'

import Table from './Table'
import { prepTableData, logGroup } from '../utils/appUtils'
import IrmaVerify from './IrmaVerify'

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

const colDef = [
    { id: 'date', label: 'Datum', width: 100, src:'created', type:"date", disablePadding: true},
    { id: 'time', label: 'Tijd', src:'created', type:"time", disablePadding: false},
    { id: 'source', label: 'Opgevraagd door', src:'source_organization', type:"string", disablePadding: false},
    { id: 'destination', label: 'Opgevraagd bij', src:'destination_organization', type:"string", disablePadding: false},
    { id: 'service', label: 'Reden', src:'service_name', type:"string", disablePadding: false }
]

const initialState = {
    dataSubjects: null,
    loggedIn: true,
    organization: null,
    jwt: null
}

class OrganizationPage extends Component {
    constructor(props) {
        super(props)

        this.state = initialState
    }

    componentDidMount() {
        this.updateOrganization(this.props)
    }

    componentWillReceiveProps(nextProps) {
        this.updateOrganization(nextProps)
    }

    updateOrganization(props) {
        let organization = props.organizations.filter((organization) => organization.name === props.match.params.name)

        if (organization.length > 0) {
            organization = organization[0]
        } else {
            organization = null
        }

        const newState = Object.assign({}, initialState, {
            organization
        })

        this.setState(newState)

        if (organization) {
            this.getDataSubjects(organization)
        }
    }

    getDataSubjects(organization) {
        axios.get(`http://${organization.insight_log_endpoint}:30080/getDataSubjects`)
        .then(response => {
            this.setState({ dataSubjects: this.convertDataSubjects(response.data.dataSubjects) })
        },(e)=>{
			console.error(e)
			this.setState({ error: true })
        })
    }

    convertDataSubjects(dataSubjects) {
        return Object.keys(dataSubjects).map((key) => {
            return {
                label: dataSubjects[key].label,
                attributes: [ key ]
            }
        })
    }

    prepData = rawData => {
        let prepData = prepTableData({
            colDef,
            rawData
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

    createTable() {
        let { data } = this.state;
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
                <h1>No information available</h1>
            )
        }
    }

    onCloseModal = () => {
        this.setState({
            modal: {
                open: false,
                data: null
            }
        })
    }

    onJWT = jwt => {
        logGroup({
            title: "Organization",
            method: "onJWT",
            jwt: jwt,
            props: this.props,
            state: this.state
        });
    }

    onCancelVerification = () =>{
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

    render() {
        if (!this.state.organization) {
            return (
                <div>Loading organization...</div>
            )
        }

        if (!this.state.loggedIn && !this.state.dataSubjects) {
            return (
                <div>Loading data subjects...</div>
            )
        }

        if (!this.state.loggedIn && this.state.dataSubjects) {
            return (
                <IrmaVerify
                    //server={`http://${this.state.organization.insight_irma_endpoint}:30080/api/v2/`}
                    server="https://demo.irmacard.org/tomcat/irma_api_server/api/v2/"
                    attributes={[{
                        label: "over18",
                        attributes: [
                            "irma-demo.MijnOverheid.ageLower.over18"
                        ]
                    }]}
                    onJWT={this.onJWT}
                    onCancel={this.onCancelVerification}
                />
            )
        }

        return (
            <React.Fragment>
                <Typography variant="title" color="primary" noWrap gutterBottom>
                    Organization: {this.state.organization.name}
                </Typography>
            </React.Fragment>
        )
    }
}

export default withStyles(styles)(withRouter(OrganizationPage))