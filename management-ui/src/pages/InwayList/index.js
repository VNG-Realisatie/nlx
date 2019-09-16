// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import PageTemplate from '../../components/PageTemplate'
import InwayList from '../../components/InwayList'

export default class InwayListPage extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            result: {},
            error: null,
        }
    }

    componentDidMount() {
        this.fetchInways()
            .then((result) => {
                this.setState({ result: result })
            })
            .catch((error) => {
                this.setState({ error: error.message })
            })
    }

    fetchInways() {
        return fetch('/api/v1/inways').then((response) => {
            if (response.ok) {
                return response.json()
            } else {
                throw new Error(
                    'Could not retrieve inway list. Please try again.',
                )
            }
        })
    }

    render() {
        return (
            <PageTemplate>
                {!this.state.error ? (
                    <InwayList result={this.state.result} />
                ) : (
                    <p data-test="error">{this.state.error}</p>
                )}
            </PageTemplate>
        )
    }
}
