import React, { Component } from 'react'
import { withRouter } from 'react-router-dom'
import { withStyles, Typography } from '@material-ui/core'

import Irma from './Irma'
import OrganizationRecords from './OrganizationRecords'

const styles = (theme) => ({
    calendarIcon: {
        fontSize: 14,
        marginBottom: -2,
        marginRight: 3,
    },
    clockIcon: {
        fontSize: 15,
        marginBottom: -3,
        marginRight: 2,
    },
})

const initialState = {
    loggedIn: false,
    organization: null,
    jwt: null,
}

class OrganizationPage extends Component {
    constructor(props) {
        super(props)

        this.state = initialState
        this.afterLogin = this.afterLogin.bind(this)
    }

    componentDidMount() {
        this.updateOrganization(this.props)
    }

    UNSAFE_componentWillReceiveProps(nextProps) {
        this.updateOrganization(nextProps)
    }

    updateOrganization(props) {
        let organization = props.organizations.filter(
            (organization) => organization.name === props.match.params.name,
        )

        if (organization.length > 0) {
            organization = organization[0]
        } else {
            organization = null
        }

        const newState = Object.assign({}, initialState, {
            organization,
        })

        this.setState(newState)
    }

    afterLogin(organization, jwt) {
        this.setState({
            loggedIn: true,
            jwt,
        })
    }

    render() {
        if (!this.state.organization) {
            return <div>Loading organization...</div>
        }

        let content
        if (!this.state.loggedIn) {
            content = (
                <Irma
                    organization={this.state.organization}
                    afterLogin={this.afterLogin}
                />
            )
        } else {
            content = (
                <OrganizationRecords
                    organization={this.state.organization}
                    jwt={this.state.jwt}
                />
            )
        }

        return (
            <div>
                <Typography variant="title" color="primary" noWrap gutterBottom>
                    Organization: {this.state.organization.name}
                </Typography>
                {content}
            </div>
        )
    }
}

export default withStyles(styles)(withRouter(OrganizationPage))
