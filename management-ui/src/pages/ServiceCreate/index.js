// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import PageTemplate from '../../components/PageTemplate'
import ServiceForm from '../../components/ServiceForm'

export default class ServiceCreate extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            error: null,
        }

        this.onSubmit = this.onSubmit.bind(this)
    }

    onSubmit(data) {
        const { history } = this.props

        this.postService(data)
            .then(() => {
                history.push('/services')
            })
            .catch((error) => {
                this.setState({ error: error.message })
            })
    }

    postService(data) {
        return fetch('/api/v1/services', {
            method: 'POST',
            body: JSON.stringify(data),
        }).then((response) => {
            if (response.ok) {
                return response.json()
            } else {
                throw new Error('Could not submit service. Please try again.')
            }
        })
    }

    render() {
        const { error } = this.state

        return (
            <PageTemplate>
                <h1>Create service</h1>
                <ServiceForm onSubmit={this.onSubmit} />
                {error ? <p data-test="error">{error}</p> : null}
            </PageTemplate>
        )
    }
}
