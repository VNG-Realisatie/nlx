// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { Link } from 'react-router-dom'
import { Button } from '../../components/Form'
import PageTemplate from '../../components/PageTemplate'
import ServiceList from '../../components/ServiceList'
import { StyledTopButtons } from './index.styles'

export default class ServiceListPage extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            result: {},
            error: null,
        }
    }

    componentDidMount() {
        this.fetchServices()
            .then((result) => {
                this.setState({ result: result })
            })
            .catch((error) => {
                this.setState({ error: error.message })
            })
    }

    fetchServices() {
        return fetch('/api/v1/services').then((response) => {
            if (response.ok) {
                return response.json()
            } else {
                throw new Error(
                    'Could not retrieve service list. Please try again.',
                )
            }
        })
    }

    render() {
        return (
            <PageTemplate>
                <StyledTopButtons>
                    <Link to="/services/create">
                        <Button data-test="create-service">
                            + Create service
                        </Button>
                    </Link>
                </StyledTopButtons>
                {!this.state.error ? (
                    <ServiceList result={this.state.result} />
                ) : (
                    <p data-test="error">{this.state.error}</p>
                )}
            </PageTemplate>
        )
    }
}
