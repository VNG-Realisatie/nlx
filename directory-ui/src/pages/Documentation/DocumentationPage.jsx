import React, { Component } from 'react'
import { RedocStandalone } from 'redoc'
import axios from 'axios'
import { Spinner } from '@commonground/design-system'

import ErrorPage from '../ErrorPage/ErrorPage'

import './DocumentationPage.scss'

class DocumentationPage extends Component {
    constructor(props) {
        super(props)

        this.state = {
            loading: true,
            type: null,
            document: null,
        }
    }

    componentDidMount() {
        const { match } = this.props

        if (!match) {
            return
        }

        axios
            .get(
                `/api/directory/get-service-api-spec/${encodeURIComponent(
                    match.params.organization_name,
                )}/${encodeURIComponent(match.params.service_name)}`,
            )
            .then((res) => {
                let document
                switch (res.data.type) {
                    case 'OpenAPI2':
                        document = JSON.parse(window.atob(res.data.document))
                        break
                    default:
                        document = null
                }

                this.setState({
                    loading: false,
                    type: res.data.type,
                    document,
                })
            })
            .catch((e) => {
                this.setState({
                    loading: false,
                    error: true,
                })
            })
    }

    render() {
        if (this.state.loading) {
            return <Spinner />
        }

        if (!this.state.document) {
            return <ErrorPage />
        }

        return (
            <RedocStandalone
                spec={this.state.document}
                options={{
                    hideDownloadButton: false,
                    theme: {
                        typography: {
                            fontSize: '14px',
                            lineHeight: '1.5',
                            fontWeightRegular: '300',
                            fontFamily: '"Muli", sans-serif',
                            smoothing: 'antialiased',
                            optimizeSpeed: true,
                            headings: {
                                fontFamily: '"Muli", sans-serif',
                            },
                            code: {
                                fontSize: '12px',
                                fontFamily: '"Fira Code", monospaced',
                                color: '#e83e8c',
                                backgroundColor: '#f8f9fa',
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
        )
    }
}

export default DocumentationPage
