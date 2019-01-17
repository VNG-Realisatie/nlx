import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Spinner } from '@commonground/design-system'

import Table from '../../components/Table'
import LogModal from '../../components/LogModal'
import { prepTableData } from '../../utils/appUtils'
import * as actionType from '../../store/actions'

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

    getTable = () => {
        let { colDef, logs } = this.props

        if (logs.length > 0) {
            return (
                <Table
                    cols={colDef}
                    onDetails={this.getDetails}
                    data={prepTableData({
                        colDef,
                        rawData: logs,
                    })}
                />
            )
        } else {
            return <p>No records - empty placeholder - IMPROVE</p>
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
        error: state.organization.logs.error,
        name: state.organization.logs.name,
        jwt: state.organization.logs.jwt,
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
