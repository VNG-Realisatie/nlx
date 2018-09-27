import React, { Component } from 'react'
import axios from 'axios'
import { withRouter } from 'react-router-dom'
import { withStyles } from '@material-ui/core'
import { Typography } from '@material-ui/core'

import Table from './Table'
import LogModal from './LogModal'
import { prepTableData } from '../utils/appUtils'
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
    { id: 'destination', label: 'Opgevraagd bij', src:'destination_organization', type:"string", disablePadding: false}
]

const initialState = {
    modal: {
        open: false,
        data: null
    },
    dataSubjects: null,
    loggedIn: true,
    organization: null,
    records: [],
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
            this.onJWT('123')
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

    onJWT(jwt) {
        this.setState({ loggedIn: true })

        axios({
            method: 'post',
            url: `http://${this.state.organization.insight_log_endpoint}:30080/fetch`,
            data: jwt
        }).then(response => {
            this.setState({records: response.data.records })
        },(e)=>{
			console.error(e)
			this.setState({ error: true })
        })
    }

    onCancelVerification = () =>{
        let { history } = this.props;
        history.push('/');
    }

    getDetails = id => {
        this.setState({
            modal: {
                open: true,
                data: this.state.records[id]
            }
        })
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
                    server={`http://${this.state.organization.insight_irma_endpoint}:30080/api/v2/`}
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

        let table
        if (this.state.records) {
            table = (
                <Table
                    cols={colDef}
                    onDetails={this.getDetails}
                    data={prepTableData({
                        colDef,
                        rawData: this.state.records
                    })}
                />
            )
        } else {
            table = (
                <p>No information available</p>
            )
        }

        const { modal } = this.state

        return (
            <div>
                <Typography variant="title" color="primary" noWrap gutterBottom>
                    Organization: {this.state.organization.name}
                </Typography>
                {table}
                { modal.data && <LogModal open={modal.open} closeModal={this.onCloseModal} data={modal.data}/> }
            </div>
        )
    }
}

export default withStyles(styles)(withRouter(OrganizationPage))