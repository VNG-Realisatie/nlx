// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React, { Component } from 'react'
import PageTemplate from '../../components/PageTemplate'
import ServiceForm from '../../components/ServiceForm'

export default class ServiceUpdate extends Component {
    constructor(props) {
        super(props)

        this.state = {
            data: null,
            error: null,
        }

        this.onSubmit = this.onSubmit.bind(this)
        this.onDelete = this.onDelete.bind(this)
    }

    componentDidMount() {
        const { match } = this.props

        this.fetchService(match.params.name)
            .then((data) => {
                this.setState({ data })
            })
            .catch((error) => {
                this.setState({ error: error.message })
            })
    }

    fetchService(name) {
        return fetch(`/api/v1/services/${encodeURIComponent(name)}`).then(
            (response) => {
                if (response.ok) {
                    return response.json()
                } else {
                    throw new Error(
                        'Could not fetch service. Please try again.',
                    )
                }
            },
        )
    }

    onSubmit(data) {
        const { history } = this.props

        this.putService(this.state.data.name, data)
            .then((data) => {
                if (data.error) {
                    this.setState({ error: data.error })
                    return
                }

                history.push('/services')
            })
            .catch((error) => {
                this.setState({ error: error.message })
            })
    }

    putService(name, data) {
        return fetch(`/api/v1/services/${encodeURIComponent(name)}`, {
            method: 'PUT',
            body: JSON.stringify(data),
        }).then((response) => {
            if (response.ok) {
                return response.json()
            } else {
                throw new Error('Could not update service. Please try again.')
            }
        })
    }

    onDelete() {
        const { history } = this.props
        const { data } = this.state

        this.deleteService(data.name)
            .then(() => {
                history.push('/services')
            })
            .catch((error) => {
                this.setState({ error: error.message })
            })
    }

    deleteService(name) {
        return fetch(`/api/v1/services/${encodeURIComponent(name)}`, {
            method: 'DELETE',
        }).then((response) => {
            if (response.ok) {
                return response.json()
            } else {
                throw new Error('Could not delete service. Please try again.')
            }
        })
    }

    render() {
        const { match } = this.props

        return (
            <PageTemplate>
                <h1>Update service {match.params.name}</h1>
                {this.state.data && (
                    <ServiceForm
                        initialValues={this.state.data}
                        onSubmit={this.onSubmit}
                        onDelete={this.onDelete}
                    />
                )}
                {this.state.error ? (
                    <p data-test="error">{this.state.error}</p>
                ) : null}
            </PageTemplate>
        )
    }
}
