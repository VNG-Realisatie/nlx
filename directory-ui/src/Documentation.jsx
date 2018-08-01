import React from 'react'
import { RedocStandalone } from 'redoc'
import axios from 'axios'
import './static/css/redoc-override.css'

export default class Doc extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            loading: true,
            type: null,
            document: null
        }
    }

    componentDidMount() {
        const { match } = this.props

        axios.get(`/api/directory/get-service-api-spec/${encodeURIComponent(match.params.organization_name)}/${encodeURIComponent(match.params.service_name)}`)
            .then((res) => {
                this.setState({
                    loading: false,
                    type: res.data.type,
                    document: window.atob(res.data.document)
                })
            })
            .catch(e => {
                this.setState({
                    loading: false,
                    error: true
                })
            })
    }

    render() {
        if (this.state.loading) {
            return (
                <div>Loading...</div>
            )
        }

        let spec
        switch (this.state.type) {
            case "OpenAPI2":
                spec = JSON.parse(this.state.document)
                break
            default:
                return false
        }

        if (!spec) {
            return (
                <div>Could not load document :(</div>
            )
        }

        return (
            <RedocStandalone
                spec={JSON.parse(this.state.document)}
                options={{
                    theme: {
                        baseFont: {
                            size: '14px',
                            lineHeight: '1.5',
                            weight: '300',
                            family: '"Muli", sans-serif',
                            smoothing: 'antialiased',
                            optimizeSpeed: true,
                        },
                        headingsFont: {
                            family: '"Muli", sans-serif',
                        },
                        code: {
                            fontSize: '12px',
                            fontFamily: '"Fira Code", monospaced',
                        },
                        colors: {
                            main: '#3d83fa',
                            success: '#00aa13',
                            redirect: '#ffa500',
                            error: '#e53935',
                            info: '#87ceeb',
                            text: '#000000',
                            code: '#e83e8c',
                            codeBg: '#f8f9fa',
                            warning: '#f1c400',
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
                            }
                        },
                        menu: {
                            width: '260px',
                            backgroundColor: '#f8f9fa',
                        },
                        rightPanel: {
                            backgroundColor: '#ffffff',
                            width: '40%',
                        }
                    },
                    scrollYOffset: '56',
                    hideLoading: true
                }}
            />
        )
    }
}
