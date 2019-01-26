import React, { Component } from 'react'
import { Switch, Route, Redirect } from 'react-router-dom'
import { Typography } from '@material-ui/core'

import { connect } from 'react-redux'
import { compose } from 'redux'
import * as actionType from '../../store/actions'

import IrmaAuthPage from './IrmaAuthPage'
import ViewRecordsPage from './ViewRecordsPage'
import ErrorPage from '../ErrorPage'

export class OrganizationPage extends Component {
    // currently loaded organization
    organization = null
    /**
     * Load selected organization based on url params from router match.
     * Extract organization from organizations and return organization name.
     * @returns {string} Organization name
     */
    getOrganization = () => {
        let { match, organizations } = this.props

        if (match.params && match.params.name && organizations) {
            let organization = organizations.filter((item) => {
                return item.name === match.params.name
            })
            if (organization.length === 1) {
                if (this.organization === organization[0]) {
                    return this.organization.name
                } else {
                    this.organization = organization[0]
                    return this.organization.name
                }
            } else {
                this.organization = null
                return null
            }
        }
    }
    /**
     * If organization not loaded in props and present in this.organization
     * it will dispatch action SELECT_ORGANIZATION
     * and redirect user to organization login page.
     */
    dispatchAction = () => {
        let { match, dispatch, history, organization } = this.props
        let propOrgName = organization.info.name
        let localOrgName = this.organization.name
        let onErrorPage = history.location.pathname.indexOf('/error') > 0

        if (
            propOrgName !== localOrgName &&
            localOrgName !== null &&
            !onErrorPage
        ) {
            dispatch({
                type: actionType.SELECT_ORGANIZATION,
                payload: this.organization,
            })

            let url = `/organization/${match.params.name}/login`
            history.push(url)
        }
    }

    render() {
        let { match } = this.props

        return (
            <div>
                <Typography variant="h4" color="primary" noWrap gutterBottom>
                    {this.getOrganization()}
                </Typography>
                <Switch>
                    <Route
                        path={`${match.url}/login`}
                        component={IrmaAuthPage}
                        {...this.props}
                    />
                    <Route
                        path={`${match.url}/view`}
                        component={ViewRecordsPage}
                        {...this.props}
                    />
                    <Route
                        path={`${match.url}/error`}
                        component={ErrorPage}
                        {...this.props}
                    />
                    <Redirect to={`${match.url}/login`} exact />
                </Switch>
            </div>
        )
    }

    componentDidMount() {
        this.dispatchAction()
    }

    componentDidUpdate() {
        this.dispatchAction()
    }
}

// -------------- REDUX CONNECTION ---------------------
/**
 * Map redux store states to local component properties
 * @param state: object, redux store object
 */
const mapStateToProps = (state) => {
    return {
        loading: state.loader.show,
        organization: state.organization,
        organizations: state.organizations.list,
    }
}

export default compose(connect(mapStateToProps))(OrganizationPage)
