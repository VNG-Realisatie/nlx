import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Spinner } from '@commonground/design-system'

import Table from '../../components/Table'
import LogModal from '../../components/LogModal'
import { prepTableData } from '../../utils/appUtils'
import * as actionType from '../../store/actions'
import NoDataMessage from '../../components/NoDataMessage'

class ViewRecordsPage extends Component {
    state = {
        error: false,
        modal: {
            open: false,
            data: null,
        },
    }
    /**
     * Send user to error page
     * @param {number} error.number error number (in same cases received from api)
     * @param {string} error.description
     */
    showError = (error) => {
        let { history, match } = this.props
        let url = match.path.replace('view', 'error')

        history.push(url, {
            error,
        })
    }

    getDetails = (id) => {
        this.setState({
            modal: {
                open: true,
                data: this.props.logs[id],
            },
        })
    }

    onCloseModal = () => {
        this.setState({
            modal: {
                open: false,
                data: null,
            },
        })
    }

    changeTableOptions = (pageDef) => {
        const { api, jwt, name } = this.props
        const { page, rowsPerPage } = pageDef
        let params = {
            page,
            rowsPerPage,
        }
        this.props.dispatch({
            type: actionType.GET_ORGANIZATION_LOGS,
            payload: {
                api,
                name,
                jwt,
                params,
            },
        })
    }

    getTable = () => {
        let { colDef, logs, pageDef } = this.props

        if (logs.length > 0) {
            let data = prepTableData({
                colDef,
                rawData: logs,
            })
            return (
                <Table
                    cols={colDef}
                    onDetails={this.getDetails}
                    onOptionsChange={this.changeTableOptions}
                    options={pageDef}
                    data={data}
                />
            )
        } else {
            return <NoDataMessage />
        }
    }

    getModal = () => {
        const { modal } = this.state
        if (modal.data) {
            return (
                <LogModal
                    open={modal.open}
                    closeModal={this.onCloseModal}
                    data={modal.data}
                />
            )
        } else {
            return null
        }
    }

    getContent = () => {
        let { loading, error } = this.props
        if (error) {
            this.showError(error)
            return null
        } else if (loading) {
            return <Spinner />
        } else {
            return (
                <React.Fragment>
                    {this.getTable()}
                    {this.getModal()}
                </React.Fragment>
            )
        }
    }

    render() {
        return <div>{this.getContent()}</div>
    }

    componentWillUnmount() {
        let { dispatch } = this.props
        dispatch({
            type: actionType.RESET_ORGANIZATION,
        })
    }
}

// -------------- REDUX CONNECTION ---------------------
/**
 * Map redux store states to local component properties
 * @param state: object, redux store object
 */
const mapStateToProps = (state) => {
    let props = {
        loading: true,
        colDef: state.organization.logs.colDef,
        logs: state.organization.logs.items,
        pageDef: state.organization.logs.pageDef,
        error: state.organization.logs.error,
        name: state.organization.logs.name,
        jwt: state.organization.logs.jwt,
        api: state.organization.logs.api,
    }
    if (
        state.organization.logs.items.length > 0 ||
        state.organization.logs.error ||
        state.organization.logs.jwt
    ) {
        props.loading = state.loader.show
    }
    return props
}

export default connect(mapStateToProps)(ViewRecordsPage)
