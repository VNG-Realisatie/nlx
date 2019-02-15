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
                const document = res.type === 'OpenAPI2' ?
                  JSON.parse(window.atob(res.document)) :
                  null;

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
                    theme: {
                        typography: {
                            fontSize: '14px',
                            lineHeight: '1.5',
                            fontWeightRegular: '300',
                            fontFamily: '"Source Sans Pro", sans-serif',
                            smoothing: 'antialiased',
                            optimizeSpeed: true,
                            headings: {
                                fontFamily: '"Source Sans Pro", sans-serif',
                            },
                            code: {
                                fontSize: '12px',
                                fontFamily: '"Fira Code", monospaced',
                                color: '#e83e8c',
                                backgroundColor: '#ffffff',
                            },
                        },
                        colors: {
                            primary: {
                                main: '#3d83fa',
                            },
                            success: {
                                main: '#00aa13',
                            },
                            warning: {
                                main: '#f1c400',
                            },
                            error: {
                                main: '#e53935',
                            },
                            text: {
                                primary: '#000000',
                            },
                            responses: {
                                redirect: '#ffa500',
                                info: '#87ceeb',
                            },
                            http: {
                                get: '#6bbd5b',
                                post: '#248fb2',
                                put: '#9b708b',
                                options: '#d3ca12',
                                patch: '#e09d43',
                                delete: '#e27a7a',
                                basic: '#999',
                                link: '#31bbb6',
                                head: '#c167e4',
                            },
                        },
                        menu: {
                            width: '260px',
                            backgroundColor: '#f8f9fa',
                        },
                        rightPanel: {
                            backgroundColor: '#ffffff',
                            width: '40%',
                        },
                    },
                    scrollYOffset: '80',
                    hideLoading: true,
                }}
              />
          </div>
        )
    }
}

export default DocumentationPage
