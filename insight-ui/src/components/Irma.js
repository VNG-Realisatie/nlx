import React, { Component } from 'react'
import { QRCode } from 'react-qr-svg'
import axios from 'axios'

let interval

export default class Irma extends Component {
    constructor(props) {
        super(props)

        this.state = {
            qrCode: "",
            error: false
        }
    }

    componentDidMount() {
        document.body.onkeyup = (e) => {
            if (e.keyCode === 32) {
                clearInterval(interval)
                this.props.afterLogin(organization, "spacebar-magic")
            }
        }


        const { organization } = this.props

        axios({
            method: 'post',
            url: `${organization.insight_log_endpoint}/generateJWT`,
            data: {
                dataSubjects: [
                    'burgerservicenummer'
                ]
            }
        }).then(response => {
            const firstJWT = response.data

            axios({
                method: 'post',
                url: `${organization.insight_irma_endpoint}/api/v2/verification/`,
                headers: { 'content-type': 'text/plain' },
                data: firstJWT
            }).then(response => {
                let irmaVerificationRequest = response.data
                const u = irmaVerificationRequest['u']

                // prepend IrmaVerifiationRequest with URL
                irmaVerificationRequest['u'] = `${organization.insight_irma_endpoint}/api/v2/verification/${u}`

                const qrCode = JSON.stringify(irmaVerificationRequest)

                this.setState({ qrCode })

                interval = setInterval(() => {
                    axios({
                        method: 'get',
                        url: `${organization.insight_irma_endpoint}/api/v2/verification/${u}/status`
                    }).then(response => {
                        if (response.data === "DONE") {
                            clearInterval(interval)
                            axios({
                                method: 'get',
                                url: `${organization.insight_irma_endpoint}/api/v2/verification/${u}/getproof`
                            }).then(response => {
                                this.props.afterLogin(organization, response.data)
                            })
                        }
                    }).catch(error => {
                        this.setState({ error: true })
                        clearInterval(interval)
                    })
                }, 1000)
            })
        })
    }

    componentWillUnmount() {
        clearInterval(interval)
    }

    render() {
        return (
            <div>
                Please login with Irma<br /><br />
                { this.state.error && (
                    <b>Something went wrong, refresh and try again</b>
                )}
                { this.state.qrCode &&
                    <QRCode
                        bgColor="#FFFFFF"
                        fgColor="#000000"
                        level="Q"
                        style={{ width: 256 }}
                        value={this.state.qrCode}
                    />
                }
            </div>
        )
    }
}