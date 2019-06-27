// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'
import { RedocStandalone } from 'redoc'
import { Spinner } from '@commonground/design-system'

import ErrorMessage from '../../components/ErrorMessage/ErrorMessage'

class DocumentationPage extends Component {
    constructor(props) {
        super(props)

        this.state = {
            loading: true,
            type: null,
            document: null,
        }
    }

    fetchServiceApiSpec(name, organization) {
        const urlSafeName = encodeURIComponent(name)
        const urlSafeOrganization = encodeURIComponent(organization)

        const url =`/api/directory/get-service-api-spec/${urlSafeOrganization}/${urlSafeName}`
        return fetch(url).then(res => res.json())
    }

    componentDidMount() {
        const { match } = this.props

        if (!match) {
            return
        }

        const { organization_name, service_name } = match.params

        this
            .fetchServiceApiSpec(service_name, organization_name)
            .then(res => {
                const document = JSON.parse(window.atob(res.document))

                this.setState({
                    loading: false,
                    type: res.type,
                    document,
                })
            })
            .catch(() => {
                this.setState({
                    loading: false,
                    error: true,
                })
            })
    }

    render() {
        const { loading, document } = this.state

        if (loading) {
            return <Spinner />
        }

        if (!document) {
            return <ErrorMessage />
        }

        return (
          <div style={({ background: '#ffffff' })}>
              <RedocStandalone
                spec={document}
                options={{
                    hideDownloadButton: false,
                    hideLoading: true,
                }}
              />
          </div>
        )
    }
}

export default DocumentationPage
