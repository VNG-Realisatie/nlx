import React, { Component } from 'react'
import axios from 'axios'
import Table from './Table'
import LogModal from './LogModal'
import { prepTableData } from '../utils/appUtils'

const colDef = [
    {
        id: 'date',
        label: 'Datum',
        width: 100,
        src: 'created',
        type: 'date',
        disablePadding: true,
    },
    {
        id: 'time',
        label: 'Tijd',
        src: 'created',
        type: 'time',
        disablePadding: false,
    },
    {
        id: 'source',
        label: 'Opgevraagd door',
        src: 'source_organization',
        type: 'string',
        disablePadding: false,
    },
    {
        id: 'destination',
        label: 'Opgevraagd bij',
        src: 'destination_organization',
        type: 'string',
        disablePadding: false,
    },
]

export default class OrganizationRecords extends Component {
    constructor(props) {
        super(props)

        this.state = {
            records: [],
            error: false,
            modal: {
                open: false,
                data: null,
            },
        }
    }

    getDetails = (id) => {
        this.setState({
            modal: {
                open: true,
                data: this.state.records[id],
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

    componentDidMount() {
        axios({
            method: 'post',
            url: `${this.props.organization.insight_log_endpoint}/fetch`,
            headers: { 'content-type': 'text/plain' },
            data: this.props.jwt,
        }).then(
            (response) => {
                this.setState({ records: response.data.records })
            },
            (e) => {
                console.error(e)
                this.setState({ error: true })
            },
        )
    }

    render() {
        let table
        if (this.state.records) {
            table = (
                <Table
                    cols={colDef}
                    onDetails={this.getDetails}
                    data={prepTableData({
                        colDef,
                        rawData: this.state.records,
                    })}
                />
            )
        } else {
            table = <p>No information available</p>
        }

        const { modal } = this.state

        return (
            <div>
                {table}
                {modal.data && (
                    <LogModal
                        open={modal.open}
                        closeModal={this.onCloseModal}
                        data={modal.data}
                    />
                )}
            </div>
        )
    }
}
